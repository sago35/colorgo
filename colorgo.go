package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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
	"NONE":    ct.None,
	"BLACK":   ct.Black,
	"RED":     ct.Red,
	"GREEN":   ct.Green,
	"YELLOW":  ct.Yellow,
	"BLUE":    ct.Blue,
	"MAGENTA": ct.Magenta,
	"CYAN":    ct.Cyan,
	"WHITE":   ct.White,
}

func makeColorRule(regexStr, color string) (ColorRule, error) {
	c, ok := ColorMap[strings.ToUpper(color)]
	if !ok {
		c = ct.None
	}

	re, err := regexp.Compile(regexStr)
	if err != nil {
		return ColorRule{}, err
	} else {
		rule := ColorRule{
			RegexStr: regexStr,
			Color:    c,
			Regex:    re,
		}
		return rule, err
	}
}

func Colorize(c *cli.Context, reader io.Reader, writer io.Writer) {
	rules := []ColorRule{}
	if len(c.Args()) > 0 {
		for i := 0; i+1 < len(c.Args()); i += 2 {
			r, err := makeColorRule(c.Args()[i], c.Args()[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "regex error:", err)
				os.Exit(1)
			} else {
				rules = append(rules, r)
			}
		}
	}

	switch c.String("input") {
	case "cp932", "shiftjis":
		reader = transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
	case "eucjp":
		reader = transform.NewReader(reader, japanese.EUCJP.NewDecoder())
	case "jis", "iso2022jp":
		reader = transform.NewReader(reader, japanese.ISO2022JP.NewDecoder())
	}

	scanner := bufio.NewScanner(reader)
	regex := initColorRules(rules)

	for scanner.Scan() {
		line := scanner.Text()
		fa := regex.FindAllIndex([]byte(line), -1)
		s := 0
		for _, x := range fa {
			fmt.Fprint(writer, line[s:x[0]])
			for _, r := range rules {
				if r.Regex.MatchString(line[x[0]:x[1]]) {
					ct.ChangeColor(r.Color, true, ct.None, false)
					break
				}
			}
			fmt.Fprint(writer, line[x[0]:x[1]])
			ct.ResetColor()
			s = x[1]
		}
		fmt.Fprint(writer, line[s:])
		fmt.Fprintln(writer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gocolor.go"
	app.Usage = "colorize stdout by regular expressions"
	app.Version = "0.0.2"
	app.Author = "sago35"
	app.Email = "sago35@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "input encoding",
			Value: "utf8",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output encoding",
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
    Default encoding is utf8 for input and output.
    Supported encodings are below.

        cp932 shiftjis eucjp utf8 jis iso2022jp

    REGEX : Regular expression

    COLOR : Color name

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
		Colorize(c, os.Stdin, os.Stdout)
	}

	app.Run(os.Args)
}
