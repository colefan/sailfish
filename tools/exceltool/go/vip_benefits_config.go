package config
type Vip_benefits struct{
	Items []*Vip_benefits_Item
}

type Vip_benefits_Item struct{
	Id int
	Name string
	Need_viplv int
	Time_limit int
	Reward [][]int
}

var InstVip_benefits *Vip_benefits

func init(){
	item := &Vip_benefits_Item{}
	item.Id = 1001
	item.Name = "V3六选一福利"
	item.Need_viplv = 3
	item.Time_limit = 7
	item.Reward = [][]int{{201,310003,100,2588},{201,310006,100,2588},{201,310037,100,2588},{201,310009,100,2588},{201,310018,100,2588},{201,310045,100,2588}}
	InstVip_benefits.Items = append(InstVip_benefits.Items,item)
}
