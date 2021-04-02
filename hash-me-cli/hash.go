// hash.go

/*
	CLI version
	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// makeHash:
func makeHash(filename string) (outStr []string, err error) {

	var fi os.FileInfo

	if fi, err = os.Stat(filename); err == nil {

		option := HR_UNIT_DEFAULT
		if opt.UseDecimal {
			option = HR_UNIT_DECIMAL
		}

		if opt.ShowFilename {
			outStr = append(outStr, strings.Join([]string{"FILE:",
				filepath.Base(filename), HumanReadableSize(float64(fi.Size()),
					option)}, "\t"))
		}

		if opt.Md4 {
			outStr = append(outStr, strings.Join([]string{"MD4:", HashMe(filename, "md4")}, "\t"))
		}
		if opt.Md5 {
			outStr = append(outStr, strings.Join([]string{"MD5:", HashMe(filename, "md5")}, "\t"))
		}
		if opt.Sha1 {
			outStr = append(outStr, strings.Join([]string{"SHA1:", HashMe(filename, "sha1")}, "\t"))
		}

		if opt.Sha256 {
			outStr = append(outStr, strings.Join([]string{"SHA256:", HashMe(filename, "sha256")}, "\t"))
		}
		if opt.Sha384 {
			outStr = append(outStr, strings.Join([]string{"SHA384:", HashMe(filename, "sha384")}, "\t"))
		}
		if opt.Sha512 {
			outStr = append(outStr, strings.Join([]string{"SHA512:", HashMe(filename, "sha512")}, "\t"))
		}

		if opt.Sha3_256 {
			outStr = append(outStr, strings.Join([]string{"SHA3_256:", HashMe(filename, "sha3-256")}, "\t"))
		}
		if opt.Sha3_384 {
			outStr = append(outStr, strings.Join([]string{"SHA3_384:", HashMe(filename, "sha3-384")}, "\t"))
		}
		if opt.Sha3_512 {
			outStr = append(outStr, strings.Join([]string{"SHA3_512:", HashMe(filename, "sha3-512")}, "\t"))
		}

		if opt.Blake2b256 {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b256:", HashMe(filename, "blake2b256")}, "\t"))
		}
		if opt.Blake2b384 {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b384:", HashMe(filename, "blake2b384")}, "\t"))
		}
		if opt.Blake2b512 {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b512:", HashMe(filename, "blake2b512")}, "\t"))
		}

		// Add newline if there is something present
		if len(outStr) > 0 {
			outStr = append(outStr, GetOsLineEnd())
		}
	}
	return
}

// doIt: do the job
func doIt() {

	var (
		ok      bool
		err     error
		results []string
		infosHash,
		content string
	)

	for _, file := range opt.Files {
		if results, err = makeHash(file); err == nil && len(results) > 0 {

			content = strings.Join(results, GetOsLineEnd())
			infosHash += content
			ok = true

			// Creation of .SUM file if needed
			if opt.SaveToFile {

				if err = ioutil.WriteFile(ExtEnsure(file, ".SUM"), []byte(content), 0644); err != nil {
					log.Printf("doIt/WriteFile: %v\n", err)
				}
			}
		}
	}

	if ok {
		fmt.Printf(infosHash)

		if opt.SUMSingleFile {

			if err = ioutil.WriteFile(ExtEnsure(opt.Output, ".SUM"), []byte(infosHash), 0644); err != nil {
				log.Printf("doIt/SUMSingleFile/WriteFile: %v\n", err)
			}
		}
	}

	if err != nil {
		log.Printf("doIt/makeHash: %v\n", err)
	}
}
