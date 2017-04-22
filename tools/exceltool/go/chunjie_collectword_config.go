package config
type Chunjie_collectword struct{
	Items []*Chunjie_collectword_Item
}

type Chunjie_collectword_Item struct{
	Id int
	Collect_item [][]int
	Rewards [][]int
}

var InstChunjie_collectword *Chunjie_collectword

func init(){
	item := &Chunjie_collectword_Item{}
	item.Id = 1
	item.Collect_item = [][]int{{201,434101,1},{201,434102,1},{201,434103,1},{201,434104,1}}
	item.Rewards = [][]int{{18,18,40000},{201,435001,4}}
	InstChunjie_collectword.Items = append(InstChunjie_collectword.Items,item)
	item = &Chunjie_collectword_Item{}
	item.Id = 2
	item.Collect_item = [][]int{{201,434105,1},{201,434106,1},{201,434107,1},{201,434107,1}}
	item.Rewards = [][]int{{4,4,10000},{201,431001,5},{1,1,20}}
	InstChunjie_collectword.Items = append(InstChunjie_collectword.Items,item)
	item = &Chunjie_collectword_Item{}
	item.Id = 3
	item.Collect_item = [][]int{{201,434101,1},{201,434102,1},{201,434105,1},{201,434106,1}}
	item.Rewards = [][]int{{2,2,25000},{201,432004,1},{1,1,20}}
	InstChunjie_collectword.Items = append(InstChunjie_collectword.Items,item)
	item = &Chunjie_collectword_Item{}
	item.Id = 4
	item.Collect_item = [][]int{{201,434103,1},{201,434104,1},{201,434107,1},{201,434107,1}}
	item.Rewards = [][]int{{2,2,25000},{201,432004,1},{201,435001,2}}
	InstChunjie_collectword.Items = append(InstChunjie_collectword.Items,item)
	item = &Chunjie_collectword_Item{}
	item.Id = 5
	item.Collect_item = [][]int{{201,434101,1},{201,434102,1},{201,434103,1},{201,434104,1},{201,434105,1},{201,434106,1},{201,434107,1},{201,434107,1}}
	item.Rewards = [][]int{{4,4,20000},{201,310005,5},{201,431001,5}}
	InstChunjie_collectword.Items = append(InstChunjie_collectword.Items,item)
}
