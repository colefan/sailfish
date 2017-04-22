package config
type Tech_order struct{
	Items []*Tech_order_Item
}

type Tech_order_Item struct{
	Order int
	Name string
}

var InstTech_order *Tech_order

func init(){
	item := &Tech_order_Item{}
	item.Order = 1
	item.Name = "attack"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 2
	item.Name = "deffence"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 3
	item.Name = "life"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 4
	item.Name = "money"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 5
	item.Name = "bullet"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 6
	item.Name = "food"
	InstTech_order.Items = append(InstTech_order.Items,item)
	item = &Tech_order_Item{}
	item.Order = 7
	item.Name = "iron"
	InstTech_order.Items = append(InstTech_order.Items,item)
}
