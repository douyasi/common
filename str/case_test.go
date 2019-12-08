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
		{"JAVAScript", "java_script"},
		{"BBC News", "bbc_news"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
		{"_Any_kind of -string", "_any_kind_of_string"},
		{"-Any-kind-of-string", "any_kind_of_string"},
		{"any kind.of. string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
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

// Testing camel snake
func TestScreamingSnake(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "FOO_BAR"},
		{"FooBar", "FOO_BAR"},
		{"Foo Bar", "FOO_BAR"},
		{" Foo Bar", "FOO_BAR"},
		{"Foo Bar ", "FOO_BAR"},
		{" Foo Bar ", "FOO_BAR"},
		{"foo", "FOO"},
		{"foo_bar", "FOO_BAR"},
		{"Foo", "FOO"},
		{"", ""},
		{"ManyManyWords", "MANY_MANY_WORDS"},
		{"manyManyWords", "MANY_MANY_WORDS"},
		{"AnyKind of_string", "ANY_KIND_OF_STRING"},
		{"JSONData", "JSON_DATA"},
		{"JAVAScript", "JAVA_SCRIPT"},
		{"BBC News", "BBC_NEWS"},
		{"userID", "USER_ID"},
		{"AAAbbb", "AA_ABBB"},
		{"_Any_kind of -string", "_ANY_KIND_OF_STRING"},
		{"-Any-kind-of-string", "ANY_KIND_OF_STRING"},
		{"any kind.of. string", "ANY_KIND_OF_STRING"},
		{"numbers2and55with000", "NUMBERS_2_AND_55_WITH_000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ScreamingSnake(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing camel snake
func TestCamelSnake(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "Foo_Bar"},
		{"FooBar", "Foo_Bar"},
		{"Foo Bar", "Foo_Bar"},
		{" Foo Bar", "Foo_Bar"},
		{"Foo Bar ", "Foo_Bar"},
		{" Foo Bar ", "Foo_Bar"},
		{"foo", "Foo"},
		{"foo_bar", "Foo_Bar"},
		{"Foo", "Foo"},
		{"", ""},
		{"ManyManyWords", "Many_Many_Words"},
		{"manyManyWords", "Many_Many_Words"},
		{"AnyKind of_string", "Any_Kind_Of_String"},
		{"JSONData", "Json_Data"},
		{"JAVAScript", "Java_Script"},
		{"BBC News", "Bbc_News"},
		{"userID", "User_Id"},
		{"AAAbbb", "Aa_Abbb"},
		{"_Any_kind of -string", "_Any_Kind_Of_String"},
		{"-Any-kind-of-string", "Any_Kind_Of_String"},
		{"any kind.of. string", "Any_Kind_Of_String"},
		{"numbers2and55with000", "Numbers_2_And_55_With_000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := CamelSnake(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing kebab
func TestKebab(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "foo-bar"},
		{"FooBar", "foo-bar"},
		{"Foo Bar", "foo-bar"},
		{" Foo Bar", "foo-bar"},
		{"Foo Bar ", "foo-bar"},
		{" Foo Bar ", "foo-bar"},
		{"foo", "foo"},
		{"foo_bar", "foo-bar"},
		{"Foo", "foo"},
		{"", ""},
		{"ManyManyWords", "many-many-words"},
		{"manyManyWords", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"JSONData", "json-data"},
		{"JAVAScript", "java-script"},
		{"BBC News", "bbc-news"},
		{"userID", "user-id"},
		{"AAAbbb", "aa-abbb"},
		{"_Any_kind of -string", "any-kind-of-string"},
		{"-Any-kind-of-string", "-any-kind-of-string"},
		{"any kind.of. string", "any-kind-of-string"},
		{"numbers2and55with000", "numbers-2-and-55-with-000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := Kebab(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing camel kebab
func TestCamelKebab(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "Foo-Bar"},
		{"FooBar", "Foo-Bar"},
		{"Foo Bar", "Foo-Bar"},
		{" Foo Bar", "Foo-Bar"},
		{"Foo Bar ", "Foo-Bar"},
		{" Foo Bar ", "Foo-Bar"},
		{"foo", "Foo"},
		{"foo_bar", "Foo-Bar"},
		{"Foo", "Foo"},
		{"", ""},
		{"ManyManyWords", "Many-Many-Words"},
		{"manyManyWords", "Many-Many-Words"},
		{"AnyKind of_string", "Any-Kind-Of-String"},
		{"JSONData", "Json-Data"},
		{"JAVAScript", "Java-Script"},
		{"BBC News", "Bbc-News"},
		{"userID", "User-Id"},
		{"AAAbbb", "Aa-Abbb"},
		{"_Any_kind of -string", "Any-Kind-Of-String"},
		{"-Any-kind-of-string", "-Any-Kind-Of-String"},
		{"any kind.of. string", "Any-Kind-Of-String"},
		{"numbers2and55with000", "Numbers-2-And-55-With-000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := CamelKebab(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing screaming kebab
func TestScreamingKebab(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "FOO-BAR"},
		{"FooBar", "FOO-BAR"},
		{"Foo Bar", "FOO-BAR"},
		{" Foo Bar", "FOO-BAR"},
		{"Foo Bar ", "FOO-BAR"},
		{" Foo Bar ", "FOO-BAR"},
		{"foo", "FOO"},
		{"foo_bar", "FOO-BAR"},
		{"Foo", "FOO"},
		{"", ""},
		{"ManyManyWords", "MANY-MANY-WORDS"},
		{"manyManyWords", "MANY-MANY-WORDS"},
		{"AnyKind of_string", "ANY-KIND-OF-STRING"},
		{"JSONData", "JSON-DATA"},
		{"JAVAScript", "JAVA-SCRIPT"},
		{"BBC News", "BBC-NEWS"},
		{"userID", "USER-ID"},
		{"AAAbbb", "AA-ABBB"},
		{"_Any_kind of -string", "ANY-KIND-OF-STRING"},
		{"-Any-kind-of-string", "-ANY-KIND-OF-STRING"},
		{"any kind.of. string", "ANY-KIND-OF-STRING"},
		{"numbers2and55with000", "NUMBERS-2-AND-55-WITH-000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ScreamingKebab(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing lower camel
func TestLowerCamel(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "fooBar"},
		{"FooBar", "fooBar"},
		{"Foo Bar", "fooBar"},
		{" Foo Bar", "fooBar"},
		{"Foo Bar ", "fooBar"},
		{" Foo Bar ", "fooBar"},
		{"foo", "foo"},
		{"foo_bar", "fooBar"},
		{"Foo", "foo"},
		{"", ""},
		{"ManyManyWords", "manyManyWords"},
		{"manyManyWords", "manyManyWords"},
		{"AnyKind of_string", "anyKindOfString"},
		{"JSONData", "jsonData"},
		{"JAVAScript", "javaScript"},
		{"BBC News", "bbcNews"},
		{"userID", "userId"},
		{"AAAbbb", "aaAbbb"},
		{"_Any_kind of -string", "anyKindOfString"},
		{"-Any-kind-of-string", "anyKindOfString"},
		{"any kind.of. string", "anyKindOfString"},
		{"numbers2and55with000", "numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := LowerCamel(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing upper camel
func TestUpperCamel(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "FooBar"},
		{"FooBar", "FooBar"},
		{"Foo Bar", "FooBar"},
		{" Foo Bar", "FooBar"},
		{"Foo Bar ", "FooBar"},
		{" Foo Bar ", "FooBar"},
		{"foo", "Foo"},
		{"foo_bar", "FooBar"},
		{"Foo", "Foo"},
		{"", ""},
		{"ManyManyWords", "ManyManyWords"},
		{"manyManyWords", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"JSONData", "JsonData"},
		{"JAVAScript", "JavaScript"},
		{"BBC News", "BbcNews"},
		{"userID", "UserId"},
		{"AAAbbb", "AaAbbb"},
		{"_Any_kind of -string", "AnyKindOfString"},
		{"-Any-kind-of-string", "AnyKindOfString"},
		{"any kind.of. string", "AnyKindOfString"},
		{"numbers2and55with000", "Numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := UpperCamel(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

// Testing uppercase
func TestDelimited(t *testing.T) {
	cases := [][]string{
		{"foo_bar", "foo.bar", "foo:bar"},
		{"FooBar", "foo.bar", "foo:bar"},
		{"Foo Bar", "foo.bar", "foo:bar"},
		{" Foo Bar", "foo.bar", "foo:bar"},
		{"Foo Bar ", "foo.bar", "foo:bar"},
		{" Foo Bar ", "foo.bar", "foo:bar"},
		{"foo", "foo", "foo"},
		{"foo_bar", "foo.bar", "foo:bar"},
		{"Foo", "foo", "foo"},
		{"", "", ""},
		{"ManyManyWords", "many.many.words", "many:many:words"},
		{"manyManyWords", "many.many.words", "many:many:words"},
		{"AnyKind of_string", "any.kind.of.string", "any:kind:of:string"},
		{"JSONData", "json.data", "json:data"},
		{"JAVAScript", "java.script", "java:script"},
		{"BBC News", "bbc.news", "bbc:news"},
		{"userID", "user.id", "user:id"},
		{"AAAbbb", "aa.abbb", "aa:abbb"},
		{"_Any_kind of -string", "any.kind.of.string", "any:kind:of:string"},
		{".Any-kind-of-string", ".any.kind.of.string", "any:kind:of:string"},
		{"any kind.of. string", "any.kind.of.string", "any:kind:of:string"},
		{":any -kind.of. string", ":any.kind.of.string", "::any:kind:of:string"},
		{"numbers2and55with000", "numbers.2.and.55.with.000", "numbers:2:and:55:with:000"},
	}
	for _, i := range cases {
		in := i[0]
		out1 := i[1]
		out2 := i[2]
		result1 := Delimited(in, '.')
		result2 := Delimited(in, ':')
		if result1 != out1 {
			t.Error("'" + in + "'('" + result1 + "' != '" + out1 + "')")
		}
		if result2 != out2 {
			t.Error("'" + in + "'('" + result2 + "' != '" + out2 + "')")
		}
	}
}
