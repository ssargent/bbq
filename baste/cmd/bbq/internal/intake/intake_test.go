package intake

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/ssargent/apis/pkg/bbq/intake/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type GrpcTest struct {
	listener *bufconn.Listener
}

func (g *GrpcTest) dialer() func(context.Context, string) (net.Conn, error) {
	g.listener = bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()

	RegisterIntake(server, nil, nil, nil, nil)

	go func() {
		if err := server.Serve(g.listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return g.listener.Dial()
	}
}

func TestIntakeServer_Session_Grpc(t *testing.T) {
	tests := map[string]struct {
		in   *pb.SessionRequest
		want func(got *pb.SessionResponse, err error)
	}{
		"success": {
			in: &pb.SessionRequest{
				Name: "test122",
			},
			want: func(got *pb.SessionResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)

			},
		},
	}

	ctx := context.Background()
	grpcTest := &GrpcTest{}
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(grpcTest.dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewIntakeServiceClient(conn)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.Session(ctx, tt.in)
			tt.want(got, err)
		})
	}
}

func TestIntakeServer_Session_Server(t *testing.T) {
	tests := map[string]struct {
		in   *pb.SessionRequest
		want func(got *pb.SessionResponse, err error)
	}{
		"success": {
			in: &pb.SessionRequest{
				Name: "test122",
			},
			want: func(got *pb.SessionResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &intakeServer{}
			ctx := context.Background()
			got, err := s.Session(ctx, tt.in)
			tt.want(got, err)
		})
	}
}
