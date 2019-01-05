package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"k8s.io/klog/glog"
	"time"
)

type ProductHandler struct{}

// product
func (server ProductHandler) CreateProduct(ctx context.Context, pbProd *pb.Product) (*pb.ProductResponse, error) {
	if pbProd.EnterpriseId == "" {
		ent, err := models.GetEnterprise(&models.Enterprise{EnterpriseName: pbProd.EnterpriseName})
		if err != nil {
			glog.Errorln(err.Error())
		}

		pbProd.EnterpriseId = ent.EnterpriseID
	}

	prod, err := models.CreateProduct(ConvertPB2Product(pbProd))

	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: ConvertProduct2PB(prod),
		Error: &pb.Error{
			Text: "success",
		},
	}, nil
}

func (server ProductHandler) DeleteProduct(ctx context.Context, prodSpec *pb.ProductSpec) (*pb.ProductResponse, error) {
	prodID := prodSpec.GetProductId()
	prodName := prodSpec.GetProductName()

	var pErr *pb.Error

	if prodID == "" && prodName == "" {
		pErr = &pb.Error{
			Text: "invalid param",
		}

		return &pb.ProductResponse{
			Error: pErr,
		}, nil
	}

	err := models.DeleteProduct(&models.Product{ProductName: prodName, ProductID: prodID})

	if err != nil {
		pErr = &pb.Error{
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Text: "success",
		}
	}

	return &pb.ProductResponse{
		Error: pErr,
	}, nil
}

func (server ProductHandler) UpdateProduct(ctx context.Context, pbProd *pb.Product) (*pb.ProductResponse, error) {
	prod := ConvertPB2Product(pbProd)

	err := models.UpdateProduct(prod)

	if err != nil {
		return &pb.ProductResponse{
			Error: &pb.Error{
				Text: err.Error(),
			},
		}, err
	}

	return &pb.ProductResponse{
		Error: &pb.Error{
			Text: "success",
		},
	}, nil

	return nil, nil
}

func (server ProductHandler) GetProduct(ctx context.Context, prodSpec *pb.ProductSpec) (*pb.ProductResponse, error) {
	prodID := prodSpec.GetProductId()
	prodName := prodSpec.GetProductName()

	var pErr *pb.Error

	if prodID == "" && prodName == "" {
		pErr = &pb.Error{
			Text: "invalid param",
		}

		return &pb.ProductResponse{
			Error: pErr,
		}, nil
	}

	ent, err := models.GetProduct(&models.Product{ProductName: prodName, ProductID: prodID})

	if err != nil {
		pErr = &pb.Error{
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Text: "success",
		}
	}

	return &pb.ProductResponse{
		Error:   pErr,
		Product: ConvertProduct2PB(ent),
	}, nil
	return nil, nil
}

func ConvertPB2Product(pbPrd *pb.Product) *models.Product {
	enp := models.Product{
		ProductID:         pbPrd.ProductId,
		ProductName:       pbPrd.ProductName,
		EnterpriseID:      pbPrd.EnterpriseId,
		HomePage:          pbPrd.HomePage,
		Email:             pbPrd.Email,
		Contacts:          pbPrd.Contacts,
		Description:       pbPrd.Desc,
		Phone:             pbPrd.Phone,
		Address:           pbPrd.Address,
		MonitorCenterPort: pbPrd.MonitorCenterPort,
		MonitorCenterHost: pbPrd.MonitorCenterHost,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	return &enp
}

func ConvertProduct2PB(prod *models.Product) *pb.Product {
	pbEnt := pb.Product{
		ProductId:         prod.ProductID,
		ProductName:       prod.ProductName,
		EnterpriseId:      prod.EnterpriseID,
		HomePage:          prod.HomePage,
		Address:           prod.Address,
		Email:             prod.Email,
		Contacts:          prod.Contacts,
		Desc:              prod.Description,
		Phone:             prod.Phone,
		MonitorCenterHost: prod.MonitorCenterHost,
		MonitorCenterPort: prod.MonitorCenterPort,
	}
	return &pbEnt
}
