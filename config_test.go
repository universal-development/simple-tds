package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	loadEnv()
	loadConfig()

	emptyValue := resolve("tomato", "curl/7.47.0")
	assert.True(t, len(emptyValue) == 0)

	url := resolve("test1", "curl/7.47.0")
	assert.True(t, strings.HasPrefix(url, "http://curl"))

	urlCurl := resolve("test1", "curl")
	assert.True(t, strings.HasPrefix(urlCurl, "http://curl"))

	urlff := resolve("test1", "Firefox2")
	assert.True(t, strings.HasPrefix(urlff, "http://firefox"))

	urlff2 := resolve("test1", "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/2")
	assert.True(t, strings.HasPrefix(urlff2, "http://firefox"))

	urlchrome := resolve("test1", "Chrome")
	assert.True(t, strings.HasPrefix(urlchrome, "http://default-domain"))
}
