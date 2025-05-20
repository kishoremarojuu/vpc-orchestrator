package orchestrator

import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

type RouteService struct {
    EC2Client *ec2.Client
}

func (r *RouteService) AddRouteToTable(ctx context.Context, routeTableID, destinationCIDR, targetID string) error {
    _, err := r.EC2Client.CreateRoute(ctx, &ec2.CreateRouteInput{
        RouteTableId:         aws.String(routeTableID),
        DestinationCidrBlock: aws.String(destinationCIDR),
        VpcPeeringConnectionId: aws.String(targetID),
    })
    if err != nil {
        return fmt.Errorf("failed to add route: %w", err)
    }
    return nil
}
