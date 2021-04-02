package appkit

type BoxType uint

const (
	BoxPrimary   BoxType = 0
	BoxSeparator BoxType = 2
	BoxCustom    BoxType = 4
)

type TitlePosition uint

const (
	NoTitle     TitlePosition = 0
	AboveTop    TitlePosition = 1
	AtTop       TitlePosition = 2
	BelowTop    TitlePosition = 3
	AboveBottom TitlePosition = 4
	AtBottom    TitlePosition = 5
	BelowBottom TitlePosition = 6
)
