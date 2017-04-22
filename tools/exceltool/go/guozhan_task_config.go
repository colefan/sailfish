package config
type Guozhan_task struct{
	Items []*Guozhan_task_Item
}

type Guozhan_task_Item struct{
	Type int
	Chance int
	Exp int
}

var InstGuozhan_task *Guozhan_task

func init(){
	item := &Guozhan_task_Item{}
	item.Type = 1
	item.Chance = 100
	item.Exp = 1250
	InstGuozhan_task.Items = append(InstGuozhan_task.Items,item)
	item = &Guozhan_task_Item{}
	item.Type = 2
	item.Chance = 0
	item.Exp = 500
	InstGuozhan_task.Items = append(InstGuozhan_task.Items,item)
	item = &Guozhan_task_Item{}
	item.Type = 3
	item.Chance = 0
	item.Exp = 500
	InstGuozhan_task.Items = append(InstGuozhan_task.Items,item)
	item = &Guozhan_task_Item{}
	item.Type = 4
	item.Chance = 0
	item.Exp = 500
	InstGuozhan_task.Items = append(InstGuozhan_task.Items,item)
}
