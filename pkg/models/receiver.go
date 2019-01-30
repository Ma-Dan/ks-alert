package models

import (
	"fmt"
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"time"
)

type Receiver struct {
	ReceiverID      string    `gorm:"type:varchar(50);primary_key"`
	ReceiverName    string    `gorm:"type:varchar(50);not null;"`
	ReceiverGroupID string    `gorm:"type:varchar(50);not null"`
	Email           string    `gorm:"type:varchar(50);not null;"`
	Phone           string    `gorm:"type:varchar(50);"`
	Wechat          string    `gorm:"type:varchar(50);"`
	CreatedAt       time.Time `gorm:"not null;"`
	UpdatedAt       time.Time `gorm:"not null;"`
}

type ReceiverGroup struct {
	Action
	ReceiverGroupID   string      `gorm:"type:varchar(50);primary_key"`
	ReceiverGroupName string      `gorm:"type:varchar(50);not null;"`
	Webhook           string      `gorm:"type:varchar(50);"`
	WebhookEnable     bool        `gorm:"type:bool;"`
	Receivers         *[]Receiver `gorm:"-"`
	Description       string      `gorm:"type:text;"`
	CreatedAt         time.Time   `gorm:"not null;"`
	UpdatedAt         time.Time   `gorm:"not null;"`
}

func (r *ReceiverGroup) Create(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupName == "" {
		return nil, Error{Text: "the receiver group name must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	if r.Receivers == nil || len(*r.Receivers) == 0 {
		return nil, Error{Text: "the receiver group must contain at least one receiver", Code: InvalidParam, Where: Caller(1, true)}
	}

	if r.ReceiverGroupID == "" {
		r.ReceiverGroupID = idutil.GetUuid36("")
	}

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%v','%s','%v','%v')", r.ReceiverGroupID, r.ReceiverGroupName,
		r.Webhook, Bool2Int[r.WebhookEnable], r.Description, r.CreatedAt, r.UpdatedAt)

	sql := "INSERT INTO receiver_groups (receiver_group_id, receiver_group_name, webhook, " +
		"webhook_enable, description, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	// create item
	sql = "INSERT INTO receivers (receiver_id, receiver_group_id, receiver_name, email, phone, wechat, created_at, updated_at) VALUES "

	receivers := *r.Receivers
	l := len(receivers)

	for i := 0; i < l; i++ {
		receiver := receivers[i]
		receiver.ReceiverID = idutil.GetUuid36("")

		item := fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%v','%v') ",
			receiver.ReceiverID, r.ReceiverGroupID, receiver.ReceiverName, receiver.Email, receiver.Phone, receiver.Wechat, receiver.CreatedAt, receiver.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return r, nil
}

func (r *ReceiverGroup) Update(tx *gorm.DB) (interface{}, error) {
	vget, err := r.Get(tx)
	if err != nil {
		return nil, err
	}

	rg := vget.(*ReceiverGroup)
	if rg.ReceiverGroupID == "" {
		return nil, Error{Text: "the receiver group id does not valid", Code: InvalidParam, Where: Caller(1, true)}
	}

	// 2. delete group
	_, err = r.Delete(tx)

	if err != nil {
		return nil, err
	}

	r.ReceiverGroupID = rg.ReceiverGroupID
	// 3. create item
	return r.Create(tx)
}

func (r *ReceiverGroup) Get(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupID == "" {
		return nil, Error{Text: "receiver group id must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	var rg ReceiverGroup

	err := tx.Model(&ReceiverGroup{}).Where("receiver_group_id=?", r.ReceiverGroupID).First(&rg).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, Error{Text: "record not found", Code: DBError, Where: Caller(1, true)}
	}

	if rg.ReceiverGroupID != "" {
		var receivers []Receiver
		//
		sql := "SELECT r.* FROM receivers as r WHERE r.receiver_group_id=?"

		if err := tx.Raw(sql, rg.ReceiverGroupID).Scan(&receivers).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
		}

		rg.Receivers = &receivers
		return &rg, nil
	}

	return nil, nil
}

func (r *ReceiverGroup) Delete(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupID == "" {
		return nil, Error{Text: "receiver group id must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	sql := "DELETE rg, r FROM receiver_groups as rg LEFT JOIN receivers as r ON rg.receiver_group_id=r.receiver_group_id WHERE rg.receiver_group_id=?"

	if err := tx.Exec(sql, r.ReceiverGroupID).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}
	return nil, nil
}
