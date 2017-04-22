package config
type Huodong_dengjilibao struct{
	Items []*Huodong_dengjilibao_Item
}

type Huodong_dengjilibao_Item struct{
	Level int
	Rewards [][]int
}

var InstHuodong_dengjilibao *Huodong_dengjilibao

func init(){
	item := &Huodong_dengjilibao_Item{}
	item.Level = 10
	item.Rewards = [][]int{{201,406002,1},{3,3,2000},{2,2,2000}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 15
	item.Rewards = [][]int{{201,405001,5},{201,431001,3},{1,1,20}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 18
	item.Rewards = [][]int{{201,406002,3},{3,3,5000},{201,431001,3}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 21
	item.Rewards = [][]int{{2,2,5000},{201,432003,3},{1,1,40}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 24
	item.Rewards = [][]int{{201,431001,4},{201,310049,10},{201,432004,2}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 27
	item.Rewards = [][]int{{3,3,10000},{201,310049,10},{2,2,10000}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 30
	item.Rewards = [][]int{{201,405002,3},{201,432003,5},{1,1,50}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 33
	item.Rewards = [][]int{{201,310017,10},{201,432004,5},{201,406002,5}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 36
	item.Rewards = [][]int{{201,435002,10},{201,310042,10},{1,1,150}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 40
	item.Rewards = [][]int{{2,2,50000},{201,432004,5},{201,310017,10}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 43
	item.Rewards = [][]int{{201,435003,20},{201,432003,5},{1,1,150}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 45
	item.Rewards = [][]int{{201,406003,5},{201,310039,20},{201,435001,10}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 47
	item.Rewards = [][]int{{201,432004,8},{3,3,100000},{201,432003,8}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 50
	item.Rewards = [][]int{{201,435001,20},{201,310049,20},{2,2,100000}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 53
	item.Rewards = [][]int{{201,431001,10},{201,435004,40},{1,1,200}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 55
	item.Rewards = [][]int{{201,406004,3},{201,435001,25},{1,1,200}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
	item = &Huodong_dengjilibao_Item{}
	item.Level = 60
	item.Rewards = [][]int{{2,2,100000},{201,435001,50},{1,1,300}}
	InstHuodong_dengjilibao.Items = append(InstHuodong_dengjilibao.Items,item)
}
