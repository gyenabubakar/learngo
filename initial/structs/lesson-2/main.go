package main

import (
	"fmt"
	"strings"
)

type song struct {
	artist string
	title  string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	afro := playlist{
		genre: "Afro-fusion",
		songs: []song{
			{artist: "Burna Boy", title: "On the Low"},
			{artist: "Burna Boy", title: "Anybody"},
		},
	}

	//song := afro.songs[0]
	afro.songs[0].title = "Odogwu"

	fmt.Printf("%-2s %-20s %-20s %s\n", "👇", "Title", "Artist", "Genre")
	fmt.Println(strings.Repeat("-", 60))

	for i, song := range afro.songs {
		fmt.Printf("%-2d %-20s %-20s %s\n", i+1, song.title, song.artist, afro.genre)
	}
}
