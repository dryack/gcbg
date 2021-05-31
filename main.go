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

package main

import (
	"fmt"
	//"github.com/DavidGamba/go-getoptions"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	utils "rcgb/lib"
	v "rcgb/vars"
)

func main() {
	var tib bool
	var gib bool
	var mib bool
	var kib bool
	var enum bool
	var prec int

	opt := utils.ProcessArgs(tib, gib, mib, kib, enum, prec)

	remaining, err := opt.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
		os.Exit(1)
	}

	if opt.Called("help") {
		fmt.Fprintf(os.Stderr, opt.Help())
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

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		utils.PrintRemaining(remaining)
	} else {
		//FIXME:  when both STDIN and args are being used, there program won't exit without a <cr>
		fmt.Println("not a tty")
		//read from STDIN (presumably a pipe)
		utils.PrintRemaining(utils.ReadFromSTDIN())
		// positional arguments if any
		utils.PrintRemaining(remaining)
	}

}
