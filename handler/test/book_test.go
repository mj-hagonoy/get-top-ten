package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mj-hagonoy/get-top-ten/handler"
)

var testInputs []map[string]interface{}

func init() {
	data, err := ioutil.ReadFile("./test-data.json")
	if err != nil {
		panic(fmt.Errorf("ioutil.ReadFile: %+v", err))
	}

	if err := json.Unmarshal(data, &testInputs); err != nil {
		panic(fmt.Errorf("json.Unmarshal: %+v", err))
	}
}

func TestGetTopTen(t *testing.T) {
	for _, input := range testInputs {
		json_data, err := json.Marshal(input)
		if err != nil {
			t.Logf("json.Marshal: expected error to be nil got %v", err)
			t.Fail()
			continue
		}

		req := httptest.NewRequest(http.MethodPost, "/top10", bytes.NewBuffer(json_data))
		w := httptest.NewRecorder()

		handler.GetTopTen(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("ioutil.ReadAll: expected error to be nil got %v", err)
			t.Fail()
			continue
		}

		if res.StatusCode != 200 {
			t.Logf("res.StatusCode: expected status to be 200 got %d", res.StatusCode)
			t.Fail()
			continue
		}

		var top10 map[string][]map[string]interface{}
		if err := json.Unmarshal(data, &top10); err != nil {
			t.Logf("json.Unmarshal: expected error to be nil got %v", err)
			t.Fail()
			continue
		}

	}
}
