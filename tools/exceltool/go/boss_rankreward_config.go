package config
type Boss_rankreward struct{
	Items []*Boss_rankreward_Item
}

type Boss_rankreward_Item struct{
	Id int
	Rank_max int
	Rank_min int
	Reward [][]int
}

var InstBoss_rankreward *Boss_rankreward

func init(){
	item := &Boss_rankreward_Item{}
	item.Id = 1001
	item.Rank_max = 1
	item.Rank_min = 1
	item.Reward = [][]int{{4,17,40000},{17,17,45},{21,21,48}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1002
	item.Rank_max = 2
	item.Rank_min = 2
	item.Reward = [][]int{{4,17,35000},{17,17,39},{21,21,42}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1003
	item.Rank_max = 3
	item.Rank_min = 5
	item.Reward = [][]int{{4,17,30000},{17,17,33},{21,21,36}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1004
	item.Rank_max = 6
	item.Rank_min = 10
	item.Reward = [][]int{{4,17,25000},{17,17,28},{21,21,30}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1005
	item.Rank_max = 11
	item.Rank_min = 15
	item.Reward = [][]int{{4,17,20000},{17,17,24},{21,21,25}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1006
	item.Rank_max = 16
	item.Rank_min = 20
	item.Reward = [][]int{{4,17,15000},{17,17,20},{21,21,21}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1007
	item.Rank_max = 21
	item.Rank_min = 30
	item.Reward = [][]int{{4,17,10000},{17,17,16},{21,21,18}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1008
	item.Rank_max = 31
	item.Rank_min = 50
	item.Reward = [][]int{{4,17,8000},{17,17,12},{21,21,15}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1009
	item.Rank_max = 51
	item.Rank_min = 75
	item.Reward = [][]int{{4,17,6000},{17,17,10},{21,21,12}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1010
	item.Rank_max = 76
	item.Rank_min = 100
	item.Reward = [][]int{{4,17,4000},{17,17,8},{21,21,10}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1011
	item.Rank_max = 101
	item.Rank_min = 200
	item.Reward = [][]int{{4,17,3000},{17,17,6},{21,21,8}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
	item = &Boss_rankreward_Item{}
	item.Id = 1012
	item.Rank_max = 201
	item.Rank_min = 1000
	item.Reward = [][]int{{4,17,2000},{17,17,4},{21,21,5}}
	InstBoss_rankreward.Items = append(InstBoss_rankreward.Items,item)
}
