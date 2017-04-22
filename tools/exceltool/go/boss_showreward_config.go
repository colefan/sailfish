package config
type Boss_showreward struct{
	Items []*Boss_showreward_Item
}

type Boss_showreward_Item struct{
	Id int
	Lv_min int
	Lv_max int
	Showreward [][]int
}

var InstBoss_showreward *Boss_showreward

func init(){
	item := &Boss_showreward_Item{}
	item.Id = 1001
	item.Lv_min = 30
	item.Lv_max = 34
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1002
	item.Lv_min = 35
	item.Lv_max = 39
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1003
	item.Lv_min = 40
	item.Lv_max = 44
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1004
	item.Lv_min = 45
	item.Lv_max = 49
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1005
	item.Lv_min = 50
	item.Lv_max = 54
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1006
	item.Lv_min = 55
	item.Lv_max = 59
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1007
	item.Lv_min = 60
	item.Lv_max = 64
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1008
	item.Lv_min = 65
	item.Lv_max = 69
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1009
	item.Lv_min = 70
	item.Lv_max = 74
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
	item = &Boss_showreward_Item{}
	item.Id = 1010
	item.Lv_min = 75
	item.Lv_max = 79
	item.Showreward = [][]int{{17,17,100},{2,2,10000},{4,4,10000},{21,21,10000}}
	InstBoss_showreward.Items = append(InstBoss_showreward.Items,item)
}
