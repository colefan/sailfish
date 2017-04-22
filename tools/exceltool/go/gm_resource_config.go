package config
type Gm_resource struct{
	Items []*Gm_resource_Item
}

type Gm_resource_Item struct{
	Id int
	Resource string
}

var InstGm_resource *Gm_resource

func init(){
	item := &Gm_resource_Item{}
	item.Id = 1
	item.Resource = "黄金"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 2
	item.Resource = "军饷"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 3
	item.Resource = "弹药"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 4
	item.Resource = "精铁"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 5
	item.Resource = "经验"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 6
	item.Resource = "功勋"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 8
	item.Resource = "战功（兵团竞技货币）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 9
	item.Resource = "部队经验"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 10
	item.Resource = "医疗点（重返德军总部医疗站内货币）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 11
	item.Resource = "军功（重返德军总部商店货币）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 12
	item.Resource = "荣誉点（三军演练货币）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 14
	item.Resource = "VIP经验"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 15
	item.Resource = "充值黄金（真金）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 16
	item.Resource = "援军令"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 17
	item.Resource = "万能碎片"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 18
	item.Resource = "国战粮食"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 201
	item.Resource = "物品"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 202
	item.Resource = "英雄"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 203
	item.Resource = "掉落包"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 204
	item.Resource = "军团资金"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 205
	item.Resource = "抽牌奖励（重返德军总部宝箱）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
	item = &Gm_resource_Item{}
	item.Id = 206
	item.Resource = "召唤援军（国战）"
	InstGm_resource.Items = append(InstGm_resource.Items,item)
}
