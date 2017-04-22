package config
type Star struct{
	Items []*Star_Item
}

type Star_Item struct{
	Star int
	Star_attribute map[string]int
}

var InstStar *Star

func init(){
	item := &Star_Item{}
	item.Star = 1
	item.Star_attribute = map[string]int{"4":0,"5":0,"6":0}
	InstStar.Items = append(InstStar.Items,item)
	item = &Star_Item{}
	item.Star = 2
	item.Star_attribute = map[string]int{"4":40,"5":16,"6":144}
	InstStar.Items = append(InstStar.Items,item)
	item = &Star_Item{}
	item.Star = 3
	item.Star_attribute = map[string]int{"4":80,"5":32,"6":288}
	InstStar.Items = append(InstStar.Items,item)
	item = &Star_Item{}
	item.Star = 4
	item.Star_attribute = map[string]int{"4":120,"5":48,"6":432}
	InstStar.Items = append(InstStar.Items,item)
	item = &Star_Item{}
	item.Star = 5
	item.Star_attribute = map[string]int{"4":160,"5":64,"6":576}
	InstStar.Items = append(InstStar.Items,item)
}
