package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func Glass() {
	window := ui.NewWindow("Калькулятор склопакета", 550, 250, false)

	boxTitle := ui.NewHorizontalBox()
	widthWindow := ui.NewLabel("Розмір вікна")
	glass := ui.NewLabel("Склопакет")
	boxTitle.Append(widthWindow, true)
	boxTitle.Append(glass, true)
	boxTitle.SetPadded(true)

	boxLabelsAndResult := ui.NewVerticalBox()
	widthLabel := ui.NewLabel("Ширина, см")
	heightLabel := ui.NewLabel("Висота, см")
	materialLabel := ui.NewLabel("Матеріал")
	boxLabelsAndResult.Append(widthLabel, true)
	boxLabelsAndResult.Append(heightLabel, true)
	boxLabelsAndResult.Append(materialLabel, true)
	boxLabelsAndResult.SetPadded(true)

	boxEntriesAndMaterialList := ui.NewVerticalBox()
	widthInput := ui.NewEntry()
	heightInput := ui.NewEntry()
	materialList := ui.NewCombobox()
	materialList.Append("Дерево")
	materialList.Append("Метал")
	materialList.Append("Металопластик")
	materialList.SetSelected(0)
	boxEntriesAndMaterialList.Append(widthInput, false)
	boxEntriesAndMaterialList.Append(heightInput, false)
	boxEntriesAndMaterialList.Append(materialList, true)
	boxEntriesAndMaterialList.SetPadded(true)

	boxGlassAndWindowsill := ui.NewVerticalBox()
	countGlass := ui.NewCombobox()
	countGlass.Append("Однокамерний") //0
	countGlass.Append("Двокамерний")
	countGlass.SetSelected(0)
	windowsill := ui.NewCheckbox("Підвіконня")
	windowsill.SetChecked(true)
	boxGlassAndWindowsill.Append(countGlass, false)
	boxGlassAndWindowsill.Append(windowsill, false)
	boxGlassAndWindowsill.SetPadded(true)

	boxResult := ui.NewHorizontalBox()
	result := ui.NewLabel("")
	button := ui.NewButton("Розрахувати")
	boxResult.Append(result, true)
	boxResult.Append(button, true)
	boxResult.SetPadded(true)

	boxCentral := ui.NewHorizontalBox()
	boxCentral.Append(boxLabelsAndResult, false)
	boxCentral.Append(boxEntriesAndMaterialList, false)
	boxCentral.Append(boxGlassAndWindowsill, false)
	boxCentral.SetPadded(true)

	box := ui.NewVerticalBox()
	box.Append(boxTitle, false)
	box.Append(boxCentral, false)
	box.Append(boxResult, false)
	box.SetPadded(true)

	window.SetChild(box)
	window.SetMargined(true)

	button.OnClicked(func(button2 *ui.Button) {

		var (
			money           float64
			valueOfMaterial float64
		)

		width, err1 := strconv.ParseFloat(widthInput.Text(), 64)
		countGlasses := countGlass.Selected()
		height, err2 := strconv.ParseFloat(heightInput.Text(), 64)
		windowsillSelect := windowsill.Checked()
		material := materialList.Selected()

		if countGlasses == 0 {
			switch material {
			case 0:
				valueOfMaterial = 2.5
			case 1:
				valueOfMaterial = 0.5
			case 2:
				valueOfMaterial = 1.5
			}
		} else {
			switch material {
			case 0:
				valueOfMaterial = 3
			case 1:
				valueOfMaterial = 1
			case 2:
				valueOfMaterial = 2
			}
		}

		if err1 == nil && err2 == nil {
			money = width * height * valueOfMaterial
			if windowsillSelect {
				money += 350
			}
			valueOfMoney := strconv.FormatFloat(money, 'f', -1, 64) + " грн"
			result.SetText(valueOfMoney)
		} else {
			result.SetText("Помилка. Перевірте, чи в полях вводу записані лише цифри")
		}
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
}
