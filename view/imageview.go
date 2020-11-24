package view

import (
	"html/template"
	"io"
	"io/ioutil"
)

type ImageView struct {
	Filename string
	AltText  string
}

func NewImageView(filename string) *ImageView {
	return &ImageView{
		filename,
		filename,
	}
}

func getTemplateFile(filename string) string {
	content, err := ioutil.ReadFile("template/" + filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func (v *ImageView) Write(w io.Writer) error {
	templateText := getTemplateFile("image.html")
	template := template.Must(template.New("imageview").Parse(templateText))
	return template.Execute(w, v)
}

// convenience function
func WriteImageView(w io.Writer, filename string) error {
	iv := NewImageView(filename)
	return iv.Write(w)
}
