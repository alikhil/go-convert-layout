package converter

import (
	"errors"
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

// Create creates new instance of converter using dictName dictionary
func Create(dictName string) (*Converter, error) {
	var dict map[string]string
	var name = strings.ToLower(dictName)
	switch name {
	case "ru":
		dict = ruDict
		break
	case "dvorak":
		dict = dvorakDict
		break
	case "by":
		dict = byDict
		break
	default:
		return nil, errors.New("dict not found")
	}
	var (
		full    []string
		reverse []string
	)
	for k, v := range dict {
		full = append(full, k, v)
		reverse = append(reverse, v, k)

		full = append(full, strings.ToUpper(k), strings.ToUpper(v))
		reverse = append(reverse, strings.ToUpper(v), strings.ToUpper(k))
	}
	return &Converter{strings.NewReplacer(full...), strings.NewReplacer(reverse...)}, nil
}
