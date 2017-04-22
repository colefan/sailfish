package config
type Specialty struct{
	Items []*Specialty_Item
}

type Specialty_Item struct{
	Id int
	Level_min int
	Level_max int
	Specialty [][]int
}

var InstSpecialty *Specialty

func init(){
	item := &Specialty_Item{}
	item.Id = 101
	item.Level_min = 1
	item.Level_max = 39
	item.Specialty = [][]int{{2,2,15000},{2,2,20000},{18,18,12000},{18,18,18000},{201,431001,2},{201,431001,3},{201,431002,2},{201,431002,3},{1,1,20},{1,1,30},{201,406002,2},{201,310045,1},{201,310049,1},{201,432004,2},{201,432004,3}}
	InstSpecialty.Items = append(InstSpecialty.Items,item)
	item = &Specialty_Item{}
	item.Id = 201
	item.Level_min = 40
	item.Level_max = 54
	item.Specialty = [][]int{{2,2,20000},{2,2,30000},{18,18,15000},{18,18,21000},{201,431001,2},{201,431001,3},{201,431002,2},{201,431002,3},{1,1,20},{1,1,30},{201,406003,2},{201,310045,1},{201,310049,1},{201,432004,2},{201,432004,4},{21,21,4},{17,17,5},{17,17,5},{201,435001,2},{201,405002,4}}
	InstSpecialty.Items = append(InstSpecialty.Items,item)
	item = &Specialty_Item{}
	item.Id = 301
	item.Level_min = 55
	item.Level_max = 69
	item.Specialty = [][]int{{2,2,35000},{2,2,50000},{18,18,20000},{18,18,30000},{201,431001,2},{201,431001,3},{201,431002,2},{201,431002,3},{1,1,40},{1,1,60},{201,406004,2},{201,310045,1},{201,310049,1},{201,432004,2},{201,432004,4},{21,21,4},{17,17,5},{17,17,5},{201,435001,2},{201,405002,8}}
	InstSpecialty.Items = append(InstSpecialty.Items,item)
	item = &Specialty_Item{}
	item.Id = 401
	item.Level_min = 70
	item.Level_max = 79
	item.Specialty = [][]int{{2,2,48000},{2,2,64000},{18,18,28000},{18,18,36000},{201,431001,2},{201,431001,3},{201,431002,2},{201,431002,3},{1,1,40},{1,1,60},{201,406004,2},{201,310045,1},{201,310049,1},{201,432004,2},{201,432004,4},{21,21,4},{17,17,5},{17,17,5},{201,435001,4},{201,405003,3}}
	InstSpecialty.Items = append(InstSpecialty.Items,item)
	item = &Specialty_Item{}
	item.Id = 501
	item.Level_min = 80
	item.Level_max = 89
	item.Specialty = [][]int{{2,2,60000},{2,2,80000},{18,18,40000},{18,18,56000},{201,431001,2},{201,431001,3},{201,431002,2},{201,431002,3},{1,1,40},{1,1,60},{201,406005,2},{201,310045,1},{201,310049,1},{201,432004,2},{201,432004,4},{21,21,4},{17,17,5},{17,17,5},{201,435001,4},{201,405003,6}}
	InstSpecialty.Items = append(InstSpecialty.Items,item)
}
