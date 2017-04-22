package config
type Limithd_rank_arena struct{
	Items []*Limithd_rank_arena_Item
}

type Limithd_rank_arena_Item struct{
	Id int
	Rank_max int
	Rank_min int
	Days3_reward [][]int
	Days7_reward [][]int
}

var InstLimithd_rank_arena *Limithd_rank_arena

func init(){
	item := &Limithd_rank_arena_Item{}
	item.Id = 90001
	item.Rank_max = 1
	item.Rank_min = 1
	item.Days3_reward = [][]int{{1,1,200},{2,2,100000},{201,310027,50}}
	item.Days7_reward = [][]int{{1,1,200},{2,2,100000},{201,310027,50}}
	InstLimithd_rank_arena.Items = append(InstLimithd_rank_arena.Items,item)
	item = &Limithd_rank_arena_Item{}
	item.Id = 90002
	item.Rank_max = 2
	item.Rank_min = 5
	item.Days3_reward = [][]int{{1,1,150},{2,2,80000},{201,310027,30}}
	item.Days7_reward = [][]int{{1,1,150},{2,2,80000},{201,310027,30}}
	InstLimithd_rank_arena.Items = append(InstLimithd_rank_arena.Items,item)
	item = &Limithd_rank_arena_Item{}
	item.Id = 90003
	item.Rank_max = 6
	item.Rank_min = 20
	item.Days3_reward = [][]int{{1,1,100},{2,2,50000},{201,310027,20}}
	item.Days7_reward = [][]int{{1,1,100},{2,2,50000},{201,310027,20}}
	InstLimithd_rank_arena.Items = append(InstLimithd_rank_arena.Items,item)
	item = &Limithd_rank_arena_Item{}
	item.Id = 90004
	item.Rank_max = 21
	item.Rank_min = 50
	item.Days3_reward = [][]int{{1,1,50},{2,2,30000},{201,310027,10}}
	item.Days7_reward = [][]int{{1,1,50},{2,2,30000},{201,310027,10}}
	InstLimithd_rank_arena.Items = append(InstLimithd_rank_arena.Items,item)
}
