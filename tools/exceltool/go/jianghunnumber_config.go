package config
type Jianghunnumber struct{
	Items []*Jianghunnumber_Item
}

type Jianghunnumber_Item struct{
	Number int
	Probability int
}

var InstJianghunnumber *Jianghunnumber

func init(){
	item := &Jianghunnumber_Item{}
	item.Number = 1
	item.Probability = 0
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 2
	item.Probability = 0
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 3
	item.Probability = 3000
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 4
	item.Probability = 3000
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 5
	item.Probability = 3000
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 6
	item.Probability = 1500
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 7
	item.Probability = 1500
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 8
	item.Probability = 1500
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 9
	item.Probability = 500
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
	item = &Jianghunnumber_Item{}
	item.Number = 10
	item.Probability = 500
	InstJianghunnumber.Items = append(InstJianghunnumber.Items,item)
}
