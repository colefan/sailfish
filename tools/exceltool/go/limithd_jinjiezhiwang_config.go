package config
type Limithd_jinjiezhiwang struct{
	Items []*Limithd_jinjiezhiwang_Item
}

type Limithd_jinjiezhiwang_Item struct{
	Id int
	Number int
	Color int
	Grade int
	Rewards [][]int
}

var InstLimithd_jinjiezhiwang *Limithd_jinjiezhiwang

func init(){
	item := &Limithd_jinjiezhiwang_Item{}
	item.Id = 50001
	item.Number = 3
	item.Color = 2
	item.Grade = 1
	item.Rewards = [][]int{{201,431001,1},{201,406002,1}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50002
	item.Number = 5
	item.Color = 2
	item.Grade = 1
	item.Rewards = [][]int{{201,431001,1},{201,406002,1}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50003
	item.Number = 3
	item.Color = 2
	item.Grade = 2
	item.Rewards = [][]int{{201,431001,2},{201,406002,2}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50004
	item.Number = 5
	item.Color = 2
	item.Grade = 2
	item.Rewards = [][]int{{201,431001,3},{201,310023,10}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50005
	item.Number = 1
	item.Color = 3
	item.Grade = 1
	item.Rewards = [][]int{{201,431001,4},{201,406003,1}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50006
	item.Number = 5
	item.Color = 3
	item.Grade = 1
	item.Rewards = [][]int{{201,431001,5},{201,310023,10}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50007
	item.Number = 1
	item.Color = 3
	item.Grade = 2
	item.Rewards = [][]int{{201,431001,6},{201,406003,2}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50008
	item.Number = 3
	item.Color = 3
	item.Grade = 2
	item.Rewards = [][]int{{201,431001,8},{201,406003,2}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50009
	item.Number = 1
	item.Color = 3
	item.Grade = 3
	item.Rewards = [][]int{{201,431001,10},{201,310023,20}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
	item = &Limithd_jinjiezhiwang_Item{}
	item.Id = 50010
	item.Number = 3
	item.Color = 3
	item.Grade = 3
	item.Rewards = [][]int{{201,431001,10},{201,406003,3}}
	InstLimithd_jinjiezhiwang.Items = append(InstLimithd_jinjiezhiwang.Items,item)
}
