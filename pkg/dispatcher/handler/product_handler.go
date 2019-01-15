package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"time"
)

type ProductHandler struct{}

// product
func (h ProductHandler) CreateProduct(ctx context.Context, pbProd *pb.Product) (*pb.ProductResponse, error) {
	if pbProd.EnterpriseId == "" {
		ent, err := models.GetEnterprise(&models.Enterprise{EnterpriseName: pbProd.EnterpriseName})
		if err != nil {
			return getProductResponse(nil, err), nil
		}
		pbProd.EnterpriseId = ent.EnterpriseID
	}

	prod, err := models.CreateProduct(ConvertPB2Product(pbProd))
	return getProductResponse(prod, err), nil
}

func getProductResponse(product *models.Product, err error) *pb.ProductResponse {
	arg := ConvertProduct2PB(product)

	var respon = pb.ProductResponse{Product: arg}
	respon.Error = ErrorWrapper(err)

	return &respon
}

func (h ProductHandler) DeleteProduct(ctx context.Context, prodSpec *pb.ProductSpec) (*pb.ProductResponse, error) {
	prodID := prodSpec.GetProductId()
	prodName := prodSpec.GetProductName()

	if prodID == "" && prodName == "" {
		return getProductResponse(nil, models.Error{
			Code: models.InvalidParam,
			Text: "product id and product name must be specified"}), nil
	}

	err := models.DeleteProduct(&models.Product{ProductName: prodName, ProductID: prodID})

	return getProductResponse(nil, err), nil
}

func (h ProductHandler) UpdateProduct(ctx context.Context, pbProd *pb.Product) (*pb.ProductResponse, error) {
	prod := ConvertPB2Product(pbProd)
	err := models.UpdateProduct(prod)
	return getProductResponse(nil, err), nil
}

func (h ProductHandler) GetProduct(ctx context.Context, prodSpec *pb.ProductSpec) (*pb.ProductResponse, error) {
	prodID := prodSpec.GetProductId()
	prodName := prodSpec.GetProductName()

	if prodID == "" && prodName == "" {
		return getProductResponse(nil, models.Error{
			Code: models.InvalidParam,
			Text: "product id and product name must be specified"}), nil
	}

	ent, err := models.GetProduct(&models.Product{ProductName: prodName, ProductID: prodID})

	return getProductResponse(ent, err), nil
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
