package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/growlog/rpc/protos"
)

func (s *AccountServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Token: "Instrument was updated",
		Status: true,
	}, nil
}

func (s *AccountServer) RefreshAccessToken(ctx context.Context, in *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	return &pb.RefreshAccessTokenResponse{
		Status: true,
		UserId: 0,
		ThingId: 0,
		ExpiresAt: nil,
	}, nil
}
