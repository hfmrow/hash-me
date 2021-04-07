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
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

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

		if mainObjects.CheckbuttonMd4.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"MD4:", HashMe(filename, "md4")}, "\t"))
		}
		if mainObjects.CheckbuttonMd5.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"MD5:", HashMe(filename, "md5")}, "\t"))
		}
		if mainObjects.CheckbuttonSha1.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA1:", HashMe(filename, "sha1")}, "\t"))
		}

		if mainObjects.CheckbuttonSha256.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA256:", HashMe(filename, "sha256")}, "\t"))
		}
		if mainObjects.CheckbuttonSha384.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA384:", HashMe(filename, "sha384")}, "\t"))
		}
		if mainObjects.CheckbuttonSha512.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA512:", HashMe(filename, "sha512")}, "\t"))
		}

		if mainObjects.CheckbuttonSha3_256.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_256:", HashMe(filename, "sha3-256")}, "\t"))
		}
		if mainObjects.CheckbuttonSha3_384.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_384:", HashMe(filename, "sha3-384")}, "\t"))
		}
		if mainObjects.CheckbuttonSha3_512.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"SHA3_512:", HashMe(filename, "sha3-512")}, "\t"))
		}

		if mainObjects.CheckbuttonBlake2b256.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b256:", HashMe(filename, "blake2b256")}, "\t"))
		}
		if mainObjects.CheckbuttonBlake2b384.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b384:", HashMe(filename, "blake2b384")}, "\t"))
		}
		if mainObjects.CheckbuttonBlake2b512.GetActive() {
			fileEntry.outStr = append(fileEntry.outStr, strings.Join([]string{"BLAKE2b512:", HashMe(filename, "blake2b512")}, "\t"))
		}
	}

	return fileEntry, err
}

// toTreview: Display to TreeView
func toTreview() {

	var (
		iter *gtk.TreeIter
		err  error
		name string
	)

	tvs.Clear()
	tvs.StoreDetach()
	defer tvs.StoreAttach()

	for _, res := range resultsHash {
		HRopt := HR_UNIT_DEFAULT
		if mainObjects.CheckbuttonUseDecimal.GetActive() {
			HRopt = HR_UNIT_DECIMAL
		}
		name = strings.Join([]string{res.base, HumanReadableSize(res.size, HRopt)}, "\t")

		if len(res.outStr) > 0 {
			for _, row := range res.outStr {
				tmpRow := strings.Split(row, "\t")
				// Convert to []interface
				iSplitted := tvs.ColValuesStringSliceToIfaceSlice(name, tmpRow[1])
				if iter, err = tvs.AddTree(
					colmap["chk"],
					colmap["dat"],
					false,
					nil,
					iSplitted...); err == nil && iter != nil {
					// Add names
					err = tvs.SetColValue(iter, colmap["mtd"], tmpRow[0])
					err = tvs.SetColValue(iter, colmap["fnm"], res.filename)
				}
			}
		}
	}
}

// toDisplay:
func toDisplay(show bool) {

	if !show {
		doIt(show)
		return
	}

	// var results []*entry
	anim, err := GetPixBufAnimation(linearProgressHorzBlue)
	if err != nil {
		log.Fatalf("GetPixBufAnimation: %s\n", err.Error())
	}
	gifImage, err := gtk.ImageNewFromAnimation(anim)
	if err != nil {
		log.Fatalf("ImageNewFromAnimation: %s\n", err.Error())
	}
	pbs = ProgressGifNew(gifImage, mainObjects.BoxMain, 1,
		func() error {
			glib.IdleAdd(func() {
				mainObjects.MainButtonDone.SetSensitive(false)
				mainObjects.GridOptions.SetSensitive(false)
			})
			doIt(show)
			if show {
				toTreview()
			}
			return nil
		},
		func() error {
			mainObjects.MainButtonDone.SetSensitive(true)
			mainObjects.GridOptions.SetSensitive(true)
			SwitchExpandStateSet(mainObjects.SwitchExpand)
			return nil
		})

	go func() {
		pbs.StartGif()
	}()
}

// doIt: do the job
func doIt(display ...bool) {

	var (
		dispOk bool = true
		err    error
		result *entry

		filesToHash,
		tmpFiles []string

		infosHash, content string
		fi                 os.FileInfo

		makeEntry = func(res *entry) (out string) {
			if mainObjects.CheckbuttonShowFilename.GetActive() {
				HRopt := HR_UNIT_DEFAULT
				if mainObjects.CheckbuttonUseDecimal.GetActive() {
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
			err = ioutil.WriteFile(ExtEnsure(f, ".SUM"), []byte(c), 0644)
			Logger.Log(err, "doIt/WriteFile")
		}
	)

	if len(display) > 0 {
		dispOk = display[0]
	}

	if filesChanged {
		// check for directories in list
		for _, f := range files {

			if fi, err = os.Stat(f); err == nil {
				if fi.IsDir() {

					// Scan directory
					depth := 0
					if mainObjects.CheckbuttonRecursiveScan.GetActive() {
						depth = -1
					}
					if tmpFiles, err = ScanDirDepth(f, depth); err == nil {
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
		resultsHash = resultsHash[:0]
		mainStatusbar.Set(fmt.Sprintf("%v", len(files)), 0)
		mainStatusbar.Set(fmt.Sprintf("%v", 0), 1)
		if mainObjects.CheckbuttonConcurrentOp.GetActive() {
			ccs = ConcurrentCalcStrucNew(runtime.NumCPU(),
				func(item interface{}) {
					if result, err = makeHash(item.(string)); err == nil && len(result.outStr) > 0 {
						result.content = strings.Join(result.outStr, GetOsLineEnd())
						resultsHash = append(resultsHash, result)
					}
					Logger.Log(err, "doIt/makeHash")
					glib.IdleAdd(func() {
						mainStatusbar.Set(fmt.Sprintf("%v", ccs.Count), 1)
					})
				})
			ccs.Run(ccs.StringSliceToIfaceSlice(files))
			ccs.Wait()

			glib.IdleAdd(func() {
				mainStatusbar.Set(fmt.Sprintf("%v", ccs.Count), 1)
			})

			sort.SliceStable(resultsHash, func(i, j int) bool {
				return resultsHash[i].filename < resultsHash[j].filename
			})
		} else {
			for idx, file := range files {
				if result, err = makeHash(file); err == nil && len(result.outStr) > 0 {
					result.content = strings.Join(result.outStr, GetOsLineEnd())
					resultsHash = append(resultsHash, result)
				}
				glib.IdleAdd(func() {
					mainStatusbar.Set(fmt.Sprintf("%v", idx+1), 1)
				})
			}
		}
		filesChanged = false
	}

	for _, res := range resultsHash {
		// Computing results for display
		infosHash += makeEntry(res)
		// Creation of .SUM file if needed
		if !dispOk {
			content = makeEntry(res)
			if mainObjects.CheckbuttonAddReminder.GetActive() {
				content += mainOptions.ReminderMessage
			}
			if mainObjects.CheckbuttonCreateFile.GetActive() {
				go writeFile(res.filename, content)
			}
		}
	}
	// Display to TextView
	if dispOk {
		glib.IdleAdd(func() {
			buff.SetText(infosHash)
		})
	} else {
		// To clipboard
		glib.IdleAdd(func() {
			clipboard.SetText(infosHash)
		})
		mainStatusbar.Set("Copied to clipboard", 1)
	}
	return
}

// func doIt(display ...bool) {

// 	var (
// 		dispOk bool = true
// 		err    error
// 		result *entry

// 		wgIn  = new(sync.WaitGroup)
// 		wgOut = new(sync.WaitGroup)

// 		count, cpuCount, cpuNb int
// 		filesToHash,
// 		tmpFiles []string

// 		infosHash, content string
// 		fi                 os.FileInfo

// 		makeEntry = func(res *entry) (out string) {
// 			if mainObjects.CheckbuttonShowFilename.GetActive() {
// 				HRopt := UNIT_DEFAULT
// 				if mainObjects.CheckbuttonUseDecimal.GetActive() {
// 					HRopt = UNIT_DECIMAL
// 				}
// 				base := strings.Join([]string{"FILE:", res.base,
// 					HumanReadableSize(res.size, HRopt)}, "\t")
// 				out += strings.Join([]string{base, res.content, GetOsLineEnd()}, GetOsLineEnd())
// 			} else {
// 				out += res.content + GetOsLineEnd()
// 			}
// 			return
// 		}
// 		writeFile = func(f, c string) {
// 			err = ioutil.WriteFile(ExtEnsure(f, ".SUM"), []byte(c), 0644)
// 			Logger.Log(err, "doIt/WriteFile")
// 		}
// 		calcHash = func(file string) {
// 			if result, err = makeHash(file); err == nil && len(result.outStr) > 0 {
// 				result.content = strings.Join(result.outStr, GetOsLineEnd())
// 				resultsHash = append(resultsHash, result)
// 			}
// 			Logger.Log(err, "doIt/makeHash")
// 			count++
// 			glib.IdleAdd(func() {
// 				mainStatusbar.Set(fmt.Sprintf("%v", count), 1)
// 			})
// 			if mainObjects.CheckbuttonConcurrentOp.GetActive() {
// 				wgOut.Done()
// 				wgIn.Done()
// 			}
// 		}
// 	)

// 	if len(display) > 0 {
// 		dispOk = display[0]
// 	}

// 	if filesChanged {
// 		// check for directories in list
// 		for _, f := range files {

// 			if fi, err = os.Stat(f); err == nil {
// 				if fi.IsDir() {

// 					// Scan directory not recursively
// 					if tmpFiles, err = ScanDirDepth(f, 0); err == nil {
// 						for _, file := range tmpFiles {

// 							if !IsExistSl(filesToHash, file) {
// 								filesToHash = append(filesToHash, file)
// 							}
// 						}
// 					}
// 					Logger.Log(err, "doIt/ScanDirDepth")
// 				} else {
// 					if !IsExistSl(filesToHash, f) {
// 						filesToHash = append(filesToHash, f)
// 					}
// 				}
// 			}
// 			Logger.Log(err, "doIt/Stat")
// 		}

// 		files = filesToHash
// 		resultsHash = resultsHash[:0]

// 		if mainObjects.CheckbuttonConcurrentOp.GetActive() {
// 			wgOut.Add(len(files))
// 			cpuCount = runtime.NumCPU()
// 			cpuNb = 1
// 			wgIn.Add(cpuCount)
// 		}
// 		for _, file := range files {
// 			if mainObjects.CheckbuttonConcurrentOp.GetActive() {
// 				go calcHash(file)

// 				if cpuNb == cpuCount {

// 					wgIn.Wait()
// 					if (len(files) - count) < cpuCount {
// 						wgIn.Add(len(files) - count)
// 					} else {
// 						wgIn.Add(cpuCount)
// 					}
// 					cpuNb = 0
// 				}
// 				cpuNb++
// 			} else {
// 				calcHash(file)
// 			}
// 		}
// 		if mainObjects.CheckbuttonConcurrentOp.GetActive() {
// 			wgOut.Wait()
// 		}
// 		sort.SliceStable(resultsHash, func(i, j int) bool {
// 			return resultsHash[i].filename < resultsHash[j].filename
// 		})
// 		filesChanged = false
// 	}

// 	for _, res := range resultsHash {

// 		infosHash += makeEntry(res)
// 		// Creation of .SUM file if needed
// 		if !dispOk {
// 			content = makeEntry(res)
// 			if mainObjects.CheckbuttonAddReminder.GetActive() {
// 				content += mainOptions.ReminderMessage
// 			}
// 			go writeFile(res.filename, content)
// 		}
// 	}
// 	// if mainObjects.CheckbuttonAddReminder.GetActive() {
// 	// 	infosHash += mainOptions.ReminderMessage
// 	// }

// 	// Display to TextView
// 	if dispOk {
// 		glib.IdleAdd(func() {
// 			buff.SetText(infosHash)
// 		})
// 	} else {
// 		// To clipboard
// 		glib.IdleAdd(func() {
// 			clipboard.SetText(infosHash)
// 		})
// 		mainStatusbar.Set("Copied to clipboard", 1)
// 	}

// 	/*else {
// 		mainStatusbar.Set("Nothing was done", 1)
// 	}*/
// 	return
// }

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
