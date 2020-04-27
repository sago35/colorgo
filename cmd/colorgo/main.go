package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sago35/colorgo"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app            = kingpin.New(`colorgo`, `colorize stdout by regular expressions`)
	inputEncoding  = app.Flag("input", "input encoding").Default("utf8").Short('i').String()
	outputEncoding = app.Flag("output", "output encoding").Default("utf8").Short('o').String()
	args           = app.Arg("args", "other args").Strings()
)

func main() {
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	err = run(*args)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(a []string) error {

	var r io.Reader
	r = os.Stdin
	switch *inputEncoding {
	case "cp932", "shiftjis":
		r = transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	case "eucjp":
		r = transform.NewReader(r, japanese.EUCJP.NewDecoder())
	case "iso2022jp":
		r = transform.NewReader(r, japanese.ISO2022JP.NewDecoder())
	default:
		// skip
	}

	var w io.Writer
	w = os.Stdout
	switch *outputEncoding {
	case "cp932", "shiftjis":
		w = transform.NewWriter(w, japanese.ShiftJIS.NewEncoder())
	case "eucjp":
		w = transform.NewWriter(w, japanese.EUCJP.NewEncoder())
	case "iso2022jp":
		w = transform.NewWriter(w, japanese.ISO2022JP.NewEncoder())
	default:
		// skip
	}

	rules, err := makeRules(a)
	if err != nil {
		return err
	}
	colorgo.Colorize(rules, r, w)
	return nil
}

func makeRules(args []string) ([]colorgo.ColorRule, error) {
	rules := []colorgo.ColorRule{}
	if len(args) > 0 {
		for i := 0; i+1 < len(args); i += 2 {
			r, err := colorgo.MakeColorRule(args[i], args[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "regex error:", err)
				os.Exit(1)
			} else {
				rules = append(rules, r)
			}
		}
	}

	return rules, nil
}
