package main

type Color struct {
	Reset   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	Gray    string
	White   string
}

func initColor() *Color {
	var color Color
	color.Reset = "\033[0m"
	color.Red = "\033[31m"
	color.Green = "\033[32m"
	color.Yellow = "\033[33m"
	color.Blue = "\033[34m"
	color.Magenta = "\033[35m"
	color.Cyan = "\033[36m"
	color.Gray = "\033[37m"
	color.White = "\033[97m"
	return &color
}
