package config
type Limithd_shuabenkuangren struct{
	Items []*Limithd_shuabenkuangren_Item
}

type Limithd_shuabenkuangren_Item struct{
	Id int
	Need_star int
	Rewards [][]int
}

var InstLimithd_shuabenkuangren *Limithd_shuabenkuangren

func init(){
	item := &Limithd_shuabenkuangren_Item{}
	item.Id = 30001
	item.Need_star = 36
	item.Rewards = [][]int{{3,3,10000},{201,405001,2},{201,431001,1}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
	item = &Limithd_shuabenkuangren_Item{}
	item.Id = 30002
	item.Need_star = 72
	item.Rewards = [][]int{{3,3,15000},{201,405001,2},{201,431001,3}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
	item = &Limithd_shuabenkuangren_Item{}
	item.Id = 30003
	item.Need_star = 108
	item.Rewards = [][]int{{3,3,20000},{201,405001,2},{201,431001,4}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
	item = &Limithd_shuabenkuangren_Item{}
	item.Id = 30004
	item.Need_star = 144
	item.Rewards = [][]int{{3,3,30000},{201,405002,2},{201,431001,5}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
	item = &Limithd_shuabenkuangren_Item{}
	item.Id = 30005
	item.Need_star = 180
	item.Rewards = [][]int{{3,3,50000},{201,405002,2},{201,431001,8}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
	item = &Limithd_shuabenkuangren_Item{}
	item.Id = 30006
	item.Need_star = 216
	item.Rewards = [][]int{{3,3,80000},{201,405002,2},{201,431001,10}}
	InstLimithd_shuabenkuangren.Items = append(InstLimithd_shuabenkuangren.Items,item)
}
