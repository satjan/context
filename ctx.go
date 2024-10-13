package context

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Parse(ctx *fiber.Ctx, req interface{}) (error, error) {
	if ctx.Method() == fiber.MethodGet || ctx.Method() == fiber.MethodDelete {
		return ParseQuery(ctx, req)
	} else {
		return ParseBody(ctx, req)
	}
}

func ParseQuery(ctx *fiber.Ctx, req interface{}) (error, error) {
	if err := ctx.QueryParser(req); err != nil {
		return JSONr(BadRequestErr, BodyParserError, ctx, nil), err
	}
	return ValidateRequest(ctx, req)
}

func ParseBody(ctx *fiber.Ctx, req interface{}) (error, error) {
	if len(ctx.Request().Body()) == 0 {
		return nil, nil
	}

	if err := ctx.BodyParser(req); err != nil {
		return JSONr(BadRequestErr, BodyParserError, ctx, nil), err
	}

	return ValidateRequest(ctx, req)
}

func ValidateRequest(ctx *fiber.Ctx, req interface{}) (error, error) {
	if err := Instance.Validate.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var errors []string
		for _, el := range errs.Translate(Instance.Trans) {
			errors = append(errors, el)
		}
		return JSONr(errors, ValidationError, ctx, nil), err
	}

	return nil, nil
}
