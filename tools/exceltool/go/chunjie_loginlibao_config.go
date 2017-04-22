package config
type Chunjie_loginlibao struct{
	Items []*Chunjie_loginlibao_Item
}

type Chunjie_loginlibao_Item struct{
	Login_day int
	Rewards [][]int
}

var InstChunjie_loginlibao *Chunjie_loginlibao

func init(){
	item := &Chunjie_loginlibao_Item{}
	item.Login_day = 1
	item.Rewards = [][]int{{1,1,100},{201,435002,5},{201,431001,10}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 2
	item.Rewards = [][]int{{2,2,50000},{201,435001,20},{18,18,80000}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 3
	item.Rewards = [][]int{{1,1,100},{201,431001,20},{201,435003,10}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 4
	item.Rewards = [][]int{{201,432003,2},{18,18,100000},{2,2,60000}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 5
	item.Rewards = [][]int{{1,1,150},{2,2,80000},{18,18,100000}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 6
	item.Rewards = [][]int{{17,17,20},{201,431001,20},{201,435001,20}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
	item = &Chunjie_loginlibao_Item{}
	item.Login_day = 7
	item.Rewards = [][]int{{1,1,288},{201,432004,5},{201,432003,5}}
	InstChunjie_loginlibao.Items = append(InstChunjie_loginlibao.Items,item)
}
