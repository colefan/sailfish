package config
type Miwu_task struct{
	Items []*Miwu_task_Item
}

type Miwu_task_Item struct{
	Num int
	Reward [][]int
}

var InstMiwu_task *Miwu_task

func init(){
	item := &Miwu_task_Item{}
	item.Num = 1
	item.Reward = [][]int{{201,406002,3},{2,2,10000},{201,431001,10}}
	InstMiwu_task.Items = append(InstMiwu_task.Items,item)
	item = &Miwu_task_Item{}
	item.Num = 3
	item.Reward = [][]int{{201,406002,5},{201,431002,10},{201,405002,5}}
	InstMiwu_task.Items = append(InstMiwu_task.Items,item)
	item = &Miwu_task_Item{}
	item.Num = 6
	item.Reward = [][]int{{201,432003,2},{5,5,200},{2,2,20000}}
	InstMiwu_task.Items = append(InstMiwu_task.Items,item)
	item = &Miwu_task_Item{}
	item.Num = 9
	item.Reward = [][]int{{201,406003,3},{5,5,300},{201,435001,10}}
	InstMiwu_task.Items = append(InstMiwu_task.Items,item)
	item = &Miwu_task_Item{}
	item.Num = 13
	item.Reward = [][]int{{201,435002,5},{5,5,500},{2,2,50000}}
	InstMiwu_task.Items = append(InstMiwu_task.Items,item)
}
