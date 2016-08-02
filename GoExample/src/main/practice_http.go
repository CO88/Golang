package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,s)
}

type Struct struct {
	Greeting 	string
	Punct 		string
	Who			string
}

func (h *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Greeting, h.Punct, h.Who)
}

