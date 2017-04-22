package config
type Yuanzhen struct{
	Items []*Yuanzhen_Item
}

type Yuanzhen_Item struct{
	Id int
	Pre_id int
	Base_reward [][]int
	Battlemap int
	Has_shop int
	Handicap int
}

var InstYuanzhen *Yuanzhen

func init(){
	item := &Yuanzhen_Item{}
	item.Id = 1001
	item.Pre_id = 0
	item.Base_reward = [][]int{{2,2,5000},{10,10,1},{205,501001,1}}
	item.Battlemap = 999100
	item.Has_shop = 0
	item.Handicap = 3000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1002
	item.Pre_id = 1001
	item.Base_reward = [][]int{{11,11,100},{10,10,2},{205,501001,1}}
	item.Battlemap = 999101
	item.Has_shop = 0
	item.Handicap = 3000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1003
	item.Pre_id = 1002
	item.Base_reward = [][]int{{2,2,6000},{10,10,3},{205,501001,1}}
	item.Battlemap = 999102
	item.Has_shop = 0
	item.Handicap = 3500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1004
	item.Pre_id = 1003
	item.Base_reward = [][]int{{2,2,7000},{10,10,4},{205,501001,1}}
	item.Battlemap = 999103
	item.Has_shop = 1
	item.Handicap = 3500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1005
	item.Pre_id = 1004
	item.Base_reward = [][]int{{11,11,200},{10,10,2},{205,501001,1}}
	item.Battlemap = 999104
	item.Has_shop = 0
	item.Handicap = 4000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1006
	item.Pre_id = 1005
	item.Base_reward = [][]int{{2,2,8000},{10,10,3},{205,501001,1}}
	item.Battlemap = 999105
	item.Has_shop = 0
	item.Handicap = 4000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1007
	item.Pre_id = 1006
	item.Base_reward = [][]int{{2,2,9000},{10,10,4},{205,501001,1}}
	item.Battlemap = 999106
	item.Has_shop = 0
	item.Handicap = 4500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1008
	item.Pre_id = 1007
	item.Base_reward = [][]int{{11,11,200},{10,10,5},{205,501001,1}}
	item.Battlemap = 999107
	item.Has_shop = 1
	item.Handicap = 4500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1009
	item.Pre_id = 1008
	item.Base_reward = [][]int{{2,2,10000},{10,10,2},{205,501001,1}}
	item.Battlemap = 999108
	item.Has_shop = 0
	item.Handicap = 5500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1010
	item.Pre_id = 1009
	item.Base_reward = [][]int{{2,2,11000},{10,10,3},{205,501001,1}}
	item.Battlemap = 999109
	item.Has_shop = 0
	item.Handicap = 5500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1011
	item.Pre_id = 1010
	item.Base_reward = [][]int{{11,11,250},{10,10,4},{205,501001,1}}
	item.Battlemap = 999110
	item.Has_shop = 1
	item.Handicap = 6500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1012
	item.Pre_id = 1011
	item.Base_reward = [][]int{{2,2,12000},{10,10,5},{205,501001,1}}
	item.Battlemap = 999111
	item.Has_shop = 0
	item.Handicap = 6500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1013
	item.Pre_id = 1012
	item.Base_reward = [][]int{{2,2,13000},{10,10,3},{205,501001,1}}
	item.Battlemap = 999112
	item.Has_shop = 0
	item.Handicap = 7000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1014
	item.Pre_id = 1013
	item.Base_reward = [][]int{{11,11,250},{10,10,4},{205,501001,1}}
	item.Battlemap = 999113
	item.Has_shop = 1
	item.Handicap = 7000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1015
	item.Pre_id = 1014
	item.Base_reward = [][]int{{2,2,14000},{10,10,5},{205,501001,1}}
	item.Battlemap = 999114
	item.Has_shop = 0
	item.Handicap = 8000
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
	item = &Yuanzhen_Item{}
	item.Id = 1016
	item.Pre_id = 1015
	item.Base_reward = [][]int{{11,11,300},{10,10,6},{205,501001,1}}
	item.Battlemap = 999115
	item.Has_shop = 0
	item.Handicap = 8500
	InstYuanzhen.Items = append(InstYuanzhen.Items,item)
}
