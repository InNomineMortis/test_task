package app

import (
	"fmt"
	"log"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host[:len(r.Host)-4], r.URL.Path)
	target := "https://" + r.Host[:len(r.Host)-4]+ "8080" + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}