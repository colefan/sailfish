package config
type Refresh struct{
	Items []*Refresh_Item
}

type Refresh_Item struct{
	Name int
	Val string
	Currency_type int
	Val_type int
}

var InstRefresh *Refresh

func init(){
	item := &Refresh_Item{}
	item.Name = 80509
	item.Val = "10,10,50,100,200,500,500,1000"
	item.Currency_type = 2
	item.Val_type = 2
	InstRefresh.Items = append(InstRefresh.Items,item)
}
