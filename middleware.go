package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("checking authentication")
			if flag {
				hf(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			hf(w, r)
		}
	}
}
