package config
type Boss_guwu struct{
	Items []*Boss_guwu_Item
}

type Boss_guwu_Item struct{
	Time int
	Consume [][]int
	Value [][]int
}

var InstBoss_guwu *Boss_guwu

func init(){
	item := &Boss_guwu_Item{}
	item.Time = 1
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,100},{10005,100}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 2
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,200},{10005,200}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 3
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,300},{10005,300}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 4
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,400},{10005,400}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 5
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,500},{10005,500}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 6
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,600},{10005,600}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 7
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,700},{10005,700}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 8
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,800},{10005,800}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 9
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,900},{10005,900}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 10
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1000},{10005,1000}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 11
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1100},{10005,1100}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 12
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1200},{10005,1200}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 13
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1300},{10005,1300}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 14
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1400},{10005,1400}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 15
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1500},{10005,1500}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 16
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1600},{10005,1600}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 17
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1700},{10005,1700}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 18
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1800},{10005,1800}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 19
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,1900},{10005,1900}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
	item = &Boss_guwu_Item{}
	item.Time = 20
	item.Consume = [][]int{{1,1,10}}
	item.Value = [][]int{{10004,2000},{10005,2000}}
	InstBoss_guwu.Items = append(InstBoss_guwu.Items,item)
}
