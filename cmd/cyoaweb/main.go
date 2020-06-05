package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the server on")
	fileName := flag.String("file", "story.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in file %s \n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := JsonStory(f)
	if err != nil {
		panic(err)
	}

	// tpl := template.Must(template.New("").Parse(storyTemplate))

	// h := NewHandler(story, WithTemplate(tpl), WithPathFunc(pathFn))
	h := NewHandler(story)
	fmt.Printf("starting the server at: %d\n", *port)

	mux := http.NewServeMux()
	mux.Handle("/", h)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// func pathFn(r *http.Request) string {
// 	path := r.URL.Path
// 	if path == "/story" || path == "/story/" {
// 		path = "/story/intro"
// 	}
// 	return path[len("/story/"):]
// }

// var storyTemplate = `
// <!DOCTYPE html>
// <html lang="en">
// <head>
//   <meta charset="UTF-8">
//   <meta name="viewport" content="width=device-width, initial-scale=1.0">
//   <title>Choose your Own Adventure</title>
// </head>
// <body>
// 	<section class="page">
// 		<h1>{{.Title}}</h1>
// 		{{range .Paragraphs}}
// 			<p>{{.}}</p>
// 		{{end}}
// 		<ul>
// 			{{range .Options}}
// 				<li> <a href="/story/{{.Chapter}}">{{.Text}}</a> </li>
// 			{{end}}
// 		</ul>
// 	</section>

// 	<style>
// 		body {
// 			font-family: helvetica, arial;
// 		}
// 		h1 {
// 			text-align:center;
// 			position:relative;
// 		}
// 		.page {
// 			width: 80%;
// 			max-width: 500px;
// 			margin: auto;
// 			margin-top: 40px;
// 			margin-bottom: 40px;
// 			padding: 80px;
// 			background: #FFFCF6;
// 			border: 1px solid #eee;
// 			box-shadow: 0 10px 6px -6px #777;
// 		}
// 		ul {
// 			border-top: 1px dotted #ccc;
// 			padding: 10px 0 0 0;
// 			-webkit-padding-start: 0;
// 		}
// 		li {
// 			padding-top: 10px;
// 		}
// 		a,
// 		a:visited {
// 			text-decoration: none;
// 			color: #6295b5;
// 		}
// 		a:active,
// 		a:hover {
// 			color: #7792a2;
// 		}
// 		p {
// 			text-indent: 1em;
// 		}
// 	</style>
// </body>
// </html>
// `
