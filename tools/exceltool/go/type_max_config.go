package config
type Type_max struct{
	Items []*Type_max_Item
}

type Type_max_Item struct{
	Type int
	Tentime_max int
	Drinktime_max int
}

var InstType_max *Type_max

func init(){
	item := &Type_max_Item{}
	item.Type = 1
	item.Tentime_max = 4
	item.Drinktime_max = 2
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 2
	item.Tentime_max = 3
	item.Drinktime_max = 1
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 3
	item.Tentime_max = 1
	item.Drinktime_max = 0
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 4
	item.Tentime_max = 5
	item.Drinktime_max = 3
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 5
	item.Tentime_max = 3
	item.Drinktime_max = 2
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 6
	item.Tentime_max = 0
	item.Drinktime_max = 2
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 7
	item.Tentime_max = 0
	item.Drinktime_max = 0
	InstType_max.Items = append(InstType_max.Items,item)
	item = &Type_max_Item{}
	item.Type = 8
	item.Tentime_max = 0
	item.Drinktime_max = 0
	InstType_max.Items = append(InstType_max.Items,item)
}
