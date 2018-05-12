package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// GetAnnotationsFromImage returns the annotations for the read image
func GetAnnotationsFromImage(ctx context.Context, client *vision.ImageAnnotatorClient, reader io.Reader) ([]string, error) {
	image, err := vision.NewImageFromReader(reader)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create image")
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to detect labels")
	}

	if len(annotations) == 0 {
		return []string{}, nil
	}

	texts := []string{}
	for _, annotation := range annotations {
		texts = append(texts, annotation.Description)
	}
	return texts, nil
}

func ParseAnnotations(annotations []string) map[string]string {
	found := make(map[string]string)
	for i, a := range annotations {
		switch a {
		case "Total:":
			fallthrough
		case "Tax:":
			found[strings.Trim(a, ":")] = annotations[i+1]
		case "Placed:":
			found["Date"] = annotations[i+1] +
				annotations[i+2] +
				annotations[i+3]
		}
	}
	return found
}

func main() {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	filename := "./images/amazon-order.jpg"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	annotations, err := GetAnnotationsFromImage(ctx, client, file)
	if err != nil {
		log.Fatalf("Failed to get annotations from file: %v", err)
	}
	fmt.Printf("%#v\n", ParseAnnotations(annotations))
}
