package utils

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func NewEC2Client() *ec2.Client {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("failed to load AWS config")
    }
    return ec2.NewFromConfig(cfg)
}