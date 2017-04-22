package config
type Areas struct{
	Items []*Areas_Item
}

type Areas_Item struct{
	Area_id int
}

var InstAreas *Areas

func init(){
	item := &Areas_Item{}
	item.Area_id = 1
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 2
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 3
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 4
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 5
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 6
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 7
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 8
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 9
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 10
	InstAreas.Items = append(InstAreas.Items,item)
	item = &Areas_Item{}
	item.Area_id = 11
	InstAreas.Items = append(InstAreas.Items,item)
}
