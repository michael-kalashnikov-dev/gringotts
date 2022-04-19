package service_test

import (
	"context"
	"github.com/mazen160/go-random"
	"github.com/michael-kalashnikov-dev/gringotts/internal/pkg/service"
	"github.com/michael-kalashnikov-dev/gringotts/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestPingServer(t *testing.T) {
	t.Parallel()

	rStr, err := random.StringRange(1, 63)
	require.NoError(t, err)

	testCases := []struct {
		name    string
		message string
		code    codes.Code
	}{
		{
			name:    "success_with_message",
			message: rStr,
			code:    codes.OK,
		},
		{
			name: "failure_no_message",
			code: codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &proto.PingRequest{
				Message: tc.message,
			}
			server := service.NewPingServer()
			res, err := server.Ping(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Message)
				require.NotEmpty(t, res.Timestamp)
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}

}
