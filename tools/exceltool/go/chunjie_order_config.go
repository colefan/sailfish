package config
type Chunjie_order struct{
	Items []*Chunjie_order_Item
}

type Chunjie_order_Item struct{
	Molude_id int
	Huodong_order int
	Startday int
	Endday int
}

var InstChunjie_order *Chunjie_order

func init(){
	item := &Chunjie_order_Item{}
	item.Molude_id = 10501
	item.Huodong_order = 1
	item.Startday = 20170127
	item.Endday = 20170202
	InstChunjie_order.Items = append(InstChunjie_order.Items,item)
	item = &Chunjie_order_Item{}
	item.Molude_id = 10502
	item.Huodong_order = 2
	item.Startday = 20170127
	item.Endday = 20170202
	InstChunjie_order.Items = append(InstChunjie_order.Items,item)
	item = &Chunjie_order_Item{}
	item.Molude_id = 10503
	item.Huodong_order = 3
	item.Startday = 20170127
	item.Endday = 20170202
	InstChunjie_order.Items = append(InstChunjie_order.Items,item)
	item = &Chunjie_order_Item{}
	item.Molude_id = 10504
	item.Huodong_order = 4
	item.Startday = 20170127
	item.Endday = 20170202
	InstChunjie_order.Items = append(InstChunjie_order.Items,item)
	item = &Chunjie_order_Item{}
	item.Molude_id = 10505
	item.Huodong_order = 5
	item.Startday = 20170127
	item.Endday = 20170202
	InstChunjie_order.Items = append(InstChunjie_order.Items,item)
}
