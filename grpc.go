package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCStatus(err error) error {
	if err == nil {
		return err
	}
	code := codes.Internal
	switch {
	case Match(err, ErrNotFound):
		code = codes.NotFound
	case Match(err, ErrPermissionDenied):
		code = codes.PermissionDenied
	}
	return status.Error(code, err.Error())
}
