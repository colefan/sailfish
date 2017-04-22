package config
type Drilling_award struct{
	Items []*Drilling_award_Item
}

type Drilling_award_Item struct{
	Id int
	Min int
	Max int
	Hour_reward [][]int
}

var InstDrilling_award *Drilling_award

func init(){
	item := &Drilling_award_Item{}
	item.Id = 1
	item.Min = 1
	item.Max = 1
	item.Hour_reward = [][]int{{12,12,100}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 2
	item.Min = 2
	item.Max = 2
	item.Hour_reward = [][]int{{12,12,90}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 3
	item.Min = 3
	item.Max = 3
	item.Hour_reward = [][]int{{12,12,80}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 4
	item.Min = 4
	item.Max = 4
	item.Hour_reward = [][]int{{12,12,78}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 5
	item.Min = 5
	item.Max = 5
	item.Hour_reward = [][]int{{12,12,76}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 6
	item.Min = 6
	item.Max = 6
	item.Hour_reward = [][]int{{12,12,74}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 7
	item.Min = 7
	item.Max = 7
	item.Hour_reward = [][]int{{12,12,72}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 8
	item.Min = 8
	item.Max = 8
	item.Hour_reward = [][]int{{12,12,70}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 9
	item.Min = 9
	item.Max = 9
	item.Hour_reward = [][]int{{12,12,68}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 10
	item.Min = 10
	item.Max = 10
	item.Hour_reward = [][]int{{12,12,66}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 11
	item.Min = 11
	item.Max = 20
	item.Hour_reward = [][]int{{12,12,63}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 12
	item.Min = 21
	item.Max = 30
	item.Hour_reward = [][]int{{12,12,60}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 13
	item.Min = 31
	item.Max = 40
	item.Hour_reward = [][]int{{12,12,55}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 14
	item.Min = 41
	item.Max = 50
	item.Hour_reward = [][]int{{12,12,50}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 15
	item.Min = 51
	item.Max = 70
	item.Hour_reward = [][]int{{12,12,45}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 16
	item.Min = 71
	item.Max = 100
	item.Hour_reward = [][]int{{12,12,40}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 17
	item.Min = 101
	item.Max = 200
	item.Hour_reward = [][]int{{12,12,35}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 18
	item.Min = 201
	item.Max = 300
	item.Hour_reward = [][]int{{12,12,25}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 19
	item.Min = 301
	item.Max = 400
	item.Hour_reward = [][]int{{12,12,20}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 20
	item.Min = 401
	item.Max = 500
	item.Hour_reward = [][]int{{12,12,17}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 21
	item.Min = 501
	item.Max = 700
	item.Hour_reward = [][]int{{12,12,15}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 22
	item.Min = 701
	item.Max = 1000
	item.Hour_reward = [][]int{{12,12,12}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 23
	item.Min = 1001
	item.Max = 2000
	item.Hour_reward = [][]int{{12,12,9}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 24
	item.Min = 2001
	item.Max = 3000
	item.Hour_reward = [][]int{{12,12,7}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 25
	item.Min = 3001
	item.Max = 4000
	item.Hour_reward = [][]int{{12,12,5}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 26
	item.Min = 4001
	item.Max = 5000
	item.Hour_reward = [][]int{{12,12,3}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 27
	item.Min = 5001
	item.Max = 7000
	item.Hour_reward = [][]int{{12,12,2}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 28
	item.Min = 7001
	item.Max = 9999
	item.Hour_reward = [][]int{{12,12,1}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
	item = &Drilling_award_Item{}
	item.Id = 29
	item.Min = 10000
	item.Max = 100000
	item.Hour_reward = [][]int{{12,12,0}}
	InstDrilling_award.Items = append(InstDrilling_award.Items,item)
}
