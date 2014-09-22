package texco

import (
	"fmt"
	"github.com/robertkrimen/otto"
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
	return "" + constant.Name + ""
}

func AddConstant(name string, value interface{}) {
	constant := Constant{name, value}
	consts = append(consts, constant)
}

func PassString(text string) string {
	Otto := otto.New()

	var processed string = text
	begin := "|"
	end := "?"
	running := true
	for running {
		pointA := strings.Index(processed, begin)
		if pointA != -1 {
			onwards := processed[pointA+1:]
			behind := processed[:pointA]
			pointB := strings.Index(onwards, end)
			if pointB != -1 {
				final := onwards[:pointB]
				//REPLACE CONSTANTS
				for _, val := range consts {
					final = strings.Replace(final, val.String(), fmt.Sprintf("%v", val.Value), -1)
				}
				//Run javascript
				finalOttoValue := ""
				Otto.Run("final = " + final)
				value, err := Otto.Get("final")
				if err == nil {
					finalOttoValue, _ = value.ToString()
					// fmt.Println(finalOttoValue + "OTTO VAL")
				}
				if finalOttoValue != "undefined" {
					final = finalOttoValue
				}
				excess := onwards[pointB+1:]
				processed = behind + final + excess
				// fmt.Println(processed)
			} else {
				running = false
			}
		} else {
			running = false
		}

	}

	for _, val := range consts {
		text = strings.Replace(text, val.String(), fmt.Sprintf("%v", val.Value), -1)
	}
	return processed
}

func PassFile(path string) (string, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// source = string(source)
	text := string(source)

	// for _, val := range consts {
	// 	text = strings.Replace(text, val.String(), fmt.Sprintf("%v", val.Value), -1)
	// }
	return PassString(text), nil
}

func ListConstants() []Constant {
	return consts
}
