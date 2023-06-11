package restful

import (
	"contentservice/pkg/interfaces/customerrors"
)

type Service[T any, dtoT any] interface {
	GetById(id int64) (*dtoT, customerrors.HTTPError)
	Create(*dtoT) (*int64, customerrors.HTTPError)
}
