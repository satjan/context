package context

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const (
	Ok               string = "0000"
	ValidationError         = "4001"
	BodyParserError         = "4004"
	UnavailableError        = "0014"
)

type ResponseType struct {
	Status bool        `json:"status"`
	Code   string      `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func OK() ResponseType {
	response := ResponseType{}
	response.Status = true
	response.Code = Ok
	return response
}

func Err(err interface{}, code string, data interface{}) ResponseType {
	response := ResponseType{}
	response.Status = false
	response.Msg = err
	response.Code = code
	response.Data = data
	return response
}

func JSONr(err interface{}, code string, c *fiber.Ctx, data interface{}) error {
	e := err
	switch v := err.(type) {
	case string:
		e = []string{v}
	}
	response := Err(e, code, data)
	return c.JSON(response)
}

func JSON(c *fiber.Ctx, data interface{}, msg interface{}) error {
	response := OK()
	response.Data = data
	if msg == "" {
		msg = "Success"
	}
	response.Msg = []string{fmt.Sprintf("%v", msg)}
	return c.JSON(response)
}
