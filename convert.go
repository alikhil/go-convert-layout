package converter

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Converter converts from one layout to english and vice versa
type Converter struct {
	fromReplacer *strings.Replacer
	toReplacer   *strings.Replacer
}

// FromEn converts from eng to converter layout
func (c *Converter) FromEn(msg string) string {
	return c.fromReplacer.Replace(msg)
}

// ToEn converts from converter layout to eng
func (c *Converter) ToEn(msg string) string {
	return c.toReplacer.Replace(msg)
}

// Create creates new instance of converter using dictName.json file in the root
func Create(dictName string) (*Converter, error) {
	var filename = dictName
	if !strings.HasSuffix(dictName, ".json") {
		filename = dictName + ".json"
	}
	var data, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var (
		dict    map[string]string
		full    []string
		reverse []string
	)
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return nil, err
	}
	for k, v := range dict {
		full = append(full, k, v)
		reverse = append(reverse, v, k)

		full = append(full, strings.ToUpper(k), strings.ToUpper(v))
		reverse = append(reverse, strings.ToUpper(v), strings.ToUpper(k))
	}
	return &Converter{strings.NewReplacer(full...), strings.NewReplacer(reverse...)}, nil
}
