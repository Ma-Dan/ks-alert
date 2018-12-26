package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateReceiverBindingGroupItem(t *testing.T) {
	Convey("test CreateReceiverBindingGroup Item", t, func() {

		productID := "product-4llxr47k7q82wz"

		severities := GetSeveritiesByProductID(productID)

		//		for _, severity := range *severities{
		//			severity.SeverityID
		//		}

		receiver01, err := CreateReceiver(&Receiver{
			ReceiverName: "aaaaaaaaaaa",
			Email:        "zlahu@foxmail.com",
			Phone:        "11111111111111",
			Wechat:       "zhangliaish",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		So(err, ShouldBeNil)

		receiver02, err := CreateReceiver(&Receiver{
			ReceiverName: "bbbbbbbbbbb",
			Email:        "zlahu@foxmail.com",
			Phone:        "222222222222222",
			Wechat:       "zhangliaish",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		So(err, ShouldBeNil)

		receiver03, err := CreateReceiver(&Receiver{
			ReceiverName: "cccccccccccc",
			Email:        "zlahu@foxmail.com",
			Phone:        "33333333333333",
			Wechat:       "zhangliaish",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		So(err, ShouldBeNil)

		receiverGroup, err := CreateReceiverGroup(&ReceiverGroup{
			ReceiverGroupName: "carman_receiver_group_01",
			Webhook:           "139.198.190.141:33333",
			WebhookEnable:     true,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		})

		receivers := []Receiver{*receiver01, *receiver02, *receiver03}
		fmt.Println(receiverGroup, receivers)
	})
}
