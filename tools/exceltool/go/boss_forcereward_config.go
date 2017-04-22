package config
type Boss_forcereward struct{
	Items []*Boss_forcereward_Item
}

type Boss_forcereward_Item struct{
	Id int
	Reward [][]int
	Green_box int
	Blue_box int
	Violet_box int
	Orange_reward int
}

var InstBoss_forcereward *Boss_forcereward

func init(){
	item := &Boss_forcereward_Item{}
	item.Id = 1
	item.Reward = [][]int{{6,6,6000}}
	item.Green_box = 100
	item.Blue_box = 60
	item.Violet_box = 25
	item.Orange_reward = 12
	InstBoss_forcereward.Items = append(InstBoss_forcereward.Items,item)
	item = &Boss_forcereward_Item{}
	item.Id = 2
	item.Reward = [][]int{{6,6,4500}}
	item.Green_box = 90
	item.Blue_box = 55
	item.Violet_box = 20
	item.Orange_reward = 10
	InstBoss_forcereward.Items = append(InstBoss_forcereward.Items,item)
	item = &Boss_forcereward_Item{}
	item.Id = 3
	item.Reward = [][]int{{6,6,3000}}
	item.Green_box = 80
	item.Blue_box = 50
	item.Violet_box = 15
	item.Orange_reward = 8
	InstBoss_forcereward.Items = append(InstBoss_forcereward.Items,item)
}
