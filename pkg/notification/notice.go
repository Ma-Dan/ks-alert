package notification

import (
	"kubesphere.io/ks-alert/pkg/executor/metric"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	"time"
)

type Notice struct {
	ResourceName          string               `json:"resource_name"`
	TriggerTime           time.Time            `json:"trigger_time"`
	Rule                  *models.AlertRule    `json:"alert_rule"`
	Metrics               *[]metric.TV         `json:"time_series_metrics"`
	CumulateReSendCount   uint32               `json:"current_resend_count"`
	MaxReSendCount        uint32               `json:"max_resend_count"`
	CurrentReSendInterval uint32               `json:"current_resend_interval"`
	NextReSendInterval    uint32               `json:"next_resend_interval"`
	SendNoticeAt          time.Time            `json:"send_notice_at"`
	FiredAlertDurations   *FiredAlertDurations `json:"fired_alert_durations"`
}

type FiredAlertDurations []Duration

type Duration struct {
	FiredAt     time.Time `json:"fired_at"`
	RecoveredAt time.Time `json:"recovered_at"`
}

type CompactedNotice struct {
	ResourceName          string               `json:"resource_name"`
	TriggerTime           time.Time            `json:"trigger_time"`
	MetricName            string               `json:"metric_name"`
	RuleName              string               `json:"rule_name"`
	ConsecutiveCount      int32                `json:"consecutive_count"`
	ConditionType         string               `json:"condition_type"`
	Threshold             float32              `json:"threshold"`
	Unit                  string               `json:"unit"`
	Period                int32                `json:"period"`
	RepeatSendType        int32                `json:"resend_type"`
	Metrics               *[]metric.TV         `json:"time_series_metrics"`
	CumulateReSendCount   uint32               `json:"current_resend_count"`
	MaxReSendCount        uint32               `json:"max_resend_count"`
	CurrentReSendInterval uint32               `json:"current_resend_interval"`
	NextReSendInterval    uint32               `json:"next_resend_interval"`
	SendNoticeAt          time.Time            `json:"send_notice_at"`
	FiredAlertDurations   *FiredAlertDurations `json:"fired_alert_durations"`
}

func (n Notice) MakeNotice(detailed bool) string {
	if detailed {
		notice := jsonutil.Marshal(n)
		return notice
	} else {
		r := n.Rule
		sn := CompactedNotice{
			Metrics:               n.Metrics,
			Threshold:             r.Threshold,
			MetricName:            r.MetricName,
			ResourceName:          n.ResourceName,
			ConditionType:         r.ConditionType,
			ConsecutiveCount:      r.ConsecutiveCount,
			Period:                r.Period,
			RepeatSendType:        r.RepeatSendType,
			RuleName:              r.AlertRuleName,
			TriggerTime:           n.TriggerTime,
			Unit:                  r.Unit,
			CumulateReSendCount:   n.CumulateReSendCount,
			CurrentReSendInterval: n.CurrentReSendInterval,
			MaxReSendCount:        n.MaxReSendCount,
			NextReSendInterval:    n.NextReSendInterval,
		}

		return jsonutil.Marshal(sn)
	}
}
