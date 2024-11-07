package main

import "net/http"

// Need to show solution

func authentificate(*http.Request) error

func Authentificate(request *http.Request) error {
	err := authentificate(request)
	if err != nil {
		return err
	}

	return nil
}
