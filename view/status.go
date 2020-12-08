package view

import (
	"encoding/json"
	"io"
)

type ReadinessView struct {
	Status string `json:"status"`
}

type HealthView struct {
	// It's now identical to ReadinessView
	// but it's to be expanded in the future
	Status string `json:"status"`
}

func (v *ReadinessView) Write(w io.Writer) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	_, errWrite := w.Write(bytes)

	return errWrite
}

func (v *HealthView) Write(w io.Writer) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	_, errWrite := w.Write(bytes)

	return errWrite
}
