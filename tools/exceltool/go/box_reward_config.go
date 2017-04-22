package config
type Box_reward struct{
	Items []*Box_reward_Item
}

type Box_reward_Item struct{
	Box_id int
	Level_min int
	Level_max int
	Rewards [][]int
}

var InstBox_reward *Box_reward

func init(){
	item := &Box_reward_Item{}
	item.Box_id = 101
	item.Level_min = 1
	item.Level_max = 34
	item.Rewards = [][]int{{5,5,100,1},{5,5,100,1},{201,431005,2,0},{6,6,150,1},{3,3,7500,0},{3,3,5000,0},{2,2,5000,0},{2,2,2500,0},{201,431001,1,0},{17,17,1,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
	item = &Box_reward_Item{}
	item.Box_id = 201
	item.Level_min = 35
	item.Level_max = 39
	item.Rewards = [][]int{{5,5,100,1},{5,5,100,1},{201,431005,2,0},{6,6,150,1},{3,3,10000,0},{201,435001,2,0},{2,2,7500,0},{201,431001,1,0},{201,435001,1,0},{17,17,1,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
	item = &Box_reward_Item{}
	item.Box_id = 301
	item.Level_min = 40
	item.Level_max = 54
	item.Rewards = [][]int{{5,5,150,1},{5,5,100,1},{201,431005,2,0},{6,6,180,1},{3,3,15000,0},{201,435001,2,0},{2,2,10000,0},{201,431001,1,0},{201,435001,1,0},{17,17,2,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
	item = &Box_reward_Item{}
	item.Box_id = 401
	item.Level_min = 55
	item.Level_max = 69
	item.Rewards = [][]int{{5,5,175,1},{5,5,125,1},{201,431005,2,0},{6,6,200,1},{3,3,25000,0},{201,435001,4,0},{2,2,15000,0},{2,2,10000,0},{201,435001,2,0},{17,17,2,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
	item = &Box_reward_Item{}
	item.Box_id = 501
	item.Level_min = 70
	item.Level_max = 79
	item.Rewards = [][]int{{5,5,200,1},{5,5,150,1},{201,431005,2,0},{6,6,250,1},{3,3,30000,0},{201,435001,4,0},{2,2,18000,0},{2,2,15000,0},{201,435001,2,0},{17,17,2,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
	item = &Box_reward_Item{}
	item.Box_id = 601
	item.Level_min = 80
	item.Level_max = 89
	item.Rewards = [][]int{{5,5,250,1},{5,5,150,1},{201,431005,2,0},{6,6,300,1},{3,3,35000,0},{201,435001,6,0},{2,2,24000,0},{2,2,18000,0},{201,435001,3,0},{17,17,2,0}}
	InstBox_reward.Items = append(InstBox_reward.Items,item)
}
