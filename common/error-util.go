package common

import (
	"fmt"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CustomError struct {
	Errors map[string]interface{} `json:"errors"`
}

func BindingValidationError(err error) CustomError {
	res := CustomError{
		Errors: make(map[string]interface{}),
	}

	for _, v := range err.(validator.ValidationErrors) {
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf(
				"{%v: %v}", v.Tag, v.Param,
			)
		} else {
			res.Errors[v.Field] = fmt.Sprintf(
				"{key: %v}", v.Tag,
			)
		}
	}

	return res
}

func NewError(key string, err error) CustomError {
	res := CustomError{
		Errors: make(map[string]interface{}),
	}
	res.Errors[key] = err.Error()
	return res
}

func CustomShouldBind(c *gin.Context, obj interface{}) error {
	method := c.Request.Method
	contentType := c.ContentType()
	b := binding.Default(method, contentType)
	return c.ShouldBindWith(obj, b)
}
