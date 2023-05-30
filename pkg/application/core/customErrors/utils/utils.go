package utils

import (
	"errors"
	"strings"
)

const ErrorMessageSeparator = "|"

func GetErrorFromString(defaultMessage string, params ...string) error {
	if len(params) == 0 {
		return errors.New(defaultMessage)
	}

	return errors.New(strings.Join(params, ErrorMessageSeparator))
}
