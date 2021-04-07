// gohObjects.go

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
	"github.com/gotk3/gotk3/gtk"
)

/* Control over all used objects from glade. */
var mainObjects *MainControlsObj

/********************************************************/
/* This section preserve user modifications on update. */
/* Main structure Declaration: You may add your own   */
/* declarations (gotk3 objects only) here.           */
/****************************************************/
type MainControlsObj struct {
	BoxMain                       *gtk.Box
	CheckbuttonAddReminder        *gtk.CheckButton
	CheckbuttonAppendFiles        *gtk.CheckButton
	CheckbuttonBlake2b256         *gtk.CheckButton
	CheckbuttonBlake2b384         *gtk.CheckButton
	CheckbuttonBlake2b512         *gtk.CheckButton
	CheckbuttonConcurrentOp       *gtk.CheckButton
	CheckbuttonCreateFile         *gtk.CheckButton
	CheckbuttonMd4                *gtk.CheckButton
	CheckbuttonMd5                *gtk.CheckButton
	CheckbuttonRecursiveScan      *gtk.CheckButton
	CheckbuttonSha1               *gtk.CheckButton
	CheckbuttonSha256             *gtk.CheckButton
	CheckbuttonSha384             *gtk.CheckButton
	CheckbuttonSha512             *gtk.CheckButton
	CheckbuttonSha3_256           *gtk.CheckButton
	CheckbuttonSha3_384           *gtk.CheckButton
	CheckbuttonSha3_512           *gtk.CheckButton
	CheckbuttonShowFilename       *gtk.CheckButton
	CheckbuttonShowSplash         *gtk.CheckButton
	CheckbuttonUseDecimal         *gtk.CheckButton
	EventBoxAppIcon               *gtk.EventBox
	EventBoxMinimize              *gtk.EventBox
	EventBoxResize                *gtk.EventBox
	GridOptions                   *gtk.Grid
	ImageAppIcon                  *gtk.Image
	LabelMethods                  *gtk.Label
	LabelOptions                  *gtk.Label
	LabelOutput                   *gtk.Label
	LabelTitle                    *gtk.Label
	MainButtonCancel              *gtk.Button
	MainButtonDone                *gtk.Button
	MainStatusbar                 *gtk.Statusbar
	mainUiBuilder                 *gtk.Builder
	MainWindow                    *gtk.Window
	ScrolledWindowTextViewDisplay *gtk.ScrolledWindow
	Stack                         *gtk.Stack
	SwitchExpand                  *gtk.Switch
	SwitchTreeView                *gtk.Switch
	TextViewDisplay               *gtk.TextView
	TreeViewDisplay               *gtk.TreeView
}

/******************************************************************/
/* This section preserve user modification on update.            */
/* GtkObjects initialisation: You may add your own declarations */
/* as you  wish, best way is to add them by grouping  same     */
/* objects names (below first declaration).                   */
/*************************************************************/
func gladeObjParser() {
	mainObjects.BoxMain = loadObject("BoxMain").(*gtk.Box)
	mainObjects.CheckbuttonAddReminder = loadObject("CheckbuttonAddReminder").(*gtk.CheckButton)
	mainObjects.CheckbuttonAppendFiles = loadObject("CheckbuttonAppendFiles").(*gtk.CheckButton)
	mainObjects.CheckbuttonBlake2b256 = loadObject("CheckbuttonBlake2b256").(*gtk.CheckButton)
	mainObjects.CheckbuttonBlake2b384 = loadObject("CheckbuttonBlake2b384").(*gtk.CheckButton)
	mainObjects.CheckbuttonBlake2b512 = loadObject("CheckbuttonBlake2b512").(*gtk.CheckButton)
	mainObjects.CheckbuttonConcurrentOp = loadObject("CheckbuttonConcurrentOp").(*gtk.CheckButton)
	mainObjects.CheckbuttonCreateFile = loadObject("CheckbuttonCreateFile").(*gtk.CheckButton)
	mainObjects.CheckbuttonMd4 = loadObject("CheckbuttonMd4").(*gtk.CheckButton)
	mainObjects.CheckbuttonMd5 = loadObject("CheckbuttonMd5").(*gtk.CheckButton)
	mainObjects.CheckbuttonRecursiveScan = loadObject("CheckbuttonRecursiveScan").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha1 = loadObject("CheckbuttonSha1").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha256 = loadObject("CheckbuttonSha256").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha384 = loadObject("CheckbuttonSha384").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha512 = loadObject("CheckbuttonSha512").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha3_256 = loadObject("CheckbuttonSha3_256").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha3_384 = loadObject("CheckbuttonSha3_384").(*gtk.CheckButton)
	mainObjects.CheckbuttonSha3_512 = loadObject("CheckbuttonSha3_512").(*gtk.CheckButton)
	mainObjects.CheckbuttonShowFilename = loadObject("CheckbuttonShowFilename").(*gtk.CheckButton)
	mainObjects.CheckbuttonShowSplash = loadObject("CheckbuttonShowSplash").(*gtk.CheckButton)
	mainObjects.CheckbuttonUseDecimal = loadObject("CheckbuttonUseDecimal").(*gtk.CheckButton)
	mainObjects.EventBoxAppIcon = loadObject("EventBoxAppIcon").(*gtk.EventBox)
	mainObjects.EventBoxMinimize = loadObject("EventBoxMinimize").(*gtk.EventBox)
	mainObjects.EventBoxResize = loadObject("EventBoxResize").(*gtk.EventBox)
	mainObjects.GridOptions = loadObject("GridOptions").(*gtk.Grid)
	mainObjects.ImageAppIcon = loadObject("ImageAppIcon").(*gtk.Image)
	mainObjects.LabelMethods = loadObject("LabelMethods").(*gtk.Label)
	mainObjects.LabelOptions = loadObject("LabelOptions").(*gtk.Label)
	mainObjects.LabelOutput = loadObject("LabelOutput").(*gtk.Label)
	mainObjects.LabelTitle = loadObject("LabelTitle").(*gtk.Label)
	mainObjects.MainButtonCancel = loadObject("MainButtonCancel").(*gtk.Button)
	mainObjects.MainButtonDone = loadObject("MainButtonDone").(*gtk.Button)
	mainObjects.MainStatusbar = loadObject("MainStatusbar").(*gtk.Statusbar)
	mainObjects.MainWindow = loadObject("MainWindow").(*gtk.Window)
	mainObjects.ScrolledWindowTextViewDisplay = loadObject("ScrolledWindowTextViewDisplay").(*gtk.ScrolledWindow)
	mainObjects.Stack = loadObject("Stack").(*gtk.Stack)
	mainObjects.SwitchExpand = loadObject("SwitchExpand").(*gtk.Switch)
	mainObjects.SwitchTreeView = loadObject("SwitchTreeView").(*gtk.Switch)
	mainObjects.TextViewDisplay = loadObject("TextViewDisplay").(*gtk.TextView)
	mainObjects.TreeViewDisplay = loadObject("TreeViewDisplay").(*gtk.TreeView)
}
