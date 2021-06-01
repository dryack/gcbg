/*The MIT License (MIT)

Copyright (c) 2021 David Ryack

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package lib

import (
	"bufio"
	"errors"
	"fmt"
	v "gcgb/vars"
	"github.com/DavidGamba/go-getoptions"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	tib      bool
	gib      bool
	mib      bool
	kib      bool
	enum     bool
	Suppress bool
	Prec     int
)

func ProcessArgs() *getoptions.GetOpt {
	Opt := getoptions.New()
	Opt.SetMode(getoptions.Bundling)
	Opt.Bool("help", false, Opt.Alias("h", "?"))
	Opt.Bool("license", false)
	Opt.Bool("version", false, Opt.Alias("V"))
	Opt.BoolVar(&tib, "tib", false, Opt.Alias("t"), Opt.Description("display in TiB"))
	Opt.BoolVar(&gib, "gib", false, Opt.Alias("g"), Opt.Description("display in GiB"))
	Opt.BoolVar(&mib, "mib", false, Opt.Alias("m"), Opt.Description("display in MiB"))
	Opt.BoolVar(&kib, "kib", false, Opt.Alias("k"), Opt.Description("display in KiB"))
	Opt.BoolVar(&enum, "enum", false, Opt.Alias("e"), Opt.Description("enumerate results"))
	Opt.BoolVar(&Suppress, "suppress", false, Opt.Alias("s"), Opt.Description("suppress SI postfix"))
	Opt.IntVar(&Prec, "precision", 2, Opt.Alias("p"), Opt.Description("show results with a precision on N decimal places"))

	return Opt
}

func DisplayHelp(opt *getoptions.GetOpt) {
	fmt.Fprintf(os.Stderr, opt.Help())
}

func CheckImmediateExitOpts(opt *getoptions.GetOpt) {
	if opt.Called("help") {
		DisplayHelp(opt)
		os.Exit(0)
	}
	if opt.Called("license") {
		fmt.Fprintf(os.Stderr, v.LicenseText)
		os.Exit(0)
	}
	if opt.Called("version") {
		fmt.Fprintf(os.Stderr, v.ProgVer)
		os.Exit(0)
	}
}

func ReadFromSTDIN() []string {
	var incoming []string
	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "ERROR: #{err}\n\n")
			}
			break
		}
		incoming = append(incoming, s)
	}
	return incoming
}

func PrintRemaining(remaining []string) {
	for i := range remaining {
		fmt.Println(strings.TrimSuffix(remaining[i], "\n"))
	}
}

func GetSIOption() (v.SIMode, error) {
	if tib == true && gib == true && mib == true && kib == true {
		return v.TGMK, nil
	} else if tib == true && gib == true && mib == true {
		return v.TGM, nil
	} else if tib && mib && kib {
		return v.TMK, nil
	} else if tib && gib && kib {
		return v.TGK, nil
	} else if tib && gib {
		return v.TG, nil
	} else if tib && mib {
		return v.TM, nil
	} else if tib && kib {
		return v.TK, nil
	} else if tib {
		return v.T, nil
	} else if gib && mib && kib {
		return v.GMK, nil
	} else if gib && mib {
		return v.GM, nil
	} else if gib && kib {
		return v.GK, nil
	} else if gib {
		return v.G, nil
	} else if mib && kib {
		return v.MK, nil
	} else if mib {
		return v.M, nil
	} else if kib {
		return v.K, nil
	} else {
		return v.G, nil
		//return -1, errors.New("unsupported SI Mode")
	}
}

func produceResult(num float64, siMode v.SIMode) ([]float64, error) {
	var a, b, c, d float64
	a = num / v.TiB
	b = num / v.GiB
	c = num / v.MiB
	d = num / v.KiB
	res := []float64{a, b, c, d}

	switch siMode {
	case 0:
		res[1] = math.NaN()
		res[2] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 1:
		res[0] = math.NaN()
		res[2] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 2:
		res[0] = math.NaN()
		res[1] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 3:
		res[0] = math.NaN()
		res[1] = math.NaN()
		res[2] = math.NaN()
		return res, nil
	case 4:
		res[2] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 5:
		res[1] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 6:
		res[1] = math.NaN()
		res[2] = math.NaN()
		return res, nil
	case 7:
		res[0] = math.NaN()
		res[3] = math.NaN()
		return res, nil
	case 8:
		res[0] = math.NaN()
		res[2] = math.NaN()
		return res, nil
	case 9:
		res[0] = math.NaN()
		res[1] = math.NaN()
		return res, nil
	case 10:
		res[3] = math.NaN()
		return res, nil
	case 11:
		res[1] = math.NaN()
		return res, nil
	case 12:
		res[2] = math.NaN()
		return res, nil
	case 13:
		res[0] = math.NaN()
		return res, nil
	case 14:
		return res, nil
	default:
		res := []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN()}
		return res, errors.New("unhandled SI Mode in produceResult()")
	}
}

func getPostFix(index int) (string, error) {
	switch index {
	case 0:
		return "TiB", nil
	case 1:
		return "GiB", nil
	case 2:
		return "MiB", nil
	case 3:
		return "KiB", nil
	}
	return "n/a", errors.New("unknown postfix in getPostFix()")
}

func DisplayResults(remaining []string, precision int, suppress bool) error {
	si, err := GetSIOption()
	if err != nil {
		return err
	}
	for i := range remaining {
		num, err := strconv.ParseFloat(remaining[i], len(remaining[i]))
		if err != nil {
			return err
		}

		res, _ := produceResult(num, si)
		for j := range res {
			if math.IsNaN(res[j]) {
				continue
			}
			// format each result with the appropriate precision
			formatted := strconv.FormatFloat(res[j], 'f', precision, 64)
			postfix := ""
			if !suppress {
				postfix, _ = getPostFix(j)
			}
			fmt.Fprintf(os.Stdout, "%s %s\n", formatted, postfix)
		}
	}
	return nil
}
