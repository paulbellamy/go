package cmd

import (
	stdLog "log"

	"github.com/spf13/cobra"
	horizon "github.com/stellar/go/services/horizon/internal"
)

var (
	config, flags = horizon.Flags()

	RootCmd = &cobra.Command{
		Use:   "horizon",
		Short: "client-facing api server for the Stellar network",
		Long:  "Client-facing API server for the Stellar network. It acts as the interface between Stellar Core and applications that want to access the Stellar network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams and more.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return horizon.NewAppFromFlags(config, flags).Serve()
		},
	}
)

func init() {
	err := flags.Init(RootCmd)
	if err != nil {
		stdLog.Fatal(err.Error())
	}
}

func Execute() error {
	return RootCmd.Execute()
}
