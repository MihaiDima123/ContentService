package interfaces

import (
	"contentservice/pkg/interfaces/restful"
)

type ContentService interface {
	restful.Service
	Configure(repository ContentRepository)
}
