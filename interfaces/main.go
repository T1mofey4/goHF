package main

import "interfaces/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	player.Record()
}

func main() {

	// var player Player = gadget.TapePlayer{}
	// mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	// playList(player, mixtape)
	// player = gadget.TapeRecorder{}
	// playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
}
