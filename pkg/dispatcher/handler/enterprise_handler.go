package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"time"
)

type EnterpriseHandler struct{}

// enterprise
func (server EnterpriseHandler) CreateEnterprise(ctx context.Context, pbEnt *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	enterprise, err := models.CreateEnterprise(ConvertPB2Enterprise(pbEnt))

	if err != nil {
		return nil, err
	}

	return &pb.EnterpriseResponse{
		Enterprise: ConvertEnterprise2PB(enterprise),
		Error: &pb.Error{
			Text: "success",
		},
	}, nil
}

func (server EnterpriseHandler) DeleteEnterprise(ctx context.Context, entSpec *pb.EnterpriseSpec) (*pb.EnterpriseResponse, error) {
	entID := entSpec.GetEnterpriseId()
	entName := entSpec.GetEnterpriseName()

	var pErr *pb.Error

	if entID == "" && entName == "" {
		pErr = &pb.Error{
			Text: "invalid param",
		}

		return &pb.EnterpriseResponse{
			Error: pErr,
		}, nil
	}

	err := models.DeleteEnterprise(&models.Enterprise{EnterpriseName: entName, EnterpriseID: entID})

	if err != nil {
		pErr = &pb.Error{
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Text: "success",
		}
	}

	return &pb.EnterpriseResponse{
		Error: pErr,
	}, nil
}

func (server EnterpriseHandler) UpdateEnterprise(ctx context.Context, pbEnt *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	ent := ConvertPB2Enterprise(pbEnt)
	err := models.UpdateEnterprise(ent)

	if err != nil {
		return &pb.EnterpriseResponse{
			Error: &pb.Error{
				Text: err.Error(),
			},
		}, err
	}

	return &pb.EnterpriseResponse{
		Error: &pb.Error{
			Text: "success",
		},
	}, nil

	return nil, nil
}
func (server EnterpriseHandler) GetEnterprise(ctx context.Context, entSpec *pb.EnterpriseSpec) (*pb.EnterpriseResponse, error) {
	entID := entSpec.GetEnterpriseId()
	entName := entSpec.GetEnterpriseName()

	var pErr *pb.Error

	if entID == "" && entName == "" {
		pErr = &pb.Error{
			Text: "invalid param",
		}

		return &pb.EnterpriseResponse{
			Error: pErr,
		}, nil
	}

	ent, err := models.GetEnterprise(&models.Enterprise{EnterpriseName: entName, EnterpriseID: entID})

	if err != nil {
		pErr = &pb.Error{
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Text: "success",
		}
	}

	return &pb.EnterpriseResponse{
		Error:      pErr,
		Enterprise: ConvertEnterprise2PB(ent),
	}, nil
	return nil, nil
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
