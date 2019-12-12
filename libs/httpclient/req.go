package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(address string) (b []byte, err error) {
	req, err := http.Get(address)
	if err != nil {
		return
	}
	defer req.Body.Close()

	b, err = ioutil.ReadAll(req.Body)
	return
}

func Post(urlAddress string, data interface{}) (r []byte, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	body := bytes.NewReader(b)
	res, err := http.Post(urlAddress, "application/json", body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err = ioutil.ReadAll(res.Body)
	return
}
