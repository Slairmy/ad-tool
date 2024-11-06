package tutorials

import (
	"AdTool/widgets"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"strings"
)

func archeryScreen(_ fyne.Window) fyne.CanvasObject {

	sqlContentEntry := widget.NewMultiLineEntry()
	sqlContentEntry.SetPlaceHolder("输入 SQL 查询...")
	sqlContentEntry.SetMinRowsVisible(10)

	// 创建行号标签
	// lineNumberLabel := widget.NewLabel("1")
	//lineNumberEntry := widget.NewMultiLineEntry()
	//lineNumberEntry.Disable() // 禁用行号输入框，防止用户修改
	//lineNumberEntry.SetMinRowsVisible(10)
	//lineNumberEntry.Scroll = 3
	//// 更新行号函数
	//updateLineNumbers := func() {
	//	// 文本输入框的总行数
	//	lines := strings.Count(sqlContentEntry.Text, "\n") + 1
	//	lineNumbers := ""
	//	for i := 1; i <= lines; i++ {
	//		lineNumbers += fmt.Sprintf("%d\n", i)
	//	}
	//	lineNumberEntry.SetText(strings.TrimRight(lineNumbers, "\n"))
	//}
	//
	//// 监听输入框文本变化
	//sqlContentEntry.OnChanged = func(s string) {
	//	updateLineNumbers()
	//}

	// 使用一个Label而不是MultiLineEntry来显示行号
	lineNumberLabel := widget.NewLabel("1")
	// lineNumberLabel.Wrapping = fyne.TextWrapWord // 设置行号换行
	// lineNumberLabel.Hide()                       // 初始化时隐藏行号，等更新

	// 更新行号函数
	updateLineNumbers := func() {
		lines := strings.Count(sqlContentEntry.Text, "\n") + 1
		lineNumbers := ""
		for i := 1; i <= lines; i++ {
			lineNumbers += fmt.Sprintf("%d\n", i)
		}
		log.Println(strings.TrimRight(lineNumbers, "\n"))
		lineNumberLabel.SetText(strings.TrimRight(lineNumbers, "\n"))
	}

	// 监听 SQL 输入框的文本变化
	sqlContentEntry.OnChanged = func(s string) {
		updateLineNumbers()
	}

	// 将行号容器放到一个垂直容器中
	lineNumberContainer := container.NewVBox(lineNumberLabel)

	// instanceShow
	dbConfigArea := widget.NewLabel("这里到时候直接显示实例,数据库")
	sqlArea := container.NewBorder(nil, dbConfigArea, lineNumberContainer, makeSelectBox(), sqlContentEntry)
	tableContainer := makeTableBox()
	return container.NewVSplit(sqlArea, tableContainer)
}

func makeTableBox() fyne.CanvasObject {
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

		return table
	}
	tableContainer.Objects = []fyne.CanvasObject{createTable()}
	return tableContainer
}

func makeSelectBox() fyne.CanvasObject {
	departmentIdEntry := widget.NewEntry()
	departmentIdEntry.PlaceHolder = "输入企业id"
	// todo 企业id框失去焦点之后自动查询店铺

	limitSelect := widget.NewSelect([]string{
		"100", "500", "1000",
	}, func(s string) {})
	limitSelect.SetSelected("100")
	button := &widget.Button{
		Text:       "SQL查询",
		Importance: widget.HighImportance,
		OnTapped: func() {
			fmt.Println("执行sql查询")
		},
	}
	// 在 Fyne 中，当 container.NewBorder 中的 right 部分加入了 VBox 容器时，
	//Resize 方法不会强制应用指定的大小。container.NewBorder 会根据布局来自动调整 right 的大小，
	//而 Resize 仅适用于独立组件，嵌入到其他布局的组件会被布局重新计算大小。
	box := container.NewVBox(departmentIdEntry, limitSelect, button)
	rightContainer := container.NewGridWrap(fyne.NewSize(200, 100), box)
	return rightContainer
}

//func makeSelectBox() fyne.CanvasObject {
//	instanceSelect := widget.NewSelect([]string{
//		"实例1", "实例2", "实例3",
//	}, func(s string) {})
//	instanceSelect.PlaceHolder = "选择一个实例"
//	databaseSelect := widget.NewSelect([]string{
//		"数据库1", "数据库2", "数据库3",
//	}, func(s string) {})
//	databaseSelect.PlaceHolder = "选择一个数据库"
//	limitSelect := widget.NewSelect([]string{
//		"100", "500", "1000",
//	}, func(s string) {})
//	limitSelect.SetSelected("100")
//
//	button := &widget.Button{
//		Text:       "SQL查询",
//		Importance: widget.HighImportance,
//		OnTapped: func() {
//			fmt.Println("执行sql查询")
//		},
//	}
//
//	return container.NewVBox(instanceSelect, databaseSelect, limitSelect, button)
//}
