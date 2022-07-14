package main

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	ErrURL  = errors.New("url is empty")
	ErrLang = errors.New("lang is empty")
)

var deploy = `
url: https://github.com/Growthers/paaticle
lang: Golang
`

type Deploy struct {
	URL  string
	Lang string
}

func ParseYaml(Yaml string) (res Deploy, err error) {
	y := Yaml
	d := make(map[string]string)
	err = yaml.Unmarshal([]byte(y), &d)
	if err != nil {
		return res, err
	}

	url, ok := d["url"]
	if !ok {
		return res, ErrURL
	}
	lang, ok := d["lang"]
	if !ok {
		return res, ErrLang
	}

	res.URL = url
	res.Lang = lang

	return res, nil
}

func main() {

	fmt.Printf("元のyaml:%+v\n", deploy)
	d, err := ParseYaml(deploy)
	if err != nil {
		panic(err)
	}

	fmt.Printf("関数から返ってきたデータ:\n%v\n%v", d.URL, d.Lang)
}
