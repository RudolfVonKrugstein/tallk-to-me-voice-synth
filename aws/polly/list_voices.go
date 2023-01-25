package polly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
)

type VoiceDesc struct {
	Name         string
	Neural       bool
	LanguageCode string
}

func ListPollyVoices(svc *polly.Polly) (*[]VoiceDesc, error) {
	var res []VoiceDesc
	// We are only interested in german and english voices
	languageCodes := []string{"en-US", "en-GB", "en-AU", "en-IN", "en-ZA", "de-DE", "de-AT"}

	for _, languageCode := range languageCodes {
		input := &polly.DescribeVoicesInput{LanguageCode: aws.String(languageCode)}
		resp, err := svc.DescribeVoices(input)
		if err != nil {
			return nil, err
		}
		for _, voice := range resp.Voices {
			neuralSupported := false
			for _, engine := range voice.SupportedEngines {
				if *engine == "neural" {
					neuralSupported = true
				}
			}
			res = append(res, VoiceDesc{
				Name:         *voice.Name,
				Neural:       neuralSupported,
				LanguageCode: *voice.LanguageCode,
			})
		}
	}
	return &res, nil
}
