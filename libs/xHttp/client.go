package xHttp

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/models"
	"github.com/yangjinguang/wechat-server/libs/logger"
	"encoding/json"
)

func Get(url string, query []models.KeyValue, header []models.KeyValue) (resHeader http.Header, body []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resHeader, body, err
	}
	return do(request, query, header)
}

func Post(url string, data map[string]interface{}, query []models.KeyValue, header []models.KeyValue) (resHeader http.Header, body []byte, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return resHeader, body, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return resHeader, body, err
	}
	return do(request, query, header)
}

func Put(url string, data map[string]interface{}, query []models.KeyValue, header []models.KeyValue) (resHeader http.Header, body []byte, err error) {
	jsonData, err := json.Marshal(data)
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return resHeader, body, err
	}
	return do(request, query, header)
}

func Delete(url string, query []models.KeyValue, header []models.KeyValue) (resHeader http.Header, body []byte, err error) {
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return resHeader, body, err
	}
	return do(request, query, header)
}

func do(request *http.Request, query []models.KeyValue, header []models.KeyValue) (resHeader http.Header, body []byte, err error) {
	c := &http.Client{}

	//request.Header.Set("Authorization", "Bearer "+h.AccessToken)
	request.Header.Set("Content-Type", "application/json")
	if query != nil {
		q := request.URL.Query()
		for _, v := range query {
			q.Add(v.Key, v.Value)
		}
		request.URL.RawQuery = q.Encode()
	}

	if header != nil {
		for _, v := range header {
			request.Header.Set(v.Key, v.Value)
		}
	}
	logger.Debug(request.URL.String())
	resp, err := c.Do(request)

	if err != nil {
		return http.Header{}, body, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	logger.Debug("statusCode:", resp.StatusCode)
	if resp.StatusCode >= 400 {
		return resp.Header, body, errors.New(fmt.Sprintf("%s\n%s", resp.Status, string(body)))
	}
	if err != nil {
		return resp.Header, body, err
	}

	return resp.Header, body, err
}
