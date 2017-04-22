package config
type Legion_upgrade struct{
	Items []*Legion_upgrade_Item
}

type Legion_upgrade_Item struct{
	Legion_level int
	Member_max int
	Upgrade_pay int
	Spot_max int
	Legion_technology_max int
}

var InstLegion_upgrade *Legion_upgrade

func init(){
	item := &Legion_upgrade_Item{}
	item.Legion_level = 1
	item.Member_max = 25
	item.Upgrade_pay = 10000
	item.Spot_max = 25
	item.Legion_technology_max = 2
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 2
	item.Member_max = 26
	item.Upgrade_pay = 400000
	item.Spot_max = 30
	item.Legion_technology_max = 4
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 3
	item.Member_max = 26
	item.Upgrade_pay = 500000
	item.Spot_max = 35
	item.Legion_technology_max = 6
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 4
	item.Member_max = 27
	item.Upgrade_pay = 900000
	item.Spot_max = 40
	item.Legion_technology_max = 8
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 5
	item.Member_max = 27
	item.Upgrade_pay = 1300000
	item.Spot_max = 45
	item.Legion_technology_max = 10
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 6
	item.Member_max = 28
	item.Upgrade_pay = 1700000
	item.Spot_max = 50
	item.Legion_technology_max = 12
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 7
	item.Member_max = 28
	item.Upgrade_pay = 2100000
	item.Spot_max = 55
	item.Legion_technology_max = 14
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 8
	item.Member_max = 29
	item.Upgrade_pay = 2800000
	item.Spot_max = 60
	item.Legion_technology_max = 16
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 9
	item.Member_max = 29
	item.Upgrade_pay = 3500000
	item.Spot_max = 63
	item.Legion_technology_max = 18
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 10
	item.Member_max = 30
	item.Upgrade_pay = 4200000
	item.Spot_max = 66
	item.Legion_technology_max = 20
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 11
	item.Member_max = 30
	item.Upgrade_pay = 4900000
	item.Spot_max = 69
	item.Legion_technology_max = 22
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 12
	item.Member_max = 31
	item.Upgrade_pay = 5600000
	item.Spot_max = 72
	item.Legion_technology_max = 24
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 13
	item.Member_max = 31
	item.Upgrade_pay = 6300000
	item.Spot_max = 75
	item.Legion_technology_max = 26
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 14
	item.Member_max = 32
	item.Upgrade_pay = 7000000
	item.Spot_max = 78
	item.Legion_technology_max = 28
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 15
	item.Member_max = 32
	item.Upgrade_pay = 8800000
	item.Spot_max = 81
	item.Legion_technology_max = 30
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 16
	item.Member_max = 33
	item.Upgrade_pay = 9900000
	item.Spot_max = 84
	item.Legion_technology_max = 32
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 17
	item.Member_max = 33
	item.Upgrade_pay = 10800000
	item.Spot_max = 87
	item.Legion_technology_max = 34
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 18
	item.Member_max = 34
	item.Upgrade_pay = 16800000
	item.Spot_max = 90
	item.Legion_technology_max = 36
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 19
	item.Member_max = 34
	item.Upgrade_pay = 24000000
	item.Spot_max = 93
	item.Legion_technology_max = 38
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
	item = &Legion_upgrade_Item{}
	item.Legion_level = 20
	item.Member_max = 35
	item.Upgrade_pay = 0
	item.Spot_max = 96
	item.Legion_technology_max = 40
	InstLegion_upgrade.Items = append(InstLegion_upgrade.Items,item)
}
