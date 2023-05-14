package interfaces

import "contentservice/pkg/interfaces/restful"

type ContentRepository interface {
	restful.Repository
	Test()
}
