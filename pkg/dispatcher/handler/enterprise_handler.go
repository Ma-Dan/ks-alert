package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"time"
)

type EnterpriseHandler struct{}

// enterprise
func (h EnterpriseHandler) CreateEnterprise(ctx context.Context, pbEnt *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	enterprise, err := models.CreateEnterprise(ConvertPB2Enterprise(pbEnt))
	return getEnterpriseResponse(enterprise, err), nil
}

func getEnterpriseResponse(enterprise *models.Enterprise, err error) *pb.EnterpriseResponse {
	arg := ConvertEnterprise2PB(enterprise)

	var respon = pb.EnterpriseResponse{Enterprise: arg}
	respon.Error = ErrorWrapper(err)

	return &respon
}

func (h EnterpriseHandler) DeleteEnterprise(ctx context.Context, entSpec *pb.EnterpriseSpec) (*pb.EnterpriseResponse, error) {
	entID := entSpec.GetEnterpriseId()
	entName := entSpec.GetEnterpriseName()
	err := models.DeleteEnterprise(&models.Enterprise{EnterpriseName: entName, EnterpriseID: entID})
	return getEnterpriseResponse(nil, err), nil
}

func (h EnterpriseHandler) UpdateEnterprise(ctx context.Context, pbEnt *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	ent := ConvertPB2Enterprise(pbEnt)
	err := models.UpdateEnterprise(ent)
	return getEnterpriseResponse(nil, err), nil
}

func (h EnterpriseHandler) GetEnterprise(ctx context.Context, entSpec *pb.EnterpriseSpec) (*pb.EnterpriseResponse, error) {
	entID := entSpec.GetEnterpriseId()
	entName := entSpec.GetEnterpriseName()
	ent, err := models.GetEnterprise(&models.Enterprise{EnterpriseName: entName, EnterpriseID: entID})
	return getEnterpriseResponse(ent, err), nil
}

func ConvertPB2Enterprise(pbEnt *pb.Enterprise) *models.Enterprise {
	enp := models.Enterprise{
		EnterpriseName: pbEnt.EnterpriseName,
		HomePage:       pbEnt.HomePage,
		Address:        pbEnt.Address,
		Email:          pbEnt.Email,
		Contacts:       pbEnt.Contacts,
		Description:    pbEnt.Desc,
		Phone:          pbEnt.Phone,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return &enp
}

func ConvertEnterprise2PB(enp *models.Enterprise) *pb.Enterprise {
	if enp == nil {
		return nil
	}
	pbEnt := pb.Enterprise{
		EnterpriseId:   enp.EnterpriseID,
		EnterpriseName: enp.EnterpriseName,
		HomePage:       enp.HomePage,
		Address:        enp.Address,
		Email:          enp.Email,
		Contacts:       enp.Contacts,
		Desc:           enp.Description,
		Phone:          enp.Phone,
	}
	return &pbEnt
}
