package main

import (
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
)

func deleteInSlice[S any](slice []S, idx int) []S {
	if idx < 0 || idx >= len(slice) {
		return slice
	}

	s := slice[:idx]

	if idx != len(slice)-1 {
		s = append(s, slice[idx+1:]...)
	}

	return s
}

func orderClicked() {
	amount, err := strconv.Atoi(strings.TrimSpace(amountInput.Text()))
	if err != nil {
		amount = 0
	}

	if bidaskGroup.Value == "bid" {
		amount = -amount
	}

	processOrder(orderAsset, orderPrice, amount)

	amountInput.SetText("")
	assetInput.SetText("")

	orderAllowed = false
}

func tickerClicked() {
	name := strings.TrimSpace(tickerInput.Text())
	if name != "" {
		myTickers = append(myTickers, ticker{
			name:   name,
			button: widget.Clickable{},
			value:  1000,
		})

	}

	tickerInput.SetText("")
}

func chatClicked() {
	prompt := strings.TrimSpace(chatInput.Text())
	if chatLog.Len() > 0 {
		chatLog.WriteString("\n")
	}
	chatLog.WriteString(prompt)
	chatLog.WriteString("\n")
	chatLog.WriteString(responses[rand.Intn(3)])

	chatInput.SetText("")
}

var responses = []string{
	"Every step you do is getting you closer to the goal",
	"No pain no gain, you know?",
	"You can't do it if you're not good enough",
}

func deleteClicked() {
	idx := slices.IndexFunc(myTickers, func(t ticker) bool {
		return t.name == orderAsset
	})

	if idx != -1 {
		myTickers = deleteInSlice(myTickers, idx)
	}

	orderAsset = ""
	orderAllowed = false
}

func cancelClicked() {
	amountInput.SetText("")
	assetInput.SetText("")
	orderAsset = ""

	orderAllowed = false
}

func processMyTickersButtons(gtx layout.Context) {
	for i := range myTickers {
		if myTickers[i].button.Clicked(gtx) {
			idx := slices.IndexFunc(myTickers, func(t ticker) bool {
				return t.name == myTickers[i].name
			})

			if idx != -1 {
				orderAllowed = true
				orderAsset = myTickers[i].name
				orderPrice = myTickers[i].value
			}

			break
		}
	}
}
