package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/pb"
	"kubesphere.io/ks-alert/pkg/stderr"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	"time"
)

type ResourceHandler struct{}

func (server ResourceHandler) CreateResource(ctx context.Context, rg *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {
	v, err := DoTransactionAction(ConvertPB2ResourceGroup(rg), MethodCreate)
	respon := getResourceGroupResponse(v, err)
	return respon, nil
}

func getResourceGroupResponse(v interface{}, err error) *pb.ResourceGroupResponse {
	var resGroup *models.ResourceGroup
	if v != nil {
		resGroup = v.(*models.ResourceGroup)
	}

	rg := ConvertResourceGroup2PB(resGroup)

	var respon = pb.ResourceGroupResponse{ResourceGroup: rg}
	respon.Error = stderr.ErrorWrapper(err)
	return &respon
}

func (server ResourceHandler) UpdateResource(ctx context.Context, rgSpec *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {
	v, err := DoTransactionAction(ConvertPB2ResourceGroup(rgSpec), MethodUpdate)
	respon := getResourceGroupResponse(v, err)
	return respon, nil
}

func (server ResourceHandler) GetResource(ctx context.Context, rgSpec *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	v, err := DoTransactionAction(&models.ResourceGroup{ResourceGroupID: rgSpec.ResourceGroupId}, MethodGet)
	respon := getResourceGroupResponse(v, err)
	return respon, nil
}

func (server ResourceHandler) DeleteResource(ctx context.Context, rg *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	v, err := DoTransactionAction(&models.ResourceGroup{ResourceGroupID: rg.ResourceGroupId}, MethodDelete)
	respon := getResourceGroupResponse(v, err)
	return respon, nil
}

func ConvertPB2ResourceGroup(rg *pb.ResourceGroup) *models.ResourceGroup {
	if rg == nil {
		return nil
	}

	return &models.ResourceGroup{
		ResourceGroupID:   rg.ResourceGroupId,
		ResourceGroupName: rg.ResourceGroupName,
		ResourceTypeID:    rg.ResourceTypeId,
		Resources:         *ConvertPB2Resource(rg.Resources),
		URIParams:         jsonutil.Marshal(rg.ResourceUriTmpl),
		Description:       rg.Desc,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func ConvertResourceGroup2PB(rg *models.ResourceGroup) *pb.ResourceGroup {
	if rg == nil {
		return nil
	}

	var v pb.ResourceUriTmpl
	jsonutil.Unmarshal(rg.URIParams, &v)

	return &pb.ResourceGroup{
		ResourceGroupId:   rg.ResourceGroupID,
		ResourceGroupName: rg.ResourceGroupName,
		ResourceTypeId:    rg.ResourceTypeID,
		Resources:         *ConvertResource2PB(rg.Resources),
		ResourceUriTmpl:   &v,
		Desc:              rg.Description,
	}
}

func ConvertResource2PB(res []*models.Resource) *[]*pb.Resource {
	l := len(res)
	var pbResources = make([]*pb.Resource, l)
	for i := 0; i < l; i++ {
		r := res[i]

		pbResources[i] = &pb.Resource{
			ResourceId:      r.ResourceID,
			ResourceGroupId: r.ResourceGroupID,
			ResourceName:    r.ResourceName,
		}
	}
	return &pbResources
}

func ConvertPB2Resource(res []*pb.Resource) *[]*models.Resource {
	l := len(res)
	var resources = make([]*models.Resource, l)
	for i := 0; i < l; i++ {
		r := res[i]

		resources[i] = &models.Resource{
			ResourceID:      r.ResourceId,
			ResourceGroupID: r.ResourceGroupId,
			ResourceName:    r.ResourceName,
			UpdatedAt:       time.Now(),
			CreatedAt:       time.Now(),
		}
	}
	return &resources
}
