package utils

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/davideginna/app-fyne/pkg/constants"
)

var (
	a         fyne.App
	themeType = constants.DarkTheme
	dt        *dao.CurrentDate
	u         *dao.UI
)

func InitUI() {
	//u = &dao.UI{}
	//setDefaultDate()
	//setDefaultTitle()
	//setDefaultThemeButton()
	//setDefaultBarLayout()
	//setDefaultContent()
	//initAndLoadDiaryList()
	//setCurrent(u.AppDate.FormattedDate)
	//setDropDown()
	//saveOnContentChange()
}

func LoadUI() fyne.CanvasObject {
	InitUI()
	leftSide := fyne.NewContainerWithLayout(layout.NewBorderLayout(u.DropDown, u.ThemeButton, nil, nil), u.DropDown, u.ThemeButton, u.BtnList)
	rightSide := fyne.NewContainerWithLayout(layout.NewBorderLayout(u.BarLayout, nil, nil, nil), u.BarLayout, u.Content)
	splitLayout := widget.NewHSplitContainer(leftSide, rightSide)
	splitLayout.Offset = constants.LayoutOffset
	return splitLayout
}

func InitAndRun() {

	a = app.NewWithID(constants.AppId)
	w := a.NewWindow(constants.Title)
	w.SetContent(LoadUI())
	w.Resize(fyne.NewSize(constants.AppWidth, constants.AppHeight))
	w.ShowAndRun()
}
