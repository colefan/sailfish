package config
type Huodong_multiple struct{
	Items []*Huodong_multiple_Item
}

type Huodong_multiple_Item struct{
	Module_id int
	Multiple int
	Type int
	Value_1  int
	Value_2 int
	Value_3 int
	Value_4 int
	Value_5 int
	Value_6 int
	Value_7 int
}

var InstHuodong_multiple *Huodong_multiple

func init(){
	item := &Huodong_multiple_Item{}
	item.Module_id = 82000
	item.Multiple = 2
	item.Type = 3
	item.Value_1  = 3
	item.Value_2 = 6
	item.Value_3 = 0
	item.Value_4 = 0
	item.Value_5 = 0
	item.Value_6 = 0
	item.Value_7 = 0
	InstHuodong_multiple.Items = append(InstHuodong_multiple.Items,item)
	item = &Huodong_multiple_Item{}
	item.Module_id = 83000
	item.Multiple = 2
	item.Type = 3
	item.Value_1  = 0
	item.Value_2 = 0
	item.Value_3 = 0
	item.Value_4 = 0
	item.Value_5 = 0
	item.Value_6 = 0
	item.Value_7 = 0
	InstHuodong_multiple.Items = append(InstHuodong_multiple.Items,item)
	item = &Huodong_multiple_Item{}
	item.Module_id = 55001
	item.Multiple = 2
	item.Type = 3
	item.Value_1  = 1
	item.Value_2 = 4
	item.Value_3 = 0
	item.Value_4 = 0
	item.Value_5 = 0
	item.Value_6 = 0
	item.Value_7 = 0
	InstHuodong_multiple.Items = append(InstHuodong_multiple.Items,item)
	item = &Huodong_multiple_Item{}
	item.Module_id = 56000
	item.Multiple = 2
	item.Type = 3
	item.Value_1  = 7
	item.Value_2 = 0
	item.Value_3 = 0
	item.Value_4 = 0
	item.Value_5 = 0
	item.Value_6 = 0
	item.Value_7 = 0
	InstHuodong_multiple.Items = append(InstHuodong_multiple.Items,item)
	item = &Huodong_multiple_Item{}
	item.Module_id = 31400
	item.Multiple = 2
	item.Type = 3
	item.Value_1  = 2
	item.Value_2 = 5
	item.Value_3 = 0
	item.Value_4 = 0
	item.Value_5 = 0
	item.Value_6 = 0
	item.Value_7 = 0
	InstHuodong_multiple.Items = append(InstHuodong_multiple.Items,item)
}
