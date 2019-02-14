package registry

import (
	"github.com/atlassian/voyager"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/asapkey"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/aws"
	ec2compute_v2 "github.com/atlassian/voyager/pkg/orchestration/wiring/ec2compute/v2"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/k8scompute"
	apik8scompute "github.com/atlassian/voyager/pkg/orchestration/wiring/k8scompute/api"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/kubeingress"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/platformdns"
	apiplatformdns "github.com/atlassian/voyager/pkg/orchestration/wiring/platformdns/api"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/postgres"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/rds"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/sqs"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/ups"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringplugin"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringutil"
)

var KnownWiringPlugins = map[voyager.ResourceType]wiringplugin.WiringPlugin{
	apik8scompute.ResourceType:  wiringutil.StatusAdapter(k8scompute.New().WireUp),
	kubeingress.ResourceType:    wiringutil.StatusAdapter(kubeingress.WireUp),
	ec2compute_v2.ResourceType:  wiringutil.StatusAdapter(ec2compute_v2.New().WireUp),
	ups.ResourceType:            wiringutil.StatusAdapter(ups.New().WireUp),
	aws.Cfn:                     aws.CfnPlugin,
	aws.DynamoDB:                aws.DynamoDBPlugin,
	aws.S3:                      aws.S3Plugin,
	postgres.ResourceType:       wiringutil.StatusAdapter(postgres.New().WireUp),
	rds.ResourceType:            wiringutil.StatusAdapter(rds.New().WireUp),
	sqs.ResourceType:            wiringutil.StatusAdapter(sqs.WireUp),
	asapkey.ResourceType:        wiringutil.StatusAdapter(asapkey.New().WireUp),
	apiplatformdns.ResourceType: wiringutil.StatusAdapter(platformdns.New().WireUp),
}
