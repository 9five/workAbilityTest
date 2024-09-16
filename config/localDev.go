//go:build localDev
// +build localDev

package config

import ()

var (
	AllowOrigin []string
)

func init() {
	AllowOrigin = []string{
		"http://localhost:8000",
	}
}
