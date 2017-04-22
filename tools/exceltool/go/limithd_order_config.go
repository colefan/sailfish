package config
type Limithd_order struct{
	Items []*Limithd_order_Item
}

type Limithd_order_Item struct{
	Molude_id int
	Huodong_order int
}

var InstLimithd_order *Limithd_order

func init(){
	item := &Limithd_order_Item{}
	item.Molude_id = 10402
	item.Huodong_order = 1
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10207
	item.Huodong_order = 2
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10208
	item.Huodong_order = 3
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10209
	item.Huodong_order = 4
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10210
	item.Huodong_order = 5
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10204
	item.Huodong_order = 6
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10201
	item.Huodong_order = 7
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10202
	item.Huodong_order = 8
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10203
	item.Huodong_order = 9
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10205
	item.Huodong_order = 10
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
	item = &Limithd_order_Item{}
	item.Molude_id = 10206
	item.Huodong_order = 11
	InstLimithd_order.Items = append(InstLimithd_order.Items,item)
}
