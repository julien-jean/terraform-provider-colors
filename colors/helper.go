package colors

import (
	"math/rand"
)

var colors = []Color{
	{
		Name: "red",
		Code: 0xFF0000,
	},
	{
		Name: "green",
		Code: 0x00FF00,
	},
	{
		Name: "blue",
		Code: 0x0000FF,
	},
	{
		Name: "white",
		Code: 0xFFF,
	},
	{
		Name: "black",
		Code: 0x000,
	},
}

func RandomColor() Color {
	pickedColor := colors[rand.Intn(len(colors))]

	return pickedColor
}
