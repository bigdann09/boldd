package products

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
)

type IProductCommand interface {
	GenerateCombinations(payload *GenerateCombinationRequest) ([]*VariantCombinationResponse, interface{})
}

type ProductCommand struct {
}

func NewProductCommand() *ProductCommand {
	return &ProductCommand{}
}

func (cmd ProductCommand) Create(payload *CreateProductRequest) {

}

func (cmd ProductCommand) GenerateCombinations(payload *GenerateCombinationRequest) ([]*VariantCombinationResponse, interface{}) {
	groups := [][]string{}
	for _, attribute := range payload.Attributes {
		groups = append(groups, attribute.Values)
	}

	if len(groups) == 0 {
		return []*VariantCombinationResponse{}, dtos.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Provide attributes and values to generate combinations",
		}
	}

	var generator func([][]string, int) [][]string
	generator = func(groups [][]string, index int) [][]string {
		fmt.Println("index", index, "length", len(groups))
		if index == len(groups) {
			return [][]string{}
		}

		result := [][]string{}
		for _, value := range groups[index] {
			for _, combination := range generator(groups, index+1) {
				result = append(result, append([]string{value}, combination...))
			}
		}
		return result
	}

	combinations := generator(groups, 0)
	fmt.Println(combinations)
	var response []*VariantCombinationResponse
	for _, combination := range combinations {
		response = append(response, &VariantCombinationResponse{Combination: combination})
	}

	return response, nil
}
