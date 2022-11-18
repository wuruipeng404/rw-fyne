/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2022/11/18 10:28
 */

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func init() {
	fontPath := findfont.List()

	for _, p := range fontPath {

		if strings.Contains(p, "simkai.ttf") {
			os.Setenv("FYNE_FONT", p)
			break
		}
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Rumple ❤️ 嘻嘻")
	myWindow.Resize(fyne.NewSize(920, 560))
	myWindow.CenterOnScreen()

	b64r := widget.NewLabel("")
	input := widget.NewEntry()

	bottomBox := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButtonWithIcon("复制结果", theme.ContentCopyIcon(), func() {
			if b64r.Text != "" {
				myWindow.Clipboard().SetContent(b64r.Text)
			}
		}),
	)

	grid := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("请输入计算文本"), input,
		widget.NewLabel("base64计算结果"), b64r,
	)

	cgrid := container.NewPadded(grid)

	border := container.NewBorder(nil, bottomBox, nil, nil, cgrid)

	// submit := container.NewCenter(container.NewHBox(widget.NewButton("提交", func() {
	// 	if input.Text != "" {
	// 		res := base64.StdEncoding.EncodeToString([]byte(input.Text))
	// 		b64r.SetText(res)
	// 	}
	// })))
	//
	// vb := container.NewVBox(grid, submit)

	tabs := container.NewPadded(container.NewAppTabs(
		container.NewTabItem("LabelmeParser", border),
		container.NewTabItem("小工具", widget.NewLabel("11111")),
	))

	// content := container.New(layout.NewGridLayout(10))

	// textc := container.New(layout.NewCenterLayout(), canvas.NewText("爱你哟", colornames.Red))

	// content := widget.NewButton("点我", func() {
	// 	myWindow.SetContent(textc)
	// })

	// content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	// })

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Clock")
//
// 	clock := widget.NewLabel("")
// 	updateTime(clock)
//
// 	w.SetContent(clock)
// 	w.Resize(fyne.NewSize(720, 720))
// 	w.SetMaster()
// 	// w2 := a.NewWindow("rumple")
// 	// w2.SetContent(widget.NewLabel("new rumple"))
// 	// w2.Resize(fyne.NewSize(100, 100))
// 	// w2.Show()
//
// 	go func() {
// 		for range time.Tick(time.Second) {
// 			updateTime(clock)
// 		}
// 	}()
//
// 	w.ShowAndRun()
// }
