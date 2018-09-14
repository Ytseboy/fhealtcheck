package config

import (
	"log"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {
	sample_endpoints := []EndPoint{}

	sample_config := Config{
		Endpoints: sample_endpoints,
	}

	dummy_endpoint := EndPoint{
		ContentRequirement: "google",
		URL:                "google.com",
	}

	sample_config.Endpoints = append(sample_endpoints, dummy_endpoint)

	conf, err := LoadConfigFromFile("sample_config.json")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	if conf.Endpoints[0].URL != sample_config.Endpoints[0].URL {
		t.Fail()
	}
}
