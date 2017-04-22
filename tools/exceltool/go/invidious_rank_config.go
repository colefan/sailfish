package config
type Invidious_rank struct{
	Items []*Invidious_rank_Item
}

type Invidious_rank_Item struct{
	Invidious_rank int
	Award [][]int
}

var InstInvidious_rank *Invidious_rank

func init(){
	item := &Invidious_rank_Item{}
	item.Invidious_rank = 1
	item.Award = [][]int{{2,2,600000},{6,6,1000}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 2
	item.Award = [][]int{{2,2,580000},{6,6,900}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 3
	item.Award = [][]int{{2,2,560000},{6,6,800}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 4
	item.Award = [][]int{{2,2,540000},{6,6,700}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 5
	item.Award = [][]int{{2,2,520000},{6,6,650}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 6
	item.Award = [][]int{{2,2,500000},{6,6,600}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 7
	item.Award = [][]int{{2,2,480000},{6,6,550}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 8
	item.Award = [][]int{{2,2,460000},{6,6,500}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 9
	item.Award = [][]int{{2,2,440000},{6,6,450}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 10
	item.Award = [][]int{{2,2,420000},{6,6,400}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 11
	item.Award = [][]int{{2,2,400000},{6,6,350}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
	item = &Invidious_rank_Item{}
	item.Invidious_rank = 12
	item.Award = [][]int{{2,2,380000},{6,6,300}}
	InstInvidious_rank.Items = append(InstInvidious_rank.Items,item)
}
