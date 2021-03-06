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
	"time"

	"github.com/gotk3/gotk3/gtk"
)

/*
 * Checkbuttons / Switch
 */
func genericMethod(chk *gtk.CheckButton) {
	filesChanged = true
	toDisplay(true)
	time.Sleep(time.Millisecond * 300)
}

func CheckbuttonShowSplashToggled(chk *gtk.CheckButton) {
	mainOptions.ShowSplash = chk.GetActive()
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

	mainOptions.SwitchExpandState = sw.GetActive()
	if sw.GetActive() {
		tvs.TreeView.ExpandAll()
	} else {
		tvs.TreeView.CollapseAll()
	}
}

/*
 * Buttons
 */
func MainButtonDoneClicked() {

	toDisplay(false)

	// anim, err := GetPixBufAnimation(linearProgressHorzBlue)
	// if err != nil {
	// 	log.Fatalf("GetPixBufAnimation: %s\n", err.Error())
	// }
	// gifImage, err := gtk.ImageNewFromAnimation(anim)
	// if err != nil {
	// 	log.Fatalf("ImageNewFromAnimation: %s\n", err.Error())
	// }
	// pbs = ProgressGifNew(gifImage, mainObjects.BoxMain, 1,
	// 	func() error {
	// 		glib.IdleAdd(func() {
	// 			mainObjects.MainButtonDone.SetSensitive(false)
	// 			mainObjects.GridOptions.SetSensitive(false)
	// 		})
	// 		doIt(false)
	// 		return nil
	// 	},
	// 	func() error {

	// 		// toTreview(results)

	// 		mainObjects.MainButtonDone.SetSensitive(true)
	// 		mainObjects.GridOptions.SetSensitive(true)
	// 		return nil
	// 	})

	// go func() {
	// 	pbs.StartGif()
	// }()

}
