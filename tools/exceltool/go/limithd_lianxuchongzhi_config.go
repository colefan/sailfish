package config
type Limithd_lianxuchongzhi struct{
	Items []*Limithd_lianxuchongzhi_Item
}

type Limithd_lianxuchongzhi_Item struct{
	Id int
	Type int
	Rewards [][]int
}

var InstLimithd_lianxuchongzhi *Limithd_lianxuchongzhi

func init(){
	item := &Limithd_lianxuchongzhi_Item{}
	item.Id = 70001
	item.Type = 1
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70002
	item.Type = 2
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70003
	item.Type = 3
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70004
	item.Type = 4
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70005
	item.Type = 5
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70006
	item.Type = 6
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70007
	item.Type = 7
	item.Rewards = [][]int{{2,2,30000},{3,3,30000},{1,1,60}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
	item = &Limithd_lianxuchongzhi_Item{}
	item.Id = 70008
	item.Type = 999
	item.Rewards = [][]int{{1,1,100},{202,10029,1}}
	InstLimithd_lianxuchongzhi.Items = append(InstLimithd_lianxuchongzhi.Items,item)
}
