package config
type Yuanjun_order struct{
	Items []*Yuanjun_order_Item
}

type Yuanjun_order_Item struct{
	Order int
	Name string
}

var InstYuanjun_order *Yuanjun_order

func init(){
	item := &Yuanjun_order_Item{}
	item.Order = 1
	item.Name = "attack"
	InstYuanjun_order.Items = append(InstYuanjun_order.Items,item)
	item = &Yuanjun_order_Item{}
	item.Order = 2
	item.Name = "defence"
	InstYuanjun_order.Items = append(InstYuanjun_order.Items,item)
	item = &Yuanjun_order_Item{}
	item.Order = 3
	item.Name = "life"
	InstYuanjun_order.Items = append(InstYuanjun_order.Items,item)
	item = &Yuanjun_order_Item{}
	item.Order = 4
	item.Name = "refresh"
	InstYuanjun_order.Items = append(InstYuanjun_order.Items,item)
}
