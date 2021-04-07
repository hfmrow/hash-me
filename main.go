// main.go

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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

// main: And at the beginning ... this part is not modified on update.
// Build options informations:
// devMode: is used in some functions to control the behavior of the program
// When software is ready to be published, this flag must be set at "false"
// that means:
// - options file will be stored in $HOME/.config/[Creat]/[softwareName],
// - translate function if used, will no more auto-update "sts" map sentences,
// - all built-in assets will be used instead of the files themselves.
//   Be aware to update assets via "Goh" and translations with "Got" before all.
func main() {

	devMode = true
	absoluteRealPath, optFilename = getAbsRealPath()

	// Initialization of assets according to the chosen mode (devMode).
	// you can set this flag to your liking without reference to devMode.
	assetsDeclarationsUseEmbedded(!devMode)

	// Create temp directory .. or not
	doTempDir = false

	/* Init & read options file */
	mainOptions = new(MainOpt) // Assignate options' structure.
	mainOptions.Read()         // Read values from options' file if exists.
	// mainOptions.Init()         // Used to reset options if its necessary (debug).

	/* Logger init. */
	Logger = Log2FileStructNew(optFilename, devMode)
	defer Logger.CloseLogger()

	/* Init gtk display */
	mainStartGtk(fmt.Sprintf("%s %s  %s %s %s",
		Name,
		Vers,
		"©"+YearCreat,
		Creat,
		LicenseAbrv),
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)
}

/*******************************************\
/* Executed before signals initialisation. */
/******************************************/
func mainApplication() {

	var err error

	screen := mainObjects.MainWindow.GetScreen()
	visual, _ := screen.GetRGBAVisual()
	mainObjects.MainWindow.SetVisual(visual)

	if wds, err = WinDecorationStructureNew(
		mainObjects.MainWindow,
		mainObjects.EventBoxResize,
		mainObjects.EventBoxMinimize,
		nil); err == nil {
		wds.SignalHandleBlockUnblock(mainObjects.ScrolledWindowTextViewDisplay, nil, nil)
		err = wds.Init()
	}
	Logger.Log(err, "mainApplication/WinDecorationStructureNew")

	txt := fmt.Sprintf(`
	` + FormatText(strings.Join([]string{Name, Vers, "©" + YearCreat, Creat, Repository}, " "), 80, true) + `
	` + FormatText(Descr, 80, true) + `
	` + FormatText(LicenseShort, 80, true))

	if mainOptions.ShowSplash {
		image, err := gtk.ImageNew()
		Logger.Log(err, "mainApplication/ImageNew")
		err = SetPict(image, hash, 64)
		Logger.Log(err, "mainApplication/SetPict")
		err = wds.SplashWin(image, txt)
		Logger.Log(err, "mainApplication/SplashWin")
	}

	/* Translate init. */
	translate = MainTranslateNew(filepath.Join(absoluteRealPath, mainOptions.LanguageFilename), devMode)
	sts = translate.Sentences

	/* Load css */
	if bytes, err := GetBytesFromVarAsset(custom); err == nil {
		CssWdgScnBytes(bytes)
	}
	Logger.Log(err, "mainApplication:GetBytesFromVarAsset")

	/* Init treeview */
	if err = initTvs(); err != nil {
		Logger.Log(err, "mainApplication/TreeViewStructureNew")
		os.Exit(1)
	}

	// Init statusbar
	mainStatusbar = StatusBarStructureNew(
		mainObjects.MainStatusbar,
		[]string{"File(s)", "Done"})

	// Init Clipboard
	clipboard, err = ClipboardNew()
	Logger.Log(err, "mainApplication/makeHash")

	/* Init D&D */
	initDnds()

	mainOptions.UpdateObjects()

	// Get buffer from textView
	buff, err = mainObjects.TextViewDisplay.GetBuffer()
	Logger.Log(err, "mainApplication/GetBuffer")

	/* Handling commandline arguments */
	if len(os.Args) > 1 {
		files = os.Args[1:]
		filesChanged = true
		toDisplay(true)
	}
}

/******************************************\
/* Executed after signals initialisation. */
/*****************************************/
func afterSignals() {}

/*************************************\
/* Executed just before closing all. */
/************************************/
func onShutdown() bool {

	var err error

	// Update 'mainOptions' with GtkObjects and save it
	if err = mainOptions.Write(); err == nil {
		// What you want to execute before closing the app.
		// Return:
		//		true for exit applicaton
		//		false does not exit application
	}
	if err != nil {
		log.Fatalf("Unexpected error on exit: %s", err.Error())
	}
	return true
}
