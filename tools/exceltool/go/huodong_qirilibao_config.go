package config
type Huodong_qirilibao struct{
	Items []*Huodong_qirilibao_Item
}

type Huodong_qirilibao_Item struct{
	Login_day int
	Rewards [][]int
}

var InstHuodong_qirilibao *Huodong_qirilibao

func init(){
	item := &Huodong_qirilibao_Item{}
	item.Login_day = 1
	item.Rewards = [][]int{{201,310049,10,1},{1,1,100,0},{3,3,20000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 2
	item.Rewards = [][]int{{201,310023,40,1},{1,1,158,0},{2,2,10000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 3
	item.Rewards = [][]int{{201,310039,10,1},{1,1,168,0},{3,3,30000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 4
	item.Rewards = [][]int{{201,310039,10,1},{1,1,188,0},{2,2,30000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 5
	item.Rewards = [][]int{{201,310039,20,1},{1,1,208,0},{3,3,40000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 6
	item.Rewards = [][]int{{201,432003,10,1},{1,1,218,0},{2,2,50000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
	item = &Huodong_qirilibao_Item{}
	item.Login_day = 7
	item.Rewards = [][]int{{201,310006,100,1},{1,1,228,0},{3,3,50000,0}}
	InstHuodong_qirilibao.Items = append(InstHuodong_qirilibao.Items,item)
}
