package products

type IProductCommand interface{}

type ProductCommand struct {
}

func NewProductCommand() *ProductCommand {
	return &ProductCommand{}
}

func (cmd ProductCommand) Create(payload *CreateProductCategory) {

}
