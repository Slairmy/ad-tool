package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func archeryScreen(_ fyne.Window) fyne.CanvasObject {

	// 初始化是一个grid布局
	sqlContentBox := makeBox()
	instanceSelectBox := makeBox()

	return container.NewGridWithColumns(2, sqlContentBox, instanceSelectBox)

	// 分为上下2块,上面是输入执行的sql,下面显示结果
	// 只有点了查询才分成2块,正常是一块
}

func makeBox() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	rect.SetMinSize(fyne.NewSize(30, 30))
	return rect
}
