package colorgo

import (
	"testing"
)

func TestInitColorRules(t *testing.T) {
	{
		rules := []ColorRule{
			{RegexStr: `foo`},
		}
		regex := initColorRules(rules)
		if regex.String() != `foo` {
			t.Fatal("regex not matched")
		}
	}

	{
		rules := []ColorRule{
			{RegexStr: `foo`},
			{RegexStr: `bar`},
		}
		regex := initColorRules(rules)
		if regex.String() != `foo|bar` {
			t.Fatal("regex not matched")
		}
	}

	{
		rules := []ColorRule{
			{RegexStr: `foo`},
			{RegexStr: `bar`},
			{RegexStr: `baz\w+`},
		}
		regex := initColorRules(rules)
		if regex.String() != `foo|bar|baz\w+` {
			t.Fatal("regex not matched")
		}
	}
}
