package grpc

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var httpToGRPC = map[int]codes.Code{
	200: codes.OK,
	400: codes.InvalidArgument,
	401: codes.PermissionDenied,
	404: codes.NotFound,
	500: codes.Internal,
	504: codes.DeadlineExceeded,
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
