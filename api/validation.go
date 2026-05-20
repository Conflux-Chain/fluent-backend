package api

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var validators = map[string]validator.Func{
	"hex": regexValidator(regexp.MustCompile(`^0x(?:[0-9a-fA-F]{2})*$`)),
}

func regexValidator(r *regexp.Regexp) validator.Func {
	return func(fl validator.FieldLevel) bool {
		field, ok := fl.Field().Interface().(string)
		return ok && r.MatchString(field)
	}
}

func init() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		logrus.Fatal("Failed to register validators due to unexpected validator type")
	}

	for tag, valFunc := range validators {
		if err := v.RegisterValidation(tag, valFunc); err != nil {
			logrus.WithError(err).WithField("tag", tag).Fatal("Failed to register validator")
		}
	}
}
