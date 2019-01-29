package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// Grpc error
type GrpcStatusError interface {
	GRPCStatus() *status.Status
	Error() string
}

// error adaptor, may convert many error type to `error` type
func ErrorWrapper(vErr interface{}, extMsg ...string) *pb.Error {
	var txt string
	var code int32
	var where string
	switch v := vErr.(type) {
	case *pb.Error:
		txt = v.Text
		code = v.Code
		where = v.Where
	case models.Error:
		txt = v.Text
		code = v.Code
		where = v.Where
	case GrpcStatusError:
		s := v.GRPCStatus()
		if s.Code() == codes.Unknown {
			code = -1
		} else if s.Code() < 20 {
			code = 4
		}
		txt = v.Error()
	case error:
		txt = v.Error()
		code = -1
	case nil:
		txt = "success"
		code = 0
	}

	if len(extMsg) != 0 {
		txt = strings.Join(extMsg, " : ") + " : " + txt
	}

	return &pb.Error{Text: txt, Code: code, Where: where}
}
