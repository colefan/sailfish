package gate

//SERVER_TYPE
const SERVER_TYPE_UNKNOWN = 0
const SERVER_TYPE_LOGIN = 1
const SERVER_TYPE_HALL = 2
const SERVER_TYPE_GAME = 10

//Route route
type Route interface {
	RouteServer(serverType int) int
}

type BaseRoute struct {
	serverBalance map[int]int
}

func (r *BaseRoute) RouteServer(serverType int) int {

	return 0
}
