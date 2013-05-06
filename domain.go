// Package domain allows for extracting the TLD.
//
// This is adapted from the PHP source at:
//   http://www.dkim-reputation.org/regdom-libs/
package domain

import (
	"fmt"
	"strings"
)

const (
	star = "*"
	bang = "!"
)

var (
	empty = tld{}
	pop   = tld{bang: empty}
)

type tld map[string]tld

// Extract the TLD.
func TLD(domain string) (string, error) {
	registered, err := Registered(domain)
	if err != nil {
		return "", err
	}
	return registered[strings.Index(registered, ".")+1:], nil
}

// Extract the base domain, including the TLD.
func Registered(domain string) (string, error) {
	parts := strings.Split(domain, ".")
	partsLen := len(parts)
	if partsLen < 2 || parts[partsLen-2] == "" {
		return "", fmt.Errorf("Invalid domain structure: %s", domain)
	}
	registered := find(parts, root)
	if registered == "" || strings.Index(registered, ".") < 0 {
		return "", fmt.Errorf("Invalid domain: %s", domain)
	}
	return registered, nil
}

func find(parts []string, current tld) string {
	const popChar = "#"
	partsLen := len(parts)
	part := ""
	if partsLen > 0 {
		part = parts[partsLen-1]
	}
	if current == nil {
		return part
	}
	if current[bang] != nil {
		return popChar
	}
	next := current[part]
	if next == nil {
		next = current[star]
	}
	if next == nil {
		return part
	}
	if partsLen > 0 {
		rest := find(parts[:partsLen-1], next)
		if rest == popChar {
			return part
		} else if len(rest) > 0 {
			return rest + "." + part
		}
	}
	return ""
}
