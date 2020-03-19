package core

import (
	"log"
	"siteanalyser/core/prcy"
)

type Site struct {
	Domain       string         `json:"domain"`
	Updated      string         `json:"updated,omitempty"`
	PrCyAnalysis *prcy.Analysis `json:"prcy,omitempty"`
}

func NewSite(domain string) Site {
	if domain == "" {
		log.Fatal("Empty domain name")
	}
	return Site{domain, "", nil}
}
