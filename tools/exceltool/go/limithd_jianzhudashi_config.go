package config
type Limithd_jianzhudashi struct{
	Items []*Limithd_jianzhudashi_Item
}

type Limithd_jianzhudashi_Item struct{
	Id int
	Building_id int
	Number int
	Building_lv int
	Rewards [][]int
}

var InstLimithd_jianzhudashi *Limithd_jianzhudashi

func init(){
	item := &Limithd_jianzhudashi_Item{}
	item.Id = 60001
	item.Building_id = 1
	item.Number = 1
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60002
	item.Building_id = 2
	item.Number = 1
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60003
	item.Building_id = 3
	item.Number = 1
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60004
	item.Building_id = 4
	item.Number = 2
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60005
	item.Building_id = 5
	item.Number = 2
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60006
	item.Building_id = 9
	item.Number = 1
	item.Building_lv = 3
	item.Rewards = [][]int{{2,2,2000},{201,405001,1}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60007
	item.Building_id = 1
	item.Number = 1
	item.Building_lv = 5
	item.Rewards = [][]int{{2,2,4000},{201,405001,2}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60008
	item.Building_id = 4
	item.Number = 2
	item.Building_lv = 5
	item.Rewards = [][]int{{2,2,4000},{201,405001,2}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60009
	item.Building_id = 5
	item.Number = 3
	item.Building_lv = 5
	item.Rewards = [][]int{{2,2,4000},{201,405001,2}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60010
	item.Building_id = 4
	item.Number = 3
	item.Building_lv = 5
	item.Rewards = [][]int{{2,2,4000},{201,405001,2}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60011
	item.Building_id = 1
	item.Number = 1
	item.Building_lv = 6
	item.Rewards = [][]int{{2,2,8000},{201,405001,2}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60012
	item.Building_id = 1
	item.Number = 1
	item.Building_lv = 8
	item.Rewards = [][]int{{2,2,20000},{201,405001,3}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60013
	item.Building_id = 1
	item.Number = 1
	item.Building_lv = 9
	item.Rewards = [][]int{{2,2,30000},{201,405001,3}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60014
	item.Building_id = 4
	item.Number = 4
	item.Building_lv = 8
	item.Rewards = [][]int{{2,2,8000},{201,405001,3}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60015
	item.Building_id = 5
	item.Number = 4
	item.Building_lv = 8
	item.Rewards = [][]int{{2,2,8000},{201,405001,3}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
	item = &Limithd_jianzhudashi_Item{}
	item.Id = 60016
	item.Building_id = 9
	item.Number = 2
	item.Building_lv = 8
	item.Rewards = [][]int{{2,2,8000},{201,405001,3}}
	InstLimithd_jianzhudashi.Items = append(InstLimithd_jianzhudashi.Items,item)
}
