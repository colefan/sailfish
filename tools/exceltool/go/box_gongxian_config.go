package config
type Box_gongxian struct{
	Items []*Box_gongxian_Item
}

type Box_gongxian_Item struct{
	Id int
	Box_num int
	Need_gongxian int
}

var InstBox_gongxian *Box_gongxian

func init(){
	item := &Box_gongxian_Item{}
	item.Id = 101
	item.Box_num = 2
	item.Need_gongxian = 75
	InstBox_gongxian.Items = append(InstBox_gongxian.Items,item)
	item = &Box_gongxian_Item{}
	item.Id = 201
	item.Box_num = 3
	item.Need_gongxian = 300
	InstBox_gongxian.Items = append(InstBox_gongxian.Items,item)
	item = &Box_gongxian_Item{}
	item.Id = 301
	item.Box_num = 3
	item.Need_gongxian = 600
	InstBox_gongxian.Items = append(InstBox_gongxian.Items,item)
	item = &Box_gongxian_Item{}
	item.Id = 401
	item.Box_num = 3
	item.Need_gongxian = 900
	InstBox_gongxian.Items = append(InstBox_gongxian.Items,item)
	item = &Box_gongxian_Item{}
	item.Id = 501
	item.Box_num = 5
	item.Need_gongxian = 1300
	InstBox_gongxian.Items = append(InstBox_gongxian.Items,item)
}
