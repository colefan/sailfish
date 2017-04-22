package config
type Lvup_require struct{
	Items []*Lvup_require_Item
}

type Lvup_require_Item struct{
	Id int
	Require [][]int
	Cost [][]int
}

var InstLvup_require *Lvup_require

func init(){
	item := &Lvup_require_Item{}
	item.Id = 1
	item.Require = [][]int{{6,0}}
	item.Cost = [][]int{{2,2,5000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 2
	item.Require = [][]int{{1,2,1},{2,2}}
	item.Cost = [][]int{{2,2,10000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 3
	item.Require = [][]int{{1,3,1}}
	item.Cost = [][]int{{2,2,20000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 4
	item.Require = [][]int{{2,3},{3,6}}
	item.Cost = [][]int{{2,2,30000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 5
	item.Require = [][]int{{1,4,1},{2,4}}
	item.Cost = [][]int{{2,2,50000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 6
	item.Require = [][]int{{2,5},{3,10}}
	item.Cost = [][]int{{2,2,80000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 7
	item.Require = [][]int{{1,4,4},{4,310}}
	item.Cost = [][]int{{2,2,100000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
	item = &Lvup_require_Item{}
	item.Id = 8
	item.Require = [][]int{{5,80},{1,5,1}}
	item.Cost = [][]int{{2,2,150000}}
	InstLvup_require.Items = append(InstLvup_require.Items,item)
}
