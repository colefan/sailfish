package config
type Limithd_leijichongzhi struct{
	Items []*Limithd_leijichongzhi_Item
}

type Limithd_leijichongzhi_Item struct{
	Id int
	Gold int
	Rewards [][]int
}

var InstLimithd_leijichongzhi *Limithd_leijichongzhi

func init(){
	item := &Limithd_leijichongzhi_Item{}
	item.Id = 40001
	item.Gold = 100
	item.Rewards = [][]int{{17,17,3,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40002
	item.Gold = 300
	item.Rewards = [][]int{{17,17,5,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40003
	item.Gold = 980
	item.Rewards = [][]int{{17,17,10,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40004
	item.Gold = 1980
	item.Rewards = [][]int{{17,17,15,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40005
	item.Gold = 3280
	item.Rewards = [][]int{{17,17,20,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40006
	item.Gold = 6480
	item.Rewards = [][]int{{17,17,25,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
	item = &Limithd_leijichongzhi_Item{}
	item.Id = 40007
	item.Gold = 12800
	item.Rewards = [][]int{{17,17,30,1}}
	InstLimithd_leijichongzhi.Items = append(InstLimithd_leijichongzhi.Items,item)
}
