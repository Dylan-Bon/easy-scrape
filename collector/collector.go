package collector

import (
	"io"
	"net/http"
)

type collector struct {
	targetAddresses []string
}

func (c collector) getBody(address string) (io.ReadCloser, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func getAllLinks() []string {}

func getAllElementsOfType(element string) []string {}

func getTableRows(tableName string) [][]string {}

func getTableColumns(tableName string) [][]string {}
