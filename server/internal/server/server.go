package server

import (
	"context"

	spv1 "buf.build/gen/go/atypicaldev/splendorapis/protocolbuffers/go/atypicaldev/splendorapis/v1"

	"connectrpc.com/connect"
)

type SplendorService struct{}

func (s *SplendorService) CreateTable(
	ctx context.Context,
	req *connect.Request[spv1.CreateTableRequest],
) (*connect.Response[spv1.CreateTableResponse], error) {
	res := connect.NewResponse(&spv1.CreateTableResponse{})

	return res, nil
}

func (s *SplendorService) JoinTable(
	ctx context.Context,
	req *connect.Request[spv1.JoinTableRequest],
) (*connect.Response[spv1.JoinTableResponse], error) {
	res := connect.NewResponse(&spv1.JoinTableResponse{})

	return res, nil
}

func (s *SplendorService) LeaveTable(
	ctx context.Context,
	req *connect.Request[spv1.LeaveTableRequest],
) (*connect.Response[spv1.LeaveTableResponse], error) {
	res := connect.NewResponse(&spv1.LeaveTableResponse{})

	return res, nil
}

func (s *SplendorService) StartGame(
	ctx context.Context,
	req *connect.Request[spv1.StartGameRequest],
) (*connect.Response[spv1.StartGameResponse], error) {
	res := connect.NewResponse(&spv1.StartGameResponse{})

	return res, nil
}
