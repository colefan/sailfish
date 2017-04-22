package config
type Legion_flag struct{
	Items []*Legion_flag_Item
}

type Legion_flag_Item struct{
	Id int
	Type int
	Flag string
}

var InstLegion_flag *Legion_flag

func init(){
	item := &Legion_flag_Item{}
	item.Id = 101
	item.Type = 1
	item.Flag = "flag_diwen_01"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 102
	item.Type = 1
	item.Flag = "flag_diwen_02"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 103
	item.Type = 1
	item.Flag = "flag_diwen_03"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 104
	item.Type = 1
	item.Flag = "flag_diwen_04"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 105
	item.Type = 1
	item.Flag = "flag_diwen_05"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 201
	item.Type = 2
	item.Flag = "flag_xingzhuang_01"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 202
	item.Type = 2
	item.Flag = "flag_xingzhuang_02"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 203
	item.Type = 2
	item.Flag = "flag_xingzhuang_03"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 204
	item.Type = 2
	item.Flag = "flag_xingzhuang_04"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 205
	item.Type = 2
	item.Flag = "flag_xingzhuang_05"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 301
	item.Type = 3
	item.Flag = "flag_huizhang_01"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 302
	item.Type = 3
	item.Flag = "flag_huizhang_02"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 303
	item.Type = 3
	item.Flag = "flag_huizhang_03"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 304
	item.Type = 3
	item.Flag = "flag_huizhang_04"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
	item = &Legion_flag_Item{}
	item.Id = 305
	item.Type = 3
	item.Flag = "flag_huizhang_05"
	InstLegion_flag.Items = append(InstLegion_flag.Items,item)
}
