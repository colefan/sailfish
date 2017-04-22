package config
type Yl_fund struct{
	Items []*Yl_fund_Item
}

type Yl_fund_Item struct{
	Id int
	Need_viplv int
	Fund_price int
	Need_level1 int
	Level1_reward [][]int
	Need_level2 int
	Level2_reward [][]int
	Need_level3 int
	Level3_reward [][]int
	Need_level4 int
	Level4_reward [][]int
	Need_level5 int
	Level5_reward [][]int
	Need_level6 int
	Level6_reward [][]int
	Need_level7 int
	Level7_reward [][]int
	Need_level8 int
	Level8_reward [][]int
}

var InstYl_fund *Yl_fund

func init(){
	item := &Yl_fund_Item{}
	item.Id = 1001
	item.Need_viplv = 2
	item.Fund_price = 1000
	item.Need_level1 = 20
	item.Level1_reward = [][]int{{1,1,300}}
	item.Need_level2 = 30
	item.Level2_reward = [][]int{{1,1,500}}
	item.Need_level3 = 40
	item.Level3_reward = [][]int{{1,1,800}}
	item.Need_level4 = 50
	item.Level4_reward = [][]int{{1,1,1000}}
	item.Need_level5 = 60
	item.Level5_reward = [][]int{{1,1,1300}}
	item.Need_level6 = 65
	item.Level6_reward = [][]int{{1,1,1600}}
	item.Need_level7 = 70
	item.Level7_reward = [][]int{{1,1,2000}}
	item.Need_level8 = 75
	item.Level8_reward = [][]int{{1,1,2500}}
	InstYl_fund.Items = append(InstYl_fund.Items,item)
}
