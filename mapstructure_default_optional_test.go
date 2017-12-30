package mapstructure

import (
	"testing"
)

func TestOptional(t *testing.T) {
	type Person struct {
		ID         int    `json:"id"`
		Name       string `json:"name,omitempty"`
		Address    string `json:"address,required"`
		IsRequired bool   `json:"is_required,omitempty"`
	}
	cases := []struct {
		input   map[string]interface{}
		wantErr bool
	}{
		{
			input: map[string]interface{}{
				"id":          10,
				"address":     "abc",
				"is_required": true,
			},
			wantErr: false,
		},
		{
			input: map[string]interface{}{
				"address":     "abc",
				"is_required": true,
			},
			wantErr: false,
		},
		{
			input: map[string]interface{}{
				"id":          10,
				"is_required": true,
			},
			wantErr: true,
		},
		{
			input: map[string]interface{}{
				"address": "abc",
			},
			wantErr: false,
		},
	}
	for _, c := range cases {
		var result Person
		config := &DecoderConfig{
			Result:        &result,
			FieldRequired: true,
			TagName:       "json",
		}
		decoder, err := NewDecoder(config)
		if err != nil {
			t.Error(err)
		}
		err = decoder.Decode(c.input)
		if c.wantErr {
			if err == nil {
				t.Error("want error but error is nil")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}

}
