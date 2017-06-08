package main

import (
	"bytes"
	"flag"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
)

func TestMakeColorRule(t *testing.T) {
	{
		cr, err := makeColorRule("", "red")
		if err != nil {
			t.Fatal("failed to call makeColorRule(): %s", err)
		}
		if cr.Regex.String() != "" {
			t.Fatal("regex not matched")
		}
		if cr.Color != ct.Red {
			t.Fatal("color not matched")
		}
	}

	{
		cr, err := makeColorRule("", "Red")
		if err != nil {
			t.Fatal("failed to call makeColorRule(): %s", err)
		}
		if cr.Regex.String() != "" {
			t.Fatal("regex not matched")
		}
		if cr.Color != ct.Red {
			t.Fatal("color not matched")
		}
	}

	{
		cr, err := makeColorRule("", "ReD")
		if err != nil {
			t.Fatal("failed to call makeColorRule(): %s", err)
		}
		if cr.Regex.String() != "" {
			t.Fatal("regex not matched")
		}
		if cr.Color != ct.Red {
			t.Fatal("color not matched")
		}
	}

	{
		cr, err := makeColorRule("", "RED")
		if err != nil {
			t.Fatal("failed to call makeColorRule(): %s", err)
		}
		if cr.Regex.String() != "" {
			t.Fatal("regex not matched")
		}
		if cr.Color != ct.Red {
			t.Fatal("color not matched")
		}
	}

	cr2, err2 := makeColorRule("", "")
	if err2 != nil {
		t.Fatal("failed to call makeColorRule(): %s", err2)
	}
	if cr2.Color != ct.None {
		t.Fatal("color not matched")
	}

}

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

func BenchmarkColorize(b *testing.B) {
	stdin := bytes.NewBufferString("foo\n")
	stdout := new(bytes.Buffer)
	set := flag.NewFlagSet("test", 0)
	set.String("input", "utf8", "")
	set.String("output", "utf8", "")
	c := cli.NewContext(nil, set, nil)

	for i := 0; i < b.N; i++ {
		Colorize(c, stdin, stdout)
	}
}
