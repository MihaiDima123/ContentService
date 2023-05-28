package customErrors

import (
	"errors"
	"strings"
)

const ErrorMessageSeparator = "|"

func getErrorFromString(defaultMessage string, params ...string) error {
	if len(params) == 0 {
		return errors.New(defaultMessage)
	}

	return errors.New(strings.Join(params, ErrorMessageSeparator))
}
