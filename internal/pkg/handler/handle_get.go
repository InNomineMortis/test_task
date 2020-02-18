package handler

import (
	"encoding/json"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//TODO answer
	answerJSON, _ := json.Marshal(answer)

	w.Write([]byte(answerJSON))

}