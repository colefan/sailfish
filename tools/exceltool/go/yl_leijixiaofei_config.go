package config
type Yl_leijixiaofei struct{
	Items []*Yl_leijixiaofei_Item
}

type Yl_leijixiaofei_Item struct{
	Id int
	Gold int
	Rewards [][]int
}

var InstYl_leijixiaofei *Yl_leijixiaofei

func init(){
	item := &Yl_leijixiaofei_Item{}
	item.Id = 1
	item.Gold = 100
	item.Rewards = [][]int{{1,1,10,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 2
	item.Gold = 300
	item.Rewards = [][]int{{1,1,30,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 3
	item.Gold = 500
	item.Rewards = [][]int{{1,1,50,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 4
	item.Gold = 1000
	item.Rewards = [][]int{{1,1,100,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 5
	item.Gold = 3000
	item.Rewards = [][]int{{1,1,300,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 6
	item.Gold = 5000
	item.Rewards = [][]int{{1,1,500,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
	item = &Yl_leijixiaofei_Item{}
	item.Id = 7
	item.Gold = 10000
	item.Rewards = [][]int{{1,1,1000,1}}
	InstYl_leijixiaofei.Items = append(InstYl_leijixiaofei.Items,item)
}
