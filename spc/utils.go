package spc

import (
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
