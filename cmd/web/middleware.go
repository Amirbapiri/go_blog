package main

import "net/http"

//LoadSession loads and saves the session on every request
func LoadSession(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
