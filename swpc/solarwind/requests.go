package solarwind

import (
	"encoding/json"
	"io"
	"net/http"
)

func rawRequest[T any](url string, rtsw *T) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &rtsw)
	if err != nil {
		return err
	}

	return nil
}
