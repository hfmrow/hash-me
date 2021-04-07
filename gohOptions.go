// gohOptions.go

/*
	Source file auto-generated on Tue, 06 Apr 2021 22:04:43 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020-21 hfmrow - Hash Me v1.2 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gotk3/gotk3/gtk"

	glco "github.com/hfmrow/genLib/crypto"
	glfs "github.com/hfmrow/genLib/files"
	glfssf "github.com/hfmrow/genLib/files/scanFileDir"
	glss "github.com/hfmrow/genLib/slices"
	glsg "github.com/hfmrow/genLib/strings"
	gltscj "github.com/hfmrow/genLib/tools/concurrentJob"
	gltsle "github.com/hfmrow/genLib/tools/log2file"
	gltsushe "github.com/hfmrow/genLib/tools/units/human_readable"

	gimc "github.com/hfmrow/gtk3Import/misc"
	gitsww "github.com/hfmrow/gtk3Import/tools/window"
	gitw "github.com/hfmrow/gtk3Import/treeview"
)

// Application infos. Only this section could be [modified during an update].
// Except for "Descr" variable, it is not recommended to manualy change
// options values relative to the project. Use GOH instead to doing this,
// or respect strictly the original applied format.
var (
	Name         = "Hash Me"
	Vers         = "v1.2"
	Descr        = "Create hash checksum for given files, allow to create .SUM files corresponding to each file. Includes Md4, Md5, Sha1, Sha256, Sha384, Sha512, Sha3_256, Sha3_384, Sha3_512, Blake2b256, Blake2b384, Blake2b512 methods."
	Creat        = "hfmrow"
	YearCreat    = "2020-21"
	LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
	LicenseAbrv  = "License (MIT)"
	Repository   = "github.com/hfmrow/hash-me"

	// Internal var declarations
	mainOptions *MainOpt
	devMode,
	doTempDir bool

	absoluteRealPath,
	optFilename,
	tempDir string

	/*
	 * Library mapping
	 */

	// Window decoration
	wds                       *gitsww.WinDecorationStructure
	WinDecorationStructureNew = gitsww.WinDecorationStructureNew

	// Treeview
	tvs                  *gitw.TreeViewStructure
	TreeViewStructureNew = gitw.TreeViewStructureNew
	columns              [][]string
	colmap               map[string]int

	// Errors handling
	Log2FileStructNew = gltsle.Log2FileStructNew
	Logger            *gltsle.Log2FileStruct

	// D&D
	dndWindow,
	dndTextView *gimc.DragNDropStruct
	DragNDropNew = gimc.DragNDropNew

	// Clipboard / Css / Statusbar
	clipboard             *gimc.Clipboard
	ClipboardNew          = gimc.ClipboardNew
	CssWdgScnBytes        = gimc.CssWdgScnBytes
	mainStatusbar         *gimc.StatusBar
	StatusBarStructureNew = gimc.StatusBarStructureNew

	// Concurrent job
	ccs                    *gltscj.ConcurrentCalcStruc
	ConcurrentCalcStrucNew = gltscj.ConcurrentCalcStrucNew

	// Files
	HumanReadableSize = gltsushe.HumanReadableSize
	HR_UNIT_DEFAULT   = gltsushe.UNIT_DEFAULT
	HR_UNIT_DECIMAL   = gltsushe.UNIT_DECIMAL
	ExtEnsure         = glfs.ExtEnsure
	HashMe            = glco.HashMe
	GetOsLineEnd      = glsg.GetOsLineEnd
	ScanDirDepth      = glfssf.ScanDirDepth

	// Slice
	IsExistSl = glss.IsExistSl

	// Misc
	files        []string
	filesChanged bool

	// TextView buffer
	buff *gtk.TextBuffer

	// Progressbar
	pbs            *gimc.ProgressBarStruct
	ProgressGifNew = gimc.ProgressGifNew

	FormatText = glsg.FormatText

	resultsHash []*entry
)

// MainOpt: This structure contains all the variables of the application, they
// will be saved when exiting and reloaded at launch.
type MainOpt struct {

	// File signature
	FileSign []string

	// Window position
	MainWinWidth,
	MainWinHeight,
	MainWinPosX,
	MainWinPosY int

	LanguageFilename string // In case where GOTranslate is used.

	ShowSplash,
	MakeOutputFile,
	RecursiveScan,
	ConcurrentOp,
	ShowFilename,
	AppendDroppedFiles,
	UseDecimal,
	Reminder,
	Md4, Md5,
	Sha1, Sha256, Sha384, Sha512,
	Sha3_256, Sha3_384, Sha3_512,
	Blake2b256, Blake2b384, Blake2b512 bool

	ReminderMessage,
	CurrentStackPage string

	SwitchStackPage,
	SwitchExpandState bool
}

// Init: Main options initialisation, Put here default values for your application.
func (opt *MainOpt) Init() {

	opt.FileSign = []string{Name, Vers, "©" + YearCreat, Creat, Repository, LicenseAbrv}
	opt.LanguageFilename = "assets/lang/eng.lang"
	opt.MainWinWidth = 800
	opt.MainWinHeight = 600

	opt.ShowSplash = true
	opt.ShowFilename = true
	opt.Reminder = true
	opt.Md5 = true
	// opt.Sha256 = true
	opt.Sha512 = true

	opt.ReminderMessage = `HowTo:	Open a command prompt and use these commands regarding your OS,
     	according to desired checksum type, MD5 | SHA256 | SHA512 ...
Win:	CertUtil -hashfile filename MD5 | SHA1 | SHA256 | SHA384 | SHA512
Linux:	md5sum filename | sha256sum filename | shasum384 filename | sha512sum filename | b2sum -l256 filename | b2sum -l512 filename
OS X:	md5 filename | shasum -a256 filename | shasum -a384 filename | shasum -a512 filename
`

}

// UpdateObjects: Options -> Objects. Put here options to assign to gtk3 objects at start
func (opt *MainOpt) UpdateObjects() {

	// With GtkApplicationWindow (does not happen with GtkWindow) I have strange behavior
	// when updating window size and position, sometimes width, height is not restored
	// successfully, I have tried to figure it out but after a few (long) times I resigned
	// myself to using a workaround method, right now I am using a timer that runs 'count'
	// times the same commands to finally get the desired result (set window's size with
	// the previously saved values).
	// count := 5
	// glib.TimeoutAdd(uint(64), func() bool {

	mainObjects.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)
	mainObjects.MainWindow.Move(opt.MainWinPosX, opt.MainWinPosY)

	// count--
	// return count > 0
	// })

	mainObjects.CheckbuttonAddReminder.SetActive(opt.Reminder)
	mainObjects.CheckbuttonMd4.SetActive(opt.Md4)
	mainObjects.CheckbuttonMd5.SetActive(opt.Md5)
	mainObjects.CheckbuttonSha1.SetActive(opt.Sha1)
	mainObjects.CheckbuttonSha256.SetActive(opt.Sha256)
	mainObjects.CheckbuttonSha384.SetActive(opt.Sha384)
	mainObjects.CheckbuttonSha512.SetActive(opt.Sha512)
	mainObjects.CheckbuttonSha3_256.SetActive(opt.Sha3_256)
	mainObjects.CheckbuttonSha3_384.SetActive(opt.Sha3_384)
	mainObjects.CheckbuttonSha3_512.SetActive(opt.Sha3_512)
	mainObjects.CheckbuttonBlake2b256.SetActive(opt.Blake2b256)
	mainObjects.CheckbuttonBlake2b384.SetActive(opt.Blake2b384)
	mainObjects.CheckbuttonBlake2b512.SetActive(opt.Blake2b512)
	mainObjects.CheckbuttonShowFilename.SetActive(opt.ShowFilename)
	mainObjects.CheckbuttonAppendFiles.SetActive(opt.AppendDroppedFiles)
	mainObjects.CheckbuttonUseDecimal.SetActive(opt.UseDecimal)
	mainObjects.CheckbuttonConcurrentOp.SetActive(opt.ConcurrentOp)
	mainObjects.CheckbuttonRecursiveScan.SetActive(opt.RecursiveScan)
	mainObjects.CheckbuttonCreateFile.SetActive(opt.MakeOutputFile)

	if len(opt.CurrentStackPage) > 0 {
		mainObjects.Stack.SetVisibleChildName(opt.CurrentStackPage)
	}
	mainObjects.SwitchTreeView.SetActive(opt.SwitchStackPage)
	mainObjects.SwitchExpand.SetActive(opt.SwitchExpandState)
	SwitchExpandStateSet(mainObjects.SwitchExpand)

	mainObjects.CheckbuttonShowSplash.SetActive(opt.ShowSplash)
}

// UpdateOptions: Objects -> Options. Put here the gtk3 objects whose
// values you want to save in the options structure on exit.
func (opt *MainOpt) UpdateOptions() {

	opt.MainWinWidth, opt.MainWinHeight = mainObjects.MainWindow.GetSize()
	opt.MainWinPosX, opt.MainWinPosY = mainObjects.MainWindow.GetPosition()

	opt.Reminder = mainObjects.CheckbuttonAddReminder.GetActive()
	opt.Md4 = mainObjects.CheckbuttonMd4.GetActive()
	opt.Md5 = mainObjects.CheckbuttonMd5.GetActive()
	opt.Sha1 = mainObjects.CheckbuttonSha1.GetActive()
	opt.Sha256 = mainObjects.CheckbuttonSha256.GetActive()
	opt.Sha384 = mainObjects.CheckbuttonSha384.GetActive()
	opt.Sha512 = mainObjects.CheckbuttonSha512.GetActive()
	opt.Sha3_256 = mainObjects.CheckbuttonSha3_256.GetActive()
	opt.Sha3_384 = mainObjects.CheckbuttonSha3_384.GetActive()
	opt.Sha3_512 = mainObjects.CheckbuttonSha3_512.GetActive()
	opt.Blake2b256 = mainObjects.CheckbuttonBlake2b256.GetActive()
	opt.Blake2b384 = mainObjects.CheckbuttonBlake2b384.GetActive()
	opt.Blake2b512 = mainObjects.CheckbuttonBlake2b512.GetActive()
	opt.ShowFilename = mainObjects.CheckbuttonShowFilename.GetActive()
	opt.AppendDroppedFiles = mainObjects.CheckbuttonAppendFiles.GetActive()
	opt.UseDecimal = mainObjects.CheckbuttonUseDecimal.GetActive()
	opt.ConcurrentOp = mainObjects.CheckbuttonConcurrentOp.GetActive()
	opt.RecursiveScan = mainObjects.CheckbuttonRecursiveScan.GetActive()
	opt.MakeOutputFile = mainObjects.CheckbuttonCreateFile.GetActive()

	opt.CurrentStackPage = mainObjects.Stack.GetVisibleChildName()
	opt.SwitchStackPage = mainObjects.SwitchTreeView.GetActive()
	opt.SwitchExpandState = mainObjects.SwitchExpand.GetActive()

	opt.ShowSplash = mainObjects.CheckbuttonShowSplash.GetActive()
}

// Read: Options from file.
func (opt *MainOpt) Read() (err error) {

	opt.Init() // Init options with defaults values

	textFileBytes, err := ioutil.ReadFile(optFilename)
	if err != nil {
		return err
	}
	return json.Unmarshal(textFileBytes, &opt)
}

// Write: Options to file
func (opt *MainOpt) Write() error {

	var out bytes.Buffer
	opt.UpdateOptions()

	jsonData, err := json.Marshal(&opt)
	if err != nil {
		return err
	} else if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
		return ioutil.WriteFile(optFilename, out.Bytes(), 0644)
	}
	return err
}
