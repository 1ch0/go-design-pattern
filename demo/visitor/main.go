package main

import "fmt"

type OrderService struct {
	orders map[int]*Order
}

func NewOrderService() IOrderService {
	return &OrderService{orders: make(map[int]*Order, 0)}
}

type IOrderService interface {
	Save(order *Order)
	Accept(visitor IOrderVisitor)
}

// Order 订单实体类，实现IOrderService 接口
type Order struct {
	ID       int
	Customer string
	City     string
	Product  string
	Quantity int
}

func (o *OrderService) Save(order *Order) {
	o.orders[order.ID] = order
}

func (o *OrderService) Accept(visitor IOrderVisitor) {
	for _, v := range o.orders {
		visitor.Visit(v)
	}
}

func NewOrder(id int, customer string, city string, product string, quantity int) *Order {
	return &Order{
		id, customer, city, product, quantity,
	}
}

type IOrderVisitor interface {
	Visit(order *Order)
	Report()
}

type CityVisitor struct {
	cities map[string]int
}

func (c *CityVisitor) Visit(order *Order) {
	item, ok := c.cities[order.City]
	if ok {
		c.cities[order.City] = item + order.Quantity
	} else {
		c.cities[order.City] = order.Quantity
	}
}

func (c *CityVisitor) Report() {
	for k, v := range c.cities {
		fmt.Printf("city: %s, value: %d\n", k, v)
	}
}

func NewCityVisitor() IOrderVisitor {
	return &CityVisitor{cities: make(map[string]int, 0)}
}

type ProductVisitor struct {
	products map[string]int
}

func (p *ProductVisitor) Visit(order *Order) {
	item, ok := p.products[order.Product]
	if ok {
		p.products[order.Product] = item + order.Quantity
	} else {
		p.products[order.Product] = order.Quantity
	}
}

func (p *ProductVisitor) Report() {
	for k, v := range p.products {
		fmt.Printf("product: %s, value: %d\n", k, v)
	}
}

func NewProductVisitor() IOrderVisitor {
	return &ProductVisitor{products: make(map[string]int, 0)}
}

func main() {
	orderService := NewOrderService()
	orderService.Save(NewOrder(1, "张三", "广州", "电视", 10))
	orderService.Save(NewOrder(2, "李四", "深圳", "冰箱", 20))
	orderService.Save(NewOrder(3, "王五", "东莞", "空调", 30))
	orderService.Save(NewOrder(4, "张三三", "广州", "空调", 10))
	orderService.Save(NewOrder(5, "李四四", "深圳", "电视", 20))
	orderService.Save(NewOrder(6, "王五五", "东莞", "冰箱", 30))

	cv := NewCityVisitor()
	orderService.Accept(cv)
	cv.Report()

	pv := NewProductVisitor()
	orderService.Accept(pv)
	pv.Report()
}
