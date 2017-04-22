package config
type Nation struct{
	Items []*Nation_Item
}

type Nation_Item struct{
	Id int
	Name string
}

var InstNation *Nation

func init(){
	item := &Nation_Item{}
	item.Id = 4
	item.Name = "text_weight00"
	InstNation.Items = append(InstNation.Items,item)
	item = &Nation_Item{}
	item.Id = 1
	item.Name = "text_weight01"
	InstNation.Items = append(InstNation.Items,item)
	item = &Nation_Item{}
	item.Id = 2
	item.Name = "text_weight02"
	InstNation.Items = append(InstNation.Items,item)
	item = &Nation_Item{}
	item.Id = 3
	item.Name = "text_weight03"
	InstNation.Items = append(InstNation.Items,item)
}
