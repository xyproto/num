package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	"github.com/xyproto/num"
)

func fracAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("please specify a fraction or a floating point number")
	}
	given := c.Args().Get(0)
	verbose := c.IsSet("verbose")
	iterations := c.Int("maxiterations")

	if strings.Contains(given, ".") {
		s, err := strconv.ParseFloat(given, 64)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Println("iterations:", iterations)
		}
		fmt.Println(num.NewFromFloat64(s, iterations))
		return nil
	}
	if strings.Count(given, ",") == 1 {
		s, err := strconv.ParseFloat(strings.Replace(given, ",", ".", 1), 64)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Println("iterations:", iterations)
		}
		fmt.Println(num.NewFromFloat64(s, iterations))
		return nil
	}
	if strings.Count(given, "/") == 1 {
		n, err := num.NewFromString(given)
		if err != nil {
			return err
		}
		fmt.Println(n)
		return nil
	}
	nf := big.NewFloat(0)
	f, b, err := nf.Parse(given, 10)
	if err != nil {
		return err
	}
	if b != 10 {
		return fmt.Errorf("unexpected base: %d", b)
	}
	r, acc := f.Rat(nil)
	if verbose {
		fmt.Println("accuracy:", acc)
	}
	n := num.NewFromRat(r)
	fmt.Println(n)
	return nil
}

// Quit with a nicely formatted error message to stderr
func quit(err error) {
	msg := err.Error()
	if !strings.HasSuffix(msg, ".") && !strings.HasSuffix(msg, "!") && !strings.Contains(msg, ":") {
		msg += "."
	}
	fmt.Fprintf(os.Stderr, "%s%s\n", strings.ToUpper(string(msg[0])), msg[1:])
	os.Exit(1)
}

func main() {
	app := cli.NewApp()

	app.Name = "fract"
	app.Usage = "convert a float to a fraction, or simplify a fraction"
	app.UsageText = "fract [options] [fraction or floating point number]"

	app.Version = "0.2"
	app.HideHelp = true

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "output version information",
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "verbose output",
		},
		cli.IntFlag{
			Name:  "maxiterations, m",
			Value: -1,
			Usage: "maximum number of interations when converting (-1 for no limit)",
		},
	}

	app.Action = fracAction
	if err := app.Run(os.Args); err != nil {
		quit(err)
	}
}
