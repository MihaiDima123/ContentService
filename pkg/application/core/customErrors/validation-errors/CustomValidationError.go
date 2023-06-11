package validationErrors

type CustomValidationError struct {
	error
	ErrorType int8
	Source    string
}

func (cve *CustomValidationError) GetErrorType() int8 {
	return cve.ErrorType
}

func (cve *CustomValidationError) GetSource() string {
	return cve.Source
}
