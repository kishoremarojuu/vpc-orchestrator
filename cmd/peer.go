package cmd

import (
    "context"
    "fmt"

    "github.com/spf13/cobra"
    "vpc-orchestrator/orchestrator"
    "vpc-orchestrator/utils"
)

var (
    sourceVpc   string
    targetVpc   string
    peerAccount string
)

func Execute(ctx context.Context) error {
    var rootCmd = &cobra.Command{
        Use:   "vpc-orchestrator",
        Short: "Automate VPC peering",
    }

    var peerCmd = &cobra.Command{
        Use:   "peer",
        Short: "Create VPC peering",
        RunE: func(cmd *cobra.Command, args []string) error {
            awsClient := utils.NewEC2Client()
            service := orchestrator.PeeringService{EC2Client: awsClient}

            connID, err := service.CreatePeering(ctx, sourceVpc, targetVpc, peerAccount)
            if err != nil {
                return err
            }

            fmt.Printf("Peering connection created successfully: %s\n", connID)
            return nil
        },
    }

    peerCmd.Flags().StringVar(&sourceVpc, "source-vpc", "", "Source VPC ID")
    peerCmd.Flags().StringVar(&targetVpc, "target-vpc", "", "Target VPC ID")
    peerCmd.Flags().StringVar(&peerAccount, "peer-account", "", "Peer AWS Account ID")

    rootCmd.AddCommand(peerCmd)
    return rootCmd.Execute()
}
