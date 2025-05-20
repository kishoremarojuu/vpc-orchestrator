package orchestrator

import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

type SecurityService struct {
    EC2Client *ec2.Client
}

func (s *SecurityService) AuthorizeSecurityGroupIngress(ctx context.Context, groupID, cidr string, port int32) error {
    input := &ec2.AuthorizeSecurityGroupIngressInput{
        GroupId: aws.String(groupID),
        IpPermissions: []ec2types.IpPermission{
            {
                IpProtocol: aws.String("tcp"),
                FromPort:   aws.Int32(port),
                ToPort:     aws.Int32(port),
                IpRanges: []ec2types.IpRange{{
                    CidrIp: aws.String(cidr),
                }},
            },
        },
    }

    _, err := s.EC2Client.AuthorizeSecurityGroupIngress(ctx, input)
    if err != nil {
        return fmt.Errorf("failed to authorize ingress: %w", err)
    }
    return nil
}
