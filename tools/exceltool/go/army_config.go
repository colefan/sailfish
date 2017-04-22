package config
type Army struct{
	Items []*Army_Item
}

type Army_Item struct{
	Id int
	Skill []int
	Bullet_index int
	Army_type int
	Type int
	Type_2 int
	Type_3 int
	Type_4 int
	Move int
	Range int
}

var InstArmy *Army

func init(){
	item := &Army_Item{}
	item.Id = 1002
	item.Skill = []int{1023,1026,1022,1024}
	item.Bullet_index = 10000
	item.Army_type = 1
	item.Type = 1
	item.Type_2 = 11
	item.Type_3 = 22
	item.Type_4 = 32
	item.Move = 2
	item.Range = 1
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1004
	item.Skill = []int{1004,1001,1002,1003}
	item.Bullet_index = 9700
	item.Army_type = 1
	item.Type = 1
	item.Type_2 = 11
	item.Type_3 = 22
	item.Type_4 = 32
	item.Move = 2
	item.Range = 1
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1005
	item.Skill = []int{1009,1011,1010,1012}
	item.Bullet_index = 9300
	item.Army_type = 2
	item.Type = 2
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 32
	item.Move = 2
	item.Range = 2
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1006
	item.Skill = []int{1005,1008,1006,1007}
	item.Bullet_index = 9500
	item.Army_type = 2
	item.Type = 2
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 32
	item.Move = 1
	item.Range = 3
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1008
	item.Skill = []int{1015,1013,1014,1016}
	item.Bullet_index = 10300
	item.Army_type = 3
	item.Type = 2
	item.Type_2 = 12
	item.Type_3 = 21
	item.Type_4 = 31
	item.Move = 2
	item.Range = 2
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1009
	item.Skill = []int{1017,1019,1018,1020}
	item.Bullet_index = 10700
	item.Army_type = 3
	item.Type = 1
	item.Type_2 = 12
	item.Type_3 = 22
	item.Type_4 = 31
	item.Move = 3
	item.Range = 1
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1011
	item.Skill = []int{1025}
	item.Bullet_index = 10000
	item.Army_type = 0
	item.Type = 1
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 31
	item.Move = 0
	item.Range = 0
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1012
	item.Skill = make([]int,0)
	item.Bullet_index = 10000
	item.Army_type = 0
	item.Type = 2
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 31
	item.Move = 0
	item.Range = 10
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1013
	item.Skill = make([]int,0)
	item.Bullet_index = 10000
	item.Army_type = 0
	item.Type = 1
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 31
	item.Move = 0
	item.Range = 0
	InstArmy.Items = append(InstArmy.Items,item)
	item = &Army_Item{}
	item.Id = 1014
	item.Skill = make([]int,0)
	item.Bullet_index = 10000
	item.Army_type = 0
	item.Type = 2
	item.Type_2 = 11
	item.Type_3 = 21
	item.Type_4 = 31
	item.Move = 0
	item.Range = 10
	InstArmy.Items = append(InstArmy.Items,item)
}
