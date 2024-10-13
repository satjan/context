package context

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Parse parses the request based on the HTTP method.
func Parse(ctx *fiber.Ctx, req interface{}) error {
	if ctx.Method() == fiber.MethodGet || ctx.Method() == fiber.MethodDelete {
		return ParseQuery(ctx, req)
	}
	return ParseBody(ctx, req)
}

// ParseQuery parses query parameters into the given request struct.
func ParseQuery(ctx *fiber.Ctx, req interface{}) error {
	if err := ctx.QueryParser(req); err != nil {
		return JSONr(BadRequestErr, BodyParserError, ctx, nil)
	}
	return ValidateRequest(ctx, req)
}

// ParseBody parses the request body into the given request struct.
func ParseBody(ctx *fiber.Ctx, req interface{}) error {
	if len(ctx.Request().Body()) == 0 {
		return nil
	}

	if err := ctx.BodyParser(req); err != nil {
		return JSONr(BadRequestErr, BodyParserError, ctx, nil)
	}

	return ValidateRequest(ctx, req)
}

// ValidateRequest validates the given request struct.
func ValidateRequest(ctx *fiber.Ctx, req interface{}) error {
	if err := Instance.Validate.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var errors []string
		for _, el := range errs.Translate(Instance.Trans) {
			errors = append(errors, el)
		}
		return JSONr(errors, ValidationError, ctx, nil)
	}
	return nil
}
