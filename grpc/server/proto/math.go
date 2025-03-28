package proto

import "context"

type mathServer struct {
	UnimplementedMathServer
}

func (s *mathServer) Sum(_ context.Context, in *NumberRequest) (*NumberReponse, error) {
	return &NumberReponse{Num: in.Num1 + in.Num2}, nil
}

func NewServer() *mathServer {
	return &mathServer{}
}
