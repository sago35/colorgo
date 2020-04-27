package colorgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
)

// ColorRule ...
type ColorRule struct {
	RegexStr string
	Regex    *regexp.Regexp
	Color    ct.Color
}

func initColorRules(rules []ColorRule) *regexp.Regexp {
	regex := []string{}
	for i := range rules {
		regex = append(regex, rules[i].RegexStr)
	}
	return regexp.MustCompile(strings.Join(regex, `|`))
}

// ColorMap ...
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

// MakeColorRule ...
func MakeColorRule(regexStr, color string) (ColorRule, error) {
	c, ok := ColorMap[strings.ToUpper(color)]
	if !ok {
		c = ct.None
	}

	re, err := regexp.Compile(regexStr)
	if err != nil {
		return ColorRule{}, err
	}

	rule := ColorRule{
		RegexStr: regexStr,
		Color:    c,
		Regex:    re,
	}
	return rule, err
}

// Colorize ...
func Colorize(rules []ColorRule, reader io.Reader, writer io.Writer) {
	regex := initColorRules(rules)

	scanner := bufio.NewScanner(reader)

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
