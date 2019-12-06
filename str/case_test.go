package str

import (
	"testing"
)

// Testing snake
func TestSnake(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "foo_bar"},
		{"FooBar", "foo_bar"},
		{"Foo Bar", "foo_bar"},
		{" Foo Bar", "foo_bar"},
		{"Foo Bar ", "foo_bar"},
		{" Foo Bar ", "foo_bar"},
		{"foo", "foo"},
		{"foo_bar", "foo_bar"},
		{"Foo", "foo"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
		// {"numbers2and55with000", "numbers_2_and_55_with_000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := Snake(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}