package main

import "github.com/eihigh/zu"

var (
	quit = zu.Input{}
)

func main() {
	zu.PushViewFunc(view)
	for zu.Next() {
		if quit.IsDown() {
			break
		}
	}
}

func popup() {
	zu.PushViewFunc(view)
	defer zu.PopView()

	for zu.Next() {
		if quit.IsDown() {
			break
		}
	}
}

func view() {

}
