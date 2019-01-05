package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"k8s.io/klog/glog"
	"time"
)

type RuntimeAlert struct {
	// this chan is used to control corresponding goroutine
	SignalSender   chan pb.AlertConfig_Signal
	AlertConfigID  string
	SignalReceiver chan int
	StatusCh       chan string
}

// goroutine status
type StatusType string

const (
	Alive     StatusType = "alive"
	Dead      StatusType = "dead"
	Unkonw    StatusType = "unknow"
	Disappear StatusType = "disappear"
)

type RuntimeAlertStatus struct {
	Status    StatusType
	timestamp time.Time
}

var CachedRuntimeAlert = make(map[string]*RuntimeAlert)

func Action(ctx context.Context, msg *pb.Message) error {

	signalx := msg.Signal

	switch signalx {
	case pb.Message_CREATE:
		// create alert by specifig alert config id within one goroutine
		fmt.Println("create alert")

	case pb.Message_STOP:
	case pb.Message_RELOAD:
	case pb.Message_OTHER:
	}
	return nil
}

func CreateRuntimeAlert(alertConfigID string) {
	// get alert config by id from DB
	alert, err := models.GetAlertBindingItem(alertConfigID)
	if err != nil {
		glog.Errorln(err.Error())
	}

	go func(alert *models.AlertBinding) {
		// get resource group, alert rule group, resource type
		alertRuleGroupID := alert.AlertRuleGroupID
		resourceGroupID := alert.ResourceGroupID
		receiverGroupID := alert.ReceiverGroupID

	}(alert)

}

func DeleteRuntimeAlert(alertConfigID string) error {
	// first step: need to delete items in related tables
	err := models.DeleteAlertBindingItem(alertConfigID)
	// if an error occured, delete runtime alert failed
	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	// second step: delete item in CachedRuntimeAlert map
	if alert, ok := CachedRuntimeAlert[alertConfigID]; ok {
		if alert != nil {
			alert.SignalSender <- pb.AlertConfig_Terminate
			for {
				sig := <-alert.SignalReceiver
				// TODO
				if sig == 0 {
					glog.Infof("terminate running alert goroutine successfully, alert_config_id is: %s", alertConfigID)
					delete(CachedRuntimeAlert, alertConfigID)
					return nil
				}

			}
		}
	}
	return nil
}

// does goroutine still alive?
func GetRuntimeAlertStatus(alertConfigID string) *RuntimeAlertStatus {
	if alert, ok := CachedRuntimeAlert[alertConfigID]; ok {
		if alert != nil {
			alert.StatusCh <- "ping"
			for {
				sig := <-alert.StatusCh
				if sig == "pong" {
					glog.Infof("alert goroutine is running, alert_config_id is: %s", alertConfigID)
					return nil
				}
			}
		}

		return &RuntimeAlertStatus{
			Status:    Alive,
			timestamp: time.Now(),
		}
	}

	return &RuntimeAlertStatus{
		Status:    Unkonw,
		timestamp: time.Now(),
	}
}
