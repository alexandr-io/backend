package grpc

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var httpToGRPC = map[int]codes.Code{
	fiber.StatusOK:                  codes.OK,
	fiber.StatusBadRequest:          codes.InvalidArgument,
	fiber.StatusUnauthorized:        codes.PermissionDenied,
	fiber.StatusNotFound:            codes.NotFound,
	fiber.StatusInternalServerError: codes.Internal,
	fiber.StatusGatewayTimeout:      codes.DeadlineExceeded,
}

// FiberErrorToGRPC transform a fiber error to a gRPC error
func FiberErrorToGRPC(err error) error {
	e, ok := err.(*fiber.Error)
	if !ok {
		return nil
	}
	err = status.Error(httpToGRPC[e.Code], e.Message)
	return err
}

// ErrorToFiber transform a gRPC error to a gRPC fiber
func ErrorToFiber(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return nil
	}
	for key, elem := range httpToGRPC {
		if elem == st.Code() {
			return fiber.NewError(key, st.Message())
		}
	}
	return nil
}
