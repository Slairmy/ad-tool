package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func archeryScreen(_ fyne.Window) fyne.CanvasObject {

	// 分为上下2块,上面是输入执行的sql,下面显示结果
	return container.NewVSplit(widget.NewLabel("SQL输入区域"), widget.NewLabel("SQL结果显示区域"))
}
