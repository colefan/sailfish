package config
type Shili_box struct{
	Items []*Shili_box_Item
}

type Shili_box_Item struct{
	Id int
	Lv_min int
	Lv_max int
	Consume [][]int
	Whitebox_pro int
	Greenbox_pro int
	Bluebox_pro int
	Violetbox_pro int
	Orangebox_pro int
	White_reward [][]int
	Green_reward [][]int
	Bule_reward [][]int
	Violet_reward [][]int
	Orange_reward [][]int
}

var InstShili_box *Shili_box

func init(){
	item := &Shili_box_Item{}
	item.Id = 1
	item.Lv_min = 1
	item.Lv_max = 39
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,1500},{2,2,2000},{18,18,1500},{18,18,2000},{6,6,100}}
	item.Green_reward = [][]int{{2,2,5000},{18,18,4000},{210,431001,3},{201,431002,5},{6,6,100}}
	item.Bule_reward = [][]int{{2,2,10000},{201,310001,2},{201,431001,5},{201,435001,3},{6,6,200}}
	item.Violet_reward = [][]int{{2,2,15000},{201,310001,5},{201,435001,5},{201,432004,2},{6,6,300}}
	item.Orange_reward = [][]int{{2,2,20000},{201,435001,10},{201,432004,4},{201,432003,5},{201,310001,15}}
	InstShili_box.Items = append(InstShili_box.Items,item)
	item = &Shili_box_Item{}
	item.Id = 2
	item.Lv_min = 40
	item.Lv_max = 49
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,2000},{2,2,4000},{18,18,3000},{18,18,6000},{6,6,100}}
	item.Green_reward = [][]int{{2,2,5000},{4,4,1000},{201,431001,3},{201,431002,5},{6,6,100}}
	item.Bule_reward = [][]int{{2,2,10000},{4,4,2000},{201,431001,5},{201,435001,3},{201,310001,2}}
	item.Violet_reward = [][]int{{2,2,15000},{4,4,3000},{201,435001,5},{201,432004,2},{201,310001,5}}
	item.Orange_reward = [][]int{{2,2,20000},{201,435001,10},{201,432004,4},{201,432003,5},{201,310001,15}}
	InstShili_box.Items = append(InstShili_box.Items,item)
	item = &Shili_box_Item{}
	item.Id = 3
	item.Lv_min = 50
	item.Lv_max = 59
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,2500},{2,2,5000},{18,18,4000},{18,18,8000},{6,6,100}}
	item.Green_reward = [][]int{{2,2,6000},{4,4,1000},{201,431001,3},{201,431002,5},{6,6,200}}
	item.Bule_reward = [][]int{{2,2,12000},{4,4,2000},{201,310001,2},{201,435001,3},{6,6,300}}
	item.Violet_reward = [][]int{{201,310001,5},{4,4,3000},{201,435001,5},{201,432004,2},{201,406007,2}}
	item.Orange_reward = [][]int{{201,310001,15},{201,435001,10},{201,432004,4},{201,406007,4},{201,432003,5}}
	InstShili_box.Items = append(InstShili_box.Items,item)
	item = &Shili_box_Item{}
	item.Id = 4
	item.Lv_min = 60
	item.Lv_max = 69
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,3000},{2,2,6000},{18,18,15000},{4,4,800},{6,6,100}}
	item.Green_reward = [][]int{{2,2,6000},{4,4,1200},{201,431001,3},{201,431002,5},{6,6,200}}
	item.Bule_reward = [][]int{{201,310001,2},{4,4,2400},{201,405003,3},{201,435001,4},{6,6,300}}
	item.Violet_reward = [][]int{{201,310001,5},{4,4,3600},{201,435001,8},{201,432004,3},{201,406008,2}}
	item.Orange_reward = [][]int{{201,310001,15},{201,435001,12},{201,432004,6},{201,406008,4},{201,435004,4}}
	InstShili_box.Items = append(InstShili_box.Items,item)
	item = &Shili_box_Item{}
	item.Id = 5
	item.Lv_min = 70
	item.Lv_max = 79
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,4000},{2,2,8000},{18,18,20000},{4,4,800},{6,6,150}}
	item.Green_reward = [][]int{{2,2,8000},{4,4,1200},{201,431001,3},{201,431002,5},{6,6,300}}
	item.Bule_reward = [][]int{{201,310001,2},{4,4,2400},{201,405003,5},{201,435001,4},{201,406009,2}}
	item.Violet_reward = [][]int{{201,310001,5},{4,4,3600},{201,435001,8},{201,432004,3},{201,406009,3}}
	item.Orange_reward = [][]int{{201,310001,15},{201,435001,12},{201,432004,6},{201,406009,5},{201,435005,6}}
	InstShili_box.Items = append(InstShili_box.Items,item)
	item = &Shili_box_Item{}
	item.Id = 6
	item.Lv_min = 80
	item.Lv_max = 90
	item.Consume = [][]int{{21,21,10}}
	item.Whitebox_pro = 4300
	item.Greenbox_pro = 3000
	item.Bluebox_pro = 1500
	item.Violetbox_pro = 800
	item.Orangebox_pro = 400
	item.White_reward = [][]int{{2,2,4000},{2,2,8000},{18,18,25000},{4,4,1000},{6,6,150}}
	item.Green_reward = [][]int{{2,2,8000},{4,4,1500},{201,431001,3},{201,431002,5},{6,6,300}}
	item.Bule_reward = [][]int{{201,310001,2},{4,4,3000},{201,405004,2},{201,435001,5},{201,406010,2}}
	item.Violet_reward = [][]int{{201,310001,5},{4,4,4500},{201,435001,10},{201,432004,3},{201,406010,3}}
	item.Orange_reward = [][]int{{201,310001,15},{201,435001,15},{201,432004,6},{201,406010,5},{201,435006,8}}
	InstShili_box.Items = append(InstShili_box.Items,item)
}
