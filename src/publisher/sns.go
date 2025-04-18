package publisher

import (
	"context"

	interfaces "github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/_interfaces"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type snsPublisher struct {
	client *sns.Client
	topic  string
}

func NewSnsPublisher(client *sns.Client, topic string) interfaces.Publisher {
	return &snsPublisher{
		client: client,
		topic:  topic,
	}
}

func (ref *snsPublisher) Publish(ctx context.Context, message string) error {
	publishInput := sns.PublishInput{
		TopicArn: aws.String(ref.topic),
		Message:  aws.String(message),
	}

	_, err := ref.client.Publish(ctx, &publishInput)
	if err != nil {
		return err
	}

	return nil
}
