package config
type Limithd_rank_kill struct{
	Items []*Limithd_rank_kill_Item
}

type Limithd_rank_kill_Item struct{
	Id int
	Rank_max int
	Rank_min int
	Days3_reward [][]int
	Days7_reward [][]int
}

var InstLimithd_rank_kill *Limithd_rank_kill

func init(){
	item := &Limithd_rank_kill_Item{}
	item.Id = 100001
	item.Rank_max = 1
	item.Rank_min = 1
	item.Days3_reward = [][]int{{1,1,500},{3,3,100000},{201,310027,50}}
	item.Days7_reward = [][]int{{1,1,500},{3,3,100000},{201,310027,50}}
	InstLimithd_rank_kill.Items = append(InstLimithd_rank_kill.Items,item)
	item = &Limithd_rank_kill_Item{}
	item.Id = 100002
	item.Rank_max = 2
	item.Rank_min = 5
	item.Days3_reward = [][]int{{1,1,250},{3,3,80000},{201,310027,30}}
	item.Days7_reward = [][]int{{1,1,250},{3,3,80000},{201,310027,30}}
	InstLimithd_rank_kill.Items = append(InstLimithd_rank_kill.Items,item)
	item = &Limithd_rank_kill_Item{}
	item.Id = 100003
	item.Rank_max = 6
	item.Rank_min = 20
	item.Days3_reward = [][]int{{1,1,100},{3,3,50000},{201,310027,20}}
	item.Days7_reward = [][]int{{1,1,100},{3,3,50000},{201,310027,20}}
	InstLimithd_rank_kill.Items = append(InstLimithd_rank_kill.Items,item)
	item = &Limithd_rank_kill_Item{}
	item.Id = 100004
	item.Rank_max = 21
	item.Rank_min = 50
	item.Days3_reward = [][]int{{18,18,100000},{3,3,30000},{201,310027,10}}
	item.Days7_reward = [][]int{{18,18,100000},{3,3,30000},{201,310027,10}}
	InstLimithd_rank_kill.Items = append(InstLimithd_rank_kill.Items,item)
}
