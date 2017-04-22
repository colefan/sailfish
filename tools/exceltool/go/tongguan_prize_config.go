package config
type Tongguan_prize struct{
	Items []*Tongguan_prize_Item
}

type Tongguan_prize_Item struct{
	Id int
	Map_id int
	Star int
	Prize [][]int
}

var InstTongguan_prize *Tongguan_prize

func init(){
	item := &Tongguan_prize_Item{}
	item.Id = 1001
	item.Map_id = 10100
	item.Star = 5
	item.Prize = [][]int{{2,2,5000},{1,1,20},{201,431001,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1002
	item.Map_id = 10100
	item.Star = 10
	item.Prize = [][]int{{2,2,5000},{1,1,30},{201,431001,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1003
	item.Map_id = 10100
	item.Star = 15
	item.Prize = [][]int{{2,2,5000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1004
	item.Map_id = 10200
	item.Star = 5
	item.Prize = [][]int{{2,2,5000},{1,1,20},{201,431001,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1005
	item.Map_id = 10200
	item.Star = 10
	item.Prize = [][]int{{2,2,5000},{1,1,30},{201,431001,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1006
	item.Map_id = 10200
	item.Star = 15
	item.Prize = [][]int{{2,2,5000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1007
	item.Map_id = 10300
	item.Star = 5
	item.Prize = [][]int{{2,2,7500},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1008
	item.Map_id = 10300
	item.Star = 10
	item.Prize = [][]int{{2,2,7500},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1009
	item.Map_id = 10300
	item.Star = 15
	item.Prize = [][]int{{2,2,7500},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1010
	item.Map_id = 10400
	item.Star = 5
	item.Prize = [][]int{{2,2,10000},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1011
	item.Map_id = 10400
	item.Star = 10
	item.Prize = [][]int{{2,2,10000},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1012
	item.Map_id = 10400
	item.Star = 15
	item.Prize = [][]int{{2,2,10000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1013
	item.Map_id = 10500
	item.Star = 5
	item.Prize = [][]int{{2,2,15000},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1014
	item.Map_id = 10500
	item.Star = 10
	item.Prize = [][]int{{2,2,15000},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1015
	item.Map_id = 10500
	item.Star = 15
	item.Prize = [][]int{{2,2,15000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1016
	item.Map_id = 10600
	item.Star = 5
	item.Prize = [][]int{{2,2,20000},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1017
	item.Map_id = 10600
	item.Star = 10
	item.Prize = [][]int{{2,2,20000},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1018
	item.Map_id = 10600
	item.Star = 15
	item.Prize = [][]int{{2,2,20000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1019
	item.Map_id = 10700
	item.Star = 5
	item.Prize = [][]int{{2,2,25000},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1020
	item.Map_id = 10700
	item.Star = 10
	item.Prize = [][]int{{2,2,25000},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1021
	item.Map_id = 10700
	item.Star = 15
	item.Prize = [][]int{{2,2,25000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1022
	item.Map_id = 10800
	item.Star = 5
	item.Prize = [][]int{{2,2,30000},{1,1,20},{201,432003,2}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1023
	item.Map_id = 10800
	item.Star = 10
	item.Prize = [][]int{{2,2,30000},{1,1,30},{201,432003,3}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
	item = &Tongguan_prize_Item{}
	item.Id = 1024
	item.Map_id = 10800
	item.Star = 15
	item.Prize = [][]int{{2,2,30000},{1,1,50},{201,432003,5}}
	InstTongguan_prize.Items = append(InstTongguan_prize.Items,item)
}
