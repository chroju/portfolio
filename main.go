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
	ghitems := feed.Items[:3]
	feed, _ = fp.ParseURL("https://chroju.hatenablog.jp/feed")
	hbitems := feed.Items[:3]

	tmpl := template.Must(template.New("index.html").Parse(htmlTemplate))
	buf := new(bytes.Buffer)
	w := io.Writer(buf)

	err := tmpl.ExecuteTemplate(w, "base", struct {
		GitHubIOEntries []*gofeed.Item
		HatenaBlogEntries []*gofeed.Item
	}{
		GitHubIOEntries: ghitems,
		HatenaBlogEntries: hbitems,
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

const htmlTemplate = `
{{define "base"}}
<!doctype html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<script src="https://kit.fontawesome.com/215264fa68.js"></script>
	<title>chroju</title>
</head>
<body>
	<h1>chroju</h1>
	<img src="https://secure.gravatar.com/avatar/542bf1e833425f6ab7bc7bd7238a4792?s=250" alt="chroju" />
	<a href="https://github.com/chroju"><i class="fab fa-github-alt"></i></a>
	<a href="https://twitter.com/chroju"><i class="fab fa-twitter"></i></a>
	<a href="https://www.instagrafa-instagramm.com/chroju"><i class="fab fa-instagram"></i></a>
	<a href="https://speakerdeck.com/chroju"><i class="fab fa-speaker-deck"></i></a>

    <h2>Who</h2>
    <dl>
        <dt>Location</dt>
            <dd>Tokyo, Japan</dd>
        <dt>Skills</dt>
            <dd>AWS / Terraform / VMware / Go / Python / bash ... etc</dd>
        <dt>Contact</dt>
            <dd>chor.chroju at gmail.com</dd>
    </dl>

    <h2>Experience</h2>
    <dl>
        <dt>TIS Inc.</dt>
            <dd>System Engineer</dd>
            <dd>2011 - 2015</dd>
        <dt>A certain company</dt>
            <dd>Web Operation Engineer</dd>
            <dd>2015 - 2019</dd>
        <dt>Freelancer</dt>
            <dd>Site Reliability Engineer</dd>
            <dd>2019 -</dd>
    </dl>

    <h2>Blogs (recent entries)</h2>
    <h3><a href="https://chroju.github.io/">the world as code</a></h3>
    <ul>
        {{range $entry := .GitHubIOEntries }}
        <li><a href="{{$entry.Link}}">{{$entry.Title}}</a></li>
        {{end}}
    </ul>

    <h3><a href="https://chroju.hatenablog.jp/">the world was not enough</a></h3>
    <ul>
        {{range $entry := .HatenaBlogEntries }}
        <li><a href="{{$entry.Link}}">{{$entry.Title}}</a></li>
        {{end}}
    </ul>

</body>
{{end}}
`
