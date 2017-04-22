package config
type Arena_prize struct{
	Items []*Arena_prize_Item
}

type Arena_prize_Item struct{
	Id int
	Win [][]int
	Lose [][]int
}

var InstArena_prize *Arena_prize

func init(){
	item := &Arena_prize_Item{}
	item.Id = 1
	item.Win = [][]int{{8,8,50}}
	item.Lose = [][]int{{8,8,25}}
	InstArena_prize.Items = append(InstArena_prize.Items,item)
	item = &Arena_prize_Item{}
	item.Id = 2
	item.Win = [][]int{{12,12,50}}
	item.Lose = [][]int{{12,12,25}}
	InstArena_prize.Items = append(InstArena_prize.Items,item)
}
