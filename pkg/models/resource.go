package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"strings"
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

type ResourceUriTmpls struct {
	ResourceUriTmpl []*ResourceUriTmpl `json:"resource_uri_tmpl,omitempty"`
}

type ResourceUriTmpl struct {
	UriTmpl      string            `json:"uri_tmpl,omitempty"`
	ResourceName []string          `json:"resource_name,omitempty"`
	PathParams   map[string]string `json:"path_params,omitempty"`
	QueryParams  string            `json:"query_params,omitempty"`
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
		return nil, Error{Text: fmt.Sprintf("type %v assert error", resGroup), Code: AssertError}
	}

	if resGroup.ResourceGroupName == "" || resGroup.ResourceTypeID == "" {
		return nil, Error{Text: "resource group name and resource type id must be specified", Code: InvalidParam}
	}

	if resGroup.Resources == nil || len(resGroup.Resources) == 0 {
		return nil, Error{Text: "resources must be specified", Code: InvalidParam}
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
		return nil, Error{Text: "at least one resource name must be specified", Code: InvalidParam}
	}

	resGroup.ResourceGroupID = idutil.GetUuid36("resource_group-")

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%v')", resGroup.ResourceGroupID, resGroup.ResourceGroupName,
		resGroup.ResourceTypeID, resGroup.Description, resGroup.URIParams, resGroup.CreatedAt, resGroup.UpdatedAt)

	sql := "INSERT INTO resource_groups (resource_group_id, resource_group_name, resource_type_id, " +
		"description, uri_params, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// create item
	l := len(resourceWithName)
	for i := 0; i < l; i++ {
		r := resourceWithName[i]
		r.ResourceID = idutil.GetUuid36("resource-")
		r.ResourceGroupID = resGroup.ResourceGroupID
	}

	if err := CreateOrUpdateResources(tx, resourceWithName); err != nil {
		return nil, err
	}

	return resGroup, nil
}

func (r ResourceGroup) Update(tx *gorm.DB, v interface{}) (interface{}, error) {

	resGroup, ok := v.(*ResourceGroup)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", resGroup), Code: AssertError}
	}

	if resGroup.ResourceGroupID == "" || resGroup.ResourceGroupName == "" {
		return nil, Error{Text: "resource group id or name must be specified", Code: InvalidParam}
	}

	// 1. get resource group first
	vget, err := r.Get(tx, v)

	if err != nil {
		return nil, err
	}

	rg := vget.(*ResourceGroup)

	if rg == nil || rg.ResourceGroupID == "" {
		return nil, Error{Text: fmt.Sprintf("resource group id: %s not exist", resGroup.ResourceGroupID), Code: InvalidParam}
	}

	// 2. update resource group
	sql := fmt.Sprintf("UPDATE resource_groups SET resource_group_name='%s',description='%s',"+
		"uri_params='%v', updated_at='%v' WHERE resource_group_id='%s'",
		resGroup.ResourceGroupName, resGroup.Description,
		resGroup.URIParams, time.Now(), resGroup.ResourceGroupID)

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// 3. delete resources or update resource
	var needDeleted, needUpdated = CompareResources(resGroup.Resources, rg.Resources, resGroup.ResourceGroupID)

	err = DeleteResources(tx, resGroup.ResourceGroupID, needDeleted)

	if err != nil {
		return nil, err
	}

	err = CreateOrUpdateResources(tx, needUpdated)

	if err != nil {
		return nil, err
	}

	return resGroup, nil
}

func (r ResourceGroup) Get(tx *gorm.DB, v interface{}) (interface{}, error) {

	rgSpec, ok := v.(*ResourceGroup)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", rgSpec), Code: AssertError}
	}

	rgID := rgSpec.ResourceGroupID

	if rgID == "" {
		return nil, Error{Text: "resource group id must be specified", Code: InvalidParam}
	}

	var rg ResourceGroup

	err := tx.Model(&ResourceGroup{}).Where("resource_group_id=?", rgSpec.ResourceGroupID).First(&rg).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, Error{Text: "record not found", Code: DBError}
	}

	if rg.ResourceGroupID != "" {
		resources, err := GetResources(tx, rg.ResourceGroupID)

		if err != nil {
			return &rg, err
		}

		rg.Resources = resources
	}

	return &rg, nil
}

func (r ResourceGroup) Delete(tx *gorm.DB, v interface{}) (interface{}, error) {
	rg, ok := v.(*ResourceGroup)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", rg), Code: AssertError}
	}

	rgID := rg.ResourceGroupID

	if rgID == "" {
		return nil, Error{Text: "resource group id must be specified", Code: InvalidParam}
	}

	sql := "DELETE rg, r FROM resource_groups as rg LEFT JOIN resources as r ON rg.resource_group_id=r.resource_group_id WHERE rg.resource_group_id=?"

	if err := tx.Debug().Exec(sql, rg.ResourceGroupID).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return nil, nil
}

// resources crud
func CreateOrUpdateResources(tx *gorm.DB, resources []*Resource) error {
	l := len(resources)
	if l == 0 {
		return nil
	}

	sql := "INSERT INTO resources (resource_id, resource_name, resource_group_id, created_at, updated_at) VALUES "

	for i := 0; i < l; i++ {
		r := (resources)[i]

		if strings.Trim(r.ResourceID, " ") == "" {
			r.ResourceID = idutil.GetUuid36("resource-")
		}

		item := fmt.Sprintf("('%s','%s','%s','%v','%v') ",
			r.ResourceID, r.ResourceName, r.ResourceGroupID, time.Now(), time.Now())

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}
	// on duplicate key update
	sql = sql + "on duplicate key update resource_name=values(resource_name),updated_at=values(updated_at)"

	if err := tx.Debug().Exec(sql).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	return nil
}

func GetResources(tx *gorm.DB, rgID string) ([]*Resource, error) {
	var resources []Resource
	sql := "SELECT r.* FROM resources as r WHERE r.resource_group_id=?"

	if err := tx.Debug().Raw(sql, rgID).Scan(&resources).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	l := len(resources)
	var ptrs = make([]*Resource, l)

	for i := 0; i < l; i++ {
		ptrs[i] = &resources[i]
	}
	return ptrs, nil
}

func CompareResources(p []*Resource, q []*Resource, rgID string) ([]*Resource, []*Resource) {
	l := len(q)

	var resID = make(map[string]int)

	for i := 0; i < l; i++ {
		resID[q[i].ResourceID] = i
	}

	l = len(p)
	var needDeleted []*Resource
	var needUpdated []*Resource
	for i := 0; i < l; i++ {
		r := p[i].ResourceID
		j, ok := resID[r]

		if r != "" && !ok {
			needDeleted = append(needDeleted, q[j])
		} else {
			p[i].ResourceGroupID = rgID
			needUpdated = append(needUpdated, p[i])
		}
	}

	return needDeleted, needUpdated
}

func DeleteResources(tx *gorm.DB, rgID string, resources []*Resource) error {

	if rgID == "" {
		return Error{Text: "resource group id must be specified", Code: InvalidParam}
	}

	if len(resources) == 0 {
		if err := tx.Debug().Exec("DELETE r FROM resources as r WHERE r.resource_group_id=?", rgID).Error; err != nil {
			return Error{Text: err.Error(), Code: DBError}
		}

	} else {
		var ids []string
		for i := 0; i < len(resources); i++ {
			r := resources[i]
			id := r.ResourceID
			if id != "" {
				ids = append(ids, id)
			}
		}

		if err := tx.Debug().Exec("DELETE r FROM resources as r WHERE r.resource_group_id=? AND r.resource_id IN (?)", rgID, ids).Error; err != nil {
			return Error{Text: err.Error(), Code: DBError}
		}
	}

	return nil
}
