package tutorials

import (
	"AdTool/widgets"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"strconv"
)

func archeryScreen(_ fyne.Window) fyne.CanvasObject {

	// 初始化是一个grid布局
	sqlContentBox := makeSQLContentBox(886, 500)
	instanceSelectBox := makeSelectBox(200, 300)
	//
	//return container.NewHBox(sqlContentBox, instanceSelectBox)

	// return container.NewGridWithColumns(2, sqlContentBox, instanceSelectBox)

	// 分为上下2块,上面是输入执行的sql,下面显示结果
	// 只有点了查询才分成2块,正常是一块

	headers := []string{"rule_name", "rule_uuid", "begin_run_at", "request_date"}
	dataRows := [][]string{
		{"广告活动有曝光，先归1", "115499914106700288", "3600", "2024-11-05"},
		{"广告活动有曝光，先归2", "115499914106700289", "7200", "2024-11-05"},
		{"广告活动有曝光，先归3", "115499914106700210", "10800", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
		{"广告活动有曝光，先归4", "115499914106700211", "14400", "2024-11-05"},
	}

	tableContainer := container.NewStack()

	createTable := func() *widget.Table {
		table := widget.NewTableWithHeaders(
			func() (int, int) {
				return len(dataRows), len(headers)
			},
			func() fyne.CanvasObject {
				return widgets.NewCellWidget("default (hopefully) large enough text", nil) // placeholder to specify width
			},
			func(id widget.TableCellID, object fyne.CanvasObject) {
				cell := object.(*widgets.CellWidget)
				row := id.Row
				col := id.Col

				cell.SetText(dataRows[row][col])
				cell.OnRightClick = func(event *fyne.PointEvent) {
					items := []*fyne.MenuItem{
						fyne.NewMenuItem(lang.X("app.copy_to_clipboard", "Copy"), func() {
							// window.Clipboard().SetContent(cell.Text)
							log.Println("复制内容: " + cell.Text)

						}),
					}
					menu := fyne.NewMenu("", items...)
					cellCanvas := fyne.CurrentApp().Driver().CanvasForObject(cell)
					widget.ShowPopUpMenuAtPosition(menu, cellCanvas, event.AbsolutePosition)
				}
				if id.Row == 0 {
					cell.Importance = widget.HighImportance
				} else {
					cell.Importance = widget.MediumImportance
				}

				cell.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
			},
		)
		table.UpdateHeader = func(id widget.TableCellID, object fyne.CanvasObject) {
			if id.Row == -1 {
				label := object.(*widget.Label)
				label.Alignment = 0
				label.SetText(headers[id.Col])
			}
			if id.Col == -1 {
				label := object.(*widget.Label)
				label.SetText(strconv.Itoa(id.Row + 1))
			}

		}
		// table.HideSeparators = true
		return table

	}

	tableContainer.Objects = []fyne.CanvasObject{createTable()}
	// return tableContainer
	sqlArea := container.NewHBox(sqlContentBox, instanceSelectBox)

	return container.NewVSplit(sqlArea, tableContainer)

}

func makeBox(w, h float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	rect.SetMinSize(fyne.NewSize(w, h))
	return rect
}

func makeSQLContentBox(w, h float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(nil)
	rect.SetMinSize(fyne.NewSize(w, h))

	// 创建多行文本输入框
	sqlContentEntry := widget.NewMultiLineEntry()
	sqlContentEntry.SetPlaceHolder("输入 SQL 查询...")
	sqlContentEntry.Resize(fyne.NewSize(w-20, h-20)) // 设置输入框尺寸并留出边距
	// 偏移一点看下
	sqlContentEntry.Move(fyne.NewPos(20, 0))

	// 使用绝对布局将背景和输入框放入容器中
	return container.NewWithoutLayout(rect, sqlContentEntry)
}

// 我想自由控制间距
func makeSelectBox(w, h float32) fyne.CanvasObject {
	instanceSelect := widget.NewSelect([]string{"实例1", "实例2", "实例3"}, func(s string) {})
	instanceSelect.PlaceHolder = "选择一个实例"
	instanceSelect.Resize(fyne.NewSize(200, 40))
	databaseSelect := widget.NewSelect([]string{"数据库1", "数据库2", "数据库3"}, func(s string) {})
	databaseSelect.PlaceHolder = "选择一个数据库"
	databaseSelect.Resize(fyne.NewSize(200, 40))
	limitSelect := widget.NewSelect([]string{"100", "500", "1000"}, func(s string) {})
	limitSelect.SetSelected("100")
	limitSelect.Resize(fyne.NewSize(200, 40))

	// 再添加一个按钮
	// selectButton := widget.NewButton("SQL查询", func() {})

	button := &widget.Button{
		Text:       "SQL查询",
		Importance: widget.HighImportance,
		OnTapped: func() {
			fmt.Println("执行sql查询")
		},
	}
	button.Resize(fyne.NewSize(100, 40))

	// 在vbox里面放一个自由布局 -- 自由布局里面的所有组件都需要自己设置位置和大小
	freeLayout := container.NewWithoutLayout(instanceSelect, databaseSelect, limitSelect, button)

	// 这里的move是相对于freeLayout这里面的位置
	instanceSelect.Move(fyne.NewPos(0, 0))                                                                              // 第一个下拉框在顶部
	databaseSelect.Move(fyne.NewPos(0, instanceSelect.Size().Height+10))                                                // 第二个下拉框在第一个下方，加10个像素间距
	limitSelect.Move(fyne.NewPos(0, instanceSelect.Size().Height+databaseSelect.Size().Height+20))                      // 第三个下拉框在第二个下方，加10个像素间距
	button.Move(fyne.NewPos(0, instanceSelect.Size().Height+databaseSelect.Size().Height+limitSelect.Size().Height+30)) // 第三个下拉框在第二个下方，加10个像素间距

	freeLayout.Resize(fyne.NewSize(200, 300))
	return freeLayout
}
