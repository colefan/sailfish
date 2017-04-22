package config
type Legion_rank struct{
	Items []*Legion_rank_Item
}

type Legion_rank_Item struct{
	Legion_rank int
	Award [][]int
}

var InstLegion_rank *Legion_rank

func init(){
	item := &Legion_rank_Item{}
	item.Legion_rank = 1
	item.Award = [][]int{{6,6,1000},{204,204,300000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 2
	item.Award = [][]int{{6,6,900},{204,204,280000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 3
	item.Award = [][]int{{6,6,800},{204,204,260000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 4
	item.Award = [][]int{{6,6,780},{204,204,240000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 5
	item.Award = [][]int{{6,6,760},{204,204,220000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 6
	item.Award = [][]int{{6,6,740},{204,204,200000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 7
	item.Award = [][]int{{6,6,720},{204,204,180000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 8
	item.Award = [][]int{{6,6,700},{204,204,160000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 9
	item.Award = [][]int{{6,6,680},{204,204,140000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 10
	item.Award = [][]int{{6,6,660},{204,204,120000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 11
	item.Award = [][]int{{6,6,640},{204,204,100000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 12
	item.Award = [][]int{{6,6,620},{204,204,80000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
	item = &Legion_rank_Item{}
	item.Legion_rank = 13
	item.Award = [][]int{{6,6,600},{204,204,60000}}
	InstLegion_rank.Items = append(InstLegion_rank.Items,item)
}
