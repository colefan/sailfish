package config
type Chunjie_leijichognzhi struct{
	Items []*Chunjie_leijichognzhi_Item
}

type Chunjie_leijichognzhi_Item struct{
	Id int
	Rmb int
	Rewards [][]int
}

var InstChunjie_leijichognzhi *Chunjie_leijichognzhi

func init(){
	item := &Chunjie_leijichognzhi_Item{}
	item.Id = 40001
	item.Rmb = 58
	item.Rewards = [][]int{{201,432004,20},{2,2,100000},{18,18,100000}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40002
	item.Rmb = 188
	item.Rewards = [][]int{{201,435001,100},{2,2,300000},{201,432003,30}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40003
	item.Rmb = 688
	item.Rewards = [][]int{{201,310012,100},{2,2,1000000},{17,17,100},{201,405003,50}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40004
	item.Rmb = 988
	item.Rewards = [][]int{{201,432004,100},{201,432003,100},{2,2,1000000},{201,431001,50}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40005
	item.Rmb = 1988
	item.Rewards = [][]int{{201,310012,100},{201,432004,100},{201,435001,200},{2,2,2000000}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40006
	item.Rmb = 3588
	item.Rewards = [][]int{{201,432004,200},{17,17,200},{1,1,5888},{2,2,5000000}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
	item = &Chunjie_leijichognzhi_Item{}
	item.Id = 40007
	item.Rmb = 5888
	item.Rewards = [][]int{{201,310012,100},{201,432004,300},{201,435001,500},{2,2,10000000}}
	InstChunjie_leijichognzhi.Items = append(InstChunjie_leijichognzhi.Items,item)
}
