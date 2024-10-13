package context

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ResponseOk(ctx *fiber.Ctx, data interface{}, err error) error {
	return Response(ctx, data, err, "")
}

func Response(ctx *fiber.Ctx, data interface{}, err error, msg string) error {
	if err == nil {
		return JSON(ctx, data, msg)
	}

	if status.Code(err) != codes.Unknown {
		return JSONr("Service unavailable", UnavailableError, ctx, data)
	}

	e, _ := status.FromError(err)

	return JSONr(e.Message(), ValidationError, ctx, data)
}
