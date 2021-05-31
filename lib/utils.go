package lib

import (
	"bufio"
	"fmt"
	"github.com/DavidGamba/go-getoptions"
	"io"
	"os"
	v "rcgb/vars"
	"strings"
)

func ProcessArgs(tib, gib, mib, kib, enum bool, prec int) *getoptions.GetOpt {
	Opt := getoptions.New()
	Opt.Bool("help", false, Opt.Alias("h", "?"))
	Opt.Bool("license", false)
	Opt.Bool("version", false, Opt.Alias("V"))
	Opt.BoolVar(&tib, "tib", false, Opt.Alias("t"), Opt.Description("display in TiB"))
	Opt.BoolVar(&gib, "gib", true, Opt.Alias("g"), Opt.Description("display in GiB"))
	Opt.BoolVar(&mib, "mib", false, Opt.Alias("m"), Opt.Description("display in MiB"))
	Opt.BoolVar(&kib, "kib", false, Opt.Alias("k"), Opt.Description("display in KiB"))
	Opt.BoolVar(&enum, "enum", false, Opt.Alias("e"), Opt.Description("enumerate results"))
	Opt.IntVar(&prec, "precision", 2, Opt.Alias("p"), Opt.Description("show results with a precision on N decimal places"))

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

func PrintRemaining(remaining []string) {
	for i := range remaining {
		fmt.Println(strings.TrimSuffix(remaining[i], "\n"))
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
