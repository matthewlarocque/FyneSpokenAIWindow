package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func appendText(str string, consoleCard *widget.Card) *widget.Label {
	currentText := consoleCard.Content.(*widget.Label).Text
	newText := currentText + "\n" + str
	return widget.NewLabel(newText)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("SpokenAI")
	myWindow.CenterOnScreen()

	text := widget.NewLabel("This is an English oral conversation practice program, powered by OpenAI. " +
		"\nPlease double check whether the configuration file is correctly configured, such as the key, before using it." +
		"\nNext, you need to select a teacher of your favorite type from the drop-down list and enter your nickname in the text box. " +
		"\nFinally, click the \"Talk\" button to start recording, click the \"Stop\" button to end the recording, " +
		"\nand click the \"Submit\" button to transmit it to GPT bot.")

	options := []string{"a patient and humble oral English teacher", "a gentle and rigorous oral English teacher",
		"a irascible and irritable oral English teacher", "a sarcastic ridicule oral English teacher"}
	combo := widget.NewSelect(options, func(s string) {})
	combo.SetSelectedIndex(0)

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Name")

	button1 := widget.NewButton("Talk", nil)
	button2 := widget.NewButton("Stop", nil)
	button3 := widget.NewButton("Submit", nil)

	button2.Disable()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Your English Name", Widget: entry},
			{Text: "Select Teacher Type", Widget: combo},
		},
	}

	consoleCard := widget.NewCard("Console logs", "", widget.NewLabel("Line 1\nLine 2\nLine 3\nLine 4"))
	consoleScroll := container.NewVScroll(consoleCard)
	consoleScroll.SetMinSize(fyne.NewSize(0, 200))

	content := container.NewVBox(
		text,
		form,
		container.NewCenter(
			container.NewHBox(
				button1,
				button2,
				button3,
			),
		),
		consoleScroll,
	)

	button1.OnTapped = func() {
		consoleCard.SetContent(appendText("Nuevo", consoleCard))
	}

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
