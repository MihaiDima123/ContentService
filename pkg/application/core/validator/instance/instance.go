package validationInstance

import "contentservice/pkg/interfaces/validator"

type ValidationInstance struct {
	validator.ValidationInstance
	InstanceParam string
	MessageParam  string
}

func (vi *ValidationInstance) Instance() string {
	return vi.InstanceParam
}

func (vi *ValidationInstance) Message() string {
	return vi.MessageParam
}

func New(instance string, message string) validator.ValidationInstance {
	vi := new(ValidationInstance)
	vi.InstanceParam = instance
	vi.MessageParam = message
	return vi
}
