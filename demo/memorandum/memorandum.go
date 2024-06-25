package memorandum

import (
	"errors"
	"fmt"
	"time"
)

type IEditor interface {
	Title(title string)
	Content(content string)
	Save()
	Undo() error
	Redo() error
	Show()
}

type Memento struct {
	title      string
	content    string
	createTime int64
}

func newMemento(title, content string) *Memento {
	return &Memento{
		title:      title,
		content:    content,
		createTime: time.Now().Unix(),
	}
}

type Editor struct {
	title    string
	content  string
	versions []*Memento
	index    int
}

func NewEditor() *Editor {
	return &Editor{
		"",
		"",
		make([]*Memento, 0),
		0,
	}
}

func (e *Editor) Title(title string) {
	e.title = title
}

func (e *Editor) Content(content string) {
	e.content = content
}

func (e *Editor) Save() {
	e.versions = append(e.versions, newMemento(e.title, e.content))
	e.index = len(e.versions) - 1
}

func (e *Editor) Undo() error {
	return e.load(e.index - 1)
}

func (e *Editor) load(i int) error {
	size := len(e.versions)
	if size <= 0 {
		return errors.New("no history versions")
	}

	if i < 0 || i >= size {
		return errors.New("no more history versions")
	}

	it := e.versions[i]
	e.title = it.title
	e.content = it.content
	e.index = i
	return nil
}

func (e *Editor) Redo() error {
	return e.load(e.index + 1)
}

func (e *Editor) Show() {
	fmt.Printf("title: %s, content: %s\n", e.title, e.content)
}
