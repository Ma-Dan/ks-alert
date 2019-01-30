package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/stderr"
	"time"
)

type SeverityHandler struct{}

// Severity
func (server SeverityHandler) CreateSeverity(ctx context.Context, pbSev *pb.Severity) (*pb.SeverityResponse, error) {

	if pbSev.SeverityEn == "" || pbSev.SeverityCh == "" || (pbSev.ProductId == "" && pbSev.SeverityId == "") {
		return getSeverityResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "severity name or product id or severity id must be specified",
		}), nil
	}

	severity, err := models.CreateSeverity(ConvertPB2Severity(pbSev))

	return getSeverityResponse(severity, err), nil
}

func getSeverityResponse(severity *models.Severity, err error) *pb.SeverityResponse {
	arg := ConvertSeverity2PB(severity)
	var respon = pb.SeverityResponse{Severity: arg}
	respon.Error = stderr.ErrorWrapper(err)

	return &respon
}

func getSeveritiesResponse(severity *[]models.Severity, err error) *pb.SeveritiesResponse {
	l := 0
	if severity != nil {
		l = len(*severity)
	}

	var pbSeverity = make([]*pb.Severity, l)

	for i := 0; i < l; i++ {
		pbSeverity[i] = ConvertSeverity2PB(&(*severity)[i])
	}

	var respon = pb.SeveritiesResponse{Severity: pbSeverity}
	respon.Error = stderr.ErrorWrapper(err)
	return &respon
}

func (server SeverityHandler) DeleteSeverity(ctx context.Context, sevSpec *pb.SeveritySpec) (*pb.SeverityResponse, error) {
	if sevSpec.SeverityId == "" {
		return getSeverityResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "severity id must be specified",
		}), nil
	}

	_, err := models.DeleteSeverity(sevSpec)
	return getSeverityResponse(nil, err), nil
}

func (server SeverityHandler) UpdateSeverity(ctx context.Context, pbSev *pb.Severity) (*pb.SeverityResponse, error) {

	if pbSev.SeverityEn == "" || pbSev.SeverityCh == "" || pbSev.SeverityId == "" {
		return getSeverityResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "severity name or product id must be specified",
		}), nil
	}

	severity, err := models.UpdateSeverity(ConvertPB2Severity(pbSev))
	return getSeverityResponse(severity, err), nil
}

func (server SeverityHandler) GetSeverity(ctx context.Context, sevSpec *pb.SeveritySpec) (*pb.SeveritiesResponse, error) {

	if sevSpec.SeverityId == "" && sevSpec.ProductId == "" && sevSpec.ProductName == "" {
		return getSeveritiesResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "severity id or product id or product name must be specified",
		}), nil
	}

	// get single Severity
	if sevSpec.SeverityId != "" {
		severity, err := models.GetSeverity(sevSpec)
		return getSeveritiesResponse(severity, err), nil
	}

	// get Severities by product_id
	var prodID = sevSpec.ProductId
	if prodID == "" && sevSpec.ProductName != "" {

		product, err := models.GetProduct(&models.Product{ProductName: sevSpec.ProductName})
		if err != nil {
			return getSeveritiesResponse(nil, err), nil
		}

		prodID = product.ProductID
	}

	sevSpec.ProductId = prodID
	severity, err := models.GetSeverity(sevSpec)
	return getSeveritiesResponse(severity, err), nil
}

func ConvertSeverity2PB(sev *models.Severity) *pb.Severity {
	if sev == nil {
		return nil
	}
	return &pb.Severity{
		SeverityId: sev.SeverityID,
		ProductId:  sev.ProductID,
		SeverityCh: sev.SeverityCh,
		SeverityEn: sev.SeverityEn,
	}
}

func ConvertPB2Severity(pbSev *pb.Severity) *models.Severity {
	if pbSev == nil {
		return nil
	}

	return &models.Severity{
		SeverityID: pbSev.SeverityId,
		ProductID:  pbSev.ProductId,
		SeverityEn: pbSev.SeverityEn,
		SeverityCh: pbSev.SeverityCh,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}
}
