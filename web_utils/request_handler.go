package web_utils

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"strings"
)

type requestHandler struct {
	saveDir string
	baseUrl string
}

func NewRequestHandler(saveDir string, baseUrl string) requestHandler {
	return requestHandler{
		saveDir: saveDir,
		baseUrl: baseUrl,
	}
}

// IsTLSCertValid checks for a TLS certification against a domain.
func IsTLSCertValid(url string) (bool, error) {
	url = GetDomain(url)
	conn, err := tls.Dial("tcp", url, nil)
	if err != nil {
		return false, err
	}
	err = conn.VerifyHostname(url)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetDomain(url string) string {
	strings.TrimPrefix(url, "https://")
	strings.TrimPrefix(url, "http://")
	strings.TrimPrefix(url, "www.")
	i := strings.Index(url, "/")
	return url[0:i]
}

// DownloadFile makes a GET request to baseUrl + path.
func (r requestHandler) DownloadFile(filename string, path string) (statusCode int, err error) {
	if r.saveDir == "" || r.baseUrl == "" {
		return 0, nil
	}

	out, err := os.Create(r.saveDir + "/" + filename)
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
