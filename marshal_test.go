package json_hidden_marshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name     string `json:"name"`
	Name2    string `json:"name2" hidden:"-"`       // skip
	Name3    string `json:"name3" hidden:"true"`    // skip
	Password string `json:"password" hidden:"mask"` // masked
	Nested   UserNested
}

type UserNested struct {
	Open   string
	Hidden string `hidden:"-"`
}

func TestMarshal(t *testing.T) {
	user := User{
		Name:     "name",
		Name2:    "name2",
		Name3:    "name3",
		Password: "password",
		Nested: UserNested{
			Open:   "open",
			Hidden: "in",
		},
	}

	out, err := Marshal(user)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	assert.JSONEq(t, string(out), `{"name":"name","password":"********","Nested":{"Open":"open"}}`)
}
