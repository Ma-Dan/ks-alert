package models

import (
	"kubesphere.io/alert-kubesphere-plugin/pkg/client"
	"kubesphere.io/alert-kubesphere-plugin/pkg/utils/idutil"
	"kubesphere.io/alert-kubesphere-plugin/pkg/utils/jsonutil"
	"time"
)

type ResourceGroup struct {
	ResourceGroupID   string    `gorm:"primary_key"`
	ResourceTypeID    string    `gorm:"not null;"`
	ResourceTypeName  string    `gorm:"type:varchar(50);"`
	ResourceGroupName string    `gorm:"type:varchar(50);not null;"`
	ResourceURITmpls  string    `gorm:"type:text;not null;"`
	CreatedBy         string    `gorm:"type:varchar(50);not null;"`
	Description       string    `gorm:"type:text;"`
	CreatedAt         time.Time `gorm:"not null;"`
	UpdatedAt         time.Time `gorm:"not null;"`
}

type ResourceType struct {
	ResourceTypeID   string `gorm:"primary_key"`
	ProductID        string `gorm:"not null;"`
	ResourceTypeName string `gorm:"type:varchar(50);not null;"`
	// MonitorCenterHost and MonitorCenterPort will override the corresponding filed in struct `product`
	MonitorCenterHost string          `gorm:"type:varchar(128);not null;"`
	MonitorCenterPort int             `gorm:"type:int;not null;default:-1;"`
	Description       string          `gorm:"type:text;"`
	Enable            bool            `gorm:"type:boolean;not null;"`
	Endpoint          string          `gorm:"type:text;not null;"`
	CreatedAt         time.Time       `gorm:"not null;"`
	UpdatedAt         time.Time       `gorm:"not null;"`
	Metrics           []Metric        `gorm:"ForeignKey:ResourceTypeID;AssociationForeignKey:MetricID"`
	ResourceGroups    []ResourceGroup `gorm:"ForeignKey:ResourceTypeID;AssociationForeignKey:ResourceGroupID"`
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

type ResourceURITmpls struct {
	ResourceURITmpl []ResourceURITmpl `json:"uri_tmpls"`
}

type ResourceURITmpl struct {
	URI       string            `json:"uri_tmpl,omitempty"`
	Params    map[string]string `json:"params,omitempty"`
	Resources []string          `json:"resources,omitempty"`
}

//func GetResourceGroups(resourceType *ResourceType) *[]ResourceGroup {
//	db, err := client.DBClient()
//	if err != nil {
//		panic(err)
//	}
//
//	var resourceGroup []ResourceGroup
//	db.Model(&Re{}).Where(resourceType).Find(&resourceGroup)
//
//	return &resourceGroup
//}

func GetResourceGroupsByResourceTypeID(resourceTypeID string) *[]ResourceGroup {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	var resourceGroups []ResourceGroup
	db.Model(&ResourceGroup{}).Where(&ResourceGroup{ResourceTypeID: resourceTypeID}).Find(&resourceGroups)

	return &resourceGroups
}

func CreateResourceGroup(resourceGroupName, resourceTypeID, resourceTypeName, createdBy, description string, resourceUri *ResourceURITmpls) error {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	var resourceType = &ResourceGroup{
		ResourceGroupID:   idutil.GetUuid36("resource_group-"),
		ResourceGroupName: resourceGroupName,
		ResourceTypeID:    resourceTypeID,
		CreatedBy:         createdBy,
		ResourceTypeName:  resourceTypeName,
		Description:       description,
		ResourceURITmpls:  jsonutil.Marshal(resourceUri),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = db.Model(&Product{}).Create(resourceType).Error
	return err
}

//func GetResourceTypes(product *Product) *[]ResourceType {
//	db, err := client.DBClient()
//	if err != nil {
//		panic(err)
//	}
//
//	var resourceTypes []ResourceType
//	db.Model(&Product{}).Where(product).Find(&resourceTypes)
//
//	return &resourceTypes
//}

func GetResourceTypeByProductID(productID string) *[]ResourceType {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	var resourceTypes []ResourceType
	db.Model(&ResourceType{}).Where(&ResourceType{ProductID: productID}).Find(&resourceTypes)

	return &resourceTypes
}

func GetAllResourceTypeCount(productID string) int {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	var count int
	db.Model(&Product{}).Where(&Product{ProductID: productID}).Count(&count)

	return count
}

func CreateSourceType(productID, resourceTypeName, description string, enable bool, resourceUri *ResourceURITmpls) error {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	var resourceType = &ResourceType{
		ResourceTypeID:   idutil.GetUuid36("resource_type-"),
		ResourceTypeName: resourceTypeName,
		ProductID:        productID,
		Enable:           enable,
		Description:      description,
		Endpoint:         jsonutil.Marshal(resourceUri),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = db.Model(&Product{}).Create(resourceType).Error
	return err
}
