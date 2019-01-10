package runtime

import (
	"context"
	"errors"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/client"
	"github.com/carmanzhang/ks-alert/pkg/executor/metric"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"k8s.io/klog/glog"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// this chan is used to control corresponding goroutine
type RuntimeAlertConfig struct {
	SigCh         chan pb.Informer_Signal
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
	firedAlerts   map[string]map[string]time.Time
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
	Map map[string]*RuntimeAlertConfig
}

var CachedRuntimeAlert = &RuntimeAlertStatus{
	Map: make(map[string]*RuntimeAlertConfig),
}

func Action(ctx context.Context, msg *pb.Informer) error {

	switch msg.Signal {
	case pb.Informer_CREATE:
		// create alert by specifig alert config id within one goroutine
		var rtAlert = &RuntimeAlertConfig{
			SigCh:     make(chan pb.Informer_Signal),
			CreatedAt: time.Now(),
		}

		err := rtAlert.reload(msg.AlertConfigId)

		if err != nil {
			return err
		}

		CachedRuntimeAlert.Lock()
		CachedRuntimeAlert.Map[msg.AlertConfigId] = rtAlert
		CachedRuntimeAlert.Unlock()

		go rtAlert.runAlert()

	case pb.Informer_TERMINATE:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return errors.New("alert config was not executor by executor")
		}

		runtimeAlert.SigCh <- pb.Informer_TERMINATE

		CachedRuntimeAlert.Lock()
		delete(CachedRuntimeAlert.Map, msg.AlertConfigId)
		CachedRuntimeAlert.Unlock()

	case pb.Informer_RELOAD:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return errors.New("alert config was not executor by executor")
		}

		runtimeAlert.SigCh <- pb.Informer_RELOAD

	default:
	}
	return nil
}

func (rtAlert *RuntimeAlertConfig) runAlert() {
	var period = 8
	d := time.Duration(time.Second * time.Duration(period))
	timer := time.NewTicker(d)
	defer timer.Stop()

	sigCh := rtAlert.SigCh
	for {
		select {
		case sig := <-sigCh:
			if sig == pb.Informer_RELOAD {
				acID := rtAlert.alertConfig.AlertConfigID
				fmt.Println("runtime rtAlert was reload, alert_config_id is: ", acID)
				var err error
				fmt.Printf("%p,%v", rtAlert, rtAlert)
				err = rtAlert.reload(acID)
				fmt.Printf("%p,%v", rtAlert, rtAlert)

				if err != nil {
					glog.Errorln(err.Error())
					rtAlert.err = err
				}

			} else if sig == pb.Informer_TERMINATE {
				fmt.Println("runtime rtAlert was terminated, alert_config_id is: ", rtAlert.alertConfig.AlertConfigID)
				return
			}

		case <-timer.C:
			// TODO need to add exception catcher function
			fmt.Println("new evalted period", len(CachedRuntimeAlert.Map), runtime.NumGoroutine())
			alertConfig := rtAlert.alertConfig

			// 0. check this alert_config's hostid, whether is consistency with this node or not
			// this step is important, because when network became unreachable, alert system will launch another node(pod)
			// to take replace of this node(pod), it means the same alert configs will be executed in a newly created node
			// if both are not consistency, goroutine will be exist.
			hostID, err := models.GetAlertConfigBindingHost(alertConfig.AlertConfigID)
			if err != nil {
				glog.Errorln(err.Error())
			}

			if hostID != option.HostID {
				return
			}

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
			rtAlert.trigger(ch)
		}

	}
}

func (rtAlert *RuntimeAlertConfig) trigger(ch chan *metric.ResourceMetrics) {
	rules := rtAlert.alertConfig.AlertRuleGroup.AlertRules
	for metricByRule := range ch {
		fmt.Println(metricByRule)
		if metricByRule != nil {
			indx := metricByRule.RuleIndx
			rule := rules[indx]
			consCount := int(rule.ConsecutiveCount)
			condition := rule.ConditionType
			threshold := rule.Threshold
			resourceMetric := metricByRule.ResourceMetric

			for resName := range resourceMetric {
				// timestamp and value
				tvs := resourceMetric[resName]
				ll := len(tvs)
				if ll < consCount {
					continue
				}

				f := true
				// `tvs` will be send to user and insert into `alert_histories` if alert fired
				tvs = tvs[ll-consCount:]
				for i := 0; i < len(tvs); i++ {
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

				ruleID := rule.AlertRuleID
				firedAlerts := rtAlert.firedAlerts
				if firedAlerts == nil {
					firedAlerts = make(map[string]map[string]time.Time)
				}

				if f {
					// add fired alert info into `firedAlerts` if not exist
					if firedRules, ok := firedAlerts[resName]; ok && firedRules != nil {
						if _, ok := firedRules[ruleID]; !ok {
							firedRules[ruleID] = time.Now()
						}

					} else {
						firedRules = make(map[string]time.Time)
						firedRules[ruleID] = time.Now()
					}

					// insert alert fired item into `alert_history`
					ah := rtAlert.makeAlertHistoryItem(indx, resName, tvs, f)
					_, err := models.CreateAlertHistory(ah)

					if err != nil {
						fmt.Println(err.Error())
					}

					fmt.Println("fired alert", rules[indx].MetricName, resName)

				} else {
					// checking whether fired alert has been recovery or not
					// remove alert fired alert info from `firedAlerts` if exist
					if firedRules, ok := firedAlerts[resName]; ok && firedRules != nil {
						if _, ok := firedRules[ruleID]; ok {
							delete(firedRules, ruleID)
						}

						if len(firedRules) == 0 {
							delete(firedAlerts, resName)
						}
					}
					// insert alert recovery item into `alert_histories`
					ah := rtAlert.makeAlertHistoryItem(indx, resName, tvs, f)
					_, err := models.CreateAlertHistory(ah)
					if err != nil {
						fmt.Println(err.Error())
					}

					fmt.Println("recovery alert", rules[indx].MetricName, resName)
				}

				// repeat send
				// check it's the time for the fired alert is going to send
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
func (rtAlert *RuntimeAlertConfig) reload(acID string) error {
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

	rtAlert.alertConfig = alertConfig
	rtAlert.resourceNames = resName
	rtAlert.uri = uri

	// TODO need to keep consistency with status before this reload
	l := len(alertRules.AlertRules)
	var freq = make([]int32, l)
	var currentFreq = make([]int32, l)
	var ruleEnable = make([]bool, l)

	for i := 0; i < l; i++ {
		rule := alertRules.AlertRules[i]
		freq[i] = rule.Period
		ruleEnable[i] = rule.Enable
	}

	rtAlert.freq = freq
	rtAlert.currentFreq = currentFreq
	rtAlert.ruleEnable = ruleEnable

	firedAlerts := rtAlert.firedAlerts
	if firedAlerts == nil {
		rtAlert.firedAlerts = make(map[string]map[string]time.Time)
	} else {
		removeOldRulesAndResources(resources, firedAlerts, alertRules)
	}

	return nil
}

// is's necessary to clean up `firedAlerts`, for the reason reloading may discard old rules and add new rules
// `firedAlerts`  saved the fired alert triggered by an existing rule on a existing resource
// `firedAlerts` field was arranged by map[resource_name](map[alert_rule_id]2019.1.1-00:00:00)
// 0. remove old resources
func removeOldRulesAndResources(resources *models.ResourceGroup, firedAlerts map[string]map[string]time.Time, alertRules *models.AlertRuleGroup) {
	l := len(resources.Resources)
	for i := 0; i < l; i++ {
		rs := resources.Resources[i]
		if _, ok := firedAlerts[rs.ResourceName]; !ok {
			delete(firedAlerts, rs.ResourceName)
		}
	}
	// 1. remove old rules
	l = len(alertRules.AlertRules)
	for k := range firedAlerts {
		currRuleMap := firedAlerts[k]
		for i := 0; i < l; i++ {
			ruleID := alertRules.AlertRules[i].AlertRuleID
			if _, ok := currRuleMap[ruleID]; !ok {
				delete(firedAlerts[k], ruleID)
			}
		}
	}
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

func (rtAlert *RuntimeAlertConfig) makeAlertHistoryItem(ruleIndx int, resName string, tvs []metric.TV, isFired bool) *models.AlertHistory {
	ac := rtAlert.alertConfig
	ruleGroup := ac.AlertRuleGroup
	resGroup := ac.ResourceGroup
	recvGroup := ac.ReceiverGroup

	ah := models.AlertHistory{}
	firedRule := ruleGroup.AlertRules[ruleIndx]

	ah.AlertConfigID = ac.AlertConfigID
	//ah.ProductID =
	ah.ReceiverGroupID = ac.ReceiverGroupID
	ah.ReceiverGroup = fmt.Sprintf("%v", recvGroup)
	ah.ReceiverGroupName = recvGroup.ReceiverGroupName

	ah.ResourceGroupID = ac.ResourceGroupID
	ah.ResourceGroupName = resGroup.ResourceGroupName
	ah.AlertedResource = resName

	ah.AlertRuleGroupID = ruleGroup.AlertRuleGroupID
	ah.TriggerMetricName = firedRule.MetricName
	ah.AlertRuleGroupName = ruleGroup.AlertRuleGroupName

	ah.SeverityID = ac.SeverityID
	ah.SeverityCh = ac.SeverityCh

	ah.RepeatSendType = uint32(firedRule.RepeatSendType)
	ah.InitRepeatSendInterval = firedRule.InitRepeatSendInterval
	ah.MaxRepeatSendCount = firedRule.MaxRepeatSendCount

	ah.CreatedAt = time.Now()
	ah.UpdatedAt = time.Now()
	lastEvalutedTime := tvs[len(tvs)-1].T
	if isFired {
		ah.AlertFiredAt = time.Unix(int64(lastEvalutedTime), 0)
	} else {
		ah.AlertRecoveryAt = time.Unix(int64(lastEvalutedTime), 0)
	}

	// TODO
	ah.CurrentRepeatSendInterval = 3
	ah.CurrentRepeatSendInterval = 3
	ah.SilenceStartAt = time.Now()
	ah.SilenceEndAt = time.Now()
	ah.Cause = ""

	ah.RequestNotificationStatus = ""
	ah.NotificationSendAt = time.Now()

	return &ah
}
