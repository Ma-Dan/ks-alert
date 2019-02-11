package notification

import (
	"context"
	"fmt"
	"kubesphere.io/ks-alert/pkg/client"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/option"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/wrappers"
	"openpitrix.io/notification/pkg/pb"
	"strings"
)

type Sender struct{}

func (s Sender) Send(receivers *[]models.Receiver, noticeStr string) string {
	emails := extractEmails(receivers)

	// TODO make it more elegant
	noticeWithReceivers := &pb.CreateNotificationRequest{
		Content:     &wrappers.StringValue{Value: noticeStr},
		AddressInfo: &wrappers.StringValue{Value: jsonutil.Marshal(map[string][]string{"email": emails})},
		ExpiredDays: &wrappers.UInt32Value{Value: 0},
		Title:       &wrappers.StringValue{Value: "here comes an alert"},
		ContentType: &wrappers.StringValue{Value: "alert"},
	}

	fmt.Println(jsonutil.Marshal(noticeWithReceivers))
	svcAddress := fmt.Sprintf("%s:%s", *option.NotificationHost, *option.NotificationPort)
	conn, err := client.GetNotificationConn(svcAddress)
	if err != nil {
		glog.Errorln(err.Error())
		return ""
	}

	nfClient := pb.NewNotificationClient(conn)
	response, err := nfClient.CreateNotification(context.Background(), noticeWithReceivers)
	if err != nil {
		glog.Errorln(err.Error())
	}
	return jsonutil.Marshal(response)
}

func extractEmails(receiver *[]models.Receiver) []string {
	var emails []string
	l := len(*receiver)

	for i := 0; i < l; i++ {
		e := strings.Trim((*receiver)[i].Email, " ")
		if e != "" {
			emails = append(emails, e)
		}
	}
	return emails
}
