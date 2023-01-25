package polly

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
	"io"
	"os"
)

func SynthesizeSpeech(service *polly.Polly, text string, voice VoiceDesc) error {
	engine := "standard"
	if voice.Neural {
		engine = "neural"
	}
	input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String(text), VoiceId: aws.String(voice.Name), Engine: aws.String(engine)}

	output, err := service.SynthesizeSpeech(input)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("outputs/%s", voice.Name), 0770)
	if err != nil {
		return err
	}
	outFile, err := os.Create(fmt.Sprintf("outputs/%s/%s.mp3", voice.Name, text))
	if err != nil {
		return err
	}

	defer outFile.Close()
	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		return err
	}
	return nil
}
