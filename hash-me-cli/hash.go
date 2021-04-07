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
	"runtime"
	"sort"
	"strings"
)

// makeHash:
type entry struct {
	base,
	filename,
	content string
	outStr []string
	size   int64
}

// makeHash:
func makeHash(filename string) (e *entry, err error) {

	var (
		fi        os.FileInfo
		fileEntry = new(entry)
	)

	if fi, err = os.Stat(filename); err == nil {

		fileEntry.size = fi.Size()
		fileEntry.base = filepath.Base(filename)
		fileEntry.filename = filename

		if opt.Md4 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"MD4:", HashMe(filename, "md4")}, "\t"))
		}
		if opt.Md5 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"MD5:", HashMe(filename, "md5")}, "\t"))
		}
		if opt.Sha1 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA1:", HashMe(filename, "sha1")}, "\t"))
		}

		if opt.Sha256 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA256:", HashMe(filename, "sha256")}, "\t"))
		}
		if opt.Sha384 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA384:", HashMe(filename, "sha384")}, "\t"))
		}
		if opt.Sha512 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA512:", HashMe(filename, "sha512")}, "\t"))
		}

		if opt.Sha3_256 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_256:", HashMe(filename, "sha3-256")}, "\t"))
		}
		if opt.Sha3_384 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_384:", HashMe(filename, "sha3-384")}, "\t"))
		}
		if opt.Sha3_512 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_512:", HashMe(filename, "sha3-512")}, "\t"))
		}

		if opt.Blake2b256 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b256:", HashMe(filename, "blake2b256")}, "\t"))
		}
		if opt.Blake2b384 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b384:", HashMe(filename, "blake2b384")}, "\t"))
		}
		if opt.Blake2b512 {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b512:", HashMe(filename, "blake2b512")}, "\t"))
		}
	} else {
		log.Printf("makeHash: %v\n", err)
	}

	return fileEntry, err
}

// doIt: do the job
func doIt() {

	var (
		err         error
		resultsHash []*entry

		infosHash,
		content string

		result *entry

		makeEntry = func(res *entry) (out string) {
			if opt.ShowFilename {
				HRopt := HR_UNIT_DEFAULT
				if opt.UseDecimal {
					HRopt = HR_UNIT_DECIMAL
				}
				base := strings.Join([]string{"FILE:", res.base,
					HumanReadableSize(res.size, HRopt)}, "\t")
				out += strings.Join([]string{base, res.content, GetOsLineEnd()}, GetOsLineEnd())
			} else {
				out += res.content + GetOsLineEnd()
			}
			return
		}
		writeFile = func(f, c string) {
			if opt.AddReminder {
				c += opt.ReminderMessage
			}
			if err = ioutil.WriteFile(ExtEnsure(f, ".SUM"), []byte(c), 0644); err != nil {
				log.Printf("doIt/WriteFile: %v\n", err)
			}
		}
	)

	ccs = ConcurrentCalcStrucNew(runtime.NumCPU(),
		func(item interface{}) {
			if result, err = makeHash(item.(string)); err == nil && len(result.outStr) > 0 {
				result.content = strings.Join(result.outStr, GetOsLineEnd())
				resultsHash = append(resultsHash, result)
			}
		})
	ccs.Run(ccs.StringSliceToIfaceSlice(opt.Files))
	ccs.Wait()

	sort.SliceStable(resultsHash, func(i, j int) bool {
		return resultsHash[i].filename < resultsHash[j].filename
	})

	for _, res := range resultsHash {
		// Computing results for display
		infosHash += makeEntry(res)
		// Creation of .SUM file if needed
		if opt.SaveToFile {
			content = makeEntry(res)
			writeFile(res.filename, content)
		}
	}

	fmt.Printf(infosHash)
	if opt.SUMSingleFile {
		writeFile(ExtEnsure(opt.OutputFilename, ".SUM"), infosHash)
	}
}
