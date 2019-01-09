package runtime

import (
	"context"
	"errors"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/client"
	"github.com/carmanzhang/ks-alert/pkg/executor/metric"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"k8s.io/klog/glog"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// this chan is used to control corresponding goroutine
type RuntimeAlert struct {
	SigCh         chan pb.Message_Signal
	StatusCh      chan string
	UpdatedAt     time.Time
	CreatedAt     time.Time
	err           error
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
type RuntimeAlertStatus struct {
	sync.RWMutex
	Map map[string]*RuntimeAlert
}

var CachedRuntimeAlert = &RuntimeAlertStatus{
	Map: make(map[string]*RuntimeAlert),
}

func Action(ctx context.Context, msg *pb.Message) error {

	switch msg.Signal {
	case pb.Message_CREATE:
		// create alert by specifig alert config id within one goroutine
		var rtAlert = &RuntimeAlert{
			SigCh:     make(chan pb.Message_Signal),
			CreatedAt: time.Now(),
		}
		// write executor's name to table `alert_config`

		err := reload(msg.AlertConfigId, rtAlert)

		if err != nil {
			return err
		}

		CachedRuntimeAlert.Lock()
		CachedRuntimeAlert.Map[msg.AlertConfigId] = rtAlert
		CachedRuntimeAlert.Unlock()

		go runAlert(rtAlert)

	case pb.Message_STOP:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return errors.New("alert config was not executor by executor")
		}

		runtimeAlert.SigCh <- pb.Message_STOP

		CachedRuntimeAlert.Lock()
		delete(CachedRuntimeAlert.Map, msg.AlertConfigId)
		CachedRuntimeAlert.Unlock()

	case pb.Message_RELOAD:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return errors.New("alert config was not executor by executor")
		}

		runtimeAlert.SigCh <- pb.Message_RELOAD

	default:
	}
	return nil
}

func runAlert(rtAlert *RuntimeAlert) {
	var period = 8
	d := time.Duration(time.Second * time.Duration(period))
	timer := time.NewTicker(d)
	defer timer.Stop()

	sigCh := rtAlert.SigCh
	for {
		select {
		case sig := <-sigCh:
			if sig == pb.Message_RELOAD {
				acID := rtAlert.alertConfig.AlertConfigID
				fmt.Println("runtime rtAlert was reload, alert_config_id is: ")
				var err error
				fmt.Printf("%p,%v", rtAlert, rtAlert)
				err = reload(acID, rtAlert)
				fmt.Printf("%p,%v", rtAlert, rtAlert)

				if err != nil {
					glog.Errorln(err.Error())
					rtAlert.err = err
				}

			} else if sig == pb.Message_STOP {
				fmt.Println("runtime rtAlert was terminated, alert_config_id is: ", rtAlert.alertConfig.AlertConfigID)
				return
			}

		case <-timer.C:

			fmt.Println("new evalted period")
			alertConfig := rtAlert.alertConfig
			// 1. rtAlert config is enable?
			//enable := isAlertEnable(alertConfig.EnableStart, alertConfig.EnableEnd)
			//if !enable {
			//	continue
			//}

			// 2. is there any rules need to execute?
			evaluatedRuleIndx := getExecutingRuleIndx(rtAlert.freq, rtAlert.currentFreq, rtAlert.ruleEnable)
			if len(evaluatedRuleIndx) == 0 {
				continue
			}

			// 3. retrieve metrics from monitoring center using evaluating rules
			rules := alertConfig.AlertRuleGroup.AlertRules
			var ch = make(chan *metric.ResourceMetrics, 15)
			getResourceMetrics(rules, evaluatedRuleIndx, rtAlert.uri, rtAlert.resourceNames, ch)
			close(ch)

			// 4. decide whether a resource trigger a rule
			trigger(ch, rules)

		}

	}
}

func trigger(ch chan *metric.ResourceMetrics, rules []*models.AlertRule) {
	for metricByRule := range ch {
		fmt.Println(metricByRule)
		if metricByRule != nil {
			indx := metricByRule.RuleIndx
			rule := rules[indx]
			consCount := int(rule.ConsecutiveCount)
			condition := rule.ConditionType
			threshold := rule.Threshold
			//matricName := metricByRule.RuleName
			resourceMetric := metricByRule.ResourceMetric
			for resName := range resourceMetric {
				// timestamp and value
				tvs := resourceMetric[resName]
				ll := len(tvs)
				if ll < consCount {
					continue
				}

				f := true
				for i := ll - consCount; i < ll; i++ {
					v, err := strconv.ParseFloat(tvs[i].V, 32)
					if err != nil {
					}
					switch condition {
					case ">=":
						f = f && (v >= float64(threshold))
					case ">":
						f = f && (v > float64(threshold))
					case "<=":
						f = f && (v <= float64(threshold))
					case "<":
						f = f && (v < float64(threshold))
					}

					if !f {
						break
					}
				}

				if f {
					// fired alert on alert_rule rules[indx] and resource resourceMetric[resName]
					fmt.Println("fired alert", rules[indx].MetricName, resName)
				} else {
					// need to change alert status according last alert status
				}

			}
		}

	}
}

// alert rules which are need to execute
func getExecutingRuleIndx(freq, currentFreq []int32, ruleEnable []bool) []int {
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
	fmt.Println("currentFreq: ", currentFreq)
	fmt.Println("Freq: ", freq)
	return evaluatedRuleIndx
}

func getResourceMetrics(rules []*models.AlertRule, evaluatedRuleIndx []int, uri string, resourceNames []string, ch chan *metric.ResourceMetrics) {
	wg := sync.WaitGroup{}
	for i := 0; i < len(evaluatedRuleIndx); i++ {
		j := evaluatedRuleIndx[i]
		metricName := rules[j].MetricName
		// period unit is Minute
		stepInMinute := rules[j].Period
		consCount := rules[j].ConsecutiveCount
		if consCount <= 0 || consCount > metric.MaxConsecutiveCount {
			consCount = metric.ConsecutiveCount
		}

		if stepInMinute <= 0 || stepInMinute > metric.MaxStep {
			stepInMinute = metric.Step
		}

		endTime := time.Now().Truncate(time.Minute).Unix()
		startTime := endTime - int64((consCount+3)*stepInMinute*60)

		wg.Add(1)
		go func() {
			metricStr := client.SendMonitoringRequest(uri, resourceNames, metricName, startTime, endTime, stepInMinute)
			resourceMetrics := metric.GetResourceTimeSeriesMetric(metricStr, metricName, startTime, endTime)
			resourceMetrics.RuleIndx = j
			fmt.Println(">>>:", resourceMetrics)
			fmt.Println(">>>:", rules)
			ch <- resourceMetrics
			wg.Done()
		}()
	}
	wg.Wait()
	//close(ch)
}

func isAlertEnable(start time.Time, end time.Time) bool {
	now := time.Now()

	if start.Before(now) && now.Before(end) {
		return true
	}

	return false
}

// reload alert config
func reload(acID string, rtAlert *RuntimeAlert) error {
	// get alert config by id from DB
	alertConfig, err := models.GetAlertConfig(&models.AlertConfig{AlertConfigID: acID})

	if err != nil {
		return err
	}

	alertRules := alertConfig.AlertRuleGroup
	if alertRules == nil || len(alertRules.AlertRules) == 0 {
		return errors.New("at least one alert rule must be specified")
	}

	receivers := alertConfig.ReceiverGroup
	if receivers == nil || len(*receivers.Receivers) == 0 {
		return errors.New("at least one receiver must be specified")
	}

	resources := alertConfig.ResourceGroup
	if resources == nil || len(resources.Resources) == 0 {
		return errors.New("at least one resource must be specified")
	}

	uri, resName, err := GetResourcesSpec(alertConfig.ResourceGroup)

	if err != nil {
		return err
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

	rtAlert.alertConfig = alertConfig
	rtAlert.resourceNames = resName
	rtAlert.uri = uri
	rtAlert.ruleEnable = ruleEnable
	rtAlert.freq = freq
	rtAlert.currentFreq = currentFreq
	return nil
}

// get resources uri and resource names
func GetResourcesSpec(resourceGroup *models.ResourceGroup) (string, []string, error) {

	var uriTmpl models.ResourceUriTmpl
	jsonutil.Unmarshal(resourceGroup.URIParams, &uriTmpl)

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

	var uriTmpls models.ResourceUriTmpls
	fmt.Println(resourceType.ResourceURITmpls)
	jsonutil.Unmarshal(resourceType.ResourceURITmpls, &uriTmpls)

	l := len(uriTmpls.ResourceUriTmpl)

	b := false

	var urlTmpl string
	for i := 0; i < l; i++ {
		u := uriTmpls.ResourceUriTmpl[i]
		// does uriParams match
		storedUriParams := u.Params
		if IsMatch(uriParams, storedUriParams) {
			urlTmpl = u.UriTmpl
			b = true
			break
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
