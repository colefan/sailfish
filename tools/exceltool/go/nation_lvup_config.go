package config
type Nation_lvup struct{
	Items []*Nation_lvup_Item
}

type Nation_lvup_Item struct{
	Nation_lv int
	Exp int
	City_money int
	City_iron int
	Killing_order int
	Tech_max int
	Yuanjun_tech_max int
	Npc_num int
	Task_exp int
	Task_gongxun int
	Task_green_box int
	Task_blue_box int
	Task_violet_box int
	Task_orange_box int
	Task_expand_exp int
	Task_expand_gongxun int
	Taskexpand_green_box int
	Taskexpand_blue_box int
	Taskexpand_violet_box int
	Taskexpand_orange_box int
	Lvup_reward [][]int
}

var InstNation_lvup *Nation_lvup

func init(){
	item := &Nation_lvup_Item{}
	item.Nation_lv = 1
	item.Exp = 2000
	item.City_money = 200
	item.City_iron = 0
	item.Killing_order = 1
	item.Tech_max = 10
	item.Yuanjun_tech_max = 5
	item.Npc_num = 0
	item.Task_exp = 1000
	item.Task_gongxun = 1200
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 800
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,50000},{6,6,2000},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 2
	item.Exp = 5000
	item.City_money = 250
	item.City_iron = 0
	item.Killing_order = 1
	item.Tech_max = 12
	item.Yuanjun_tech_max = 10
	item.Npc_num = 0
	item.Task_exp = 1000
	item.Task_gongxun = 1200
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 800
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,100000},{6,6,2500},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 3
	item.Exp = 9000
	item.City_money = 300
	item.City_iron = 0
	item.Killing_order = 1
	item.Tech_max = 16
	item.Yuanjun_tech_max = 15
	item.Npc_num = 0
	item.Task_exp = 1000
	item.Task_gongxun = 1400
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 900
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,150000},{6,6,3000},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 4
	item.Exp = 15000
	item.City_money = 360
	item.City_iron = 100
	item.Killing_order = 1
	item.Tech_max = 20
	item.Yuanjun_tech_max = 20
	item.Npc_num = 0
	item.Task_exp = 1000
	item.Task_gongxun = 1400
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 900
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,200000},{6,6,3500},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 5
	item.Exp = 25000
	item.City_money = 420
	item.City_iron = 120
	item.Killing_order = 1
	item.Tech_max = 24
	item.Yuanjun_tech_max = 25
	item.Npc_num = 1
	item.Task_exp = 1000
	item.Task_gongxun = 1600
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 1000
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,250000},{6,6,4000},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 6
	item.Exp = 40000
	item.City_money = 500
	item.City_iron = 140
	item.Killing_order = 1
	item.Tech_max = 28
	item.Yuanjun_tech_max = 30
	item.Npc_num = 1
	item.Task_exp = 1000
	item.Task_gongxun = 1600
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 1000
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,300000},{6,6,4500},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 7
	item.Exp = 80000
	item.City_money = 600
	item.City_iron = 160
	item.Killing_order = 1
	item.Tech_max = 32
	item.Yuanjun_tech_max = 35
	item.Npc_num = 1
	item.Task_exp = 1000
	item.Task_gongxun = 1800
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 1200
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,350000},{6,6,5000},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 8
	item.Exp = 150000
	item.City_money = 700
	item.City_iron = 180
	item.Killing_order = 1
	item.Tech_max = 36
	item.Yuanjun_tech_max = 40
	item.Npc_num = 1
	item.Task_exp = 1000
	item.Task_gongxun = 1800
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 1200
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,400000},{6,6,5500},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
	item = &Nation_lvup_Item{}
	item.Nation_lv = 9
	item.Exp = 150000
	item.City_money = 800
	item.City_iron = 200
	item.Killing_order = 1
	item.Tech_max = 40
	item.Yuanjun_tech_max = 45
	item.Npc_num = 1
	item.Task_exp = 1000
	item.Task_gongxun = 2100
	item.Task_green_box = 28
	item.Task_blue_box = 16
	item.Task_violet_box = 6
	item.Task_orange_box = 4
	item.Task_expand_exp = 600
	item.Task_expand_gongxun = 1400
	item.Taskexpand_green_box = 18
	item.Taskexpand_blue_box = 10
	item.Taskexpand_violet_box = 4
	item.Taskexpand_orange_box = 2
	item.Lvup_reward = [][]int{{3,3,450000},{6,6,6000},{1,1,100}}
	InstNation_lvup.Items = append(InstNation_lvup.Items,item)
}
