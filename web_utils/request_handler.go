package web_utils

import (
	"io"
	"net/http"
	"os"
)

type requestHandler struct {
	saveDir string
	baseUrl string
}

func newRequestHandler(saveDir string, baseUrl string) requestHandler {
	return requestHandler{
		saveDir: saveDir,
		baseUrl: baseUrl,
	}
}

func (r requestHandler) downloadFile(filename string, path string) (status int, err error) {
	if r.saveDir == "" || r.baseUrl == "" {
		return 0, nil
	}

	out, err := os.Create(r.baseUrl + "/" + filename)
	if err != nil {
		return 0, err
	}

	resp, err := http.Get(r.baseUrl + path)
	if err != nil || resp.StatusCode != 200 {
		return resp.StatusCode, err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, err
}
