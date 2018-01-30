package main

import (
	"os"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nulab/go-typetalk/typetalk"
	"log"
	"strconv"
	"fmt"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudwatchlogsevt"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context, event *cloudwatchlogsevt.Event) (Response, error) {

	token := os.Getenv("TYPETALK_TOKEN")

	talkId, err := strconv.Atoi(os.Getenv("TYPETALK_TALK_ID"))
	if err != nil {
		log.Fatal(err)
	}

	topicId, err := strconv.Atoi(os.Getenv("TYPETALK_TOPIC_ID"))
	if err != nil {
		log.Fatal(err)
	}

	//re := regexp.MustCompile(`elasticbeanstalk/(.*?)/var/log/.*`)
	//environmentName := re.FindStringSubmatch(event.Records[0].LogGroup)[1]

	messages := make([]string, len(event.Records))
	for i, record := range event.Records {
		messages[i] = record.LogEvent.Message
	}

	// TypetalkのPOSTサイズ上限があるので長い分は抜く
	r := []rune(strings.Join(messages, "\n"))
	message := fmt.Sprintf("```\n%s\n```", string(r[0:3900]))

	// 追いかけやすいように便利なまとめ機能にもいれる
	options := typetalk.PostMessageOptions{}
	options.TalkIds = []int{talkId}

	client := typetalk.NewClient(nil).SetTypetalkToken(token)

	client.Messages.PostMessage(ctx, topicId, message, &options)

	return Response{
		Message: "posted awslogs to Typetalk.",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
