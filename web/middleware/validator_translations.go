package middleware

import (
	"com.gientech/selection/pkg/validate"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		var trans ut.Translator
		if locale == "" {
			trans = validate.Trans
		} else {
			trans = validate.GetTrans(locale)
		}
		c.Set("trans", trans)
		c.Next()
	}
}
