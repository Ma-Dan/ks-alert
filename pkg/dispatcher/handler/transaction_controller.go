package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/pkg/errors"
	"k8s.io/klog/glog"
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

func DoTransactionAction(v interface{}, tp TP, method string) (interface{}, error) {

	// transaction begin
	db, e := dbutil.DBClient()
	if e != nil {
		return nil, e
	}

	tx := db.Begin()

	var res interface{}
	var err error
	res, err = CallReflect(v, method, tx)

	//switch tp {
	//case AlertConfig:
	//	res, err = CallReflect(v, method, tx)
	//case RuleGroup:
	//	res, err = CallReflect(v, method, tx)
	//case ReceiverGroup:
	//	res, err = CallReflect(v, method, tx)
	//case ResourceGroup:
	//	res, err = CallReflect(v, method, tx)
	//}

	if err != nil {
		tx.Rollback()
		glog.Errorln(err.Error())
		return nil, err
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
