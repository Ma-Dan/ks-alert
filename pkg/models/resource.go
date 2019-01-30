package models

import (
	"fmt"
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Resource struct {
	Action
	ResourceID      string    `gorm:"primary_key"`
	ResourceName    string    `gorm:"type:varchar(50);"`
	ResourceGroupID string    `gorm:"not null;"`
	CreatedAt       time.Time `gorm:"not null;"`
	UpdatedAt       time.Time `gorm:"not null;"`
}

type ResourceGroup struct {
	Action
	ResourceGroupID   string      `gorm:"primary_key"`
	ResourceGroupName string      `gorm:"type:varchar(50);not null;"`
	ResourceTypeID    string      `gorm:"type:varchar(50);"`
	Resources         []*Resource `gorm:"-"`
	URIParams         string      `gorm:"type:text;not null;"`
	Description       string      `gorm:"type:text;"`
	CreatedAt         time.Time   `gorm:"not null;"`
	UpdatedAt         time.Time   `gorm:"not null;"`
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
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	var tp ResourceType
	db = db.Model(&ResourceType{}).Where(resourceType).First(&tp)

	if db.RecordNotFound() {
		return nil, nil
	}

	if db.Error != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return &tp, nil
}

func CreateResourceType(resourceType *ResourceType) (*ResourceType, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	resourceType.ResourceTypeID = idutil.GetUuid36("")

	err = db.Model(&ResourceType{}).Create(resourceType).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return resourceType, nil

}

func UpdateResourceType(resourceType *ResourceType) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	if resourceType.ResourceTypeID != "" {
		err = db.Model(resourceType).Where("resource_type_id = ?", resourceType.ResourceTypeID).Update(resourceType).Error
	} else if resourceType.ProductID != "" && resourceType.ResourceTypeName != "" {
		err = db.Model(resourceType).Where("product_id = ? and resource_type_name = ? ", resourceType.ProductID, resourceType.ResourceTypeName).Update(resourceType).Error
	}

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}

func DeleteResourceType(resourceType *ResourceType) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	err = db.Delete(resourceType).Error

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}

func (r *ResourceGroup) Create(tx *gorm.DB) (interface{}, error) {
	if r.ResourceGroupName == "" || r.ResourceTypeID == "" {
		return nil, Error{Text: "resource group name and resource type id must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	if r.Resources == nil || len(r.Resources) == 0 {
		return nil, Error{Text: "resources must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	var resourceWithName []*Resource
	resources := r.Resources
	for i := 0; i < len(resources); i++ {
		rs := resources[i]
		if rs.ResourceName != "" {
			resourceWithName = append(resourceWithName, rs)
		}
	}

	if len(resourceWithName) == 0 {
		return nil, Error{Text: "at least one resource name must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	r.ResourceGroupID = idutil.GetUuid36("")

	// create group
	item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%v')", r.ResourceGroupID, r.ResourceGroupName,
		r.ResourceTypeID, r.Description, r.URIParams, r.CreatedAt, r.UpdatedAt)

	sql := "INSERT INTO resource_groups (resource_group_id, resource_group_name, resource_type_id, " +
		"description, uri_params, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	// create item
	l := len(resourceWithName)
	for i := 0; i < l; i++ {
		rs := resourceWithName[i]
		rs.ResourceID = idutil.GetUuid36("")
		rs.ResourceGroupID = r.ResourceGroupID
	}

	if err := CreateOrUpdateResources(tx, resourceWithName); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *ResourceGroup) Update(tx *gorm.DB) (interface{}, error) {
	if r.ResourceGroupID == "" || r.ResourceGroupName == "" {
		return nil, Error{Text: "resource group id or name must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	// 1. get resource group first
	vget, err := r.Get(tx)

	if err != nil {
		return nil, err
	}

	rg := vget.(*ResourceGroup)

	if rg == nil || rg.ResourceGroupID == "" {
		return nil, Error{Text: fmt.Sprintf("resource group id: %s not exist", r.ResourceGroupID), Code: InvalidParam, Where: Caller(1, true)}
	}

	// 2. update resource group
	sql := fmt.Sprintf("UPDATE resource_groups SET resource_group_name='%s',description='%s',"+
		"uri_params='%v', updated_at='%v' WHERE resource_group_id='%s'",
		r.ResourceGroupName, r.Description,
		r.URIParams, time.Now(), r.ResourceGroupID)

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	// 3. delete resources or update resource
	var needDeleted = CompareResources(r.Resources, rg.Resources)

	err = DeleteResources(tx, needDeleted)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(r.Resources); i++ {
		r.Resources[i].ResourceGroupID = r.ResourceGroupID
		if r.Resources[i].ResourceID == "" {
			r.Resources[i].ResourceID = idutil.GetUuid36("")
		}
	}

	err = CreateOrUpdateResources(tx, r.Resources)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *ResourceGroup) Get(tx *gorm.DB) (interface{}, error) {
	rgID := r.ResourceGroupID

	if rgID == "" {
		return nil, Error{Text: "resource group id must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	var rg ResourceGroup

	err := tx.Model(&ResourceGroup{}).Where("resource_group_id=?", rgID).First(&rg).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	exist := tx.RecordNotFound()

	if exist {
		return nil, Error{Text: "record not found", Code: DBError, Where: Caller(1, true)}
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

func (r *ResourceGroup) Delete(tx *gorm.DB) (interface{}, error) {
	rgID := r.ResourceGroupID

	if rgID == "" {
		return nil, Error{Text: "resource group id must be specified", Code: InvalidParam, Where: Caller(1, true)}
	}

	sql := "DELETE rg, r FROM resource_groups as rg LEFT JOIN resources as r ON rg.resource_group_id=r.resource_group_id WHERE rg.resource_group_id=?"

	if err := tx.Exec(sql, r.ResourceGroupID).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
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
			r.ResourceID = idutil.GetUuid36("")
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

	if err := tx.Exec(sql).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}

func GetResources(tx *gorm.DB, rgID string) ([]*Resource, error) {
	var resources []Resource
	sql := "SELECT r.* FROM resources as r WHERE r.resource_group_id=?"

	if err := tx.Raw(sql, rgID).Scan(&resources).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	l := len(resources)
	var ptrs = make([]*Resource, l)

	for i := 0; i < l; i++ {
		ptrs[i] = &resources[i]
	}
	return ptrs, nil
}

func CompareResources(p []*Resource, q []*Resource) []*Resource {
	l := len(q)

	var resID = make(map[string]int)

	for i := 0; i < l; i++ {
		resID[q[i].ResourceID] = i
	}

	l = len(p)
	for i := 0; i < l; i++ {
		r := p[i].ResourceID
		if r != "" {
			delete(resID, r)
		}
	}

	var needDeleted []*Resource
	for _, i := range resID {
		needDeleted = append(needDeleted, q[i])
	}

	return needDeleted
}

func DeleteResources(tx *gorm.DB, resources []*Resource) error {
	if len(resources) == 0 {
		return nil
	}

	var ids []string
	for i := 0; i < len(resources); i++ {
		r := resources[i]
		id := r.ResourceID
		if id != "" {
			ids = append(ids, id)
		}
	}

	if err := tx.Exec("DELETE r FROM resources as r WHERE r.resource_id IN (?)", ids).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}
