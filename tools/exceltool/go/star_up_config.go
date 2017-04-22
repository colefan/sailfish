package config
type Star_up struct{
	Items []*Star_up_Item
}

type Star_up_Item struct{
	Star int
	Cost int
}

var InstStar_up *Star_up

func init(){
	item := &Star_up_Item{}
	item.Star = 1
	item.Cost = 20000
	InstStar_up.Items = append(InstStar_up.Items,item)
	item = &Star_up_Item{}
	item.Star = 2
	item.Cost = 100000
	InstStar_up.Items = append(InstStar_up.Items,item)
	item = &Star_up_Item{}
	item.Star = 3
	item.Cost = 300000
	InstStar_up.Items = append(InstStar_up.Items,item)
	item = &Star_up_Item{}
	item.Star = 4
	item.Cost = 500000
	InstStar_up.Items = append(InstStar_up.Items,item)
	item = &Star_up_Item{}
	item.Star = 5
	item.Cost = 1000000
	InstStar_up.Items = append(InstStar_up.Items,item)
}
