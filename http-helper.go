package u_push

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPPostJSON with json []byte
func httpPostJSON(reqUrl string, headers map[string]string, payload io.Reader, reqTimeout time.Duration) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "text/json"
	return httpRequest("POST", reqUrl, headers, payload, reqTimeout)
}

func httpRequest(method string, reqUrl string, headers map[string]string, payload io.Reader, reqTimeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest(method, reqUrl, payload)
	if err != nil {
		return nil, errors.New("make http request error: " + err.Error())
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	/*ctx := req.Context()
	nextCtx, _ := context.WithDeadline(ctx, time.Now().Add(HttpReqTimeout))
	req.WithContext(nextCtx)*/

	client := &http.Client{
		Timeout: reqTimeout,
	}
	resp, err := client.Do(req)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
			}
		}()
	}

	if err != nil {
		return nil, errors.New("do http request error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ready http response error: " + err.Error())
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf(fmt.Sprintf("HttpStatusCode:%d, Desc:%s", resp.StatusCode, string(body)))
	}

	return body, nil
}
