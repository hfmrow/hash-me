// miscFunct.go

/*
	Source file auto-generated on Sat, 19 Dec 2020 15:06:30 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020 H.F.M - Hash Me github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// updateStatusbar:
func updateStatusbar() {
	mainStatusbar.CleanAll()
	mainStatusbar.Set(fmt.Sprintf("%v", len(files)), 0)
}

// makeHash:
func makeHash(filename string) (outStr []string, err error) {

	var (
		fi   os.FileInfo
		iter *gtk.TreeIter
	)

	if fi, err = os.Stat(filename); err == nil {

		HRopt := UNIT_DEFAULT
		if mainObjects.CheckbuttonUseDecimal.GetActive() {
			HRopt = UNIT_DECIMAL
		}

		fName := strings.Join([]string{"FILE:",
			filepath.Base(filename), HumanReadableSize(float64(fi.Size()),
				HRopt)}, "\t")

		if mainObjects.CheckbuttonMd4.GetActive() {
			outStr = append(outStr, strings.Join([]string{"MD4:", HashMe(filename, "md4")}, "\t"))
		}
		if mainObjects.CheckbuttonMd5.GetActive() {
			outStr = append(outStr, strings.Join([]string{"MD5:", HashMe(filename, "md5")}, "\t"))
		}
		if mainObjects.CheckbuttonSha1.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA1:", HashMe(filename, "sha1")}, "\t"))
		}

		if mainObjects.CheckbuttonSha256.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA256:", HashMe(filename, "sha256")}, "\t"))
		}
		if mainObjects.CheckbuttonSha384.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA384:", HashMe(filename, "sha384")}, "\t"))
		}
		if mainObjects.CheckbuttonSha512.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA512:", HashMe(filename, "sha512")}, "\t"))
		}

		if mainObjects.CheckbuttonSha3_256.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA3_256:", HashMe(filename, "sha3-256")}, "\t"))
		}
		if mainObjects.CheckbuttonSha3_384.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA3_384:", HashMe(filename, "sha3-384")}, "\t"))
		}
		if mainObjects.CheckbuttonSha3_512.GetActive() {
			outStr = append(outStr, strings.Join([]string{"SHA3_512:", HashMe(filename, "sha3-512")}, "\t"))
		}

		if mainObjects.CheckbuttonBlake2b256.GetActive() {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b256:", HashMe(filename, "blake2b256")}, "\t"))
		}
		if mainObjects.CheckbuttonBlake2b384.GetActive() {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b384:", HashMe(filename, "blake2b384")}, "\t"))
		}
		if mainObjects.CheckbuttonBlake2b512.GetActive() {
			outStr = append(outStr, strings.Join([]string{"BLAKE2b512:", HashMe(filename, "blake2b512")}, "\t"))
		}

		/* Display to TreeView */
		if len(outStr) > 0 {

			name := strings.Split(fName, "\t")[1:]
			for _, row := range outStr {

				tmpRow := strings.Split(row, "\t")
				// Convert to []interface
				iSplitted := tvs.ColValuesStringSliceToIfaceSlice(strings.Join(name, " "), tmpRow[1])

				if iter, err = tvs.AddTree(

					colmap["chk"],
					colmap["dat"],
					false,
					nil,
					iSplitted...); err == nil && iter != nil {
					// Add hash method
					err = tvs.SetColValue(iter, colmap["mtd"], tmpRow[0])
					err = tvs.SetColValue(iter, colmap["fnm"], filename)
				}
			}

			// Prepend filename if requested
			if mainObjects.CheckbuttonShowFilename.GetActive() {
				outStr = append([]string{fName}, outStr...)
			}

			// Add newline if there is something present (for further usage)
			if len(outStr) > 0 {
				outStr = append(outStr, GetOsLineEnd())
			}
		}
	}
	return
}

// doIt: do the job
func doIt(display ...bool) {

	var (
		ok, dispOk bool = false, true
		err        error
		results,
		filesToHash,
		tmpFiles []string

		infosHash, content string
		fi                 os.FileInfo

		// defer function.
		attachAndExpand = func() {
			tvs.StoreAttach()
			SwitchExpandStateSet(mainObjects.SwitchExpand)
		}
	)

	if len(display) > 0 {
		dispOk = display[0]
	}

	// check for directories in list

	for _, f := range files {

		if fi, err = os.Stat(f); err == nil {
			if fi.IsDir() {

				// Scan directory not recursively
				if tmpFiles, err = ScanDirDepth(f, 0); err == nil {
					for _, file := range tmpFiles {

						if !IsExistSl(filesToHash, file) {
							filesToHash = append(filesToHash, file)
						}
					}
				}
				Logger.Log(err, "doIt/ScanDirDepth")
			} else {
				if !IsExistSl(filesToHash, f) {
					filesToHash = append(filesToHash, f)
				}
			}
		}
		Logger.Log(err, "doIt/Stat")
	}

	files = filesToHash

	// Prepare treeview for display
	tvs.Clear()
	tvs.StoreDetach()
	defer attachAndExpand()

	for _, file := range files {
		if results, err = makeHash(file); err == nil && len(results) > 0 {

			content = strings.Join(results, GetOsLineEnd())
			infosHash += content
			ok = true

			// Creation of .SUM file if needed
			if mainObjects.CheckbuttonCreateFile.GetActive() {
				if mainObjects.CheckbuttonAddReminder.GetActive() {
					content += mainOptions.ReminderMessage
				}
				ioutil.WriteFile(ExtEnsure(file, ".SUM"), []byte(content), 0644)
			}
		}
		Logger.Log(err, "doIt/makeHash/in")
	}

	if mainObjects.CheckbuttonAddReminder.GetActive() {
		infosHash += mainOptions.ReminderMessage
	}

	if ok {

		// Display to TextView
		if dispOk {

			Logger.Log(err, "doIt/TextViewDisplay.GetBuffer")
			buff.SetText(infosHash)

			// buff.ApplyTagByName("non-editable", buff.GetStartIter(), buff.GetEndIter())

		} else {

			// To clipboard
			clipboard.SetText(infosHash)

		}

		mainStatusbar.Set("Copied to clipboard", 1)
	} else {
		mainStatusbar.Set("Nothing was done", 1)
	}

	Logger.Log(err, "doIt/makeHash/out")
}

/* Experimental Transparent window TESTs */

var alphaSupported = false

func Transparent(widget *gtk.Widget, alpha float64) {

	// Needed for transparency
	widget.SetAppPaintable(true)

	widget.Connect("screen-changed", func(wdgt *gtk.Widget, oldScreen *gdk.Screen, userData ...interface{}) {
		changed(wdgt)
	})
	widget.Connect("draw", func(wdgt *gtk.Widget, context *cairo.Context) {
		exposeDraw(wdgt, context, alpha)
	})

	changed(widget)
}

func changed(widget *gtk.Widget) {
	screen, _ := widget.GetScreen()
	visual, _ := screen.GetRGBAVisual()

	if visual != nil {
		alphaSupported = true
	} else {
		println("Alpha not supported")
		alphaSupported = false
	}

	widget.SetVisual(visual)
}

func exposeDraw(wdgt *gtk.Widget, ctx *cairo.Context, alpha float64) {
	if alphaSupported {
		ctx.SetSourceRGBA(0.0, 0.0, 0.0, alpha)
	} else {
		ctx.SetSourceRGB(0.0, 0.0, 0.0)
	}

	ctx.SetOperator(cairo.OPERATOR_SOURCE)
	ctx.Paint()
}

/*

var alphaSupported = false

func TransparentWindow(win *gtk.Window, alpha float64) {

	// Needed for transparency
	win.SetAppPaintable(true)

	win.Connect("screen-changed", func(widget *gtk.Widget, oldScreen *gdk.Screen, userData ...interface{}) {
		screenChanged(widget)
	})
	win.Connect("draw", func(window *gtk.Window, context *cairo.Context) {
		exposeDraw(window, context, alpha)
	})

	screenChanged(&win.Widget)
}

func screenChanged(widget *gtk.Widget) {
	screen, _ := widget.GetScreen()
	visual, _ := screen.GetRGBAVisual()

	if visual != nil {
		alphaSupported = true
	} else {
		println("Alpha not supported")
		alphaSupported = false
	}

	widget.SetVisual(visual)
}

func exposeDraw(w *gtk.Window, ctx *cairo.Context, alpha float64) {
	if alphaSupported {
		ctx.SetSourceRGBA(0.0, 0.0, 0.0, alpha)
	} else {
		ctx.SetSourceRGB(0.0, 0.0, 0.0)
	}

	ctx.SetOperator(cairo.OPERATOR_SOURCE)
	ctx.Paint()
}

*/
