package main

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {

	version := getVersion()

	// 文字列が数字のn.n.nでないとエラー
	matched, err := regexp.MatchString(`^[0-9]+.[0-9]+.[0-9]+`, string(version))
	if err != nil {
		t.Errorf("%v", err)
	} else if !matched {
		t.Errorf("version should n.n.n format: %s", string(version))
	}
}
