package spc

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// parseInt returns nil or an int
//
// If the string is "UNK" or "", it returns nil
// Otherwise it returns the int
func parseInt(s string) (*int, error) {
	s = strings.TrimSpace(s)

	if s == "UNK" || s == "" {
		return nil, nil
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func fetch(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("unexpected response: %s", resp.Status)

	}

	return resp.Body, nil
}
