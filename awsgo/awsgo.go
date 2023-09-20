package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InitializeAws() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-"))

	if err != nil {
		panic("Error loading AWS settings ./aws/aws" + err.Error())
	}
}
