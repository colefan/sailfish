package config
type Skill_index struct{
	Items []*Skill_index_Item
}

type Skill_index_Item struct{
	Id int
	Skill_id []int
}

var InstSkill_index *Skill_index

func init(){
	item := &Skill_index_Item{}
	item.Id = 1000101
	item.Skill_id = []int{102701}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000102
	item.Skill_id = []int{102702}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000103
	item.Skill_id = []int{102703}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000104
	item.Skill_id = []int{102704}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000105
	item.Skill_id = []int{102705}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000201
	item.Skill_id = []int{100401}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000202
	item.Skill_id = []int{100402}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000203
	item.Skill_id = []int{100403}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000204
	item.Skill_id = []int{100404}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000205
	item.Skill_id = []int{100405}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000301
	item.Skill_id = []int{103301,105501}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000302
	item.Skill_id = []int{103302,105502}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000303
	item.Skill_id = []int{103303,105503}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000304
	item.Skill_id = []int{103304,105504}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000305
	item.Skill_id = []int{103305,105505}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002401
	item.Skill_id = []int{101401}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002402
	item.Skill_id = []int{101402}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002403
	item.Skill_id = []int{101403}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002404
	item.Skill_id = []int{101404}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002405
	item.Skill_id = []int{101405}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000401
	item.Skill_id = []int{104201}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000402
	item.Skill_id = []int{104202}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000403
	item.Skill_id = []int{104203}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000404
	item.Skill_id = []int{104204}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000405
	item.Skill_id = []int{104205}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000901
	item.Skill_id = []int{101101,105901}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000902
	item.Skill_id = []int{101102,105902}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000903
	item.Skill_id = []int{101103,105903}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000904
	item.Skill_id = []int{101104,105904}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000905
	item.Skill_id = []int{101105,105905}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000501
	item.Skill_id = []int{101001}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000502
	item.Skill_id = []int{101002}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000503
	item.Skill_id = []int{101003}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000504
	item.Skill_id = []int{101004}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000505
	item.Skill_id = []int{101005}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000601
	item.Skill_id = []int{100701}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000602
	item.Skill_id = []int{100702}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000603
	item.Skill_id = []int{100703}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000604
	item.Skill_id = []int{100704}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000605
	item.Skill_id = []int{100705}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000701
	item.Skill_id = []int{103601}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000702
	item.Skill_id = []int{103602}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000703
	item.Skill_id = []int{103603}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000704
	item.Skill_id = []int{103604}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000705
	item.Skill_id = []int{103605}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000801
	item.Skill_id = []int{100901}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000802
	item.Skill_id = []int{100902}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000803
	item.Skill_id = []int{100903}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000804
	item.Skill_id = []int{100904}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1000805
	item.Skill_id = []int{100905}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001201
	item.Skill_id = []int{101601}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001202
	item.Skill_id = []int{101602}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001203
	item.Skill_id = []int{101603}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001204
	item.Skill_id = []int{101604}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001205
	item.Skill_id = []int{101605}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001301
	item.Skill_id = []int{100101}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001302
	item.Skill_id = []int{100102}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001303
	item.Skill_id = []int{100103}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001304
	item.Skill_id = []int{100104}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001305
	item.Skill_id = []int{100105}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001601
	item.Skill_id = []int{102901}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001602
	item.Skill_id = []int{102902}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001603
	item.Skill_id = []int{102903}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001604
	item.Skill_id = []int{102904}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001605
	item.Skill_id = []int{102905}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001701
	item.Skill_id = []int{104901}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001702
	item.Skill_id = []int{104902}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001703
	item.Skill_id = []int{104903}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001704
	item.Skill_id = []int{104904}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001705
	item.Skill_id = []int{104905}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001801
	item.Skill_id = []int{101301}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001802
	item.Skill_id = []int{101302}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001803
	item.Skill_id = []int{101303}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001804
	item.Skill_id = []int{101304}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001805
	item.Skill_id = []int{101305}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002501
	item.Skill_id = []int{101501}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002502
	item.Skill_id = []int{101502}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002503
	item.Skill_id = []int{101503}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002504
	item.Skill_id = []int{101504}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002505
	item.Skill_id = []int{101505}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001901
	item.Skill_id = []int{104401,106101}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001902
	item.Skill_id = []int{104402,106102}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001903
	item.Skill_id = []int{104403,106103}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001904
	item.Skill_id = []int{104404,106104}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1001905
	item.Skill_id = []int{104405,106105}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002601
	item.Skill_id = []int{101801}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002602
	item.Skill_id = []int{101802}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002603
	item.Skill_id = []int{101803}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002604
	item.Skill_id = []int{101804}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002605
	item.Skill_id = []int{101805}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002001
	item.Skill_id = []int{101701}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002002
	item.Skill_id = []int{101702}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002003
	item.Skill_id = []int{101703}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002004
	item.Skill_id = []int{101704}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002005
	item.Skill_id = []int{101705}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002101
	item.Skill_id = []int{101201}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002102
	item.Skill_id = []int{101202}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002103
	item.Skill_id = []int{101203}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002104
	item.Skill_id = []int{101204}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002105
	item.Skill_id = []int{101205}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002201
	item.Skill_id = []int{102601}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002202
	item.Skill_id = []int{102602}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002203
	item.Skill_id = []int{102603}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002204
	item.Skill_id = []int{102604}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002205
	item.Skill_id = []int{102605}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002701
	item.Skill_id = []int{101901,107001}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002702
	item.Skill_id = []int{101902,107002}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002703
	item.Skill_id = []int{101903,107003}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002704
	item.Skill_id = []int{101904,107004}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002705
	item.Skill_id = []int{101905,107005}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002301
	item.Skill_id = []int{105201}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002302
	item.Skill_id = []int{105202}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002303
	item.Skill_id = []int{105203}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002304
	item.Skill_id = []int{105204}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002305
	item.Skill_id = []int{105205}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003601
	item.Skill_id = []int{100301}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003602
	item.Skill_id = []int{100302}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003603
	item.Skill_id = []int{100303}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003604
	item.Skill_id = []int{100304}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003605
	item.Skill_id = []int{100305}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003701
	item.Skill_id = []int{100201}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003702
	item.Skill_id = []int{100202}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003703
	item.Skill_id = []int{100203}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003704
	item.Skill_id = []int{100204}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003705
	item.Skill_id = []int{100205}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004001
	item.Skill_id = []int{102101}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004002
	item.Skill_id = []int{102102}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004003
	item.Skill_id = []int{102103}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004004
	item.Skill_id = []int{102104}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004005
	item.Skill_id = []int{102105}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003801
	item.Skill_id = []int{100501}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003802
	item.Skill_id = []int{100502}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003803
	item.Skill_id = []int{100503}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003804
	item.Skill_id = []int{100504}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003805
	item.Skill_id = []int{100505}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003901
	item.Skill_id = []int{108001,102001}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003902
	item.Skill_id = []int{108002,102002}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003903
	item.Skill_id = []int{108003,102003}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003904
	item.Skill_id = []int{108004,102004}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1003905
	item.Skill_id = []int{108005,102005}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004201
	item.Skill_id = []int{104701,100000}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004202
	item.Skill_id = []int{104702,100000}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004203
	item.Skill_id = []int{104703,100000}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004204
	item.Skill_id = []int{104704,100000}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004205
	item.Skill_id = []int{104705,100000}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004301
	item.Skill_id = []int{102201,106001}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004302
	item.Skill_id = []int{102202,106002}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004303
	item.Skill_id = []int{102203,106003}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004304
	item.Skill_id = []int{102204,106004}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004305
	item.Skill_id = []int{102205,106005}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004401
	item.Skill_id = []int{104101}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004402
	item.Skill_id = []int{104102}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004403
	item.Skill_id = []int{104103}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004404
	item.Skill_id = []int{104104}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004405
	item.Skill_id = []int{104105}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004901
	item.Skill_id = []int{100801}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004902
	item.Skill_id = []int{100802}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004903
	item.Skill_id = []int{100803}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004904
	item.Skill_id = []int{100804}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004905
	item.Skill_id = []int{100805}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002901
	item.Skill_id = []int{105301}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002902
	item.Skill_id = []int{105302}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002903
	item.Skill_id = []int{105303}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002904
	item.Skill_id = []int{105304}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1002905
	item.Skill_id = []int{105305}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004501
	item.Skill_id = []int{105401}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004502
	item.Skill_id = []int{105402}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004503
	item.Skill_id = []int{105403}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004504
	item.Skill_id = []int{105404}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
	item = &Skill_index_Item{}
	item.Id = 1004505
	item.Skill_id = []int{105405}
	InstSkill_index.Items = append(InstSkill_index.Items,item)
}
