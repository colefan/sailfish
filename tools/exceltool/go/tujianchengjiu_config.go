package config
type Tujianchengjiu struct{
	Items []*Tujianchengjiu_Item
}

type Tujianchengjiu_Item struct{
	Id int
	Front_id int
	Achievement_value int
	Consume [][]int
	Target int
	Attribute [][]int
}

var InstTujianchengjiu *Tujianchengjiu

func init(){
	item := &Tujianchengjiu_Item{}
	item.Id = 1000
	item.Front_id = 0
	item.Achievement_value = 0
	item.Consume = make([][]int,0)
	item.Target = 1
	item.Attribute = [][]int{{10004,0},{10005,0},{10006,0}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1001
	item.Front_id = 0
	item.Achievement_value = 50
	item.Consume = [][]int{{2,2,20000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,80},{10005,0},{10006,0}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1002
	item.Front_id = 1001
	item.Achievement_value = 120
	item.Consume = [][]int{{2,2,50000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,80},{10005,80},{10006,0}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1003
	item.Front_id = 1002
	item.Achievement_value = 200
	item.Consume = [][]int{{2,2,80000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,80},{10005,80},{10006,80}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1004
	item.Front_id = 1003
	item.Achievement_value = 300
	item.Consume = [][]int{{2,2,100000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,180},{10005,80},{10006,80}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1005
	item.Front_id = 1004
	item.Achievement_value = 400
	item.Consume = [][]int{{2,2,120000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,180},{10005,180},{10006,80}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1006
	item.Front_id = 1005
	item.Achievement_value = 500
	item.Consume = [][]int{{2,2,140000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,180},{10005,180},{10006,180}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1007
	item.Front_id = 1006
	item.Achievement_value = 600
	item.Consume = [][]int{{2,2,160000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,280},{10005,180},{10006,180}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1008
	item.Front_id = 1007
	item.Achievement_value = 750
	item.Consume = [][]int{{2,2,180000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,280},{10005,280},{10006,180}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1009
	item.Front_id = 1008
	item.Achievement_value = 900
	item.Consume = [][]int{{2,2,200000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,280},{10005,280},{10006,280}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1010
	item.Front_id = 1009
	item.Achievement_value = 1000
	item.Consume = [][]int{{2,2,250000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,430},{10005,280},{10006,280}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1011
	item.Front_id = 1010
	item.Achievement_value = 1100
	item.Consume = [][]int{{2,2,250000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,430},{10005,430},{10006,280}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1012
	item.Front_id = 1011
	item.Achievement_value = 1200
	item.Consume = [][]int{{2,2,250000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,430},{10005,430},{10006,430}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1013
	item.Front_id = 1012
	item.Achievement_value = 1300
	item.Consume = [][]int{{2,2,400000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,630},{10005,430},{10006,430}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1014
	item.Front_id = 1013
	item.Achievement_value = 1400
	item.Consume = [][]int{{2,2,400000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,630},{10005,630},{10006,430}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1015
	item.Front_id = 1014
	item.Achievement_value = 1500
	item.Consume = [][]int{{2,2,400000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,630},{10005,630},{10006,630}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1016
	item.Front_id = 1015
	item.Achievement_value = 1650
	item.Consume = [][]int{{2,2,600000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,830},{10005,630},{10006,630}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1017
	item.Front_id = 1016
	item.Achievement_value = 1800
	item.Consume = [][]int{{2,2,600000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,830},{10005,830},{10006,630}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
	item = &Tujianchengjiu_Item{}
	item.Id = 1018
	item.Front_id = 1017
	item.Achievement_value = 2000
	item.Consume = [][]int{{2,2,600000}}
	item.Target = 1
	item.Attribute = [][]int{{10004,830},{10005,830},{10006,830}}
	InstTujianchengjiu.Items = append(InstTujianchengjiu.Items,item)
}
