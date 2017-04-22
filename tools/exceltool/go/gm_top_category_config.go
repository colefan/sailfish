package config
type Gm_top_category struct{
	Items []*Gm_top_category_Item
}

type Gm_top_category_Item struct{
	Id int
	Enum_value string
	Kingdom string
}

var InstGm_top_category *Gm_top_category

func init(){
	item := &Gm_top_category_Item{}
	item.Id = 1
	item.Enum_value = "top_earning"
	item.Kingdom = "earning"
	InstGm_top_category.Items = append(InstGm_top_category.Items,item)
	item = &Gm_top_category_Item{}
	item.Id = 2
	item.Enum_value = "top_expenditure"
	item.Kingdom = "expenditure"
	InstGm_top_category.Items = append(InstGm_top_category.Items,item)
}
