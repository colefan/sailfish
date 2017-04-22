package config
type Tujianliebiao struct{
	Items []*Tujianliebiao_Item
}

type Tujianliebiao_Item struct{
	Id int
	Level int
	Hero_id [][]int
}

var InstTujianliebiao *Tujianliebiao

func init(){
	item := &Tujianliebiao_Item{}
	item.Id = 1001
	item.Level = 30
	item.Hero_id = [][]int{{10042},{10016}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1002
	item.Level = 30
	item.Hero_id = [][]int{{10020},{10023}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1003
	item.Level = 30
	item.Hero_id = [][]int{{10017},{10049}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1004
	item.Level = 30
	item.Hero_id = [][]int{{10021},{10009},{10018}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1005
	item.Level = 30
	item.Hero_id = [][]int{{10022},{10003},{10037}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1006
	item.Level = 30
	item.Hero_id = [][]int{{10002},{10045},{10006}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1007
	item.Level = 30
	item.Hero_id = [][]int{{10017},{10036},{10049},{10019}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1008
	item.Level = 30
	item.Hero_id = [][]int{{10007},{10023},{10049},{10042}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1009
	item.Level = 30
	item.Hero_id = [][]int{{10002},{10013},{10036},{10029}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1010
	item.Level = 30
	item.Hero_id = [][]int{{10027},{10026},{10019},{10005}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
	item = &Tujianliebiao_Item{}
	item.Id = 1011
	item.Level = 30
	item.Hero_id = [][]int{{10007},{10039},{10043},{10012}}
	InstTujianliebiao.Items = append(InstTujianliebiao.Items,item)
}
