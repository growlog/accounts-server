package controllers

import (
	"context"
	// "log"

	tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/growlog/rpc/protos"
	"github.com/growlog/accounts-server/internal/utils"
)

func (s *AccountServer) VerifyAccessToken(ctx context.Context, in *pb.VerifyAccessTokenRequest) (*pb.VerifyAccessTokenResponse, error) {
    // Get the access token string which we will verify.
	accessToken := in.Token

	// Plug in our string into our access token handler and validate.
	claims, err := utils.VerifyAccessTokenString(accessToken)

	// CASE 1 OF 2: ERROR
	//---------------------
	if err != nil {
		// Return a response indicating that there the verification failed.
		return &pb.VerifyAccessTokenResponse{
			Status: false,
			UserId: 0,
			ThingId: 0,
			ExpiresAt: nil,
		}, nil
	}

    // CASE 2 OF 2: SUCCESS
    //---------------------
	// Convert from int64 to `protocol buffer timestamp` object.
	ts := &tspb.Timestamp{
		Seconds: claims.ExpiresAt,
		Nanos: 0,
	}

    // Return a response indicating the verification passed and here are the
	// contents of the claim.
	return &pb.VerifyAccessTokenResponse{
		Status: true,
		UserId: claims.UserId,
		ThingId: claims.ThingId,
		ExpiresAt: ts,
	}, nil
}
