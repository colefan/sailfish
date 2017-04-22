package config
type Chongzhilibao struct{
	Items []*Chongzhilibao_Item
}

type Chongzhilibao_Item struct{
	Id int
	Name string
	Pay int
	Time_limit int
	Reward [][]int
}

var InstChongzhilibao *Chongzhilibao

func init(){
	item := &Chongzhilibao_Item{}
	item.Id = 1001
	item.Name = ""
	item.Pay = 1
	item.Time_limit = 999
	item.Reward = [][]int{{2,2,20000},{3,3,40000},{201,405001,20},{201,431001,10},{202,10007,1}}
	InstChongzhilibao.Items = append(InstChongzhilibao.Items,item)
}
