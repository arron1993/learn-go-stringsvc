package main

import (
    "errors"
    "strings"
)

var ErrEmpty = errors.New("Empty String")

type stringService struct{}

type StringService interface {
    Uppercase(string) (string, error)
    Count(string) int
}

func (stringService) Uppercase(s string) (string, error) {
    if s == "" {
        return "", ErrEmpty
    }
    return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
    return len(s)
}
