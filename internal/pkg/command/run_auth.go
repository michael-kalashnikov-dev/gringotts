package command

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/michael-kalashnikov-dev/gringotts/internal/pkg/service"
	"github.com/michael-kalashnikov-dev/gringotts/pkg/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

// runAuthCmd represents the run command
var runAuthCmd = &cobra.Command{
	Use:   "run",
	Short: "Start running of specified server",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("grpc").Value.String() == "true" {
			grpcAddress := cmd.Flags().Lookup("grpcAddress").Value.String()
			log.Printf("start gRPC server on address %s", grpcAddress)

			grpcServer := grpc.NewServer()
			pingServer := service.NewPingServer()
			proto.RegisterPingServiceServer(grpcServer, pingServer)

			listener, err := net.Listen("tcp", grpcAddress)
			cobra.CheckErr(err)

			err = grpcServer.Serve(listener)
			cobra.CheckErr(err)
		} else {
			cobra.CheckErr(cmd.Help())
		}
	},
}

func init() {
	authCmd.AddCommand(runAuthCmd)

	flagenv.New(
		runAuthCmd.Flags(),
		flagenv.Bool(),
		flagenv.WithFlagName("grpc"),
		flagenv.WithShorthand("-"),
		flagenv.WithUsage("indicates if gRPC server should start running"),
	)

	flagenv.New(
		runAuthCmd.Flags(),
		flagenv.String(),
		flagenv.WithFlagName("grpcAddress"),
		flagenv.WithUsage("gRPC server address example: 0.0.0.0:8000"),
	)
}
