package config
type Zhengdianlingjiang struct{
	Items []*Zhengdianlingjiang_Item
}

type Zhengdianlingjiang_Item struct{
	ID int
	Level_limit int
	Level_max int
	Open_time int
	End_time int
	Rewards [][]int
}

var InstZhengdianlingjiang *Zhengdianlingjiang

func init(){
	item := &Zhengdianlingjiang_Item{}
	item.ID = 1
	item.Level_limit = 28
	item.Level_max = 34
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,27000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 13
	item.Level_limit = 35
	item.Level_max = 44
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,44000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 2
	item.Level_limit = 45
	item.Level_max = 54
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,80000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 3
	item.Level_limit = 55
	item.Level_max = 64
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,120000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 4
	item.Level_limit = 65
	item.Level_max = 74
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,160000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 5
	item.Level_limit = 75
	item.Level_max = 84
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,210000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 6
	item.Level_limit = 85
	item.Level_max = 100
	item.Open_time = 1230
	item.End_time = 1300
	item.Rewards = [][]int{{18,18,260000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 7
	item.Level_limit = 28
	item.Level_max = 34
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,27000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 14
	item.Level_limit = 35
	item.Level_max = 44
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,44000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 8
	item.Level_limit = 45
	item.Level_max = 54
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,80000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 9
	item.Level_limit = 55
	item.Level_max = 64
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,120000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 10
	item.Level_limit = 65
	item.Level_max = 74
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,160000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 11
	item.Level_limit = 75
	item.Level_max = 84
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,210000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
	item = &Zhengdianlingjiang_Item{}
	item.ID = 12
	item.Level_limit = 85
	item.Level_max = 100
	item.Open_time = 1930
	item.End_time = 2030
	item.Rewards = [][]int{{18,18,260000}}
	InstZhengdianlingjiang.Items = append(InstZhengdianlingjiang.Items,item)
}
