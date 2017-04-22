package config
type Gm_item struct{
	Items []*Gm_item_Item
}

type Gm_item_Item struct{
	Id int
	Item string
}

var InstGm_item *Gm_item

func init(){
	item := &Gm_item_Item{}
	item.Id = 1
	item.Item = "军械"
	InstGm_item.Items = append(InstGm_item.Items,item)
	item = &Gm_item_Item{}
	item.Id = 2
	item.Item = "宝物"
	InstGm_item.Items = append(InstGm_item.Items,item)
	item = &Gm_item_Item{}
	item.Id = 3
	item.Item = "将魂"
	InstGm_item.Items = append(InstGm_item.Items,item)
	item = &Gm_item_Item{}
	item.Id = 4
	item.Item = "其他消耗品"
	InstGm_item.Items = append(InstGm_item.Items,item)
	item = &Gm_item_Item{}
	item.Id = 5
	item.Item = "英雄"
	InstGm_item.Items = append(InstGm_item.Items,item)
}
