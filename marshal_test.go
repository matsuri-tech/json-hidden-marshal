package json_hidden_marshal

import "testing"

type User struct {
	Name     string `json:"name"`
	Name2    string `json:"name2" hidden:"-"`       // skip
	Name3    string `json:"name3" hidden:"true"`    // skip
	Password string `json:"password" hidden:"mask"` // masked
}

func TestMarshal(t *testing.T) {
	user := User{
		Name:     "name",
		Name2:    "name2",
		Name3:    "name3",
		Password: "password",
	}

	out, err := Marshal(user)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	if string(out) != `{"name":"name","password":"********"}` {
		t.Errorf("%+v\n", string(out))
	}
}
