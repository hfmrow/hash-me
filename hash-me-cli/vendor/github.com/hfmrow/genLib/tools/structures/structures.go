// structures.go

/*
	Source file auto-generated on Wed, 02 Sep 2020 20:25:56 using Gotk3ObjHandler v1.5 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020 H.F.M - Basic configuration structure's methods that hold.
	Read, Write, Comparaison and File Signature capabilities.
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package structures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// StructureHandler: Structure that hold some methods used to manage
// option/configuration disigned structures.
type StructureHandler struct {
	// Used to hold previous structure content/state for comparison.
	lastjsonString string
}

func StructureHandlerNew() (sh *StructureHandler) {
	return new(StructureHandler)
}

// StructRead: structure from Json file
func (sh *StructureHandler) StructRead(structure interface{}, filename string) error {

	textFileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(textFileBytes, &structure)
}

// StructWrite: structure to Json file
func (sh *StructureHandler) StructWrite(structure interface{}, filename string) error {

	var out bytes.Buffer

	jsonData, err := json.Marshal(structure)
	if err != nil {
		return err
	} else if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
		return ioutil.WriteFile(filename, out.Bytes(), 0644)
	}
	return err
}

// StructHasChanged: Determine whether the structure was changed
// since the last check.
func (sh *StructureHandler) StructHasChanged(structure interface{}) (changed bool) {

	jsonData, err := json.Marshal(&structure)
	if err == nil {

		jsonString := string(jsonData)
		if len(sh.lastjsonString) > 0 {

			if jsonString != sh.lastjsonString {
				changed = true
			}
		}
		sh.lastjsonString = jsonString

	} else {
		log.Printf("StructChanged: %v", err)
	}

	return
}

// StructSignMake: create file signature using provided arguments
func (sh *StructureHandler) StructSignMake(Name, Vers, YearCreat, Creat, Repository, LicenseAbrv string) []string {

	return []string{Name, Vers, "©" + YearCreat, Creat, Repository, LicenseAbrv}
}

// StructSignVersChk: Checking for version 'major, minor, micro'
// accept input as string and/or interger. Return true whether the
// given 'vers'  >= 'min'
func (sh *StructureHandler) StructSignVersChk(vers, min interface{}) bool {

	var (
		remDupSpce = func(inStr string) string {

			//	to match 2 or more whitespace symbols inside a string
			remInside := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
			return strings.TrimSpace(remInside.ReplaceAllString(inStr, " "))
		}

		splitAtNum = func(inStr string) (outInt []int) {

			//	to match non numeric inside a string
			toSplit := regexp.MustCompile(`[[:alpha:][:punct:]]`)
			spaceSepared := toSplit.ReplaceAllString(inStr, " ")
			outText := strings.Split(remDupSpce(spaceSepared), " ")

			for _, val := range outText {
				iVal, err := strconv.Atoi(val)
				if err != nil {
					log.Println("StructSignatureVersChk", err)
				}
				outInt = append(outInt, iVal)
			}
			return
		}
		actV = splitAtNum(fmt.Sprintf("%v", vers))
		minV = splitAtNum(fmt.Sprintf("%v", min))
	)

	for idx, aVal := range actV {
		if idx >= len(minV) {
			break
		}
		mVal := minV[idx]
		switch {
		case mVal < aVal:
			return true
		case mVal == aVal:
			continue
		case mVal > aVal:
			return false
		}
	}
	return true
}
