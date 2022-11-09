package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyId_WhenCreateNewOrder_thenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
	// _, err := NewOrder("", 10, 1)
	// if err == nil {
	// 	t.Error("Expected an error, but did not get one")
	// }
}

func TestGivenEmptyPrice_WhenCreateNewOrder_thenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")

}

func TestGivenEmptyTax_WhenCreateNewOrder_thenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")

}

func TestGivenValidParams_WhenCreateNewOrder_thenShouldReceiveCreateOrder(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.0}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())

}

func TestGivenValidParams_WhenCallNewOrderFunc_thenShouldReceiveCreateOrder(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)

}
func TestGivenPriceAndTax_WhenCalculatePrice_ShouldReceiveFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculatePrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
