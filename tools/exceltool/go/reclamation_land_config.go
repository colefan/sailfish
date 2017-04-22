package config
type Reclamation_land struct{
	Items []*Reclamation_land_Item
}

type Reclamation_land_Item struct{
	Reclamation_num  int
	Command_lv int
	Cost_silver  int
	Cost_iron  int
}

var InstReclamation_land *Reclamation_land

func init(){
	item := &Reclamation_land_Item{}
	item.Reclamation_num  = 1
	item.Command_lv = 1
	item.Cost_silver  = 0
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 2
	item.Command_lv = 1
	item.Cost_silver  = 0
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 3
	item.Command_lv = 1
	item.Cost_silver  = 3000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 4
	item.Command_lv = 3
	item.Cost_silver  = 30000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 5
	item.Command_lv = 4
	item.Cost_silver  = 30000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 6
	item.Command_lv = 5
	item.Cost_silver  = 60000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 7
	item.Command_lv = 6
	item.Cost_silver  = 60000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 8
	item.Command_lv = 7
	item.Cost_silver  = 90000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 9
	item.Command_lv = 8
	item.Cost_silver  = 90000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 10
	item.Command_lv = 9
	item.Cost_silver  = 90000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 11
	item.Command_lv = 10
	item.Cost_silver  = 120000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 12
	item.Command_lv = 11
	item.Cost_silver  = 120000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 13
	item.Command_lv = 12
	item.Cost_silver  = 120000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 14
	item.Command_lv = 13
	item.Cost_silver  = 150000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 15
	item.Command_lv = 14
	item.Cost_silver  = 150000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 16
	item.Command_lv = 15
	item.Cost_silver  = 150000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 17
	item.Command_lv = 16
	item.Cost_silver  = 180000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 18
	item.Command_lv = 17
	item.Cost_silver  = 180000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
	item = &Reclamation_land_Item{}
	item.Reclamation_num  = 19
	item.Command_lv = 18
	item.Cost_silver  = 180000
	item.Cost_iron  = 0
	InstReclamation_land.Items = append(InstReclamation_land.Items,item)
}
