package polly

import (
	"github.com/aws/aws-sdk-go/service/polly"
	"voice-synth/aws"
)

func PollyService(profile string, region string) *polly.Polly {
	sess := aws.AwsSession(profile, region)
	return polly.New(sess)
}
