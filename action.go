package main

import (
	"bytes"
	"encoding/json"
	"github.com/jpicht/logger"
	"io"
	"net/http"
)

type action struct {
	Action   string      `yaml:"action"`
	URL      string      `yaml:"url"`
	Method   string      `yaml:"method"`
	Expected []int       `yaml:"expected"`
	Data     interface{} `yaml:"data"`
}

func (a *action) Run(r *recipe, l logger.Logger) bool {
	tmp, _ := json.Marshal(a.Data)
	req, err := http.NewRequest(a.Method, r.url(a.URL), bytes.NewBuffer(tmp))
	if err != nil {
		l.Errorf("Creating request failed: '%s'", err)
		return false
	}

	l.Infof("%10s %s", req.Method, req.URL)

	resp, err := http.DefaultClient.Do(req)
	if err != nil && err != io.EOF {
		l.Errorf("Executing request failed: '%s', %#v", err, resp)
		return false
	}

	for i := range a.Expected {
		if a.Expected[i] == resp.StatusCode {
			return true
		}
	}

	l.Errorf("Unexpected result: %d (expected: %#v)", resp.StatusCode, a.Expected)

	return false
}
