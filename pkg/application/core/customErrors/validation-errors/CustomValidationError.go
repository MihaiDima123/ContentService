package validationErrors

type CustomValidationError struct {
	error
	ErrorType int8
	Source    string
	Title     string
}

func (cve *CustomValidationError) GetErrorType() int8 {
	return cve.ErrorType
}

func (cve *CustomValidationError) GetSource() string {
	return cve.Source
}

func (cve *CustomValidationError) GetTitle() string {
	return cve.Title
}
