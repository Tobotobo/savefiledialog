// savefiledialog.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package savefiledialog

import (
	"github.com/Tobotobo/commondialogs"
	"github.com/lxn/win"
)

type saveFileDialog struct {
	Owner                 win.HWND
	Title                 string
	Filter                string
	FilterIndex           int
	FilePath              string
	InitialDirPath        string
	IsHideOverwritePrompt bool
}

type SaveFileDialog struct {
	InnerValue saveFileDialog
}

func New() *SaveFileDialog {
	return &SaveFileDialog{
		InnerValue: saveFileDialog{
			Owner:                 0,
			Title:                 "名前を付けて保存",
			Filter:                "すべてのファイル(*.*)|*.*",
			FilterIndex:           1,
			FilePath:              "",
			InitialDirPath:        "",
			IsHideOverwritePrompt: false,
		},
	}
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Show() (filePath string, accepted bool) {
	wdlg := new(commondialogs.FileDialog)
	wdlg.Title = dlg.InnerValue.Title
	wdlg.Filter = dlg.InnerValue.Filter
	wdlg.FilterIndex = dlg.InnerValue.FilterIndex
	wdlg.FilePath = dlg.InnerValue.FilePath
	wdlg.InitialDirPath = dlg.InnerValue.InitialDirPath
	wdlg.Flags = 0

	if dlg.InnerValue.IsHideOverwritePrompt {
		wdlg.Flags &= ^uint32(win.OFN_OVERWRITEPROMPT)
	} else {
		wdlg.Flags |= uint32(win.OFN_OVERWRITEPROMPT)
	}

	ok, err := wdlg.ShowSave(dlg.InnerValue.Owner)
	if err != nil {
		panic(err)
	}
	dlg.InnerValue.FilePath = wdlg.FilePath

	return wdlg.FilePath, ok
}

func Show() (filePath string, accepted bool) {
	dlg := New()
	return dlg.Show()
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Owner(owner win.HWND) *SaveFileDialog {
	dlg.InnerValue.Owner = owner
	return dlg
}

func Owner(owner win.HWND) *SaveFileDialog {
	dlg := New()
	return dlg.Owner(owner)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Title(title string) *SaveFileDialog {
	dlg.InnerValue.Title = title
	return dlg
}

func Title(title string) *SaveFileDialog {
	dlg := New()
	return dlg.Title(title)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Filter(filter string, index ...int) *SaveFileDialog {
	dlg.InnerValue.Filter = filter
	if len(index) > 0 {
		dlg.InnerValue.FilterIndex = index[0]
	}
	return dlg
}

func Filter(filter string, index ...int) *SaveFileDialog {
	dlg := New()
	return dlg.Filter(filter, index...)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) InitFilePath(path string) *SaveFileDialog {
	dlg.InnerValue.FilePath = path
	return dlg
}

func InitFilePath(path string) *SaveFileDialog {
	dlg := New()
	return dlg.InitFilePath(path)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) InitDirPath(path string) *SaveFileDialog {
	dlg.InnerValue.InitialDirPath = path
	return dlg
}

func InitDirPath(path string) *SaveFileDialog {
	dlg := New()
	return dlg.InitDirPath(path)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) HideOverwritePrompt() *SaveFileDialog {
	dlg.InnerValue.IsHideOverwritePrompt = true
	return dlg
}

func HideOverwritePrompt() *SaveFileDialog {
	dlg := New()
	return dlg.HideOverwritePrompt()
}
