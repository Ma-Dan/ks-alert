package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
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

type ReceiverBindingGroup struct {
	ReceiverID      string    `gorm:"type:varchar(50);primary_key"`
	ReceiverGroupID string    `gorm:"type:varchar(50);primary_key" json:"-"`
	CreatedAt       time.Time `gorm:"not null;"`
	UpdatedAt       time.Time `gorm:"not null;"`
}

type ReceiverGroup struct {
	Action
	ReceiverGroupID   string      `gorm:"type:varchar(50);primary_key" json:"-"`
	ReceiverGroupName string      `gorm:"type:varchar(50);not null;" json:"receiver_group_name"`
	Webhook           string      `gorm:"type:varchar(50);" json:"webhook, omitempty"`
	WebhookEnable     bool        `gorm:"type:bool;" json:"webhook_enable, omitempty"`
	Receivers         *[]Receiver `gorm:"-" json:"receivers"`
	Description       string      `gorm:"type:text;" json:"desc"`
	CreatedAt         time.Time   `gorm:"not null;" json:"-"`
	UpdatedAt         time.Time   `gorm:"not null;" json:"-"`
}

func (r *ReceiverGroup) Create(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupName == "" {
		return nil, Error{Text: "the receiver group name must be specified", Code: InvalidParam}
	}

	if r.Receivers == nil || len(*r.Receivers) == 0 {
		return nil, Error{Text: "the receiver group must contain at least one receiver", Code: InvalidParam}
	}

	r.ReceiverGroupID = idutil.GetUuid36("receiver_group-")

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%v','%s','%v','%v')", r.ReceiverGroupID, r.ReceiverGroupName,
		r.Webhook, Bool2Int[r.WebhookEnable], r.Description, r.CreatedAt, r.UpdatedAt)

	sql := "INSERT INTO receiver_groups (receiver_group_id, receiver_group_name, webhook, " +
		"webhook_enable, description, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// create item
	sql = "INSERT INTO receivers (receiver_id, receiver_name, email, phone, wechat, created_at, updated_at) VALUES "

	receivers := *r.Receivers
	l := len(receivers)

	var recvIds []string
	var createdRrecvIds []string

	for i := 0; i < l; i++ {
		receiver := receivers[i]

		// TODO need to validate the receiver_id exist
		//if receiver.ReceiverID != "" {
		//	recvIds = append(recvIds, receiver.ReceiverID)
		//	continue
		//}

		recvId := idutil.GetUuid36("rule_id-")
		createdRrecvIds = append(createdRrecvIds, recvId)

		receiver.ReceiverID = recvId

		item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%v') ",
			receiver.ReceiverID, receiver.ReceiverName, receiver.Email, receiver.Phone, receiver.Wechat, receiver.CreatedAt, receiver.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if len(createdRrecvIds) > 0 {
		fmt.Println(sql)

		if err := tx.Exec(sql).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError}
		}
	}

	sql = "INSERT INTO receiver_binding_groups (receiver_id, receiver_group_id, created_at, updated_at) VALUES "

	recvIds = append(recvIds, createdRrecvIds...)

	for i := 0; i < len(recvIds); i++ {

		item := fmt.Sprintf("('%s','%s','%v','%v') ",
			recvIds[i], r.ReceiverGroupID, time.Now(), time.Now())

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return r, nil
}

// TODO this function is not feasible
func (r *ReceiverGroup) Update(tx *gorm.DB) (interface{}, error) {

	if r.ReceiverGroupID == "" {
		return nil, Error{Text: "the receiver group id must be specified", Code: InvalidParam}
	}

	// 1. get group first
	vget, err := r.Get(tx)

	if err != nil {
		return nil, err
	}

	rg := vget.(*ReceiverGroup)

	if rg == nil || rg.ReceiverGroupID == "" {
		return nil, Error{Text: fmt.Sprintf("resource group id: %s not exist", r.ReceiverGroupID), Code: InvalidParam}
	}

	// 2. delete group
	_, err = r.Delete(tx)

	if err != nil {
		return nil, err
	}

	// 3. create item
	createDate := rg.CreatedAt
	r.CreatedAt = createDate

	return r.Create(tx)
}

func (r *ReceiverGroup) Get(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupID == "" {
		return nil, Error{Text: "receiver group id must be specified", Code: InvalidParam}
	}

	var rg ReceiverGroup

	err := tx.Model(&ReceiverGroup{}).Where("receiver_group_id=?", r.ReceiverGroupID).First(&rg).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, Error{Text: "record not found", Code: DBError}
	}

	if rg.ReceiverGroupID != "" {
		var receivers []Receiver
		//
		sql := "SELECT r.* FROM receiver_binding_groups as rb LEFT JOIN receivers as r ON rb.receiver_id=r.receiver_id WHERE rb.receiver_group_id=?"

		if err := tx.Debug().Raw(sql, rg.ReceiverGroupID).Scan(&receivers).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError}
		}

		//err := tx.Find(&receivers, "receiver_group_id=?", rg.ReceiverGroupID).Error
		//
		//if err != nil {
		//	tx.Rollback()
		//	return nil, err
		//}
		rg.Receivers = &receivers

		return &rg, nil
	}

	return nil, nil
}

func (r *ReceiverGroup) Delete(tx *gorm.DB) (interface{}, error) {
	if r.ReceiverGroupID == "" {
		return nil, Error{Text: "receiver group id must be specified", Code: InvalidParam}
	}

	sql := "DELETE rg, rb FROM receiver_groups as rg LEFT JOIN receiver_binding_groups as rb ON rg.receiver_group_id=rb.receiver_group_id WHERE rg.receiver_group_id=?"

	if err := tx.Exec(sql, r.ReceiverGroupID).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// delete receiver which is not in any receiver group
	sql = "DELETE r FROM receivers as r LEFT JOIN receiver_binding_groups as rb ON r.receiver_id=rb.receiver_id WHERE rb.receiver_group_id IS NULL;"

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return nil, nil
}
