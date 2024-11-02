package main

import (
	"AdTool/tutorials"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

func getIcon(i int) fyne.Resource {
	switch i % 3 {
	case 1:
		return theme.HomeIcon()
	case 2:
		return theme.MailSendIcon()
	default:
		return theme.MediaVideoIcon()
	}
}

func main() {

	a := app.NewWithID("ad-tool")
	w := a.NewWindow("Ad Tool")

	// 设置为主程序,关联了就关闭应用
	w.SetMaster()

	// 堆栈容器
	content := container.NewStack()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An intro")
	// 文本换行模式
	intro.Wrapping = fyne.TextWrapWord

	// 设置结构
	setTutorial := func(t tutorials.Tutorial) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
	split.Offset = 0.2

	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

// 构建侧边的功能栏
func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	// 返回当前正在运行的应用程序的实例
	currentApp := fyne.CurrentApp()

	tree := &widget.Tree{
		// 不需要子级
		ChildUIDs: func(uid string) []string {
			return tutorials.TutorialIndex[uid]
		},
		IsBranch: func(uid widget.TreeNodeID) bool {
			children, ok := tutorials.TutorialIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
			t, ok := tutorials.Tutorials[uid]
			if !ok {
				fyne.LogError("Miss tutorial panel: "+uid, nil)
				return
			}
			node.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid widget.TreeNodeID) {
			if t, ok := tutorials.Tutorials[uid]; ok {
				currentApp.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	// Preferences 方法用于访问当前应用程序的偏好设置,比如窗口大小颜色等等
	if loadPrevious {
		currentPref := currentApp.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			currentApp.Settings().SetTheme(&forcedVariant{
				Theme:   theme.DefaultTheme(),
				variant: theme.VariantDark,
			})
		}),
		widget.NewButton("Light", func() {
			currentApp.Settings().SetTheme(&forcedVariant{
				Theme:   theme.DefaultTheme(),
				variant: theme.VariantLight,
			})
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
