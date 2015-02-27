package main

import (
	"bufio"
	"code.google.com/p/mahonia"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
	"os"
	"regexp"
	"strings"
)

type ColorRule struct {
	RegexStr string
	Regex    *regexp.Regexp
	Color    ct.Color
}

func initColorRules(rules []ColorRule) *regexp.Regexp {
	regex := []string{}
	for i, _ := range rules {
		regex = append(regex, rules[i].RegexStr)
	}
	return regexp.MustCompile(strings.Join(regex, `|`))
}

var ColorMap map[string]ct.Color = map[string]ct.Color{
	"None":    ct.None,
	"Black":   ct.Black,
	"Red":     ct.Red,
	"Green":   ct.Green,
	"Yellow":  ct.Yellow,
	"Blue":    ct.Blue,
	"Magenta": ct.Magenta,
	"Cyan":    ct.Cyan,
	"White":   ct.White,
}

func makeColorRule(regexStr, color string) ColorRule {
	c, ok := ColorMap[color]
	if !ok {
		c = ct.None
	}

	rule := ColorRule{
		RegexStr: regexStr,
		Color:    c,
		Regex:    regexp.MustCompile(regexStr),
	}
	return rule
}

func main() {
	app := cli.NewApp()
	app.Name = "gocolor.go"
	app.Usage = "文字コードを変換しつつ、正規表現で色づけを行う"
	app.Version = "0.0.1"
	app.Author = "sago35"
	app.Email = "sago35@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "入力文字コード",
			Value: "utf8",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "出力文字コード",
			Value: "utf8",
		},
	}
	cli.AppHelpTemplate = `NAME:
    {{.Name}} - {{.Usage}}

USAGE:
    {{.Name}} [OPTIONS] [ REGEX  COLOR ]*

VERSION:
    {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
    {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
    {{.Email}}{{end}}{{end}}

OPTIONS:
    {{range .Flags}}{{.}}
    {{end}}

OTHER:
    入出力文字コードは、デフォルト値は入力はutf8、出力はutf8となっています
    以下を設定可能です

        cp932 shijtjis eucjp utf8

    REGEXは、色づけを行う正規表現
    COLORは、以下の色の名前を指定可能です

        None
        Black
        Red
        Green
        Yellow
        Blue
        Magenta
        Cyan
        White
`

	app.Action = func(c *cli.Context) {
		rules := []ColorRule{}
		if len(c.Args()) > 0 {
			for i := 0; i + 1 < len(c.Args()); i += 2 {
				rules = append(rules, makeColorRule(c.Args()[i], c.Args()[i+1]))
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		regex := initColorRules(rules)

		for scanner.Scan() {
			line := mahonia.NewDecoder(c.String("input")).ConvertString(scanner.Text())
			fa := regex.FindAllIndex([]byte(line), -1)
			s := 0
			for _, x := range fa {
				fmt.Print(mahonia.NewEncoder(c.String("output")).ConvertString(line[s:x[0]]))
				for _, r := range rules {
					if r.Regex.MatchString(line[x[0]:x[1]]) {
						ct.ChangeColor(r.Color, true, ct.None, false)
						break
					}
				}
				fmt.Print(mahonia.NewEncoder(c.String("output")).ConvertString(line[x[0]:x[1]]))
				ct.ResetColor()
				s = x[1]
			}
			fmt.Print(mahonia.NewEncoder(c.String("output")).ConvertString(line[s:]))
			fmt.Println()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	app.Run(os.Args)
}
