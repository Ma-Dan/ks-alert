package models

import (
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"time"
)

type Receiver struct {
	ReceiverID   string    `gorm:"type:varchar(50);primary_key" json:"receiver_id, omitempty"`
	ReceiverName string    `gorm:"type:varchar(50);not null;" json:"receiver_name, omitempty"`
	Email        string    `gorm:"type:varchar(50);not null;" json:"email, omitempty"`
	Phone        string    `gorm:"type:varchar(50);" json:"phone, omitempty"`
	Wechat       string    `gorm:"type:varchar(50);" json:"wechat, omitempty"`
	CreatedAt    time.Time `gorm:"not null;" json:"-"`
	UpdatedAt    time.Time `gorm:"not null;" json:"-"`
}

type ReceiverGroup struct {
	ReceiverGroupID   string     `gorm:"type:varchar(50);primary_key" json:"-"`
	ReceiverGroupName string     `gorm:"type:varchar(50);not null;" json:"receiver_group_name"`
	Webhook           string     `gorm:"type:varchar(50);" json:"webhook, omitempty"`
	WebhookEnable     bool       `gorm:"type:bool;" json:"webhook_enable, omitempty"`
	Receivers         []Receiver `gorm:"-" json:"receivers"`
	Description       string     `gorm:"type:text;" json:"desc"`
	CreatedAt         time.Time  `gorm:"not null;" json:"-"`
	UpdatedAt         time.Time  `gorm:"not null;" json:"-"`
}

func CreateReceiverGroup(receiverGroup *ReceiverGroup) (*ReceiverGroup, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	receiverGroup.ReceiverGroupID = idutil.GetUuid36("receiver_group-")

	err = db.Model(&ReceiverGroup{}).Create(receiverGroup).Error

	return receiverGroup, err
}

func CreateReceivers(receivers *[]Receiver) (*[]Receiver, error) {
	var createdReceiver []Receiver

	for i := 0; i < len(*receivers); i++ {
		receiver, err := CreateReceiver(&(*receivers)[i])
		if err != nil {
			return &createdReceiver, err
		}
		createdReceiver = append(createdReceiver, *receiver)
	}

	return &createdReceiver, nil
}

func CreateReceiver(receiver *Receiver) (*Receiver, error) {

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	if receiver.ReceiverID == "" {
		receiver.ReceiverID = idutil.GetUuid36("receiver-")
		err = db.Model(&Receiver{}).Create(receiver).Error
		return receiver, err
	}

	return receiver, nil
}
