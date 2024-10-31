package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderFacade(t *testing.T) {
	// 实例化各个服务实现
	inventoryService := &InventoryServiceImpl{}
	paymentService := &PaymentServiceImpl{}
	shippingService := &ShippingServiceImpl{}

	// 使用门面处理订单
	orderFacade := NewOrderFacade(inventoryService, paymentService, shippingService)
	err := orderFacade.PlaceOrder("item123", 99.99)
	assert.NoError(t, err)
}
