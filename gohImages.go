// gohImages.go

/*
	Source file auto-generated on Tue, 06 Apr 2021 22:04:43 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020-21 hfmrow - Hash Me v1.2 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/**********************************************************/
/* This section preserve user modifications on update.   */
/* Images declarations, used to initialize objects with */
/* The SetPict() func, accept both kind of variables:  */
/* filename or []byte content in case of using        */
/* embedded binary data. The variables names are the */
/* same. "assetsDeclarationsUseEmbedded(bool)" func */
/* could be used to toggle between filenames and   */
/* embedded binary type. See SetPict()            */
/* declaration to learn more on how to use it.   */
/************************************************/
func assignImages() {
	SetPict(mainObjects.ImageAppIcon, hash, 16)
	// SetPict(mainObjects.MainButtonCancel, logout48, 18)
	// SetPict(mainObjects.MainButtonDone, tickIcon48, 18)
	SetPict(mainObjects.MainWindow, hash)
}

/**********************************************************/
/* This section is rewritten on assets update.           */
/* Assets var declarations, this step permit to make a  */
/* bridge between the differents types used, string or */
/* []byte, and to simply switch from one to another.  */
/*****************************************************/
var mainGlade interface{}              // assets/glade/main.glade
var custom interface{}                 // assets/css/custom.css
var hash interface{}                   // assets/images/hash.png
var linearProgressHorzBlue interface{} // assets/images/linear-progress-horz-blue.gif
