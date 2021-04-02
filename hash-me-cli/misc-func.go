// misc-func.go

/*
	CLI version
	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func chkCmdLine() error {

	var (
		err    error
		length int
		arg    string
		fi     os.FileInfo
	)
	// Default depth scandir
	depth := 0

	if length = len(os.Args); length == 1 {

		dispHelp()
		os.Exit(0)
	}

	// getting arguments.
	for idx := 1; idx < length; idx++ {

		switch os.Args[idx] {

		case "-h", "--help":
			dispHelp()
		case "-s", "--sum":
			opt.SaveToFile = true
		case "-n", "--noname":
			opt.ShowFilename = false
		case "-d", "--decimal":
			opt.UseDecimal = true
		case "-r", "--recursive":
			depth = -1
		case "-Md4":
			opt.Md4 = true
		case "-Md5":
			opt.Md5 = true
		case "-Sha1":
			opt.Sha1 = true
		case "-Sha256":
			opt.Sha256 = true
		case "-Sha384":
			opt.Sha384 = true
		case "-Sha512":
			opt.Sha512 = true
		case "-Sha3_256":
			opt.Sha3_256 = true
		case "-Sha3_384":
			opt.Sha3_384 = true
		case "-Sha3_512":
			opt.Sha3_512 = true
		case "-Blake2b256":
			opt.Blake2b256 = true
		case "-Blake2b384":
			opt.Blake2b384 = true
		case "-Blake2b512":
			opt.Blake2b512 = true
		case "-i", "--input":

			for idx+1 < length {

				idx++
				arg = os.Args[idx]
				if arg == "." {

					arg, err = os.Getwd()
				}

				if fi, err = os.Stat(arg); err == nil {

					if fi.IsDir() {

						opt.Files, err = ScanDirDepth(arg, depth)
						break
					}

					opt.Files = append(opt.Files, arg)
				} else {

					idx--
					err = nil
					break
				}
			}
		case "-o", "--output":

			idx++
			if idx < length {

				opt.SaveToFile = false
				opt.SUMSingleFile = true
				opt.Output = os.Args[idx]
			} else {

				idx--
				err = fmt.Errorf("Bad argument: '%v', but no filename to export to.", os.Args[idx])
			}

		default:

			dispHelp()
			err = fmt.Errorf("Bad argument: %v", os.Args[idx])
		}
	}
	return err
}

/*
./hash-me-cli -Md4 -Md5 -Sha1 -Sha256 -Sha384 -Sha512 -Sha3_256 -Sha3_384 -Sha3_512 -Blake2b256 -Blake2b384 -Blake2b512 -i hash-me-cli
*/

func dispHelp() {
	exec, _ := os.Executable()
	exec = filepath.Base(exec)
	fmt.Printf(`
` + FormatText(strings.Join([]string{Name, Vers, "©" + YearCreat, Creat, Repository}, " "), 80, true) + `
` + FormatText(Descr, 80, true) + `
` + FormatText(LicenseShort, 80, true) + `

usage:
` + exec + ` [OPTION] [HASH-METHODS]... -i [FILES]... -o [OUTPUT-FILE]

[OPTION]

-h, --help	This help.

-s, --sum	Add a '.SUM' file for each given file(s) including choosen
		hash methods.

-r, --recursive	If a directory is given as input file (or dot), determine
		whether	a recursive scan will be done or not, default NOT.

-n --noname	Disable filename display.

-d --decimal	Use decimal for file size calculation. (1000 instead of 1024)

[HASH-METHODS]

	Uses one or more hash methods to calculate the checksum of the
	given file(s).

-Md4, -Md5,
-Sha1, -Sha256, -Sha384, -Sha512,
-Sha3_256, -Sha3_384, -Sha3_512,
-Blake2b256, -Blake2b384, -Blake2b512

	BLAKE hashing methods are impressive in terms of their
	reliability and speed compared to SHAs. Not yet really popular.
	They're present here not for public distribution use but rather
	for personal use on large files (e.g: comparison of .mkv).

[INPUT-FILES]

-i, --input	Indicate that the following files should be considered to be
		hashed.

-o, --output	Specify the name of the file in which you want to place the
		collected checksums. This replaces the "-sum" option, so you
		can only use one or the other.

Example:

` + exec + ` -s -MD5 -SHA1 -i data.json main.go blue.book

	Will create '.SUM' file for 'data.json' 'main.go' 'blue.book',
	including they're MD5 and SHA1 checksum.

` + exec + ` -SHA256 -BLAKE2b256 -i data.json main.go blue.book -o output.SUM

	Calculate the SHA256, SHA512 and BLAKE2b256 checksum for the given
	'data.json' 'main.go' 'blue.book' files and fill a single 'output.SUM'
	file with the informations gathered.
`)
}
