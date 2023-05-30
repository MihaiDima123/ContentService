package restful

import (
	"contentservice/pkg/interfaces/customerrors"
)

type Service[T any] interface {
	GetById(id int64) (*T, customerrors.HTTPError)
	Create(T) (int64, customerrors.HTTPError)
}
