package qs3

import (
	"context"
	"fmt"
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/cos"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/user"
)

func NewQs3Service() *Service {
	return &Service{}
}

var project = "qclouds3"

type Service struct{}

func (s *Service) CreateBucket(name string) (string, error) {
	ctx := context.Background()
	stackName := name
	progName := createBucket(name)

	_ = os.Setenv("TENCENTCLOUD_REGION", "ap-guangzhou")
	_ = os.Setenv("TENCENTCLOUD_SECRET_KEY", "xxxxxxx")
	_ = os.Setenv("TENCENTCLOUD_SECRET_ID", "xxxxxxxx")

	stack, err := auto.NewStackInlineSource(ctx, stackName, project, progName)
	if err != nil {
		// if stack already exists, 409
		if auto.IsCreateStack409Error(err) {
			fmt.Printf("stack %s already exists\n", stackName)
			return "", nil
		}

		fmt.Printf("stack creation failed: %v", err)
		return "", nil
	}

	// deploy the stack
	// we'll write all of the update logs to st	out so we can watch requests get processed
	tmpfile, _ := os.OpenFile("/tmp/stacklog", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	upRes, err := stack.Up(ctx, optup.ProgressStreams(tmpfile))
	if err != nil {
		fmt.Printf("update stack file %s failed: %v", stackName, err)
		return "", nil
	}

	return upRes.Outputs["bucketName"].Value.(string), nil
}

func createBucket(bucketName string) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		info, err := user.GetInfo(ctx, nil, nil)
		if err != nil {
			return err
		}
		bucket, err := cos.NewBucket(ctx, bucketName, &cos.BucketArgs{
			Acl:    pulumi.String("private"),
			Bucket: pulumi.String(fmt.Sprintf("%v%v%v", bucketName, "-", info.AppId)),
		})
		if err != nil {
			return err
		}

		ctx.Export("bucketName", bucket.ID())
		ctx.Export("Acl", bucket.Acl.ToStringPtrOutput())
		return nil
	}
}
