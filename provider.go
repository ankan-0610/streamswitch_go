package main

import (
	"time"
)

// Provider represents a text provider with a URL
type Provider struct {
	Name    string
	URL     string
	Latency time.Duration
}

// Providers list with their URLs
var providers = []Provider{
	{Name: "Provider1", URL: "http://localhost:5000"},
	{Name: "Provider2", URL: "http://localhost:5001"},
}

// selectProvider simply returns the current provider based on providerIndex
func selectProvider(providerIndex int) *Provider {
	return &providers[providerIndex]
}