package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mj-hagonoy/get-top-ten/handler"
)

func TestGetTopTen(t *testing.T) {
	values := map[string]string{"data": "hello world!\nhello!Hello!!!!!\nhi hi\n"}
	json_data, err := json.Marshal(values)
	if err != nil {
		t.Errorf("json.Marshal: expected error to be nil got %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/top10", bytes.NewBuffer(json_data))
	w := httptest.NewRecorder()

	handler.GetTopTen(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll: expected error to be nil got %v", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode: expected status to be 200 got %d", res.StatusCode)
	}

	var top10 map[string][]map[string]interface{}
	if err := json.Unmarshal(data, &top10); err != nil {
		t.Errorf("json.Unmarshal: expected error to be nil got %v", err)
	}

	if _, ok := top10["top10"]; !ok {
		t.Error("expected field 'top10' to exists")
	}
}
