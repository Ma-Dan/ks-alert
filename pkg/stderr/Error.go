package stderr

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"google.golang.org/grpc/status"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Success = iota
	InvalidParam
	AssertError
	DBError
	RuntimeError
	GrpcError
	OtherError
)

type Error struct {
	Code  int32
	Text  string
	Where string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, where: %s", e.Code, e.Text, e.Where)
}

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
	case Error:
		txt = v.Text
		code = v.Code
		where = v.Where
	case GrpcStatusError:
		//s := v.GRPCStatus()
		//if s.Code() == codes.Unknown {
		//	code = -1
		//} else if s.Code() < 20 {
		//	code = 4
		//}
		code = GrpcError
		txt = v.Error()
	case error:
		txt = v.Error()
		code = OtherError
	case nil:
		txt = "success"
		code = Success
	}

	if len(extMsg) != 0 {
		txt = strings.Join(extMsg, " : ") + " : " + txt
	}

	return &pb.Error{Text: txt, Code: code, Where: where}
}

func Caller(calldepth int, short bool) string {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	} else if short {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}
