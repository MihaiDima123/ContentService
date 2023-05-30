package restful

import (
	"contentservice/pkg/interfaces/errors"
)

type Service[T any] interface {
	GetById(id int64) (*T, errors.HTTPError)
	Create(T) (int64, errors.HTTPError)
}
