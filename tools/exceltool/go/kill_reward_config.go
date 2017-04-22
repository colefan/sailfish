package config
type Kill_reward struct{
	Items []*Kill_reward_Item
}

type Kill_reward_Item struct{
	Id int
	Fightingcapacity_pro int
	Gongxun_pro int
}

var InstKill_reward *Kill_reward

func init(){
	item := &Kill_reward_Item{}
	item.Id = 1
	item.Fightingcapacity_pro = 0
	item.Gongxun_pro = 50000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 2
	item.Fightingcapacity_pro = 4000
	item.Gongxun_pro = 40000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 3
	item.Fightingcapacity_pro = 5000
	item.Gongxun_pro = 30000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 4
	item.Fightingcapacity_pro = 6000
	item.Gongxun_pro = 20000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 5
	item.Fightingcapacity_pro = 7000
	item.Gongxun_pro = 15000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 6
	item.Fightingcapacity_pro = 8000
	item.Gongxun_pro = 12000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 7
	item.Fightingcapacity_pro = 9000
	item.Gongxun_pro = 10000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 8
	item.Fightingcapacity_pro = 11000
	item.Gongxun_pro = 9000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 9
	item.Fightingcapacity_pro = 12000
	item.Gongxun_pro = 8000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 10
	item.Fightingcapacity_pro = 13000
	item.Gongxun_pro = 7000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 11
	item.Fightingcapacity_pro = 14000
	item.Gongxun_pro = 6000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 12
	item.Fightingcapacity_pro = 15500
	item.Gongxun_pro = 5000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 13
	item.Fightingcapacity_pro = 17000
	item.Gongxun_pro = 4000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 14
	item.Fightingcapacity_pro = 18500
	item.Gongxun_pro = 3000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 15
	item.Fightingcapacity_pro = 20000
	item.Gongxun_pro = 2000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 16
	item.Fightingcapacity_pro = 30000
	item.Gongxun_pro = 1000
	InstKill_reward.Items = append(InstKill_reward.Items,item)
	item = &Kill_reward_Item{}
	item.Id = 17
	item.Fightingcapacity_pro = 40000
	item.Gongxun_pro = 0
	InstKill_reward.Items = append(InstKill_reward.Items,item)
}
