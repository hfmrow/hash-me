// gohSignals.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 17:02:29 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020-21 hfmrow - Hash Me v1.1 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/********************************************************/
/* This section preserve user modifications on update. */
/* Signals & Property implementations:                */
/* initialize signals used by gtk objects ...        */
/****************************************************/
func signalsPropHandler() {
	mainObjects.CheckbuttonAddReminder.Connect("notify", blankNotify)
	mainObjects.CheckbuttonAppendFiles.Connect("notify", blankNotify)
	mainObjects.CheckbuttonBlake2b256.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonBlake2b384.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonBlake2b512.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonCreateFile.Connect("notify", blankNotify)
	mainObjects.CheckbuttonMd4.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonMd5.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha1.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha256.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha384.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha512.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha3_256.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha3_384.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonSha3_512.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonShowFilename.Connect("toggled", toDisplay)
	mainObjects.CheckbuttonShowSplash.Connect("toggled", CheckbuttonShowSplashToggled)
	mainObjects.CheckbuttonUseDecimal.Connect("toggled", toDisplay)
	mainObjects.EventBoxAppIcon.Connect("notify", blankNotify)
	mainObjects.EventBoxMinimize.Connect("notify", blankNotify)
	mainObjects.EventBoxResize.Connect("notify", blankNotify)
	mainObjects.ImageAppIcon.Connect("notify", blankNotify)
	mainObjects.LabelMethods.Connect("notify", blankNotify)
	mainObjects.LabelOptions.Connect("notify", blankNotify)
	mainObjects.LabelOutput.Connect("notify", blankNotify)
	mainObjects.LabelTitle.Connect("notify", blankNotify)
	mainObjects.MainButtonCancel.Connect("clicked", windowDestroy)
	mainObjects.MainButtonDone.Connect("clicked", MainButtonDoneClicked)
	mainObjects.MainStatusbar.Connect("notify", blankNotify)
	mainObjects.MainWindow.Connect("notify", blankNotify)
	mainObjects.ScrolledWindowTextViewDisplay.Connect("notify", blankNotify)
	mainObjects.Stack.Connect("notify", blankNotify)
	mainObjects.SwitchExpand.Connect("state-set", SwitchExpandStateSet)
	mainObjects.SwitchTreeView.Connect("state-set", SwitchTreeViewStateSet)
	mainObjects.TextViewDisplay.Connect("notify", blankNotify)
	mainObjects.TreeViewDisplay.Connect("notify", blankNotify)
}
