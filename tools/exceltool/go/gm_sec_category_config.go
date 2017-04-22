package config
type Gm_sec_category struct{
	Items []*Gm_sec_category_Item
}

type Gm_sec_category_Item struct{
	Id int
	Enum_value string
	Phylum string
}

var InstGm_sec_category *Gm_sec_category

func init(){
	item := &Gm_sec_category_Item{}
	item.Id = 1
	item.Enum_value = "sec_libao"
	item.Phylum = "礼包"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 2
	item.Enum_value = "sec_nation"
	item.Phylum = "国战"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 3
	item.Enum_value = "sec_arena"
	item.Phylum = "兵团竞技"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 4
	item.Enum_value = "sec_zhuxian_pt"
	item.Phylum = "普通副本"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 5
	item.Enum_value = "sec_zhuxian_jy"
	item.Phylum = "精英副本"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 6
	item.Enum_value = "sec_crusade"
	item.Phylum = "风云战役"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 7
	item.Enum_value = "sec_zhuxian_ptttjs"
	item.Phylum = "主线普通副本通关奖励"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 8
	item.Enum_value = "sec_zhuxian_jyttjs"
	item.Phylum = "主线精英副本通关奖励"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 9
	item.Enum_value = "sec_tavern"
	item.Phylum = "火车站"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 10
	item.Enum_value = "sec_trade"
	item.Phylum = "军需站"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 11
	item.Enum_value = "sec_training"
	item.Phylum = "兵团演练"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 12
	item.Enum_value = "sec_yuanzheng"
	item.Phylum = "重返德军总部"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 13
	item.Enum_value = "sec_arenatop"
	item.Phylum = "三军演练"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 14
	item.Enum_value = "sec_shop"
	item.Phylum = "商店"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 15
	item.Enum_value = "sec_activity"
	item.Phylum = "活动"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 16
	item.Enum_value = "sec_recharge"
	item.Phylum = "充值面板（常规充值）"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
	item = &Gm_sec_category_Item{}
	item.Id = 17
	item.Enum_value = "sec_recharge_first"
	item.Phylum = "充值面板（首次赠送）"
	InstGm_sec_category.Items = append(InstGm_sec_category.Items,item)
}
