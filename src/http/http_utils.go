package http

import (
	"io"
	"net/http"
)

func HttpGet(url string) (string, error) {

	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body), nil
}
