package problemDetailImpl

import problemDetail "contentservice/pkg/interfaces/problem-detail"

type ProblemDetailImpl struct {
	TitleProp      string `json:"title"`
	StatusProp     int    `json:"status"`
	InstanceProp   string `json:"instance"`
	DetailProp     string `json:"detail"`
	PropertiesProp any    `json:"properties"`
}

func (pd *ProblemDetailImpl) Title(title string) problemDetail.ProblemDetail {
	pd.TitleProp = title
	return pd
}

func (pd *ProblemDetailImpl) Status(status int) problemDetail.ProblemDetail {
	pd.StatusProp = status
	return pd
}

func (pd *ProblemDetailImpl) Instance(instance string) problemDetail.ProblemDetail {
	pd.InstanceProp = instance
	return pd
}

func (pd *ProblemDetailImpl) Detail(detail string) problemDetail.ProblemDetail {
	pd.DetailProp = detail
	return pd
}

func (pd *ProblemDetailImpl) Properties(props any) problemDetail.ProblemDetail {
	pd.PropertiesProp = props
	return pd
}

func New() problemDetail.ProblemDetail {
	return new(ProblemDetailImpl)
}
