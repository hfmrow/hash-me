// objHandler.go

/*
	Source file auto-generated on Sat, 19 Dec 2020 15:06:30 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func CheckbuttonShowSplashToggled(chk *gtk.CheckButton) {
	mainOptions.ShowSplash = chk.GetActive()
}

func toDisplay() {
	mainObjects.CheckbuttonAddReminder.SetActive(false)
	mainObjects.CheckbuttonCreateFile.SetActive(false)

	doIt(true)

	mainObjects.CheckbuttonAddReminder.SetActive(mainOptions.Reminder)
	mainObjects.CheckbuttonCreateFile.SetActive(mainOptions.SaveToFile)
}

func MainButtonDoneClicked() {

	doIt(false)
}

func SwitchTreeViewStateSet(sw *gtk.Switch) {

	if sw.GetActive() {
		mainObjects.Stack.SetVisibleChildName("StackTreeView")
		mainObjects.SwitchExpand.SetVisible(true)
	} else {
		mainObjects.Stack.SetVisibleChildName("StackTextView")
		mainObjects.SwitchExpand.SetVisible(false)
	}
}

func SwitchExpandStateSet(sw *gtk.Switch) {

	if sw.GetActive() {
		tvs.TreeView.ExpandAll()
	} else {
		tvs.TreeView.CollapseAll()
	}
}
