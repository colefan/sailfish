package config
type Huodong_tuijian struct{
	Items []*Huodong_tuijian_Item
}

type Huodong_tuijian_Item struct{
	ID int
	Charactor_id int
	Type int
	Rewards [][]int
	Buy_limit int
	Prize int
}

var InstHuodong_tuijian *Huodong_tuijian

func init(){
	item := &Huodong_tuijian_Item{}
	item.ID = 1
	item.Charactor_id = 10009
	item.Type = 1
	item.Rewards = [][]int{{201,310009,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
	item = &Huodong_tuijian_Item{}
	item.ID = 2
	item.Charactor_id = 10003
	item.Type = 2
	item.Rewards = [][]int{{201,310003,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
	item = &Huodong_tuijian_Item{}
	item.ID = 3
	item.Charactor_id = 10039
	item.Type = 2
	item.Rewards = [][]int{{201,310039,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
	item = &Huodong_tuijian_Item{}
	item.ID = 4
	item.Charactor_id = 10016
	item.Type = 3
	item.Rewards = [][]int{{201,310016,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
	item = &Huodong_tuijian_Item{}
	item.ID = 5
	item.Charactor_id = 10037
	item.Type = 3
	item.Rewards = [][]int{{201,310037,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
	item = &Huodong_tuijian_Item{}
	item.ID = 6
	item.Charactor_id = 10018
	item.Type = 3
	item.Rewards = [][]int{{201,310018,5}}
	item.Buy_limit = 4
	item.Prize = 250
	InstHuodong_tuijian.Items = append(InstHuodong_tuijian.Items,item)
}
