package config
type Huodong_zaixianjiangli struct{
	Items []*Huodong_zaixianjiangli_Item
}

type Huodong_zaixianjiangli_Item struct{
	Online_prize_id int
	Times int
	Rewards [][]int
}

var InstHuodong_zaixianjiangli *Huodong_zaixianjiangli

func init(){
	item := &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 1
	item.Times = 30
	item.Rewards = [][]int{{201,402001,1},{201,405001,2},{201,403005,1}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 2
	item.Times = 60
	item.Rewards = [][]int{{201,402001,2},{201,405002,1},{201,403005,1}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 3
	item.Times = 120
	item.Rewards = [][]int{{201,431002,2},{1,1,20},{201,431001,1}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 4
	item.Times = 300
	item.Rewards = [][]int{{201,402001,3},{201,431002,2},{201,403005,2}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 5
	item.Times = 600
	item.Rewards = [][]int{{1,1,30},{201,431002,2},{201,403005,2}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 6
	item.Times = 900
	item.Rewards = [][]int{{201,402001,4},{201,431001,1},{201,405002,2}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
	item = &Huodong_zaixianjiangli_Item{}
	item.Online_prize_id = 7
	item.Times = 1800
	item.Rewards = [][]int{{1,1,50},{201,431001,1},{201,403005,2}}
	InstHuodong_zaixianjiangli.Items = append(InstHuodong_zaixianjiangli.Items,item)
}
