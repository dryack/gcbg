package vars

const (
	Kib         = 1024.0
	MiB         = 1048576.0
	GiB         = 1073742000.0
	TiB         = 1099512000000.0
	ProgVer     = "rgb 0.03"
	AuthAddy    = "<git.lamashtu@gmail.com>"
	LicenseText = "The MIT License (MIT)\nCopyright (c) 2013-2018 David Ryack\n\nPermission is hereby granted, free of " +
		"charge, to any person obtaining a copy of\nthis software and associated documentation files (the \"Software\")," +
		" to deal in\nthe Software without restriction, including without limitation the rights to\nuse, copy, " +
		"modify, merge, publish, distribute, sublicense, and/or sell copies of\nthe Software, and to permit persons " +
		"to whom the Software is furnished to do so,\nsubject to the following conditions:\n\nThe above copyright " +
		"notice and this permission notice shall be included in all\ncopies or substantial portions of the Software." +
		"\n \nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT " +
		"NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS\nFOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN" +
		" NO EVENT SHALL THE AUTHORS OR\nCOPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, " +
		"WHETHER\nIN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN\nCONNECTION WITH THE " +
		"SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.\""
)

var precision uint8 = 2

type SIMode int8

const (
	TGMK SIMode = iota
	TMK
	TGK
	TG
	TM
	TK
	T
	GMK
	GK
	G
	MK
	M
	K
)
