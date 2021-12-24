package view

import (
	"fmt"
	"testing"

	"github.com/mpstewart/go-cgol/internal/cgol"
)

func TestGameModel_View(t *testing.T) {
	m := InitialModel(0, 0)
	got := MustView("View doesn't panic", t, m.View)
	want := fmt.Sprintf(
		"%s%s",
		m.boardView(),
		m.statusBarView(),
	)
	if got != want {
		t.Errorf("GameModel.View() = %v, want %v", got, want)
	}
}

func TestGameModel_boardView(t *testing.T) {
	m := InitialModel(2, 2)
	got := m.boardView()
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

func TestGameModel_statusBarView(t *testing.T) {
	type fields struct {
		currentBoard *cgol.Board
		paused       bool
		tickRate     int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test 1",
			fields: fields{
				currentBoard: &cgol.Board{Height: 1, Width: 2},
				paused:       false,
				tickRate:     int64(250),
			},
			want: "h: 1, w: 2, Hz: 4.00, paused: false",
		},
		{
			name: "test 2",
			fields: fields{
				currentBoard: &cgol.Board{Height: 100, Width: 200},
				paused:       true,
				tickRate:     int64(500),
			},
			want: "h: 100, w: 200, Hz: 2.00, paused: true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := GameModel{
				currentBoard: tt.fields.currentBoard,
				paused:       tt.fields.paused,
				tickRate:     tt.fields.tickRate,
			}
			if got := m.statusBarView(); got != tt.want {
				t.Errorf("GameModel.statusBarView() = %v, want %v", got, tt.want)
			}
		})
	}
}
