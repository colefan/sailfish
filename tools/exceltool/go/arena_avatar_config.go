package config
type Arena_avatar struct{
	Items []*Arena_avatar_Item
}

type Arena_avatar_Item struct{
	ID int
}

var InstArena_avatar *Arena_avatar

func init(){
	item := &Arena_avatar_Item{}
	item.ID = 1
	InstArena_avatar.Items = append(InstArena_avatar.Items,item)
	item = &Arena_avatar_Item{}
	item.ID = 2
	InstArena_avatar.Items = append(InstArena_avatar.Items,item)
}
