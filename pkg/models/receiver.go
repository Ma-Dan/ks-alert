package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

func (r ReceiverGroup) Create(tx *gorm.DB, v interface{}) (interface{}, error) {
	recvGroup, ok := v.(*ReceiverGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", recvGroup)
	}

	recvGroup.ReceiverGroupID = idutil.GetUuid36("receiver_group-")

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%v','%s','%v','%v')", recvGroup.ReceiverGroupID, recvGroup.ReceiverGroupName,
		recvGroup.Webhook, Bool2Int[recvGroup.WebhookEnable], recvGroup.Description, recvGroup.CreatedAt, recvGroup.UpdatedAt)

	sql := "INSERT INTO receiver_groups (receiver_group_id, receiver_group_name, webhook, " +
		"webhook_enable, description, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return nil, err
	}

	// create item
	sql = "INSERT INTO receivers (receiver_id, receiver_name, email, phone, wechat, created_at, updated_at) VALUES "

	receivers := *recvGroup.Receivers
	l := len(receivers)

	var recvIds []string
	var createdRrecvIds []string

	for i := 0; i < l; i++ {
		r := receivers[i]

		// TODO need to validate the receiver_id exist
		//if r.ReceiverID != "" {
		//	recvIds = append(recvIds, r.ReceiverID)
		//	continue
		//}

		recvId := idutil.GetUuid36("rule_id-")
		createdRrecvIds = append(createdRrecvIds, recvId)

		r.ReceiverID = recvId

		item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%v') ",
			r.ReceiverID, r.ReceiverName, r.Email, r.Phone, r.Wechat, r.CreatedAt, r.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if len(createdRrecvIds) > 0 {
		fmt.Println(sql)

		if err := tx.Exec(sql).Error; err != nil {
			return nil, err
		}
	}

	sql = "INSERT INTO receiver_binding_groups (receiver_id, receiver_group_id, created_at, updated_at) VALUES "

	recvIds = append(recvIds, createdRrecvIds...)

	for i := 0; i < len(recvIds); i++ {

		item := fmt.Sprintf("('%s','%s','%v','%v') ",
			recvIds[i], recvGroup.ReceiverGroupID, time.Now(), time.Now())

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if err := tx.Exec(sql).Error; err != nil {
		return nil, err
	}

	return recvGroup, nil
}

// TODO this function is not feasible
func (r ReceiverGroup) Update(tx *gorm.DB, v interface{}) (interface{}, error) {

	recvGroup, ok := v.(*ReceiverGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", recvGroup)
	}

	// 1. get group first
	vget, err := r.Get(tx, v)

	if err != nil {
		return nil, err
	}

	rg := vget.(*ReceiverGroup)

	if rg == nil || rg.ReceiverGroupID == "" {
		return nil, errors.Errorf("resource group id: %s not exist", recvGroup.ReceiverGroupID)
	}

	// 2. delete group
	_, err = r.Delete(tx, v)

	if err != nil {
		return nil, err
	}

	// 3. create item
	createDate := rg.CreatedAt
	recvGroup.CreatedAt = createDate

	return r.Create(tx, recvGroup)
}

func (r ReceiverGroup) Get(tx *gorm.DB, v interface{}) (interface{}, error) {

	recvGroup, ok := v.(*ReceiverGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", recvGroup)
	}

	var rg ReceiverGroup

	err := tx.Model(&ReceiverGroup{}).Where("receiver_group_id=?", recvGroup.ReceiverGroupID).First(&rg).Error

	if err != nil {
		return nil, err
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, errors.New("record not found")
	}

	if rg.ReceiverGroupID != "" {
		var receivers []Receiver
		//
		sql := "SELECT r.* FROM receiver_binding_groups as rb LEFT JOIN receivers as r ON rb.receiver_id=r.receiver_id WHERE rb.receiver_group_id=?"

		if err := tx.Exec(sql, rg.ReceiverGroupID).Find(&receivers).Error; err != nil {
			return nil, err
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

func (r ReceiverGroup) Delete(tx *gorm.DB, v interface{}) (interface{}, error) {
	recvGroup, ok := v.(*ReceiverGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", recvGroup)
	}

	sql := "DELETE rg, rb FROM receiver_groups as rg LEFT JOIN receiver_binding_groups as rb ON rg.receiver_group_id=rb.receiver_group_id WHERE rg.receiver_group_id=?"

	if err := tx.Exec(sql, recvGroup.ReceiverGroupID).Error; err != nil {
		return nil, err
	}

	// delete receiver which is not in any receiver group
	sql = "DELETE r FROM receivers as r LEFT JOIN receiver_binding_groups as rb ON r.receiver_id=rb.receiver_id WHERE rb.receiver_group_id IS NULL;"

	if err := tx.Exec(sql).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
