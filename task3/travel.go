package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func Travel() {
	window := ui.NewWindow("Розрахувати відпочинок", 400, 250, false)

	boxLabels := ui.NewHorizontalBox()
	countryLabel := ui.NewLabel("Країна")
	seasonLabel := ui.NewLabel("Пора року")
	dayLabel := ui.NewLabel("Кількість днів")
	boxLabels.Append(countryLabel, true)
	boxLabels.Append(seasonLabel, true)
	boxLabels.Append(dayLabel, true)
	boxLabels.SetPadded(true)

	boxLists := ui.NewHorizontalBox()
	countryList := ui.NewCombobox()
	countryList.Append("Болгарія")
	countryList.Append("Німеччина")
	countryList.Append("Польща")
	countryList.SetSelected(0)
	seasonList := ui.NewCombobox()
	seasonList.Append("літо")
	seasonList.Append("зима")
	seasonList.SetSelected(0)
	dayList := ui.NewCombobox()
	for i := 1; i <= 7; i++ {
		dayList.Append(strconv.Itoa(i))
	}
	dayList.SetSelected(0)
	boxLists.Append(countryList, false)
	boxLists.Append(seasonList, false)
	boxLists.Append(dayList, false)
	boxLists.SetPadded(true)

	boxCheckbox := ui.NewVerticalBox()
	guide := ui.NewCheckbox("Індивідуальний гід")
	guide.SetChecked(true)
	luxuryHotel := ui.NewCheckbox("Проживання в люкс номері")
	luxuryHotel.SetChecked(true)
	button := ui.NewButton("Розрахувати вартість відпустки")
	boxCheckbox.Append(guide, false)
	boxCheckbox.Append(luxuryHotel, false)
	boxCheckbox.Append(button, false)
	boxCheckbox.SetPadded(true)

	boxResult := ui.NewHorizontalBox()
	resultLabel := ui.NewLabel("Вартість = ")
	resultMoneyLabel := ui.NewLabel("")
	boxResult.Append(resultLabel, false)
	boxResult.Append(resultMoneyLabel, false)
	boxResult.SetPadded(true)
	boxResult.Hide()

	box := ui.NewVerticalBox()
	box.Append(boxLabels, false)
	box.Append(boxLists, false)
	box.Append(boxCheckbox, false)
	box.Append(boxResult, false)
	box.SetPadded(true)

	window.SetChild(box)
	window.SetMargined(true)

	button.OnClicked(func(*ui.Button) {
		country := countryList.Selected()
		season := seasonList.Selected()
		days := float64(dayList.Selected() + 1)
		guideSelect := guide.Checked()
		luxury := luxuryHotel.Checked()
		var valueOfMoney float64 = 0
		if season == 0 {
			switch country {
			case 0:
				valueOfMoney = 100
			case 1:
				valueOfMoney = 160
			case 2:
				valueOfMoney = 120
			}
		} else {
			switch country {
			case 0:
				valueOfMoney = 150
			case 1:
				valueOfMoney = 200
			case 2:
				valueOfMoney = 180
			}
		}

		valueOfMoney *= days
		if guideSelect {
			valueOfMoney += 50 * days
		}
		if luxury {
			valueOfMoney += valueOfMoney * 0.2
		}
		moneyStr := strconv.FormatFloat(valueOfMoney, 'f', -1, 64) + " $"
		boxResult.Show()
		resultMoneyLabel.SetText(moneyStr)
	})

	window.Show()
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}
