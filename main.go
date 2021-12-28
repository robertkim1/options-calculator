package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"optionsCalculator/blackScholes"
	"strconv"
)

type numericalEntry struct {
	widget.Entry
}
func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}
func (e *numericalEntry) TypedRune(r rune) {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		e.Entry.TypedRune(r)
	}
}
func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func main() {

	//var spotPrice float64 = 97.39
	//var strikePrice float64 = 97
	//var time float64 = .0712328767
	//var volatility float64 = .4906
	//rfRate := .0141
	//call price output: 5.32
	//put price output: 4.83

	myApp := app.New()
	mainWindow := myApp.NewWindow("Options Calculator")
	mainWindow.Resize(fyne.NewSize(600, 600))

	typeOfOptionLabel := widget.NewLabel("Type of Option")
	var optionTypeString string
	optionTypeRead := widget.NewSelect([]string{"Call Option","Put Option"}, func(value string) {
		optionTypeString = value
	})
	spotPriceLabel := widget.NewLabel("Current Share Price")
	spotPrice := newNumericalEntry()
	spotPrice.SetPlaceHolder("Enter current share price")
	strikePriceLabel := widget.NewLabel("Strike Price of Option")
	strikePrice := newNumericalEntry()
	strikePrice.SetPlaceHolder("Enter strike price")
	timeLabel := widget.NewLabel("Time until Expiration in Years")
	time := newNumericalEntry()
	time.SetPlaceHolder("Enter time until expiration (years)")
	volatilityLabel := widget.NewLabel("Volatility")
	volatility := newNumericalEntry()
	volatility.SetPlaceHolder("Enter volatility percentage (decimal)")
	rfRateLabel := widget.NewLabel("Risk-Free Interest Rate")
	rfRate := newNumericalEntry()
	rfRate.SetPlaceHolder("Enter current risk-free interest rate")
	output := widget.NewLabel("")

	calculate := widget.NewButton("Calculate", func() {
		spotPriceFloat, _ := strconv.ParseFloat(spotPrice.Text, 64)
		strikePriceFloat, _ := strconv.ParseFloat(strikePrice.Text, 64)
		timeFloat, _ := strconv.ParseFloat(time.Text, 64)
		volatilityFloat, _ := strconv.ParseFloat(volatility.Text, 64)
		rfRateFloat, _ := strconv.ParseFloat(rfRate.Text, 64)
		d1 := blackScholes.D1(spotPriceFloat, strikePriceFloat, rfRateFloat, volatilityFloat, timeFloat)
		d2 := blackScholes.D2(d1, volatilityFloat, timeFloat)
		if optionTypeString == "Call Option" {
			callPrice := blackScholes.CallPrice(d1, spotPriceFloat, d2, strikePriceFloat, rfRateFloat, timeFloat)
			output.SetText(fmt.Sprintf("$%.2f", callPrice))
		} else {
			putPrice := blackScholes.PutPrice(d1, spotPriceFloat, d2, strikePriceFloat, rfRateFloat, timeFloat)
			output.SetText(fmt.Sprintf("$%.2f", putPrice))
		}
	})

	content := container.NewVBox(typeOfOptionLabel, optionTypeRead, spotPriceLabel, spotPrice,strikePriceLabel, strikePrice, timeLabel, time, volatilityLabel, volatility, rfRateLabel, rfRate, calculate, output)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()

}
