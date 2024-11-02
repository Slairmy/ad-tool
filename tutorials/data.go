package tutorials

import "fyne.io/fyne/v2"

// Tutorial 定义功能栏集合
type Tutorial struct {
	Title, Intro string

	View func(w fyne.Window) fyne.CanvasObject
}

var Tutorials = map[string]Tutorial{
	"welcome": {
		Title: "Welcome",
		Intro: "",
		View:  welcomeScreen,
	},
	"archery": {
		Title: "Archery-SQL查询",
		Intro: "",
		View:  archeryScreen,
	},
}

// TutorialIndex 暂时就2个功能
var TutorialIndex = map[string][]string{
	"": {"welcome", "archery"},
}
