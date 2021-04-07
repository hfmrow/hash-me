// initDnd.go

/*
	Source file auto-generated on Mon, 21 Dec 2020 15:29:04 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

func initDnds() {

	// Main window
	dndWindow = DragNDropNew(mainObjects.MainWindow, nil, func() {
		var ok bool
		filenames := *dndWindow.FilesList

		if mainObjects.CheckbuttonAppendFiles.GetActive() {
			for _, file := range filenames {

				if !IsExistSl(files, file) {
					ok = true
					files = append(files, file)
				}
			}
		} else {
			ok = true
			files = filenames
		}

		filesChanged = ok
		toDisplay(true)
	})

	// TextView (don't know why, it override the main window DnD ...)
	dndTextView = DragNDropNew(mainObjects.TextViewDisplay, nil, func() {
		var ok bool
		filenames := *dndTextView.FilesList

		if mainObjects.CheckbuttonAppendFiles.GetActive() {
			for _, file := range filenames {

				if !IsExistSl(files, file) {
					ok = true
					files = append(files, file)
				}
			}
		} else {
			ok = true
			files = filenames
		}
		filesChanged = ok
		toDisplay(true)
	})
}
