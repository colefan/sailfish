package config
type Limithd_rank_power struct{
	Items []*Limithd_rank_power_Item
}

type Limithd_rank_power_Item struct{
	Id int
	Rank_max int
	Rank_min int
	Days3_reward [][]int
	Days7_reward [][]int
}

var InstLimithd_rank_power *Limithd_rank_power

func init(){
	item := &Limithd_rank_power_Item{}
	item.Id = 80001
	item.Rank_max = 1
	item.Rank_min = 1
	item.Days3_reward = [][]int{{1,1,800},{201,435001,50},{201,310013,50}}
	item.Days7_reward = [][]int{{1,1,800},{201,435001,50},{201,310013,50}}
	InstLimithd_rank_power.Items = append(InstLimithd_rank_power.Items,item)
	item = &Limithd_rank_power_Item{}
	item.Id = 80002
	item.Rank_max = 2
	item.Rank_min = 5
	item.Days3_reward = [][]int{{1,1,400},{201,435001,30},{201,310013,30}}
	item.Days7_reward = [][]int{{1,1,400},{201,435001,30},{201,310013,30}}
	InstLimithd_rank_power.Items = append(InstLimithd_rank_power.Items,item)
	item = &Limithd_rank_power_Item{}
	item.Id = 80003
	item.Rank_max = 6
	item.Rank_min = 20
	item.Days3_reward = [][]int{{1,1,100},{201,435001,20},{201,310013,20}}
	item.Days7_reward = [][]int{{1,1,100},{201,435001,20},{201,310013,20}}
	InstLimithd_rank_power.Items = append(InstLimithd_rank_power.Items,item)
	item = &Limithd_rank_power_Item{}
	item.Id = 80004
	item.Rank_max = 21
	item.Rank_min = 50
	item.Days3_reward = [][]int{{2,2,100000},{201,435001,10},{201,310013,10}}
	item.Days7_reward = [][]int{{2,2,100000},{201,435001,10},{201,310013,10}}
	InstLimithd_rank_power.Items = append(InstLimithd_rank_power.Items,item)
}
