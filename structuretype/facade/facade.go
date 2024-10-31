/*
假设我们有一个复杂的订单系统，其中包含了库存检查、支付处理、物流安排等多个服务。我们可以通过门面模式为这些服务提供一个简化的接口，
使得调用者无需关注这些服务的内部实现细节，只需调用门面接口来完成订单处理
*/
package facade

import "fmt"

// 定义库存接口
type InventoryService interface {
	CheckStock(itemID string) bool
}

// 定义支付接口
type PaymentService interface {
	ProcessPayment(amount float64) bool
}

// 定义物流接口
type ShippingService interface {
	ArrangeShipping(itemID string) bool
}

// 具体实现库存服务
type InventoryServiceImpl struct{}

func (s *InventoryServiceImpl) CheckStock(itemID string) bool {
	fmt.Println("检查库存...")
	// 假设库存检查通过
	return true
}

// 具体实现支付服务
type PaymentServiceImpl struct{}

func (s *PaymentServiceImpl) ProcessPayment(amount float64) bool {
	fmt.Println("处理支付...")
	// 假设支付成功
	return true
}

// 具体实现物流服务
type ShippingServiceImpl struct{}

func (s *ShippingServiceImpl) ArrangeShipping(itemID string) bool {
	fmt.Println("安排物流...")
	// 假设物流安排成功
	return true
}

// 定义订单门面接口
type OrderFacade interface {
	PlaceOrder(itemID string, amount float64) error
}

// 实现订单门面
type OrderFacadeImpl struct {
	inventoryService InventoryService
	paymentService   PaymentService
	shippingService  ShippingService
}

// NewOrderFacade 创建一个新的 OrderFacade 实例
func NewOrderFacade(inventory InventoryService, payment PaymentService, shipping ShippingService) OrderFacade {
	return &OrderFacadeImpl{
		inventoryService: inventory,
		paymentService:   payment,
		shippingService:  shipping,
	}
}

// PlaceOrder 通过门面处理订单
func (f *OrderFacadeImpl) PlaceOrder(itemID string, amount float64) error {
	if !f.inventoryService.CheckStock(itemID) {
		return fmt.Errorf("库存不足")
	}

	if !f.paymentService.ProcessPayment(amount) {
		return fmt.Errorf("支付失败")
	}

	if !f.shippingService.ArrangeShipping(itemID) {
		return fmt.Errorf("物流安排失败")
	}

	fmt.Println("订单处理成功！")
	return nil
}
