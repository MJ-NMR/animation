package main

import (
	"fmt"
	"os"
	"sync"
)

const (
	frame = 60
	scale = 10
	h     = 9 * scale
	w     = 16 * scale
)

func main() {
	var wg sync.WaitGroup
	for num := range frame {
		wg.Add(1)
		ppm(num, &wg)
	}

	wg.Wait()
}

func ppm(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	outputf := fmt.Sprintf("output/image-%02d.ppm", num)
	f, err := os.Create(outputf)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer f.Close()
	fmt.Printf("start on %02d \n", num)

	fmt.Fprint(f, "P6\n")
	fmt.Fprintf(f, "%d %d\n", w, h)
	fmt.Fprintf(f, "255\n")
	for y := range h {
		for range w {
			if ((y+num)/scale)%2 == 0 {
				f.Write([]byte{0xff, 0x00, 0x00})
			} else {
				f.Write([]byte{0x00, 0xff, 0x00})
			}
		}
	}
	fmt.Printf("Genatare %v\n", outputf)
}
