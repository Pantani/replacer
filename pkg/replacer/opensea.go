package replacer

import (
	"bytes"
	"html/template"
)

const (
	TEMPLATE = `<a href="{{.Url}}">{{.Name}}</a>`
)

type Html string
type Names []Name

type Name struct {
	Name    string `json:"name" bson:"name"`
	Url     string `json:"url,omitempty" bson:"url"`
	Snippet string `json:"-" `
}

func (n *Name) GenerateSnippet() error {
	tpl := template.New("url")
	tpl, err := tpl.Parse(TEMPLATE)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	err = tpl.Execute(&out, n)
	if err != nil {
		return err
	}
	n.Snippet = out.String()
	return nil
}

type NameRequest struct {
	Url string `json:"url"`
}

type Status struct {
	Status bool `json:"status"`
}

type Error struct {
	Message string `json:"message"`
}
