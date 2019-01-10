package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"time"
)

type ResourceTypeHandler struct{}

// resource type
func (server ResourceTypeHandler) CreateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {

	if resourceType.ProductId == "" {
		return &pb.ResourceTypeResponse{
			Error: &pb.Error{

				Text: "product_id must be specified",
			},
		}, nil
	}

	prod, err := models.CreateResourceType(ConvertPB2ResourceType(resourceType))

	if err != nil {
		return nil, err
	}

	return &pb.ResourceTypeResponse{
		ResourceType: ConvertResourceType2PB(prod),
		Error: &pb.Error{

			Text: "success",
		},
	}, nil
}

func (server ResourceTypeHandler) UpdateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	if resourceType.ProductId == "" {
		return &pb.ResourceTypeResponse{
			Error: &pb.Error{

				Text: "product_id must be specified",
			},
		}, nil
	}

	err := models.UpdateResourceType(ConvertPB2ResourceType(resourceType))

	if err != nil {
		return nil, err
	}

	return &pb.ResourceTypeResponse{
		Error: &pb.Error{

			Text: "success",
		},
	}, nil
}

func (server ResourceTypeHandler) GetResourceType(ctx context.Context, resourceTypeSpec *pb.ResourceTypeSpec) (*pb.ResourceTypeResponse, error) {
	typeID := resourceTypeSpec.ResourceTypeId

	prodID := resourceTypeSpec.ProductId
	typeName := resourceTypeSpec.ResourceTypeName

	var pErr = &pb.Error{

		Text: "success",
	}

	var resourceType *models.ResourceType

	if typeID != "" {
		tp, err := models.GetResourceType(&models.ResourceType{ResourceTypeID: typeID})
		if err != nil {
			pErr = &pb.Error{

				Text: err.Error(),
			}
		}
		resourceType = tp
	} else if prodID != "" && typeName != "" {
		tp, err := models.GetResourceType(&models.ResourceType{ProductID: prodID, ResourceTypeName: typeName})
		if err != nil {
			pErr = &pb.Error{

				Text: err.Error(),
			}
		}
		resourceType = tp
	} else {
		pErr = &pb.Error{
			Text: "invalid param",
		}
	}

	return &pb.ResourceTypeResponse{
		Error:        pErr,
		ResourceType: ConvertResourceType2PB(resourceType),
	}, nil
}

func (server ResourceTypeHandler) DeleteResourceType(ctx context.Context, resourceTypeSpec *pb.ResourceTypeSpec) (*pb.ResourceTypeResponse, error) {
	typeID := resourceTypeSpec.ResourceTypeId

	prodID := resourceTypeSpec.ProductId
	typeName := resourceTypeSpec.ResourceTypeName

	var pErr = &pb.Error{

		Text: "success",
	}

	if typeID != "" {
		err := models.DeleteResourceType(&models.ResourceType{ResourceTypeID: typeID})
		if err != nil {
			pErr = &pb.Error{

				Text: err.Error(),
			}
		}
	} else if prodID != "" && typeName != "" {
		err := models.DeleteResourceType(&models.ResourceType{ProductID: prodID, ResourceTypeName: typeName})
		if err != nil {
			pErr = &pb.Error{

				Text: err.Error(),
			}
		}
	} else {
		pErr = &pb.Error{
			Text: "invalid param",
		}
	}

	return &pb.ResourceTypeResponse{
		Error: pErr,
	}, nil
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
