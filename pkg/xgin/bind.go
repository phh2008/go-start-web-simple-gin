package xgin

import (
	"com.gientech/selection/pkg/validate"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

// ValidError 详细校验看 https://github.com/go-playground/validator
type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func ShouldBind(c *gin.Context, v interface{}) (bool, ValidErrors) {
	err := c.ShouldBind(v)
	if err != nil {
		return TranslateValidationError(c, err)
	}
	return true, nil
}

func ShouldBindJSON(c *gin.Context, v interface{}) (bool, ValidErrors) {
	err := c.ShouldBindJSON(v)
	if err != nil {
		return TranslateValidationError(c, err)
	}
	return true, nil
}

func ShouldBindQuery(c *gin.Context, v interface{}) (bool, ValidErrors) {
	err := c.ShouldBindQuery(v)
	if err != nil {
		return TranslateValidationError(c, err)
	}
	return true, nil
}

func TranslateValidationError(c *gin.Context, err error) (bool, ValidErrors) {
	v := c.Value("trans")
	var errs ValidErrors
	trans, ok := v.(ut.Translator)
	if !ok || trans == nil {
		trans = validate.Trans
	}
	validErrs, ok := err.(val.ValidationErrors)
	if !ok {
		errs = append(errs, &ValidError{
			Key:     "",
			Message: err.Error(),
		})
		return false, errs
	}
	for key, value := range validErrs.Translate(trans) {
		errs = append(errs, &ValidError{
			Key:     key,
			Message: value,
		})
	}
	return false, errs
}
