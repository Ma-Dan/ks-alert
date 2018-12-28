package models

import (
	"errors"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"time"
	"github.com/golang/glog"
)

type Resource struct {
	ResourceID      string    `gorm:"primary_key" json:"-"`
	ResourceName    string    `gorm:"type:varchar(50);" json:"resource_name"`
	ResourceTypeID  string    `gorm:"type:varchar(50);" json:"resource_type_id"`
	ResourceGroupID string    `gorm:"not null;" json:"-"`
	CreatedAt       time.Time `gorm:"not null;" json:"-"`
	UpdatedAt       time.Time `gorm:"not null;" json:"-"`
}

type ResourceGroup struct {
	ResourceGroupID   string     `gorm:"primary_key" json:"-"`
	ResourceGroupName string     `gorm:"type:varchar(50);not null;" json:"resource_group_name"`
	Resources         []Resource `gorm:"-" json:"resources"`

	URIParams       Params `gorm:"-" json:"resource_uri_params, omitempty"`
	URIParamsString string `gorm:"type:text;not null;" json:"-"`

	Description string    `gorm:"type:text;" json:"desc"`
	CreatedAt   time.Time `gorm:"not null;" json:"-"`
	UpdatedAt   time.Time `gorm:"not null;" json:"-"`
}

type ResourceType struct {
	ResourceTypeID   string `gorm:"primary_key;not null;"`
	ProductID        string `gorm:"not null;"`
	ResourceTypeName string `gorm:"type:varchar(50);not null;"`
	Description      string `gorm:"type:text;"`
	Enable           bool   `gorm:"type:boolean;not null;"`

	// MonitorCenterHost and MonitorCenterPort will override the corresponding filed in struct `product`
	MonitorCenterHost string `gorm:"type:varchar(128);"`
	MonitorCenterPort int32    `gorm:"type:int;"` //default:-1;
	// ResourceURITmpls struct -> json
	ResourceURITmpls string `gorm:"type:text;not null;"`

	CreatedAt      time.Time       `gorm:"not null;"`
	UpdatedAt      time.Time       `gorm:"not null;"`
	Metrics        []Metric        `gorm:"ForeignKey:ResourceTypeID;AssociationForeignKey:MetricID"`
	ResourceGroups []ResourceGroup `gorm:"ForeignKey:ResourceTypeID;AssociationForeignKey:ResourceGroupID"`
}

// the specific resource type has its own endpoint, can be expressed by URI
/**
{
	"uri_tmpls": [{
			"uri_tmpl": "namespces/{ns_name}/pods",
			"params": ["ns_name"]
		},
		{
			"uri_tmpl": "namespces/{ns_name}/pods/{pod_name}",
			"params": ["ns_name", "pod_name"]
		},
		{
			"uri_tmpl": "nodes/{node_id}/pods",
			"params": ["node_id"]
		},
		{
			"uri_tmpl": "nodes/{node_id}/pods/{pod_name}",
			"params": {"node_id":"", "pod_name":""},
			"resources":[]
		},
	]
}
*/

type URI string
type Params map[string]string
type Resources []string

type ResourceURITmpl struct {
	URI       URI       `json:"uri_tmpl,omitempty"`
	Params    Params    `json:"params,omitempty"`
	Resources Resources `json:"resource_name,omitempty"`
}

type ResourceURITmpls struct {
	ResourceURITmpl []ResourceURITmpl `json:"resource_uri_tmpl"`
}

//func GetResourceGroups(resourceType *ResourceType) *[]ResourceGroup {
//	db, err := dbutil.DBClient()
//	if err != nil {
//		panic(err)
//	}
//
//	var resourceGroup []ResourceGroup
//	db.Model(&Re{}).Where(resourceType).Find(&resourceGroup)
//
//	return &resourceGroup
//}

//func GetResourceGroupsByResourceTypeID(resourceTypeID string) *[]ResourceGroup {
//	db, err := dbutil.DBClient()
//	if err != nil {
//		panic(err)
//	}
//
//	var resourceGroups []ResourceGroup
//	db.Model(&ResourceGroup{}).Where(&ResourceGroup{ResourceTypeID: resourceTypeID}).Find(&resourceGroups)
//
//	return &resourceGroups
//}

func CreateResources(resources *[]Resource, resourceGroup *ResourceGroup, uriParams *Params) error {

	if resourceGroup == nil || resourceGroup.ResourceGroupID == "" {
		return errors.New("resource group create field")
	}

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	resourceGroupID := resourceGroup.ResourceGroupID

	for _, res := range *resources {

		if res.ResourceTypeID == "" {
			return errors.New("create resource field, resource type is not given")
		}

		res.ResourceID = idutil.GetUuid36("resource-")
		res.ResourceGroupID = resourceGroupID
		res.UpdatedAt = time.Now()

		err = db.Model(&Resource{}).Create(&res).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateResourceGroup(resourceGroupName, description string) (*ResourceGroup, error) {

	if resourceGroupName == "" {
		return nil, errors.New("resource Group Name is not given")
	}

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	var resourceGroup = &ResourceGroup{
		ResourceGroupID:   idutil.GetUuid36("resource_group-"),
		ResourceGroupName: resourceGroupName,
		Description:       description,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = db.Model(&ResourceGroup{}).Create(resourceGroup).Error
	return resourceGroup, err
}

//func GetResourceTypes(product *Product) *[]ResourceType {
//	db, err := dbutil.DBClient()
//	if err != nil {
//		panic(err)
//	}
//
//	var resourceTypes []ResourceType
//	db.Model(&Product{}).Where(product).Find(&resourceTypes)
//
//	return &resourceTypes
//}

func GetResourceType(resourceType *ResourceType) (*ResourceType, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var tp ResourceType
	err = db.Model(&ResourceType{}).Where(resourceType).Find(&tp).Error

	if err != nil {
		return nil, err
	}
	return &tp, nil
}

func CreateResourceType(resourceType *ResourceType) (*ResourceType, error){
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	resourceType.ResourceTypeID = idutil.GetUuid36("resource_type-")

	err = db.Model(&ResourceType{}).Create(resourceType).Error
	return resourceType, err

}

func UpdateResourceType(resourceType *ResourceType) (error){
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	if resourceType.ResourceTypeID != "" {
		err = db.Model(resourceType).Where("resource_type_id = ?", resourceType.ResourceTypeID).Update(resourceType).Error
	}else if resourceType.ProductID != "" && resourceType.ResourceTypeName != "" {
		err = db.Model(resourceType).Where("product_id = ? and resource_type_name = ? ", resourceType.ProductID, resourceType.ResourceTypeName).Update(resourceType).Error
	}

	return err
}

func DeleteResourceType(resourceType *ResourceType) (error){
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	err = db.Delete(resourceType).Error
	return err
}