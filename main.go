package main

import (
	"bytes"
	"html/template"
	"io"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://chroju.github.io/atom.xml")
	items := feed.Items

	tmpl := template.Must(template.ParseFiles("test.html"))
	buf := new(bytes.Buffer)
	w := io.Writer(buf)

	err := tmpl.ExecuteTemplate(w, "base", struct {
		Title string
	}{
		Title: items[0].Title,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(buf.Bytes()),
	}, nil
}

func main() {
	lambda.Start(handler)
}
