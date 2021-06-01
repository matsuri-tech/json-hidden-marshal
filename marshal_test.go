package json_hidden_marshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	cases := []struct {
		in       interface{}
		expected string
	}{
		{
			// hidden:- or hidden:true to skip, hidden:mask to mask
			in: struct {
				Name     string `json:"name"`
				Name2    string `json:"name2" hidden:"-"`       // skip
				Name3    string `json:"name3" hidden:"true"`    // skip
				Password string `json:"password" hidden:"mask"` // masked
			}{
				Name:     "foo",
				Password: "password",
			},
			expected: `{"name":"foo","password":"********"}`,
		},
		{
			// nested type
			in: struct {
				Nested struct {
					Open   string `json:"open"`
					Hidden string `json:"hidden" hidden:"-"`
				} `json:"nested"`
			}{
				Nested: struct {
					Open   string `json:"open"`
					Hidden string `json:"hidden" hidden:"-"`
				}{
					Open:   "open",
					Hidden: "hidden",
				},
			},
			expected: `{"nested":{"open":"open"}}`,
		},
		{
			// Without json tags
			in: struct {
				Name     string
				Password string `hidden:"-"`
			}{
				Name:     "foo",
				Password: "password",
			},
			expected: `{"Name":"foo"}`,
		},
	}

	for _, c := range cases {
		out, err := Marshal(c.in)
		if err != nil {
			t.Errorf("%+v\n", err)
		}

		assert.JSONEq(t, string(out), c.expected)
	}
}
