package factorymethod

type audio struct {
	format
}

func newAudio() IFormat {
	return &audio{
		format: format{
			multimedia: "Audio format is playing",
		},
	}
}
