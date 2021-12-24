package view

import (
	"fmt"
	"testing"
)

func TestGameModel_View(t *testing.T) {
	m := InitialModel(0, 0)
	MustView("View doesn't panic", t, m.View)
}

func TestGameModel_BoardView(t *testing.T) {
	m := InitialModel(2, 2)
	got := m.BoardView()
	want := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		"┌─────┐",
		"│     │",
		"│     │",
		"└─────┘",
	)
	if got != want {
		t.Errorf("Board looks wrong")
		t.Logf("Got:\n%s", got)
		t.Logf("Want:\n%s", want)
	}
}

func MustView(name string, t *testing.T, fn func() string) string {
	var s string
	t.Run(name, func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Caught a panic: %s", r)
			}
		}()
		s = fn()
	})

	return s
}
