package config
type Chunjie_lianxuchongzhi struct{
	Items []*Chunjie_lianxuchongzhi_Item
}

type Chunjie_lianxuchongzhi_Item struct{
	Id int
	Type int
	Day int
	Rewards [][]int
}

var InstChunjie_lianxuchongzhi *Chunjie_lianxuchongzhi

func init(){
	item := &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50001
	item.Type = 1
	item.Day = 1
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50002
	item.Type = 1
	item.Day = 2
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50003
	item.Type = 1
	item.Day = 3
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50004
	item.Type = 1
	item.Day = 4
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50005
	item.Type = 1
	item.Day = 5
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50006
	item.Type = 1
	item.Day = 6
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50007
	item.Type = 1
	item.Day = 7
	item.Rewards = [][]int{{1,1,120},{2,2,100000},{201,435002,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50008
	item.Type = 2
	item.Day = 3
	item.Rewards = [][]int{{1,1,300},{2,2,500000},{201,435003,5}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50009
	item.Type = 2
	item.Day = 5
	item.Rewards = [][]int{{1,1,500},{2,2,500000},{201,431001,30}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
	item = &Chunjie_lianxuchongzhi_Item{}
	item.Id = 50010
	item.Type = 2
	item.Day = 7
	item.Rewards = [][]int{{1,1,800},{2,2,500000},{201,435003,10}}
	InstChunjie_lianxuchongzhi.Items = append(InstChunjie_lianxuchongzhi.Items,item)
}
