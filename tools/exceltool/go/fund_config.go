package config
type Fund struct{
	Items []*Fund_Item
}

type Fund_Item struct{
	Fund_id int
	Fund_name string
	Need_viplv int
	Fund_price int
	Time_limit int
	Day1_reward [][]int
	Day2_reward [][]int
	Day3_reward [][]int
	Day4_reward [][]int
	Day5_reward [][]int
	Day6_reward [][]int
	Day7_reward [][]int
}

var InstFund *Fund

func init(){
	item := &Fund_Item{}
	item.Fund_id = 1001
	item.Fund_name = "text_V1_fund"
	item.Need_viplv = 1
	item.Fund_price = 500
	item.Time_limit = 3
	item.Day1_reward = [][]int{{1,1,300}}
	item.Day2_reward = [][]int{{1,1,200}}
	item.Day3_reward = [][]int{{1,1,200}}
	item.Day4_reward = [][]int{{1,1,200}}
	item.Day5_reward = [][]int{{1,1,200}}
	item.Day6_reward = [][]int{{1,1,200}}
	item.Day7_reward = [][]int{{1,1,200},{202,10019,1}}
	InstFund.Items = append(InstFund.Items,item)
	item = &Fund_Item{}
	item.Fund_id = 1002
	item.Fund_name = "text_V5_fund"
	item.Need_viplv = 5
	item.Fund_price = 1500
	item.Time_limit = 3
	item.Day1_reward = [][]int{{1,1,900}}
	item.Day2_reward = [][]int{{1,1,600}}
	item.Day3_reward = [][]int{{1,1,600}}
	item.Day4_reward = [][]int{{1,1,600}}
	item.Day5_reward = [][]int{{1,1,600}}
	item.Day6_reward = [][]int{{1,1,600}}
	item.Day7_reward = [][]int{{1,1,600},{202,10005,1}}
	InstFund.Items = append(InstFund.Items,item)
}
