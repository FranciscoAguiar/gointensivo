package usecase

import "github.com/FranciscoAguiar/gointensivo/internal/order/entity"

type GetTotalOutPutDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutPutDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutPutDTO{Total: total}, nil
}
