package main

import (
	"github.com/jpicht/logger"
)

type recipe struct {
	Name    string    `yaml:"name"`
	BaseURL string    `yaml:"baseUrl"`
	Actions []*action `yaml:"actions"`
}

func (r *recipe) url(suffix string) string {
	return r.BaseURL + suffix
}

func (r *recipe) Run(l logger.Logger) {
	l.Infof("Starting test '%s'", r.Name)

	for actionI := range r.Actions {
		if !r.Actions[actionI].Run(r, l) {
			l.Error("Action failed.")
			return
		}
	}
	l.Info("All is well.")
}
