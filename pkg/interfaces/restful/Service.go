package restful

import (
	"contentservice/pkg/application/core/customErrors/interfaces"
)

type Service[T any] interface {
	GetById(id int64) (*T, interfaces.HTTPError)
	Create(T) (int64, interfaces.HTTPError)
}
