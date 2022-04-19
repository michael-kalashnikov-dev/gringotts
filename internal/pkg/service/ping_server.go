package service

import (
	"context"
	"fmt"
	"github.com/michael-kalashnikov-dev/gringotts/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type PingServer struct {
	proto.UnimplementedPingServiceServer
}

func NewPingServer() *PingServer {
	return &PingServer{}
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

// Ping its unary rpc for pinging
func (server *PingServer) Ping(
	ctx context.Context,
	req *proto.PingRequest,
) (*proto.PingResponse, error) {
	message := req.GetMessage()
	if len(message) < 1 {
		return nil, status.Error(codes.InvalidArgument, "message cannot be empty!")
	}

	timestamp := timestamppb.Now()
	log.Printf("receive a ping request with message: %s\n", message)
	log.Println(timestamp)

	// some heavy processing
	time.Sleep(6 * time.Second)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	res := &proto.PingResponse{
		Message:   fmt.Sprintf("Ping: %s", message),
		Timestamp: timestamp,
	}
	log.Println("before response")
	return res, nil
}
