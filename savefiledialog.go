package savefiledialog

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type SaveFileDialog struct {
	OwnerForm             *walk.Form
	TitleText             string
	FilterText            string
	FilterIndex           int
	FilePath              string
	InitialDirPath        string
	IsHideOverwritePrompt bool
}

func New() *SaveFileDialog {
	return &SaveFileDialog{
		OwnerForm:             nil,
		TitleText:             "名前を付けて保存",
		FilterText:            "すべてのファイル(*.*)|*.*",
		FilterIndex:           1,
		FilePath:              "",
		InitialDirPath:        "",
		IsHideOverwritePrompt: false,
	}
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Show() (accepted bool, filePath string) {
	wdlg := new(walk.FileDialog)
	wdlg.Title = dlg.TitleText
	wdlg.Filter = dlg.FilterText
	wdlg.FilterIndex = dlg.FilterIndex
	wdlg.FilePath = dlg.FilePath
	wdlg.InitialDirPath = dlg.InitialDirPath
	wdlg.Flags = 0

	if dlg.IsHideOverwritePrompt {
		wdlg.Flags &= ^uint32(win.OFN_OVERWRITEPROMPT)
	} else {
		wdlg.Flags |= uint32(win.OFN_OVERWRITEPROMPT)
	}

	var owner walk.Form = nil
	if dlg.OwnerForm != nil {
		owner = *dlg.OwnerForm
	}

	ok, err := wdlg.ShowSave(owner)
	if err != nil {
		panic(err)
	}

	return ok, wdlg.FilePath
}

func Show() (accepted bool, filePath string) {
	dlg := New()
	return dlg.Show()
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Owner(owner *walk.Form) *SaveFileDialog {
	dlg.OwnerForm = owner
	return dlg
}

func Owner(owner *walk.Form) *SaveFileDialog {
	dlg := New()
	return dlg.Owner(owner)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Title(title string) *SaveFileDialog {
	dlg.TitleText = title
	return dlg
}

func Title(title string) *SaveFileDialog {
	dlg := New()
	return dlg.Title(title)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) Filter(filter string, index ...int) *SaveFileDialog {
	dlg.FilterText = filter
	if len(index) > 0 {
		dlg.FilterIndex = index[0]
	}
	return dlg
}

func Filter(filter string, index ...int) *SaveFileDialog {
	dlg := New()
	return dlg.Filter(filter, index...)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) InitFilePath(path string) *SaveFileDialog {
	dlg.FilePath = path
	return dlg
}

func InitFilePath(path string) *SaveFileDialog {
	dlg := New()
	return dlg.InitFilePath(path)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) InitDirPath(path string) *SaveFileDialog {
	dlg.InitialDirPath = path
	return dlg
}

func InitDirPath(path string) *SaveFileDialog {
	dlg := New()
	return dlg.InitDirPath(path)
}

// ----------------------------------------------------------------

func (dlg *SaveFileDialog) HideOverwritePrompt() *SaveFileDialog {
	dlg.IsHideOverwritePrompt = true
	return dlg
}

func HideOverwritePrompt() *SaveFileDialog {
	dlg := New()
	return dlg.HideOverwritePrompt()
}
