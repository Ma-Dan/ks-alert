package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/pkg/errors"
	"k8s.io/klog/glog"
	"net/http"
	"reflect"
)

type TP string

const (
	AlertConfig   TP = "AlertConfig"
	RuleGroup     TP = "RuleGroup"
	ReceiverGroup TP = "ReceiverGroup"
	ResourceGroup TP = "ResourceGroup"
)

const (
	MethodCreate = "Create"
	MethodUpdate = "Update"
	MethodGet    = "Get"
	MethodDelete = "Delete"
)

func DoTransactionAction(v interface{}, tp TP, method string, bol ...bool) (interface{}, error) {

	// transaction begin
	db, e := dbutil.DBClient()
	if e != nil {
		return nil, e
	}

	tx := db.Begin()

	var res interface{}
	var err error

	if tp == AlertConfig {
		var ruleGroupResponse, recvGroupResponse, resGroupResponse interface{}

		alertConfig := v.(*models.AlertConfig)
		if len(bol) == 0 {

			var ruleGroup interface{}
			if method == http.MethodPut || method == http.MethodPost {
				ruleGroup = alertConfig.AlertRuleGroup
			} else {
				ruleGroup = &pb.AlertRuleGroupSpec{AlertRuleGroupId: alertConfig.AlertRuleGroup.AlertRuleGroupID}
			}
			ruleGroupResponse, err = CallReflect(models.AlertRuleGroup{}, method, tx, ruleGroup)

			if err != nil {
				tx.Rollback()
				glog.Errorln(err.Error())
				return nil, err
			}

			var receiverGroup interface{}
			receiverGroup = alertConfig.ReceiverGroup
			//if method == http.MethodPut || method == http.MethodPost {
			//} else {
			//	receiverGroup = &pb.ReceiverGroupSpec{ReceiverGroupId: alertConfig.ReceiverGroup.ReceiverGroupID}
			//}
			recvGroupResponse, err = CallReflect(models.ReceiverGroup{}, method, tx, receiverGroup)

			if err != nil {
				tx.Rollback()
				glog.Errorln(err.Error())
				return nil, err
			}

			var resourceGroup interface{}
			resourceGroup = alertConfig.ResourceGroup
			//if method == http.MethodPut || method == http.MethodPost {
			//} else {
			//	resourceGroup = &pb.ResourceGroupSpec{ResourceGroupId: alertConfig.ResourceGroup.ResourceGroupID}
			//}
			resGroupResponse, err = CallReflect(models.ResourceGroup{}, method, tx, resourceGroup)

			if err != nil {
				tx.Rollback()
				glog.Errorln(err.Error())
				return nil, err
			}
		}

		if method == http.MethodPut || method == http.MethodPost {
			alertConfig.AlertRuleGroupID = ruleGroupResponse.(*models.AlertRuleGroup).AlertRuleGroupID
			alertConfig.ResourceGroupID = resGroupResponse.(*models.ResourceGroup).ResourceGroupID
			alertConfig.ReceiverGroupID = recvGroupResponse.(*models.ReceiverGroup).ReceiverGroupID
		}

		alertConfigResponse, err := CallReflect(models.AlertConfig{}, method, tx, alertConfig)

		if err != nil {
			tx.Rollback()
			glog.Errorln(err.Error())
			return nil, err
		}

		res = []interface{}{alertConfigResponse, ruleGroupResponse, recvGroupResponse, resGroupResponse}

	}

	if tp == RuleGroup {
		// create rule group
		res, err = CallReflect(models.AlertRuleGroup{}, method, tx, v)

		if err != nil {
			tx.Rollback()
			glog.Errorln(err.Error())
		}
	}

	if tp == ReceiverGroup {
		// receiver group
		res, err = CallReflect(models.ReceiverGroup{}, method, tx, v)

		if err != nil {
			tx.Rollback()
			glog.Errorln(err.Error())
		}
	}

	if tp == ResourceGroup {
		// create resource group
		res, err = CallReflect(models.ResourceGroup{}, method, tx, v)

		if err != nil {
			tx.Rollback()
			glog.Errorln(err.Error())
		}
	}

	// TODO need to validate closing db connection
	// transaction end
	tx.Commit()

	return res, err
}

func CallReflect(any interface{}, method string, args ...interface{}) (interface{}, error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	v := reflect.ValueOf(any).MethodByName(method)

	if v.String() == "<invalid Value>" {
		return nil, errors.New("invalid Value")
	}

	values := v.Call(inputs)

	if len(values) == 1 {
		return values[0].Interface(), nil

	} else if len(values) == 2 {

		if values[1].Interface() == nil {
			return values[0].Interface(), nil
		}

		e := values[1].Interface().(error)
		return values[0].Interface(), e
	}

	return nil, nil
}
