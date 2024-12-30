package server

import (
	"atypicaldev/splendor-go/internal/repository"
	"context"
	"fmt"
	"math/rand/v2"
	"strings"

	spv1 "buf.build/gen/go/atypicaldev/splendorapis/protocolbuffers/go/atypicaldev/splendorapis/v1"

	"connectrpc.com/connect"
)

type SplendorService struct {
	Repo repository.SplendorRepository
}

var adjectives []string = []string{
	"splendid",
	"glorious",
	"bad",
	"maleficent",
	"lazy",
	"industrious",
}

var nouns = []string{
	"dog",
	"river",
	"cloud",
	"folder",
	"sea",
	"idea",
	"framework",
}

func getName() string {
	ai := rand.N(len(adjectives))
	ni := rand.N(len(nouns))

	builder := strings.Builder{}

	builder.WriteString(adjectives[ai])
	builder.WriteString("-")

	builder.WriteString(nouns[ni])
	return builder.String()
}

func (s *SplendorService) CreateTable(
	ctx context.Context,
	req *connect.Request[spv1.CreateTableRequest],
) (*connect.Response[spv1.CreateTableResponse], error) {
	repo := s.Repo

	msg := req.Msg
	fmt.Printf("CreateTable request: %v\n", msg)
	tableData, err := repo.CreateTable(ctx, getName())
	if err != nil {
		return nil, err
	}

	tableId := tableData.TableID.String()
	table := spv1.Table_builder{
		TableId: &tableId,
		Players: []*spv1.Player{
			{Id: msg.CreatorId},
		},
	}.Build()
	res := connect.NewResponse(&spv1.CreateTableResponse{
		Table: table,
	})

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
