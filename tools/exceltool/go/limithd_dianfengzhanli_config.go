package config
type Limithd_dianfengzhanli struct{
	Items []*Limithd_dianfengzhanli_Item
}

type Limithd_dianfengzhanli_Item struct{
	Id int
	Power int
	Rewards [][]int
}

var InstLimithd_dianfengzhanli *Limithd_dianfengzhanli

func init(){
	item := &Limithd_dianfengzhanli_Item{}
	item.Id = 10001
	item.Power = 3000
	item.Rewards = [][]int{{2,2,3000},{201,431001,1},{1,1,20}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
	item = &Limithd_dianfengzhanli_Item{}
	item.Id = 10002
	item.Power = 5000
	item.Rewards = [][]int{{2,2,5000},{201,431001,2},{1,1,30}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
	item = &Limithd_dianfengzhanli_Item{}
	item.Id = 10003
	item.Power = 8000
	item.Rewards = [][]int{{2,2,8000},{201,431001,2},{1,1,50}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
	item = &Limithd_dianfengzhanli_Item{}
	item.Id = 10004
	item.Power = 11000
	item.Rewards = [][]int{{201,310017,10},{201,431001,3},{1,1,100}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
	item = &Limithd_dianfengzhanli_Item{}
	item.Id = 10005
	item.Power = 14000
	item.Rewards = [][]int{{2,2,13000},{201,310017,10},{1,1,200}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
	item = &Limithd_dianfengzhanli_Item{}
	item.Id = 10006
	item.Power = 16000
	item.Rewards = [][]int{{2,2,15000},{201,310017,10},{1,1,300}}
	InstLimithd_dianfengzhanli.Items = append(InstLimithd_dianfengzhanli.Items,item)
}
