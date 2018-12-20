package models

import (
	"k8s.io/klog/glog"
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
	// ignore this field because it will be appeared in table `ReceiverBindingGroup`
	SeverityID string `gorm:"-"  json:"severity_id"`
}

// Association table
type ReceiverBindingGroup struct {
	ReceiverID      string `gorm:"type:varchar(50);primary_key"`
	ReceiverGroupID string `gorm:"type:varchar(50);primary_key"`
	SeverityID      string `gorm:"type:varchar(5);not null;"`
	//Webhook      string    `gorm:"type:varchar(50);"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
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

func CreateReceiverBindingGroupItem(receivers *[]Receiver, receiverGroup *ReceiverGroup) error {

	var err error

	if receiverGroup.ReceiverGroupID == "" {
		receiverGroup, err = CreateReceiverGroup(receiverGroup)
	}

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	for _, receiver := range *receivers {
		var err error
		var createdReceiver *Receiver
		var receiverID = receiver.ReceiverID

		if receiverID == "" {
			// need to create this user, otherwise this user is exists
			createdReceiver, err = CreateReceiver(&receiver)
			receiverID = createdReceiver.ReceiverID
		}

		if err != nil {
			glog.Errorln(err.Error())
			return err
		}

		// Create item in table `ReceiverBindingGroup`

		err = db.Model(&ReceiverBindingGroup{}).Create(&ReceiverBindingGroup{
			ReceiverGroupID: receiverGroup.ReceiverGroupID,
			ReceiverID:      receiverID,
			SeverityID:      receiver.SeverityID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}).Error

		if err != nil {
			glog.Errorln(err.Error())
			return err
		}
	}

	return nil
}

func CreateReceivers(receivers *[]Receiver) (*[]Receiver, error) {
	var createdReceiver []Receiver

	for i:=0; i< len(*receivers); i++ {
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
