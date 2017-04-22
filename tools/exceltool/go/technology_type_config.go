package config
type Technology_type struct{
	Items []*Technology_type_Item
}

type Technology_type_Item struct{
	Id int
	Army_type int
}

var InstTechnology_type *Technology_type

func init(){
	item := &Technology_type_Item{}
	item.Id = 1
	item.Army_type = 1
	InstTechnology_type.Items = append(InstTechnology_type.Items,item)
	item = &Technology_type_Item{}
	item.Id = 2
	item.Army_type = 2
	InstTechnology_type.Items = append(InstTechnology_type.Items,item)
	item = &Technology_type_Item{}
	item.Id = 3
	item.Army_type = 3
	InstTechnology_type.Items = append(InstTechnology_type.Items,item)
}
