package config
type Miwu_base struct{
	Items []*Miwu_base_Item
}

type Miwu_base_Item struct{
	Area int
	Npc_id int
	Miwu_city int
}

var InstMiwu_base *Miwu_base

func init(){
	item := &Miwu_base_Item{}
	item.Area = 1001
	item.Npc_id = 9020
	item.Miwu_city = 5001
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 1002
	item.Npc_id = 9028
	item.Miwu_city = 5002
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 1003
	item.Npc_id = 9032
	item.Miwu_city = 5003
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 2001
	item.Npc_id = 9020
	item.Miwu_city = 5004
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 2002
	item.Npc_id = 9028
	item.Miwu_city = 5005
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 2003
	item.Npc_id = 9032
	item.Miwu_city = 5006
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 3001
	item.Npc_id = 9020
	item.Miwu_city = 5007
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 3002
	item.Npc_id = 9028
	item.Miwu_city = 5008
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 3003
	item.Npc_id = 9032
	item.Miwu_city = 5009
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 4001
	item.Npc_id = 9025
	item.Miwu_city = 5010
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 4002
	item.Npc_id = 9035
	item.Miwu_city = 5011
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 4003
	item.Npc_id = 9035
	item.Miwu_city = 5012
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
	item = &Miwu_base_Item{}
	item.Area = 4004
	item.Npc_id = 9035
	item.Miwu_city = 5013
	InstMiwu_base.Items = append(InstMiwu_base.Items,item)
}
