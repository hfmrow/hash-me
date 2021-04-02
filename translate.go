// translate.go

// File generated on Mon, 21 Dec 2020 18:50:49 using Gotk3 Objects Translate v1.5 2019-20 H.F.M

/*
* 	This program comes with absolutely no warranty.
*	See the The MIT License (MIT) for details:
*	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

// initGtkObjectsText: read translations from structure and set them to objects.
func (trans *MainTranslate) initGtkObjectsText() {
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonAddReminder.Widget, "CheckbuttonAddReminder")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonAppendFiles.Widget, "CheckbuttonAppendFiles")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonBlake2b256.Widget, "CheckbuttonBlake2b256")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonBlake2b384.Widget, "CheckbuttonBlake2b384")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonBlake2b512.Widget, "CheckbuttonBlake2b512")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonCreateFile.Widget, "CheckbuttonCreateFile")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonMd4.Widget, "CheckbuttonMd4")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonMd5.Widget, "CheckbuttonMd5")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha1.Widget, "CheckbuttonSha1")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha256.Widget, "CheckbuttonSha256")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha384.Widget, "CheckbuttonSha384")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha512.Widget, "CheckbuttonSha512")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha3_256.Widget, "CheckbuttonSha3_256")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha3_384.Widget, "CheckbuttonSha3_384")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSha3_512.Widget, "CheckbuttonSha3_512")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonShowFilename.Widget, "CheckbuttonShowFilename")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonUseDecimal.Widget, "CheckbuttonUseDecimal")
	trans.setTextToGtkObjects(&mainObjects.LabelMethods.Widget, "LabelMethods")
	trans.setTextToGtkObjects(&mainObjects.LabelOptions.Widget, "LabelOptions")
	trans.setTextToGtkObjects(&mainObjects.LabelOutput.Widget, "LabelOutput")
	trans.setTextToGtkObjects(&mainObjects.LabelTitle.Widget, "LabelTitle")
	trans.setTextToGtkObjects(&mainObjects.MainButtonCancel.Widget, "MainButtonCancel")
	trans.setTextToGtkObjects(&mainObjects.MainButtonDone.Widget, "MainButtonDone")
	trans.setTextToGtkObjects(&mainObjects.MainStatusbar.Widget, "MainStatusbar")
	trans.setTextToGtkObjects(&mainObjects.SwitchExpand.Widget, "SwitchExpand")
	trans.setTextToGtkObjects(&mainObjects.SwitchTreeView.Widget, "SwitchTreeView")
	trans.setTextToGtkObjects(&mainObjects.TextViewDisplay.Widget, "TextViewDisplay")
	trans.setTextToGtkObjects(&mainObjects.TreeViewDisplay.Widget, "TreeViewDisplay")
}
// Translations structure declaration. To be used in main application.
var translate = new(MainTranslate)

// sts: some sentences/words used in the application. Mostly used in Development mode.
// You must add there all sentences used in your application. Or not ...
// They'll be added to language file each time application started
// when "devMode" is set at true.
var sts = map[string]string{
	`cancel`: `Cancel`,
	`dat`: `Data`,
	`mtd`: `Hash method`,
	`no`: `No`,
	`yes`: `Yes`,
}


// Translations structure with methods
type MainTranslate struct {
	// Public
	ProgInfos    progInfo
	Language     language
	Options      parsingFlags
	ObjectsCount int
	Objects      []object
	Sentences    map[string]string
	// Private
	objectsLoaded bool
}

// MainTranslateNew: Initialise new translation structure and assign language file content to GtkObjects.
// devModeActive, indicate that the new sentences must be added to previous language file.
func MainTranslateNew(filename string, devModeActive ...bool) (mt *MainTranslate) {
	var err error
	mt = new(MainTranslate)
	if err = mt.read(filename); err == nil {
		mt.initGtkObjectsText()
		if len(devModeActive) != 0 {
			if devModeActive[0] {
				mt.Sentences = sts
				err := mt.write(filename)
				if err != nil {
					fmt.Printf("%s\n%s\n", "Cannot write actual sentences to language file.", err.Error())
				}
			}
		}
	} else {
		fmt.Printf("%s\n%s\n", "Error loading language file !\nNot an error when you just creating from glade Xml or GOH project file.", err.Error())
	}
	return
}

// readFile: language file.
func (trans *MainTranslate) read(filename string) (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(filename); err == nil {
		if err = json.Unmarshal(textFileBytes, &trans); err == nil {
			trans.objectsLoaded = true
		}
	}
	return
}

// Write json datas to file
func (trans *MainTranslate) write(filename string) (err error) {
	var out bytes.Buffer
	var jsonData []byte
	if jsonData, err = json.Marshal(&trans); err == nil && trans.objectsLoaded {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(filename, out.Bytes(), 0644)
		}
	}
	return
}

type parsingFlags struct {
	SkipLowerCase,
	SkipEmptyLabel,
	SkipEmptyName,
	DoBackup bool
}

type progInfo struct {
	Name,
	Version,
	Creat,
	MainObjStructName,
	GladeXmlFilename,
	TranslateFilename,
	ProjectRootDir,
	GohProjFile string
}

type language struct {
	LangNameLong,
	LangNameShrt,
	Author,
	Date,
	Updated string
	Contributors []string
}

type object struct {
	Class,
	Id,
	Label,
	Tooltip,
	Text,
	Uri,
	Comment string
	LabelMarkup,
	LabelWrap,
	TooltipMarkup bool
	Idx int
}

// Define available property within objects
type propObject struct {
	Class string
	Label,
	LabelMarkup,
	LabelWrap,
	Tooltip,
	TooltipMarkup,
	Text,
	Uri bool
}

// Property that exists for Gtk3 Object ...	(Used for Class capability)
var propPerObjects = []propObject{
	{Class: "GtkButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkMenuButton", Label: true, Tooltip: true, TooltipMarkup: true},

	// {Class: "GtkToolButton", Label: true, Tooltip: true, TooltipMarkup: true},    // Deprecated since 3.10
	// {Class: "GtkImageMenuItem", Label: true, Tooltip: true, TooltipMarkup: true}, // Deprecated since 3.10

	{Class: "GtkMenuItem", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkCheckMenuItem", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkRadioMenuItem", Label: true, Tooltip: true, TooltipMarkup: true},

	{Class: "GtkToggleButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkLabel", Label: true, LabelMarkup: true, Tooltip: true, TooltipMarkup: true, LabelWrap: true},
	{Class: "GtkSpinButton", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkEntry", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkCheckButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkProgressBar", Tooltip: true, TooltipMarkup: true, Text: true},
	{Class: "GtkSearchBar", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkImage", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkRadioButton", Label: true, LabelMarkup: false, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkComboBoxText", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkComboBox", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkLinkButton", Label: true, Tooltip: true, TooltipMarkup: true, Uri: true},
	{Class: "GtkSwitch", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkTreeView", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkFileChooserButton", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkTextView", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkSourceView", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkStatusbar", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkScrolledWindow", Tooltip: true, TooltipMarkup: true},
}

// setTextToGtkObjects: read translations from structure and set them to object.
// like this: setTextToGtkObjects(&mainObjects.TransLabelHint.Widget, "TransLabelHint")
func (trans *MainTranslate) setTextToGtkObjects(obj *gtk.Widget, objectId string) {
	for _, currObject := range trans.Objects {
		if currObject.Id == objectId {
			for _, props := range propPerObjects {
				if currObject.Class == props.Class {
					if props.Label {
						obj.SetProperty("label", currObject.Label)
						if props.LabelMarkup {
							obj.SetProperty("use-markup", currObject.LabelMarkup)
							obj.SetProperty("label", strings.ReplaceAll(currObject.Label, "&", "&amp;"))
						}
					}
					if props.LabelWrap {
						obj.SetProperty("wrap", currObject.LabelWrap)
					}
					if props.Tooltip && !currObject.TooltipMarkup {
						obj.SetProperty("tooltip_text", currObject.Tooltip)
					}
					if props.Tooltip && currObject.TooltipMarkup {
						obj.SetProperty("tooltip_markup", strings.ReplaceAll(currObject.Tooltip, "&", "&amp;"))
					}
					if props.Text {
						obj.SetProperty("text", currObject.Text)
					}
					if props.Uri {
						obj.SetProperty("uri", currObject.Uri)
					}
					break
				}
			}
			break
		}
	}
}
