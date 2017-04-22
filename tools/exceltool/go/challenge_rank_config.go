package config
type Challenge_rank struct{
	Items []*Challenge_rank_Item
}

type Challenge_rank_Item struct{
	Id int
	Rank_max int
	Rank_min int
	Reward [][]int
}

var InstChallenge_rank *Challenge_rank

func init(){
	item := &Challenge_rank_Item{}
	item.Id = 1001
	item.Rank_max = 1
	item.Rank_min = 1
	item.Reward = [][]int{{1,1,600},{201,432003,30}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1002
	item.Rank_max = 2
	item.Rank_min = 2
	item.Reward = [][]int{{1,1,500},{201,432003,20}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1003
	item.Rank_max = 3
	item.Rank_min = 3
	item.Reward = [][]int{{1,1,400},{201,432003,20}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1004
	item.Rank_max = 4
	item.Rank_min = 5
	item.Reward = [][]int{{1,1,350},{201,432003,10}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1005
	item.Rank_max = 6
	item.Rank_min = 10
	item.Reward = [][]int{{1,1,300},{201,432003,10}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1006
	item.Rank_max = 11
	item.Rank_min = 20
	item.Reward = [][]int{{1,1,250},{201,432003,10}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1007
	item.Rank_max = 21
	item.Rank_min = 30
	item.Reward = [][]int{{1,1,200},{201,432003,10}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1008
	item.Rank_max = 31
	item.Rank_min = 50
	item.Reward = [][]int{{1,1,150},{201,432003,10}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1009
	item.Rank_max = 51
	item.Rank_min = 100
	item.Reward = [][]int{{1,1,100},{201,432003,5}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1010
	item.Rank_max = 101
	item.Rank_min = 200
	item.Reward = [][]int{{1,1,50},{201,432003,5}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1011
	item.Rank_max = 201
	item.Rank_min = 1000
	item.Reward = [][]int{{2,2,50000}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
	item = &Challenge_rank_Item{}
	item.Id = 1012
	item.Rank_max = 1001
	item.Rank_min = 10000
	item.Reward = [][]int{{2,2,250}}
	InstChallenge_rank.Items = append(InstChallenge_rank.Items,item)
}
