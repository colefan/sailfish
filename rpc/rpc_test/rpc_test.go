package rpc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/colefan/sailfish"
)

type MathService struct {
}

type MathServiceReq struct {
	A int
	B int
}

type MathServiceResp struct {
	Sum int
}

func (s *MathService) Add(req MathServiceReq, resp *MathServiceResp) error {
	resp.Sum = req.A + req.B
	return nil
}

func TestRPC(t *testing.T) {
	s := sailfish.CreateRPCServer("127.0.0.1:10101")
	s.RegisterService(new(MathService))
	go s.Serve()
	// defer s.()
	time.Sleep(1 * time.Second)

	c, err := sailfish.CreateRPCClient("127.0.0.1:10101")
	if err != nil {
		t.FailNow()
	}
	req := MathServiceReq{A: 10, B: 20}
	var resp MathServiceResp
	err = c.Call("MathService.Add", req, &resp)
	if err != nil {
		t.Fatalf("call :%v", err)
		// t.FailNow()
	}
	if resp.Sum != 30 {
		t.Fatalf("result :%v", resp.Sum)
		// t.Failed()
	}

	fmt.Printf(" %d + %d = %d \n", req.A, req.B, resp.Sum)

}
