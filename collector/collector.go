package collector

import (
	"io"
	"net/http"
	"strings"
)

type collector struct {
	targetAddresses []string
}

func (c collector) GetBody(address string) ([]byte, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c collector) GetAllLinks(address string) ([]string, error) {
	body, err := c.GetBody(address)
	if err != nil {
		return nil, err
	}
	//href="blah.com/xyz" note two brackets
	bodyString := string(body)
	var links []string

	for strings.Contains(bodyString, "href") {
		firstIndex := strings.Index(bodyString, "href=\"")
		secondIndex := strings.Index(bodyString[firstIndex+6:], "\"") + firstIndex + 6
		links = append(links, bodyString[firstIndex+6:secondIndex])
		bodyString = bodyString[secondIndex+1:] // TODO potential out of bounds w/ bad html
	}
	return links, err
}

func GetAllElementsOfType(element string) []string {}

func GetTableRows(tableName string) [][]string {}

func GetTableColumns(tableName string) [][]string {}
