package tick

import "fmt"

func ExampleTick_Span() {

	t := Tick{}
	for i := 0; i < 9; i++ {

		// We recommend to call Advance at the beginning of the loop.
		// So the tick starts with 1.
		t.Advance(1)
		e0 := t.Elapsed()

		// Chained spans.
		t.Span(0, 3, func(t Tick) { // tick 0, 1, 2
			e1 := t.Elapsed()
			fmt.Println(e0, e1)
		}).Span(2, 4, func(t Tick) { // tick 5, 6, 7, 8 (3, 4 are skipped)
			e1 := t.Elapsed()
			fmt.Println(e0, e1)
		})

		// Another span.
		t.Span(2, 4, func(t Tick) { // tick 2, 3, 4, 5
			fmt.Println(e0, "another span")
		})
	}

	// Output:
	// 1 1
	// 2 2
	// 2 another span
	// 3 another span
	// 4 another span
	// 5 0
	// 5 another span
	// 6 1
	// 7 2
	// 8 3
}

func ExampleTick_From() {

	t := Tick{}
	for i := 0; i < 5; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.From(2, func(t Tick) { // tick 2, 3, 4, 5, ...
			e1 := t.Elapsed()
			fmt.Println(e0, e1)
		})
	}

	// Output:
	// 2 0
	// 3 1
	// 4 2
	// 5 3
}

func ExampleTick_Once() {

	t := Tick{}
	for i := 0; i < 5; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.Span(0, 3, func(t Tick) { // tick 0, 1, 2
			e1 := t.Elapsed()
			fmt.Println(e0, e1)
		}).Once(0, func() { // tick 3
			fmt.Println(e0)
		})
	}

	// Output:
	// 1 1
	// 2 2
	// 3
}

func ExampleTick_Repeat() {

	t := Tick{}
	for i := 0; i < 10; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.Repeat(2, 3, func(n int, t Tick) { // tick 2, 3, 4; 5, 6, 7; 8, 9, 10;
			e1 := t.Elapsed()
			fmt.Println(e0, n, e1)
		})
	}

	// Output:
	// 2 0 0
	// 3 0 1
	// 4 0 2
	// 5 1 0
	// 6 1 1
	// 7 1 2
	// 8 2 0
	// 9 2 1
	// 10 2 2
}

func ExampleTick_Every() {

	t := Tick{}
	for i := 0; i < 9; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.Every(1, 3, func(n int) { // tick 1, 4, 7
			fmt.Println(e0, n)
		})
	}

	// Output:
	// 1 0
	// 4 1
	// 7 2
}

func ExampleTick_Every_advance() {

	t := Tick{}
	for i := 0; i < 9; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(9)
		e0 := t.Elapsed()

		// The callback will be called in the first loop after the original tick has passed.

		t.Every(10, 30, func(n int) { // tick 18, 45, 72 (original tick: 10, 40, 70)
			fmt.Println(e0, n)
		})
	}

	// Output:
	// 18 0
	// 45 1
	// 72 2
}

func ExampleTick_RepeatFor() {

	t := Tick{}
	for i := 0; i < 10; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.RepeatFor(1, 3, 2, func(n int, t Tick) { // tick 1, 2, 3; 4, 5, 6
			e1 := t.Elapsed()
			fmt.Println(e0, n, e1)
		}).Once(0, func() {
			fmt.Println(e0, "done")
		})
	}

	// Output:
	// 1 0 0
	// 2 0 1
	// 3 0 2
	// 4 1 0
	// 5 1 1
	// 6 1 2
	// 7 done
}

func ExampleTick_EveryFor() {

	t := Tick{}
	for i := 0; i < 10; i++ {

		// We recommend to call Advance at the beginning of the loop.
		t.Advance(1)
		e0 := t.Elapsed()

		t.EveryFor(1, 2, 3, func(n int) { // tick 1, 3, 5
			fmt.Println(e0, n)
		}).Once(0, func() {
			fmt.Println(e0, "done")
		})
	}

	// Output:
	// 1 0
	// 3 1
	// 5 2
	// 7 done
}
