package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type Resource struct {
	Action
	ResourceID      string    `gorm:"primary_key" json:"-"`
	ResourceName    string    `gorm:"type:varchar(50);" json:"resource_name"`
	ResourceGroupID string    `gorm:"not null;" json:"-"`
	CreatedAt       time.Time `gorm:"not null;" json:"-"`
	UpdatedAt       time.Time `gorm:"not null;" json:"-"`
}

type ResourceGroup struct {
	Action
	ResourceGroupID   string      `gorm:"primary_key" json:"-"`
	ResourceGroupName string      `gorm:"type:varchar(50);not null;" json:"resource_group_name"`
	ResourceTypeID    string      `gorm:"type:varchar(50);" json:"resource_type_id"`
	Resources         []*Resource `gorm:"-" json:"resources"`
	URIParams         string      `gorm:"type:text;not null;" json:"-"`
	Description       string      `gorm:"type:text;" json:"desc"`
	CreatedAt         time.Time   `gorm:"not null;" json:"-"`
	UpdatedAt         time.Time   `gorm:"not null;" json:"-"`
}

type ResourceType struct {
	ResourceTypeID   string `gorm:"primary_key;not null;"`
	ProductID        string `gorm:"not null;"`
	ResourceTypeName string `gorm:"type:varchar(50);not null;"`
	Description      string `gorm:"type:text;"`
	Enable           bool   `gorm:"type:boolean;not null;"`

	// MonitorCenterHost and MonitorCenterPort will override the corresponding filed in struct `product`
	MonitorCenterHost string `gorm:"type:varchar(128);"`
	MonitorCenterPort int32  `gorm:"type:int;"` //default:-1;
	ResourceURITmpls  string `gorm:"type:text;not null;"`

	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}

func GetResourceType(resourceType *ResourceType) (*ResourceType, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var tp ResourceType
	db = db.Model(&ResourceType{}).Where(resourceType).First(&tp)

	if db.RecordNotFound() {
		return nil, nil
	}

	if db.Error != nil {
		return nil, err
	}

	return &tp, nil
}

func CreateResourceType(resourceType *ResourceType) (*ResourceType, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	resourceType.ResourceTypeID = idutil.GetUuid36("resource_type-")

	err = db.Model(&ResourceType{}).Create(resourceType).Error
	return resourceType, err

}

func UpdateResourceType(resourceType *ResourceType) error {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	if resourceType.ResourceTypeID != "" {
		err = db.Model(resourceType).Where("resource_type_id = ?", resourceType.ResourceTypeID).Update(resourceType).Error
	} else if resourceType.ProductID != "" && resourceType.ResourceTypeName != "" {
		err = db.Model(resourceType).Where("product_id = ? and resource_type_name = ? ", resourceType.ProductID, resourceType.ResourceTypeName).Update(resourceType).Error
	}

	return err
}

func DeleteResourceType(resourceType *ResourceType) error {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	err = db.Delete(resourceType).Error
	return err
}

func (r ResourceGroup) Create(tx *gorm.DB, v interface{}) (interface{}, error) {
	resGroup, ok := v.(*ResourceGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", resGroup)
	}

	var resourceWithName []*Resource
	resources := resGroup.Resources
	for i := 0; i < len(resources); i++ {
		r := resources[i]
		if r.ResourceName != "" {
			resourceWithName = append(resourceWithName, r)
		}
	}

	if len(resourceWithName) == 0 {
		return nil, errors.New("at least one resource name must be specified")
	}

	resGroup.ResourceGroupID = idutil.GetUuid36("resource_group-")

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%v')", resGroup.ResourceGroupID, resGroup.ResourceGroupName,
		resGroup.ResourceTypeID, resGroup.Description, resGroup.URIParams, resGroup.CreatedAt, resGroup.UpdatedAt)

	sql := "INSERT INTO resource_groups (resource_group_id, resource_group_name, resource_type_id, " +
		"description, uri_params, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// create item
	sql = "INSERT INTO resources (resource_id, resource_name, resource_group_id, created_at, updated_at) VALUES "

	l := len(resourceWithName)
	for i := 0; i < l; i++ {
		r := resourceWithName[i]

		resId := idutil.GetUuid36("resource-")
		r.ResourceID = resId
		r.ResourceGroupID = resGroup.ResourceGroupID

		item := fmt.Sprintf("('%s','%s','%s','%v','%v') ",
			r.ResourceID, r.ResourceName, resGroup.ResourceGroupID, r.CreatedAt, r.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if err := tx.Debug().Exec(sql).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}

func (r ResourceGroup) Update(tx *gorm.DB, v interface{}) (interface{}, error) {

	resGroup, ok := v.(*ResourceGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", resGroup)
	}

	// 1. get resource group first
	vget, err := r.Get(tx, v)

	if err != nil {
		return nil, err
	}

	rg := vget.(*ResourceGroup)

	if rg == nil || rg.ResourceGroupID == "" {
		return nil, errors.Errorf("resource group id: %s not exist", resGroup.ResourceGroupID)
	}

	// 2. update resource group
	// update alert rule group
	sql := fmt.Sprintf("UPDATE alert_rule_groups SET alert_rule_group_name='%s',description='%s',"+
		"system_rule='%v', updated_at='%v' WHERE alert_rule_group_id='%s'",
		ruleGroup.AlertRuleGroupName, ruleGroup.Description,
		Bool2Int[ruleGroup.SystemRule],
		ruleGroup.UpdatedAt, ruleGroup.AlertRuleGroupID)

	fmt.Println(sql)

	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. delete resource group
	_, err = r.Delete(tx, v)

	if err != nil {
		return nil, err
	}

	// 3. create resource group
	createDate := rg.CreatedAt
	resGroup.CreatedAt = createDate

	return r.Create(tx, resGroup)
}

func (r ResourceGroup) Get(tx *gorm.DB, v interface{}) (interface{}, error) {

	rgSpec, ok := v.(*ResourceGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", rgSpec)
	}

	var rg ResourceGroup

	err := tx.Model(&ResourceGroup{}).Where("resource_group_id=?", rgSpec.ResourceGroupID).First(&rg).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, errors.New("record not found")
	}

	if rg.ResourceGroupID != "" {
		var resources []Resource
		sql := "SELECT r.* FROM resources as r WHERE r.resource_group_id=?"

		if err := tx.Debug().Raw(sql, rg.ResourceGroupID).Scan(&resources).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		l := len(resources)
		var ptrs = make([]*Resource, l)

		for i := 0; i < l; i++ {
			ptrs[i] = &resources[i]
		}

		rg.Resources = ptrs
		return &rg, nil
	}

	return nil, nil
}

func (r ResourceGroup) Delete(tx *gorm.DB, v interface{}) (interface{}, error) {
	rg, ok := v.(*ResourceGroup)

	if !ok {
		return nil, errors.Errorf("type %v assert error", rg)
	}

	sql := "DELETE rg, r FROM resource_groups as rg LEFT JOIN resources as r ON rg.resource_group_id=r.resource_group_id WHERE rg.resource_group_id=?"

	if err := tx.Debug().Exec(sql, rg.ResourceGroupID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}

// batch resources crud
func CreateResources(tx *gorm.DB, resources *[]*Resource) error {
	// create item
	sql := "INSERT INTO resources (resource_id, resource_name, resource_group_id, created_at, updated_at) VALUES "

	l := len(*resources)
	for i := 0; i < l; i++ {
		r := (*resources)[i]

		resId := idutil.GetUuid36("resource-")
		r.ResourceID = resId

		item := fmt.Sprintf("('%s','%s','%s','%v','%v') ",
			r.ResourceID, r.ResourceName, r.ResourceGroupID, r.CreatedAt, r.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	if err := tx.Debug().Exec(sql).Error; err != nil {
		return err
	}

	return nil
}

func UpdateResources(tx *gorm.DB, v interface{}) (interface{}, error) {
}

func GetResources(tx *gorm.DB, v interface{}) (interface{}, error) {
}

func DeleteResources(tx *gorm.DB, v interface{}) (interface{}, error) {
}
