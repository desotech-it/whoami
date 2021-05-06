package view

import "io"

type View interface {
	Write(w io.Writer) error
	WriteAsText(w io.Writer)
}
