package validation

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmptyData        = errors.New("empty data")
	ErrStringValidation = errors.New("string validation error")
	ErrGstValidation    = errors.New("gst validation error")
)

func ValidateString(input string) error {
	str := strings.TrimSpace(input)
	if str == "" {
		return ErrEmptyData
	}
	return nil
}

func ValidateGST(input string) error {
	gst := strings.ToUpper(strings.TrimSpace(input))
	if gst == "" {
		return ErrEmptyData
	}
	r := regexp.MustCompile("^[0-9]{2}[A-Z]{5}[0-9]{4}[A-Z]{1}[1-9A-Z]{1}Z[0-9A-Z]{1}$")
	matched := r.Match([]byte(gst))
	if !matched {
		return ErrGstValidation
	}
	return nil
}
