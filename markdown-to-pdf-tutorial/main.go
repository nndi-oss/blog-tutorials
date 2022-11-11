package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	pdf "github.com/stephenafamo/goldmark-pdf"
	"github.com/yuin/goldmark"
)

func main() {
	markdown := goldmark.New(
		goldmark.WithRenderer(
			pdf.New(
				pdf.WithTraceWriter(os.Stdout),
				pdf.WithContext(context.Background()),
				pdf.WithImageFS(os.DirFS(".")),
				// pdf.WithLinkColor("cc4578"),
				pdf.WithHeadingFont(pdf.GetTextFont("Arial", pdf.FontCourier)),
				pdf.WithBodyFont(pdf.GetTextFont("Arial", pdf.FontCourier)),
				pdf.WithCodeFont(pdf.GetCodeFont("Arial", pdf.FontCourier)),
			),
		),
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "cache")
		indexHtmlTemplate := template.Must(template.ParseFiles("index.html"))
		indexHtmlTemplate.ExecuteTemplate(w, "index.html", nil)
	})

	mux.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		source, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		if err := markdown.Convert([]byte(source), &buf); err != nil {
			http.Error(w, "failed to generate pdf", http.StatusInternalServerError)
			return
		}
		base64Encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
		dataURL := fmt.Sprintf("data:application/octet-stream;base64,%s", base64Encoded)
		w.Write([]byte(dataURL))
	})

	http.ListenAndServe(":8001", mux)
}
