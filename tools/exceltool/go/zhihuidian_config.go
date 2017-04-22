package config
type Zhihuidian struct{
	Items []*Zhihuidian_Item
}

type Zhihuidian_Item struct{
	ID int
	Price int
}

var InstZhihuidian *Zhihuidian

func init(){
	item := &Zhihuidian_Item{}
	item.ID = 1
	item.Price = 6
	InstZhihuidian.Items = append(InstZhihuidian.Items,item)
	item = &Zhihuidian_Item{}
	item.ID = 2
	item.Price = 3
	InstZhihuidian.Items = append(InstZhihuidian.Items,item)
	item = &Zhihuidian_Item{}
	item.ID = 3
	item.Price = 20
	InstZhihuidian.Items = append(InstZhihuidian.Items,item)
}
