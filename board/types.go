package board

type Board [Size][Size]int
type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

var MoveNames = map[Move]string{
	Up:    "Up",
	Down:  "Down",
	Left:  "Left",
	Right: "Right",
}

const (
	Size      = 4
	MaxRounds = 30
)
