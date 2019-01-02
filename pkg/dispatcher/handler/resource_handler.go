package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"github.com/pkg/errors"
	"k8s.io/klog/glog"
	"time"
)

type ResourceHandler struct{}

func (server ResourceHandler) CreateResource(ctx context.Context, rg *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {

	if rg.ResourceGroupName == "" || rg.ResourceTypeId == "" {
		return nil, errors.New("resource group name and resource type id must be specified")
	}

	if rg.Resources == nil || len(rg.Resources) == 0 {
		return nil, errors.New("resources must be specified")
	}

	_, err := DoTransactionAction(ConvertPB2ResourceGroup(rg), ResourceGroup, MethodCreate)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.ResourceGroupResponse{}, nil
}

func (server ResourceHandler) DeleteResource(ctx context.Context, rg *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	rgID := rg.ResourceGroupId

	if rgID == "" {
		return nil, errors.New("resource group id must be specified")
	}

	_, err := DoTransactionAction(&models.ResourceGroup{ResourceGroupID: rgID}, ResourceGroup, MethodDelete)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.ResourceGroupResponse{}, nil
}

func (server ResourceHandler) UpdateResource(ctx context.Context, rgSpec *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {
	if rgSpec.ResourceGroupId == "" {
		return nil, errors.New("resource group id must be specified")
	}

	rg := ConvertPB2ResourceGroup(rgSpec)

	v, err := DoTransactionAction(rg, ResourceGroup, MethodUpdate)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var respon *models.ResourceGroup

	if v != nil {
		respon = v.(*models.ResourceGroup)
	}

	return &pb.ResourceGroupResponse{
		ResourceGroup: ConvertResourceGroup2PB(respon),
	}, nil

}

func (server ResourceHandler) GetResource(ctx context.Context, rgSpec *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	rgID := rgSpec.ResourceGroupId

	if rgID == "" {
		return nil, errors.New("resource group id must be specified")
	}

	rg, err := DoTransactionAction(&models.ResourceGroup{ResourceGroupID: rgID}, ResourceGroup, MethodGet)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var respon *models.ResourceGroup

	if rg != nil {
		respon = rg.(*models.ResourceGroup)
	}

	return &pb.ResourceGroupResponse{
		ResourceGroup: ConvertResourceGroup2PB(respon),
	}, nil
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
		URIParams:         jsonutil.Marshal(rg.ResourceUriTmpls),
		Description:       rg.Desc,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func ConvertResourceGroup2PB(rg *models.ResourceGroup) *pb.ResourceGroup {
	if rg == nil {
		return nil
	}

	var v pb.ResourceUriTmpls
	jsonutil.Unmarshal(rg.URIParams, &v)

	return &pb.ResourceGroup{
		ResourceGroupId:   rg.ResourceGroupID,
		ResourceGroupName: rg.ResourceGroupName,
		ResourceTypeId:    rg.ResourceTypeID,
		Resources:         *ConvertResource2PB(rg.Resources),
		ResourceUriTmpls:  &v,
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
