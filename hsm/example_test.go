package hsm

import (
	"fmt"
)

func Example_callback() {
	enter := func(h *HSM) { fmt.Println("enter " + h.Current()) }
	exit := func(h *HSM) { fmt.Println("exit " + h.Current()) }

	scenes := []*State{
		{Name: "/", Enter: enter, Exit: exit},
		{Name: "/title", Enter: enter, Exit: exit},
		{Name: "/game", Enter: enter, Exit: exit},
		{Name: "/game/start", Enter: enter, Exit: exit},
		{Name: "/game/main", Enter: enter, Exit: exit},
	}

	fmt.Println("---")
	scene := NewHSM(scenes, "/title", 1)
	fmt.Println("---")
	scene.Change("/game/start")
	fmt.Println("---")
	scene.Change("/game/main")
	fmt.Println("---")
	scene.Change("/title")

	// Output:
	// ---
	// enter /
	// enter /title
	// ---
	// exit /title
	// enter /game
	// enter /game/start
	// ---
	// exit /game/start
	// enter /game/main
	// ---
	// exit /game/main
	// exit /game
	// enter /title
}

func Example_update() {
	update := func(h *HSM) {
		fmt.Printf("%d: update %s %d\n", h.TickOf("/").Elapsed(), h.Current(), h.Tick().Elapsed())
	}
	draw := func(h *HSM) {
		fmt.Printf("%d: draw %s %d\n", h.TickOf("/").Elapsed(), h.Current(), h.Tick().Elapsed())
	}

	scenes := []*State{
		{Name: "/title", Update: update, Draw: draw},
		{Name: "/game", Update: update, Draw: draw},
		{Name: "/game/start", Update: update, Draw: draw},
		{Name: "/game/main", Update: update, Draw: draw},
	}

	fmt.Println("---")
	scene := NewHSM(scenes, "/title", 1)
	scene.Draw()
	scene.Update()
	scene.Draw()
	scene.Update()

	fmt.Println("---")
	scene.Change("/game/start")
	scene.Draw()
	scene.Update()
	scene.Draw()
	scene.Update()

	fmt.Println("---")
	scene.Change("/game/main")
	scene.Draw()
	scene.Update()
	scene.Draw()
	scene.Update()

	fmt.Println("---")
	scene.Change("/title")
	scene.Draw()
	scene.Update()
	scene.Draw()
	scene.Update()

	// Output:
	// ---
	// 0: draw /title 0
	// 1: update /title 1
	// 1: draw /title 1
	// 2: update /title 2
	// ---
	// 2: draw /game/start 0
	// 3: update /game/start 1
	// 3: draw /game/start 1
	// 4: update /game/start 2
	// ---
	// 4: draw /game/main 0
	// 5: update /game/main 1
	// 5: draw /game/main 1
	// 6: update /game/main 2
	// ---
	// 6: draw /title 0
	// 7: update /title 1
	// 7: draw /title 1
	// 8: update /title 2
}
