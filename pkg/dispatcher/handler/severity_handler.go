package handler

import (
	"context"
	"errors"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/golang/glog"
	"time"
)

type SeverityHandler struct{}

// Severity
func (server SeverityHandler) CreateSeverity(ctx context.Context, pbSev *pb.Severity) (*pb.SeverityResponse, error) {

	if pbSev.SeverityEn == "" || pbSev.SeverityCh == "" || pbSev.ProductId == "" {
		return nil, errors.New("severity name or product id must be specified")
	}

	severity, err := models.CreateSeverity(ConvertPB2Severity(pbSev))

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.SeverityResponse{
		Severity: ConvertSeverity2PB(severity),
	}, nil
}

func (server SeverityHandler) DeleteSeverity(ctx context.Context, sevSpec *pb.SeveritySpec) (*pb.SeverityResponse, error) {
	if sevSpec.SeverityId == "" {
		return nil, errors.New("severity id must be specified")
	}

	_, err := models.DeleteSeverity(sevSpec)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.SeverityResponse{}, nil
}

func (server SeverityHandler) UpdateSeverity(ctx context.Context, pbSev *pb.Severity) (*pb.SeverityResponse, error) {

	if pbSev.SeverityEn == "" || pbSev.SeverityCh == "" || pbSev.SeverityId == "" {
		return nil, errors.New("severity name or id must be specified")
	}

	severity, err := models.UpdateSeverity(ConvertPB2Severity(pbSev))

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.SeverityResponse{
		Severity: ConvertSeverity2PB(severity),
	}, nil
}

func (server SeverityHandler) GetSeverity(ctx context.Context, sevSpec *pb.SeveritySpec) (*pb.SeveritiesResponse, error) {

	if sevSpec.SeverityId == "" && sevSpec.ProductId == "" && sevSpec.ProductName == "" {
		return nil, errors.New("severity id or product id or product name must be specified")
	}

	// get single Severity
	if sevSpec.SeverityId != "" {
		severity, err := models.GetSeverity(sevSpec)

		if err != nil {
			return nil, err
		}

		if severity != nil || len(*severity) > 0 {
			return &pb.SeveritiesResponse{
				Severity: []*pb.Severity{ConvertSeverity2PB(&(*severity)[0])},
			}, nil
		}
	}

	// get Severities by product_id
	if sevSpec.ProductId == "" && sevSpec.ProductName != "" {

		product, err := models.GetProduct(&models.Product{ProductName: sevSpec.ProductName})
		if err != nil {
			return nil, err
		}

		sevSpec.ProductId = product.ProductID

		severity, err := models.GetSeverity(sevSpec)
		if err != nil {
			return nil, err
		}

		if severity == nil {
			return nil, nil
		}

		l := len(*severity)
		var respon = make([]*pb.Severity, l)

		for i := 0; i < l; i++ {
			respon[i] = ConvertSeverity2PB(&(*severity)[i])
		}

		return &pb.SeveritiesResponse{
			Severity: respon,
		}, nil

	}

	return nil, nil
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
