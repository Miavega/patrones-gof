package factorymethod

import "fmt"

type format struct {
	multimedia string
}

func (f *format) Play() {
	fmt.Println(f.multimedia)
}
