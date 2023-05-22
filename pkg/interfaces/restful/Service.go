package restful

import "contentservice/pkg/application/customErrors"

type Service[T interface{}] interface {
	GetById(id int64) (*T, customErrors.HTTPError)
	Create(T) (int64, customErrors.HTTPError)
}
