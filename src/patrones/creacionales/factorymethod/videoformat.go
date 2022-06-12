package factorymethod

type video struct {
	format
}

func newVideo() IFormat {
	return &video{
		format: format{
			multimedia: "Video format is playing",
		},
	}
}
