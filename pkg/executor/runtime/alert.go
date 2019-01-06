package runtime

import (
	"context"
	"errors"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	. "github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"k8s.io/klog/glog"
	"regexp"
	"strings"
	"time"
)

// this chan is used to control corresponding goroutine
type RuntimeAlert struct {
	SigCh         chan pb.Message_Signal
	StatusCh      chan string
	AlertConfigID string
	UpdatedAt     time.Time
	CreatedAt     time.Time
	err           error
}

type AlertInfo struct {
	acID          string
	alertConfig   *models.AlertConfig
	uri           string
	resourceNames []string
	freq          []int32
	currentFreq   []int32
	ruleEnable    []bool
}

// goroutine status
type StatusType string

const (
	Alive  StatusType = "alive"
	Dead   StatusType = "dead"
	Unkonw StatusType = "unknow"
)

type AlertStatus struct {
	Status    StatusType
	timestamp time.Time
}

// key is alert config id
// value is alert runtime parameters included channels
var CachedRuntimeAlert = make(map[string]*RuntimeAlert)

func Action(ctx context.Context, msg *Message) error {

	signalx := msg.Signal

	switch signalx {
	case Message_CREATE:
		// create alert by specifig alert config id within one goroutine
		fmt.Println("create alert")

		alertInfo, err := reload(msg.AlertConfigId)

		if err != nil {
			return err
		}

		alert := &RuntimeAlert{
			SigCh:     make(chan pb.Message_Signal),
			CreatedAt: time.Now(),
		}

		CachedRuntimeAlert[msg.AlertConfigId] = alert

		go runAlert(alertInfo, alert)

	case Message_STOP:

	case Message_RELOAD:

	case Message_OTHER:
	}
	return nil
}

func runAlert(alertInfo *AlertInfo, alert *RuntimeAlert) {

	var period = 1
	d := time.Duration(time.Minute * time.Duration(period))
	timer := time.NewTicker(d)
	defer timer.Stop()

	sigCh := alert.SigCh
	for {
		select {
		case sig := <-sigCh:
			if sig == pb.Message_RELOAD {
				var err error
				alertInfo, err = reload(alert.AlertConfigID)

				if err != nil {
					glog.Errorln(err.Error())
					alert.err = err
				}

			} else if sig == pb.Message_STOP {
				glog.Infoln("runtime alert was terminated, alert_config_id is: ", alert.AlertConfigID)
				sigCh <- pb.Message_STOP
				return
			}

		case <-timer.C:
			// retrieve metrics from moniting center
			alertConfig := alertInfo.alertConfig
			enable := isAlertEnable(alertConfig.EnableStart, alertConfig.EnableEnd)
			if !enable {
				continue
			}

			freq := alertInfo.freq
			currentFreq := alertInfo.currentFreq
			ruleEnable := alertInfo.ruleEnable

			l := len(freq)
			var evaluatedRuleIndx []int
			for i := 0; i < l; i++ {
				if !ruleEnable[i] {
					continue
				}

				currentFreq[i] = (currentFreq[i] + 1) % freq[i]
				if currentFreq[i] == 0 {
					// means this rule whill be evaluated in turn
					evaluatedRuleIndx = append(evaluatedRuleIndx, i)
				}
			}

			if len(evaluatedRuleIndx) == 0 {
				continue
			}

			// get metrics name from evaluated rule
			metricNames := getMetricsName(evaluatedRuleIndx, alertConfig.AlertRuleGroup.AlertRules)

		}

	}
}

func getMetricsName(index []int, rules []*models.AlertRule) []string {
	var metricNames []string
	for i := 0; i < len(index); i++ {
		metricNames = append(metricNames, rules[i].MetricName)
	}

	return metricNames
}

func isAlertEnable(start time.Time, end time.Time) bool {
	now := time.Now()

	if start.Before(now) && now.Before(end) {
		return true
	}

	return false
}

// reload alert config
func reload(acID string) (*AlertInfo, error) {
	// get alert config by id from DB
	alertConfig, err := models.GetAlertConfig(&models.AlertConfig{AlertConfigID: acID})

	if err != nil {
		return nil, err
	}

	alertRules := alertConfig.AlertRuleGroup
	if alertRules == nil || len(alertRules.AlertRules) == 0 {
		return nil, errors.New("at least one alert rule must be specified")
	}

	receivers := alertConfig.ReceiverGroup
	if receivers == nil || len(*receivers.Receivers) == 0 {
		return nil, errors.New("at least one receiver must be specified")
	}

	resources := alertConfig.ResourceGroup
	if resources == nil || len(resources.Resources) == 0 {
		return nil, errors.New("at least one resource must be specified")
	}

	uri, resName, err := getResourcesSpec(alertConfig.ResourceGroup)

	if err != nil {
		return nil, err
	}

	l := len(alertRules.AlertRules)
	var freq = make([]int32, l)
	var currentFreq = make([]int32, l)
	var ruleEnable = make([]bool, l)

	for i := 0; i < l; i++ {
		rule := alertRules.AlertRules[i]
		freq[i] = rule.Period
		ruleEnable[i] = rule.Enable
	}

	return &AlertInfo{
		acID:          acID,
		alertConfig:   alertConfig,
		resourceNames: resName,
		uri:           uri,
		ruleEnable:    ruleEnable,
		freq:          freq,
		currentFreq:   currentFreq,
	}, nil
}

// get resources uri and resource names
func getResourcesSpec(resourceGroup *models.ResourceGroup) (string, []string, error) {

	var uriTmpl models.ResourceUriTmpl
	jsonutil.Unmarshal(resourceGroup.URIParams, &uriTmpl)

	if uriTmpl.UriTmpl == "" {
		return "", nil, errors.New("resource uri tmpl must be specified")
	}

	uriParams := uriTmpl.Params
	// find uriTmpls by resource type id
	resourceType, err := models.GetResourceType(&models.ResourceType{ResourceTypeID: resourceGroup.ResourceTypeID})
	if err != nil {
		return "", nil, err
	}

	monitorHost := strings.Trim(resourceType.MonitorCenterHost, " ")
	monitorPort := resourceType.MonitorCenterPort
	if monitorHost == "" || monitorPort == 0 {
		// get monitoring host:port by product id
		prod, err := models.GetProduct(&models.Product{ProductID: resourceType.ProductID})
		if err != nil {
			return "", nil, err
		}

		monitorHost = strings.Trim(prod.MonitorCenterHost, " ")
		monitorPort = prod.MonitorCenterPort
	}

	if monitorHost == "" || monitorPort == 0 {
		return "", nil, errors.New("monitoring center must be specified")
	}

	var uriTmpls []*models.ResourceUriTmpl
	jsonutil.Unmarshal(resourceType.ResourceURITmpls, &uriTmpls)

	l := len(uriTmpls)

	b := false

	var urlTmpl string
	for i := 0; i < l; i++ {
		u := uriTmpls[i]
		if u != nil {
			// does uriParams match
			storedUriParams := u.Params
			if IsMatch(uriParams, storedUriParams) {
				urlTmpl = u.UriTmpl
				b = true
				break
			}
		}
	}

	if !b {
		return "", nil, errors.New("resource uri parameters dose not match any existing resource uri template")
	}

	uri, err := AssembeURLPrefix(monitorHost, monitorPort, urlTmpl, uriParams)

	if err != nil {
		return "", nil, err
	}

	resources := resourceGroup.Resources
	l = len(resources)
	var resNames []string

	for i := 0; i < l; i++ {
		if resources[i] != nil {
			resNames = append(resNames, resources[i].ResourceName)
		}
	}

	return uri, resNames, nil
}

func AssembeURLPrefix(host string, port int32, uriTmpl string, params map[string]string) (string, error) {
	r, err := regexp.Compile(`\{\w+\}`)
	if err != nil {
		// compile error
		return "", err
	}

	uri := r.ReplaceAllStringFunc(uriTmpl, func(s string) string {
		s = strings.Trim(s, "{")
		s = strings.Trim(s, "}")
		return params[s]
	})

	uri = fmt.Sprintf("%s:%d%s", host, port, uri)
	return uri, nil
}

func IsMatch(p map[string]string, q map[string]string) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {

		if len(p) == len(q) {
			for k := range p {
				if _, ok := q[k]; !ok {
					return false
				}
			}

		} else {
			return false
		}

	} else {
		return false
	}

	return true
}

//
//func DeleteRuntimeAlert(alertConfigID string) error {
//	// first step: need to delete items in related tables
//	err := models.DeleteAlertBindingItem(alertConfigID)
//	// if an error occured, delete runtime alert failed
//	if err != nil {
//		glog.Errorln(err.Error())
//		return err
//	}
//
//	// second step: delete item in CachedRuntimeAlert map
//	if alert, ok := CachedRuntimeAlert[alertConfigID]; ok {
//		if alert != nil {
//			alert.SignalSender <- pb.AlertConfig_Terminate
//			for {
//				sig := <-alert.SignalReceiver
//				// TODO
//				if sig == 0 {
//					glog.Infof("terminate running alert goroutine successfully, alert_config_id is: %s", alertConfigID)
//					delete(CachedRuntimeAlert, alertConfigID)
//					return nil
//				}
//
//			}
//		}
//	}
//	return nil
//}
//
//// does goroutine still alive?
//func GetRuntimeAlertStatus(alertConfigID string) *RuntimeAlertStatus {
//	if alert, ok := CachedRuntimeAlert[alertConfigID]; ok {
//		if alert != nil {
//			alert.StatusCh <- "ping"
//			for {
//				sig := <-alert.StatusCh
//				if sig == "pong" {
//					glog.Infof("alert goroutine is running, alert_config_id is: %s", alertConfigID)
//					return nil
//				}
//			}
//		}
//
//		return &RuntimeAlertStatus{
//			Status:    Alive,
//			timestamp: time.Now(),
//		}
//	}
//
//	return &RuntimeAlertStatus{
//		Status:    Unkonw,
//		timestamp: time.Now(),
//	}
//}
