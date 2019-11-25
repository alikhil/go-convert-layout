package converter

import (
	"testing"
)

func TestConverter_FromEn(t *testing.T) {
	type fields struct {
		c *Converter
	}
	type args struct {
		msg string
	}

	var ruC, err = Create("ru")
	if err != nil {
		t.Fatal(err)
	}
	dvorakC, err := Create("dvorak")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"russian", fields{ruC}, args{"ghbdtn"}, "привет"},
		{"dvorak eng", fields{dvorakC}, args{"jdpps"}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.c
			if got := c.FromEn(tt.args.msg); got != tt.want {
				t.Errorf("Converter.FromEn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_ToEn(t *testing.T) {
	type fields struct {
		c *Converter
	}
	type args struct {
		msg string
	}

	var ruC, err = Create("ru")
	if err != nil {
		t.Fatal(err)
	}
	dvorakC, err := Create("dvorak")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"russian", fields{ruC}, args{"руддщ"}, "hello"},
		{"dvorak eng", fields{dvorakC}, args{"d.nnr"}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.c
			if got := c.ToEn(tt.args.msg); got != tt.want {
				t.Errorf("Converter.ToEn() = %v, want %v", got, tt.want)
			}
		})
	}
}
