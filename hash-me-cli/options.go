// options.go

/*
	CLI version
	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	glco "github.com/hfmrow/genLib/crypto"
	glfs "github.com/hfmrow/genLib/files"
	glfssf "github.com/hfmrow/genLib/files/scanFileDir"
	glss "github.com/hfmrow/genLib/slices"
	glsg "github.com/hfmrow/genLib/strings"
	gltsss "github.com/hfmrow/genLib/tools/structures"
	gltsushe "github.com/hfmrow/genLib/tools/units/human_readable"
)

// Application infos. Only this section could be [modified during an update].
// Except for "Descr" variable, it is not recommended to manualy change
// options values relative to the project. Use GOH instead to doing this,
// or respect strictly the original applied format.
var (
	Name         = "Hash Me CLI version"
	Vers         = "v1.1"
	Descr        = "Create hash checksum for given files, allow to create .SUM files corresponding to each file. Includes Md4, Md5, Sha1, Sha256, Sha384, Sha512, Sha3_256, Sha3_384, Sha3_512, Blake2b256, Blake2b384, Blake2b512 methods."
	Creat        = "H.F.M"
	YearCreat    = "2020-21"
	LicenseShort = "This program comes with absolutely no warranty. See The License (MIT) for details: https://opensource.org/licenses/mit-license.php"
	LicenseAbrv  = "License (MIT)"
	Repository   = "https://github.com/hfmrow/hash-me/hash-me-cli"

	// Internal var declarations
	opt *MainOpt

	/*
	 * Library mapping
	 */

	// Files
	HumanReadableSize = gltsushe.HumanReadableSize
	HR_UNIT_DECIMAL   = gltsushe.UNIT_DECIMAL
	HR_UNIT_DEFAULT   = gltsushe.UNIT_DEFAULT
	ExtEnsure         = glfs.ExtEnsure
	HashMe            = glco.HashMe
	GetOsLineEnd      = glsg.GetOsLineEnd
	GetAbsRealPath    = glfs.GetAbsRealPath
	ScanDirDepth      = glfssf.ScanDirDepth

	ToCamel    = glsg.ToCamel
	FormatText = glsg.FormatText

	// Slice
	IsExistSl = glss.IsExistSl
)

type MainOpt struct {
	// Structure handler
	sh       *gltsss.StructureHandler
	FileSign []string

	ShowFilename,
	UseDecimal,

	SaveToFile,
	SUMSingleFile,

	Md4, Md5,
	Sha1, Sha256, Sha384, Sha512,
	Sha3_256, Sha3_384, Sha3_512,
	Blake2b256, Blake2b384, Blake2b512 bool

	Files  []string
	Output string
}

func (o *MainOpt) Init() {

	o.ShowFilename = true
}
