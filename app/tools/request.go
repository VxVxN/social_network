package tools

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnmarshalRequest(body io.ReadCloser, structReq interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if json.Unmarshal(bodyBytes, &structReq); err != nil {
		return err
	}
	return nil
}
