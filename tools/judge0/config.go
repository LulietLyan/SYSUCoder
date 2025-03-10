package judge0

import (
	"SYSUCODER/boot/configuration"
	"net/http"
)

var (
	config configuration.JudgeConf
	preUrl string
)

func About() (*http.Response, error) {
	url := preUrl + "/about"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-rapidapi-key", config.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}
