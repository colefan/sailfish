package gate

import "sailfish/network"

//Route route
type Route interface {
	RouteServer(serverType int, pack network.PackInf)
}

type BaseRoute struct {
	// serverBalance map[int]int
}

func (r *BaseRoute) RouteServer(ProxyServerNode int, pack network.PackInf) {

}
