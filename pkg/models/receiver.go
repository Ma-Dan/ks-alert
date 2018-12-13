package models

import (
	"k8s.io/klog/glog"
	"kubesphere.io/ks-alert/pkg/client"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"time"
)

type Receiver struct {
	ReceiverID   string    `gorm:"type:varchar(50);primary_key"`
	ReceiverName string    `gorm:"type:varchar(50);not null;"`
	Email        string    `gorm:"type:varchar(50);not null;"`
	Phone        string    `gorm:"type:varchar(50);"`
	Wechat       string    `gorm:"type:varchar(50);"`
	CreatedAt    time.Time `gorm:"not null;"`
	UpdatedAt    time.Time `gorm:"not null;"`
	// ignore this field because it will be appeared in table `ReceiverBindingGroup`
	SeverityID string `gorm:"-"`
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
	ReceiverGroupID   string    `gorm:"type:varchar(50);primary_key"`
	ReceiverGroupName string    `gorm:"type:varchar(50);not null;"`
	Webhook           string    `gorm:"type:varchar(50);"`
	WebhookEnable     bool      `gorm:"type:bool;"`
	CreatedAt         time.Time `gorm:"not null;"`
	UpdatedAt         time.Time `gorm:"not null;"`
}

func CreateReceiverGroup(receiverGroup *ReceiverGroup) (*ReceiverGroup, error) {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	receiverGroup.ReceiverGroupID = idutil.GetUuid36("receiver_group-")

	err = db.Model(&ReceiverGroup{}).Create(receiverGroup).Error

	return receiverGroup, err
}

func CreateReceiverBindingGroupItem(receivers *[]Receiver, receiverGroup *ReceiverGroup) error {

	receiverGroup, err := CreateReceiverGroup(receiverGroup)
	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	for _, receiver := range *receivers {
		var err error
		var createdReceiver *Receiver
		if receiver.ReceiverID == "" {
			// need to create this user, otherwise this user is exists
			createdReceiver, err = CreateReceiver(&receiver)
		}

		if err != nil {
			glog.Errorln(err.Error())
			return err
		}

		// Create item in table `ReceiverBindingGroup`

		err = db.Model(&ReceiverBindingGroup{}).Create(&ReceiverBindingGroup{
			ReceiverGroupID: receiverGroup.ReceiverGroupID,
			ReceiverID:      createdReceiver.ReceiverID,
			SeverityID:      createdReceiver.SeverityID,
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

func CreateReceiver(receiver *Receiver) (*Receiver, error) {

	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	receiver.ReceiverID = idutil.GetUuid36("receiver-")

	err = db.Model(&Receiver{}).Create(receiver).Error

	return receiver, err
}
