package main

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	ErrYaml = errors.New("yaml is empty")
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

	if d["url"] == "" || d["lang"] == "" {
		return res, ErrYaml
	}
	res.URL = d["url"]
	res.Lang = d["lang"]

	return
}

func main() {

	fmt.Printf("元のyaml:%+v\n", deploy)
	d, err := ParseYaml(deploy)
	if err != nil {
		panic(err)
	}

	fmt.Printf("関数から返ってきたデータ:\n%v\n%v", d.URL, d.Lang)
}
