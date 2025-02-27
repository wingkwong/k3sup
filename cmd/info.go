package cmd

import (
	"fmt"

	"github.com/alexellis/k3sup/cmd/apps"
	"github.com/spf13/cobra"
)

func MakeInfo() *cobra.Command {

	info := &cobra.Command{
		Use:   "info",
		Short: "Find info about a Kubernetes app",
		Long:  "Find info about how to use the installed Kubernetes app",
		Example: `  k3sup app info [APP]
k3sup app info openfaas
k3sup app info inlets-operator
k3sup app info mongodb
k3sup app info
k3sup app info --help`,
		SilenceUsage: true,
	}

	info.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("You can get info about: openfaas, inlets-operator, mongodb")
			return nil
		}

		if len(args) != 1 {
			return fmt.Errorf("you can only get info about exactly one installed app")
		}

		appName := args[0]

		switch appName {
		case "inlets-operator":
			fmt.Printf("Info for app: %s\n", appName)
			fmt.Println(apps.InletsOperatorInfoMsg)
		case "openfaas":
			fmt.Printf("Info for app: %s\n", appName)
			fmt.Println(apps.OpenFaaSInfoMsg)
		case "mongodb":
			fmt.Printf("Info for app: %s\n", appName)
			fmt.Println(apps.MongoDBInfoMsg)
		case "metrics-server":
			fmt.Printf("Info for app: %s\n", appName)
			fmt.Println(apps.MetricsInfoMsg)
		default:
			return fmt.Errorf("no info available for app: %s", appName)
		}

		return nil
	}

	return info
}
