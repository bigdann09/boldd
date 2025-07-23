package attributes

type IAttributeCommand interface {
	Create(payload *CreateAttributeRequest) interface{}
}

type AttributeCommand struct {
}

func NewAttributeCommand() *AttributeCommand {
	return &AttributeCommand{}
}

func (cmd AttributeCommand) Create(payload *CreateAttributeRequest) interface{} {
	return nil
}
