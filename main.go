package texco

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	consts []Constant
)

type Constant struct {
	Name  string
	Value interface{}
}

func (constant Constant) String() string {
	return "=" + constant.Name + "="
}

func AddConstant(name string, value interface{}) {
	constant := Constant{name, value}
	consts = append(consts, constant)
}

func PassFile(path string) (string, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// source = string(source)
	text := string(source)

	for _, val := range consts {
		text = strings.Replace(text, val.String(), fmt.Sprintf("%v", val.Value), -1)
	}
	return text, nil
}
