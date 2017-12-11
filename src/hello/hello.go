package main

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/polly"
import "fmt"


func main() {
  session, err := session.NewSession()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  service := polly.New(session)

  speechInput := &polly.SynthesizeSpeechInput{
    OutputFormat: aws.String("mp3"),
    Text: aws.String("This is my sample text"),
    VoiceId: aws.String("Mathieu"),
  }

  output, err := service.SynthesizeSpeech(speechInput)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(output)
}
