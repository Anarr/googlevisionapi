package googlevisionapi

import (
	"context"
	"errors"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

//AuthWithCredentials auth app with credentials file
func AuthWithCredentials(file string) {
	//set env variable for pass google auth
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", file)
}

//DetectSafeSearchURI detect nsfw content(safe search) from remote image
func DetectSafeSearchURI(file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)

	props, err := client.DetectSafeSearch(ctx, image, nil)

	if err != nil {
		return err
	}

	fmt.Println("Result", props)

	fmt.Println("Safe Search properties:")
	fmt.Println("Adult:", props.Adult)
	fmt.Println("Medical:", props.Medical)
	fmt.Println("Racy:", props.Racy)
	fmt.Println("Spoofed:", props.Spoof)
	fmt.Println("Violence:", props.Violence)
	return nil
}

//DetectLabels detect labels from image
func DetectLabels(file string) ([]string, err) {
	var labels []string
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		return labels, err
	}

	f, err := os.Open(file)

	if err != nil {
		return labels, err
	}

	defer f.Close()

	image, err := vision.NewImageFromReader(f)

	if err != nil {
		return labels, err
	}

	annotations, err := client.DetectLabels(ctx, image, nil, 10)

	if err != nil {
		return labels, err
	}

	if len(annotations) == 0 {
		return labels, errors.New("can not detect labels")
	}

	for _, annotation := range annotations {
		labels = append(labels, annotation.String())
	}

	return labels, nil
}

//DetectLabelsURI detect labels from remote image
func DetectLabelsURI(file string) ([]string, error) {
	var labels []string
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		return labels, err
	}

	image := vision.NewImageFromURI(file)

	annotations, err := client.DetectLabels(ctx, image, nil, 10)

	if err != nil {
		return labels, err
	}

	if len(annotations) == 0 {
		return labels, errors.New("can not detect labels")
	}

	for _, annotation := range annotations {
		labels = append(labels, annotation.String())
	}

	return labels, nil
}
