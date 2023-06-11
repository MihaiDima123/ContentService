package problemDetailImpl

import (
	"contentservice/pkg/interfaces/customerrors"
	problemDetail "contentservice/pkg/interfaces/problem-detail"
	"net/http"
)

// ProblemDetailImpl Problem detail implementation to return my custom errors
type ProblemDetailImpl struct {
	TitleProp      string                     `json:"title"`
	StatusProp     int                        `json:"status"`
	InstanceProp   string                     `json:"instance"`
	DetailProp     string                     `json:"detail"`
	PropertiesProp []customerrors.CustomError `json:"properties"`
}

func (pd *ProblemDetailImpl) Title(title string) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd.TitleProp = title
	return pd
}

func (pd *ProblemDetailImpl) Status(status int) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd.StatusProp = status
	return pd
}

func (pd *ProblemDetailImpl) Instance(instance string) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd.InstanceProp = instance
	return pd
}

func (pd *ProblemDetailImpl) Detail(detail string) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd.DetailProp = detail
	return pd
}

func (pd *ProblemDetailImpl) Properties(props []customerrors.CustomError) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd.PropertiesProp = props
	return pd
}

func New() problemDetail.ProblemDetail[customerrors.CustomError] {
	return new(ProblemDetailImpl)
}

func NewOfHttpError(error customerrors.HTTPError) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd := new(ProblemDetailImpl)
	pd.Status(error.GetStatus())
	pd.Detail(error.Error())

	return pd
}

func NewOfValidationError(err customerrors.ValidationError) problemDetail.ProblemDetail[customerrors.CustomError] {
	pd := new(ProblemDetailImpl)
	pd.Status(http.StatusBadRequest)
	pd.Detail(err.Error())
	pd.Title(err.GetTitle())
	pd.Instance(err.GetSource())
	return pd
}
