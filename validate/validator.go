package validate

import (
	"github.com/labstack/echo"
	"net/http"
)

type Context struct {
	echo.Context
}

func (c *Context) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request:"+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validation Error:"+err.Error())
	}

	return nil
}
