package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type textExtractor interface {
	getText(reader io.Reader) (string, error)
}

type googleExtractor struct{}

func (g googleExtractor) getText(reader io.Reader) (string, error) {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", err
	}
	image, err := vision.NewImageFromReader(reader)
	if err != nil {
		return "", err
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", err
	}

	if len(annotations) == 0 {
		return "No text found!", nil
	}

	return annotations[0].Description, nil
}

type awsExtractor struct{}

func (g awsExtractor) getText(reader io.Reader) (string, error) {
	sess := session.Must(session.NewSession())
	config := &aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	rekognitionService := rekognition.New(sess, config)
	image := rekognition.Image{Bytes: []byte(bytes)}
	output, err := rekognitionService.DetectText(&rekognition.DetectTextInput{
		Image: &image,
	})
	if err != nil {
		return "", err
	}
	lines := []string{}
	for _, td := range output.TextDetections {
		if *td.Type != "LINE" {
			break
		}
		lines = append(lines, *td.DetectedText)
	}
	return strings.Join(lines, "\n"), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <image>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	var extractor textExtractor
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != "" {
		extractor = googleExtractor{}
	} else {
		extractor = awsExtractor{}
	}
	text, err := extractor.getText(file)
	if err != nil {
		log.Fatalf("Failed to get text from file: %v", err)
	}
	fmt.Printf("%s\n", text)
}
