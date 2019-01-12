package notification

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	Success = "Success"
	Failure = "Failure"
	Unknow  = "Unknow"
)

type Sender struct{}

type SendStatus struct {
	ReceiverID string
	Status     string
	Message    string
	Timestamp  int64
}

type SendStatusMap map[string]SendStatus

var senderClient = &http.Client{}
var notificationAddress = "http://139.198.190.141:8082/notice"

func sendRequest(epurl string, email string, noticeStr string) (string, error) {
	response, err := senderClient.PostForm(epurl, url.Values{"email": {email}, "notice": {noticeStr}})
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return string(contents), err
		}

		return string(contents), nil
	}
}

func (s Sender) Send(receivers *[]models.Receiver, noticeStr string) *SendStatusMap {
	l := len(*receivers)

	var ch = make(chan *SendStatus, l)
	wg := sync.WaitGroup{}
	for i, _ := range *receivers {
		wg.Add(1)
		go func(ch chan *SendStatus, i int) {
			r := (*receivers)[i]
			recvID := r.ReceiverID
			email := r.Email
			fmt.Println(email)
			_, err := sendRequest(notificationAddress, email, noticeStr)
			var status *SendStatus
			if err == nil {
				status = &SendStatus{
					Timestamp:  time.Now().Unix(),
					Status:     Success,
					ReceiverID: recvID,
				}
			} else {
				status = &SendStatus{
					Timestamp:  time.Now().Unix(),
					Status:     Failure,
					Message:    err.Error(),
					ReceiverID: recvID,
				}
			}

			ch <- status
			wg.Done()
		}(ch, i)

	}

	wg.Wait()
	close(ch)

	statusMap := make(SendStatusMap, l)

	for c := range ch {
		statusMap[c.ReceiverID] = *c
	}

	return &statusMap
}
