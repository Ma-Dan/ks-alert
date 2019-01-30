package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"time"
)

type ResourceTypeHandler struct{}

// resource type
func (server ResourceTypeHandler) CreateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {

	if resourceType.ProductId == "" {
		return getResourceTypeResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "product_id must be specified"}), nil
	}

	prod, err := models.CreateResourceType(ConvertPB2ResourceType(resourceType))
	return getResourceTypeResponse(prod, err), nil
}

func getResourceTypeResponse(resourceType *models.ResourceType, err error) *pb.ResourceTypeResponse {
	arg := ConvertResourceType2PB(resourceType)
	var respon = pb.ResourceTypeResponse{ResourceType: arg}
	respon.Error = stderr.ErrorWrapper(err)

	return &respon
}

func (server ResourceTypeHandler) UpdateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	if resourceType.ProductId == "" {
		return getResourceTypeResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "product_id must be specified"}), nil
	}

	err := models.UpdateResourceType(ConvertPB2ResourceType(resourceType))

	return getResourceTypeResponse(nil, err), nil
}

func (server ResourceTypeHandler) GetResourceType(ctx context.Context, resourceTypeSpec *pb.ResourceTypeSpec) (*pb.ResourceTypeResponse, error) {
	typeID := resourceTypeSpec.ResourceTypeId

	prodID := resourceTypeSpec.ProductId
	typeName := resourceTypeSpec.ResourceTypeName

	if typeID != "" {
		tp, err := models.GetResourceType(&models.ResourceType{ResourceTypeID: typeID})
		return getResourceTypeResponse(tp, err), nil
	} else if prodID != "" && typeName != "" {
		tp, err := models.GetResourceType(&models.ResourceType{ProductID: prodID, ResourceTypeName: typeName})
		return getResourceTypeResponse(tp, err), nil
	} else {
		return getResourceTypeResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "invalid params"}), nil
	}
}

func (server ResourceTypeHandler) DeleteResourceType(ctx context.Context, resourceTypeSpec *pb.ResourceTypeSpec) (*pb.ResourceTypeResponse, error) {
	typeID := resourceTypeSpec.ResourceTypeId

	prodID := resourceTypeSpec.ProductId
	typeName := resourceTypeSpec.ResourceTypeName

	var err error

	if typeID != "" {
		err = models.DeleteResourceType(&models.ResourceType{ResourceTypeID: typeID})
	} else if prodID != "" && typeName != "" {
		err = models.DeleteResourceType(&models.ResourceType{ProductID: prodID, ResourceTypeName: typeName})
	} else {
		err = stderr.Error{
			Code: stderr.InvalidParam,
			Text: "invalid params"}
	}
	return getResourceTypeResponse(nil, err), nil
}

func ConvertPB2ResourceType(pbPrd *pb.ResourceType) *models.ResourceType {

	enp := models.ResourceType{
		ProductID:         pbPrd.ProductId,
		ResourceTypeName:  pbPrd.ResourceTypeName,
		MonitorCenterHost: pbPrd.MonitorCenterHost,
		MonitorCenterPort: pbPrd.MonitorCenterPort,
		Description:       pbPrd.Desc,
		ResourceTypeID:    pbPrd.ResourceTypeId,
		ResourceURITmpls:  jsonutil.Marshal(pbPrd.ResourceUriTmpl),
		Enable:            pbPrd.Enable,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	return &enp
}

func ConvertResourceType2PB(prod *models.ResourceType) *pb.ResourceType {
	if prod == nil {
		return nil
	}

	uriTmpls := pb.ResourceUriTmpls{}

	jsonutil.Unmarshal(prod.ResourceURITmpls, &uriTmpls)

	pbEnt := pb.ResourceType{
		ProductId:         prod.ProductID,
		ResourceTypeName:  prod.ResourceTypeName,
		ResourceTypeId:    prod.ResourceTypeID,
		MonitorCenterHost: prod.MonitorCenterHost,
		MonitorCenterPort: prod.MonitorCenterPort,
		Enable:            prod.Enable,
		ResourceUriTmpl:   &uriTmpls,
		Desc:              prod.Description,
	}

	return &pbEnt
}
