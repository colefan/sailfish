package config
type Boss_damagereward struct{
	Items []*Boss_damagereward_Item
}

type Boss_damagereward_Item struct{
	Id int
	Lv_min int
	Lv_max int
	Damage int
	Reward [][]int
}

var InstBoss_damagereward *Boss_damagereward

func init(){
	item := &Boss_damagereward_Item{}
	item.Id = 1001
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 200000
	item.Reward = [][]int{{2,2,5000},{18,18,5000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1002
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 900000
	item.Reward = [][]int{{2,2,7000},{18,18,7000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1003
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 1900000
	item.Reward = [][]int{{2,2,9000},{18,18,9000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1004
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 2900000
	item.Reward = [][]int{{2,2,9000},{18,18,9000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1005
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 3900000
	item.Reward = [][]int{{2,2,10000},{18,18,10000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1006
	item.Lv_min = 30
	item.Lv_max = 34
	item.Damage = 4900000
	item.Reward = [][]int{{2,2,10000},{18,18,10000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1007
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 300000
	item.Reward = [][]int{{2,2,7000},{18,18,7000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1008
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 1100000
	item.Reward = [][]int{{2,2,10000},{18,18,10000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1009
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 2200000
	item.Reward = [][]int{{2,2,13000},{18,18,13000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1010
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 3400000
	item.Reward = [][]int{{2,2,13000},{18,18,13000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1011
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 4500000
	item.Reward = [][]int{{2,2,14000},{18,18,14000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1012
	item.Lv_min = 35
	item.Lv_max = 39
	item.Damage = 5700000
	item.Reward = [][]int{{2,2,14000},{18,18,14000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1013
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 300000
	item.Reward = [][]int{{2,2,9000},{18,18,9000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1014
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 1300000
	item.Reward = [][]int{{2,2,13000},{18,18,13000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1015
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 2600000
	item.Reward = [][]int{{2,2,16000},{18,18,16000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1016
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 3900000
	item.Reward = [][]int{{2,2,16000},{18,18,16000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1017
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 5300000
	item.Reward = [][]int{{2,2,18000},{18,18,18000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1018
	item.Lv_min = 40
	item.Lv_max = 44
	item.Damage = 6600000
	item.Reward = [][]int{{2,2,18000},{18,18,18000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1019
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 400000
	item.Reward = [][]int{{2,2,12000},{18,18,12000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1020
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 1600000
	item.Reward = [][]int{{2,2,17000},{18,18,17000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1021
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 3200000
	item.Reward = [][]int{{2,2,22000},{18,18,22000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1022
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 4800000
	item.Reward = [][]int{{2,2,22000},{18,18,22000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1023
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 6400000
	item.Reward = [][]int{{2,2,24000},{18,18,24000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1024
	item.Lv_min = 45
	item.Lv_max = 49
	item.Damage = 8100000
	item.Reward = [][]int{{2,2,24000},{18,18,24000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1025
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 700000
	item.Reward = [][]int{{2,2,15000},{18,18,15000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1026
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 2300000
	item.Reward = [][]int{{2,2,21000},{18,18,21000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1027
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 4700000
	item.Reward = [][]int{{2,2,27000},{18,18,27000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1028
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 7100000
	item.Reward = [][]int{{2,2,27000},{18,18,27000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1029
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 9500000
	item.Reward = [][]int{{2,2,30000},{18,18,30000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1030
	item.Lv_min = 50
	item.Lv_max = 54
	item.Damage = 11900000
	item.Reward = [][]int{{2,2,30000},{18,18,30000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1031
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 800000
	item.Reward = [][]int{{2,2,18000},{18,18,18000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1032
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 2900000
	item.Reward = [][]int{{2,2,25000},{18,18,25000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1033
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 5900000
	item.Reward = [][]int{{2,2,32000},{18,18,32000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1034
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 8800000
	item.Reward = [][]int{{2,2,32000},{18,18,32000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1035
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 11800000
	item.Reward = [][]int{{2,2,36000},{18,18,36000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1036
	item.Lv_min = 55
	item.Lv_max = 59
	item.Damage = 14700000
	item.Reward = [][]int{{2,2,36000},{18,18,36000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1037
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 1000000
	item.Reward = [][]int{{2,2,21000},{18,18,21000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1038
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 3600000
	item.Reward = [][]int{{2,2,29000},{18,18,29000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1039
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 7200000
	item.Reward = [][]int{{2,2,38000},{18,18,38000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1040
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 10800000
	item.Reward = [][]int{{2,2,38000},{18,18,38000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1041
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 14400000
	item.Reward = [][]int{{2,2,42000},{18,18,42000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1042
	item.Lv_min = 60
	item.Lv_max = 64
	item.Damage = 18000000
	item.Reward = [][]int{{2,2,42000},{18,18,42000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1043
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 1200000
	item.Reward = [][]int{{2,2,24000},{18,18,24000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1044
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 4200000
	item.Reward = [][]int{{2,2,34000},{18,18,34000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1045
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 8500000
	item.Reward = [][]int{{2,2,43000},{18,18,43000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1046
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 12800000
	item.Reward = [][]int{{2,2,43000},{18,18,43000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1047
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 17100000
	item.Reward = [][]int{{2,2,48000},{18,18,48000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1048
	item.Lv_min = 65
	item.Lv_max = 69
	item.Damage = 21400000
	item.Reward = [][]int{{2,2,48000},{18,18,48000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1049
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 1500000
	item.Reward = [][]int{{2,2,27000},{18,18,27000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1050
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 5000000
	item.Reward = [][]int{{2,2,38000},{18,18,38000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1051
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 10000000
	item.Reward = [][]int{{2,2,49000},{18,18,49000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1052
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 15000000
	item.Reward = [][]int{{2,2,49000},{18,18,49000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1053
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 20100000
	item.Reward = [][]int{{2,2,54000},{18,18,54000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1054
	item.Lv_min = 70
	item.Lv_max = 74
	item.Damage = 25100000
	item.Reward = [][]int{{2,2,54000},{18,18,54000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1055
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 2100000
	item.Reward = [][]int{{2,2,30000},{18,18,30000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1056
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 7100000
	item.Reward = [][]int{{2,2,42000},{18,18,42000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1057
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 14300000
	item.Reward = [][]int{{2,2,54000},{18,18,54000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1058
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 21500000
	item.Reward = [][]int{{2,2,54000},{18,18,54000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1059
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 28700000
	item.Reward = [][]int{{2,2,60000},{18,18,60000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1060
	item.Lv_min = 75
	item.Lv_max = 79
	item.Damage = 35900000
	item.Reward = [][]int{{2,2,60000},{18,18,60000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1061
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 2400000
	item.Reward = [][]int{{2,2,33000},{18,18,33000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1062
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 8300000
	item.Reward = [][]int{{2,2,46000},{18,18,46000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1063
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 16600000
	item.Reward = [][]int{{2,2,59000},{18,18,59000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1064
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 24900000
	item.Reward = [][]int{{2,2,59000},{18,18,59000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1065
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 33000000
	item.Reward = [][]int{{2,2,66000},{18,18,66000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1066
	item.Lv_min = 80
	item.Lv_max = 84
	item.Damage = 41000000
	item.Reward = [][]int{{2,2,66000},{18,18,66000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1067
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 2000000
	item.Reward = [][]int{{2,2,36000},{18,18,36000},{21,21,2}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1068
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 9000000
	item.Reward = [][]int{{2,2,50000},{18,18,50000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1069
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 18000000
	item.Reward = [][]int{{2,2,65000},{18,18,65000},{21,21,3}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1070
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 28000000
	item.Reward = [][]int{{2,2,65000},{18,18,65000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1071
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 37000000
	item.Reward = [][]int{{2,2,72000},{18,18,72000},{21,21,4}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
	item = &Boss_damagereward_Item{}
	item.Id = 1072
	item.Lv_min = 85
	item.Lv_max = 89
	item.Damage = 47000000
	item.Reward = [][]int{{2,3,72000},{18,18,72000},{21,21,5}}
	InstBoss_damagereward.Items = append(InstBoss_damagereward.Items,item)
}
