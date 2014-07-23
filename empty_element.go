package ace

import "io"

// emptyElement represents an empty element.
type emptyElement struct {
	elementBase
}

// Does nothing.
func (e *emptyElement) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newEmpty creates and returns an empty element.
func newEmptyElement(ln *line, rslt *result, parent element) *emptyElement {
	return &emptyElement{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
