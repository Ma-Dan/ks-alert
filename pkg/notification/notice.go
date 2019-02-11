package notification

import (
	"github.com/carmanzhang/ks-alert/pkg/executor/metric"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"time"
)

type Notice struct {
	ResourceName          string            `json:"resource_name"`
	TriggerTime           time.Time         `json:"trigger_time"`
	Rule                  *models.AlertRule `json:"alert_rule"`
	Metrics               *[]metric.TV      `json:"time_series_metrics"`
	CumulateReSendCount   uint32            `json:"current_resend_count"`
	MaxReSendCount        uint32            `json:"max_resend_count"`
	CurrentReSendInterval uint32            `json:"current_resend_interval"`
	NextReSendInterval    uint32            `json:"next_resend_interval"`
	SendNoticeAt          time.Time         `json:"send_notice_at"`
}

type FiredAlertDurations []Duration

type Duration struct {
	FiredAt     time.Time `json:"fired_at"`
	RecoveredAt time.Time `json:"recovered_at"`
}

type CompactedNotice struct {
	ResourceName          string       `json:"resource_name,omitempty"`
	TriggerTime           *time.Time   `json:"trigger_time,omitempty"`
	MetricName            string       `json:"metric_name,omitempty"`
	RuleName              string       `json:"rule_name,omitempty"`
	ConsecutiveCount      int32        `json:"consecutive_count,omitempty"`
	ConditionType         string       `json:"condition_type,omitempty"`
	Threshold             float32      `json:"threshold,omitempty"`
	Unit                  string       `json:"unit,omitempty"`
	Period                int32        `json:"period,omitempty"`
	RepeatSendType        int32        `json:"resend_type,omitempty"`
	Metrics               *[]metric.TV `json:"time_series_metrics,omitempty"`
	CumulateReSendCount   uint32       `json:"current_resend_count,omitempty"`
	MaxReSendCount        uint32       `json:"max_resend_count,omitempty"`
	CurrentReSendInterval uint32       `json:"current_resend_interval,omitempty"`
	NextReSendInterval    uint32       `json:"next_resend_interval,omitempty"`
	SendNoticeAt          *time.Time   `json:"send_notice_at,omitempty"`
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
			TriggerTime:           &n.TriggerTime,
			Unit:                  r.Unit,
			CumulateReSendCount:   n.CumulateReSendCount,
			CurrentReSendInterval: n.CurrentReSendInterval,
			MaxReSendCount:        n.MaxReSendCount,
			NextReSendInterval:    n.NextReSendInterval,
		}

		return jsonutil.Marshal(sn)
	}
}
