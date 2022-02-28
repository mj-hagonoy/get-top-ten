package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mj-hagonoy/get-top-ten/pkg/book"
)

type TopTenRequest struct {
	Data string `json:"data"`
}

type TopTenResponse struct {
	Data []book.Rank `json:"top10"`
}

func GetTopTen(w http.ResponseWriter, r *http.Request) {
	bytesData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		return
	}
	var body TopTenRequest
	if err := json.Unmarshal(bytesData, &body); err != nil {
		http.Error(w, "unable to parse body", http.StatusBadRequest)
		return
	}
	var response TopTenResponse
	response.Data = book.GetTopTenWords(body.Data)

	byteResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(byteResponse)
}
