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
	utils "gcgb/lib"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	opt := utils.ProcessArgs()
	// early bail if there are no args
	if len(os.Args) < 2 {
		fmt.Println("ERROR: no arguments provided\n")
		utils.DisplayHelp(opt)
		os.Exit(1)
	}
	remaining, err := opt.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
		utils.DisplayHelp(opt)
		os.Exit(1)
	}

	// options that result in immediate exit, such as --version or --help
	utils.CheckImmediateExitOpts(opt)
	prec := utils.Prec
	suppress := utils.Suppress

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		//utils.PrintRemaining(remaining)
		err := utils.DisplayResults(remaining, prec, suppress)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
			utils.DisplayHelp(opt)
			os.Exit(1)
		}
	} else {
		//read from STDIN (presumably a pipe)
		fromStdin := utils.ReadFromSTDIN()
		// positional arguments if any
		remaining = append(fromStdin, remaining...)
		err := utils.DisplayResults(remaining, prec, suppress)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
			utils.DisplayHelp(opt)
			os.Exit(1)
		}
	}

}
