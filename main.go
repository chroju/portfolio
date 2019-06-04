package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://chroju.github.io/atom.xml")
	items := feed.Items

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       items[0].Title,
	}, nil
}

func main() {
	lambda.Start(handler)
}
