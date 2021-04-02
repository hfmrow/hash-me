// initTreeView.go

/*
	Source file auto-generated on Mon, 21 Dec 2020 04:17:28 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
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

// initTreeViewPackageContent: Init ListStore
func initTvs() (err error) {

	columns = [][]string{
		{"chk", "text"}, // Will be hidden
		{sts[`mtd`], "text"},
		{sts[`dat`], "text"},
		{"filename", "text"}, // Will be hidden
	}
	colmap = map[string]int{
		`chk`: 0,
		`mtd`: 1,
		`dat`: 2,
		`fnm`: 3,
	}

	if tvs, err = TreeViewStructureNew(mainObjects.TreeViewDisplay, false, false); err != nil {
		return
	}

	tvs.AddColumns(columns, false, true, false, true, false, true)

	tvs.Columns[colmap["dat"]].Editable = true
	tvs.Columns[colmap["chk"]].Visible = false
	tvs.Columns[colmap["fnm"]].Visible = false

	if err = tvs.StoreSetup(new(gtk.TreeStore)); err == nil {

		tvs.TreeView.SetExpanderColumn(tvs.Columns[colmap["dat"]].Column)
	}

	Logger.Log(err, "initTvs")
	return
}
