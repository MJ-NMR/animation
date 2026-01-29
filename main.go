package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

const (
	frame = 60
	scale = 10
	h     = 9 * scale
	w     = 16 * scale
)

func main() {
	if err := os.Mkdir("output", 0755); err != nil && os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for num := range frame {
		wg.Add(1)
		ppm(num, &wg)
	}
	wg.Wait()

	cmd := exec.Command("ffmpeg", "-i", "output/image-%02d.ppm", "-r", "60", "output/output.mp4")
	if err := cmd.Run(); err != nil {
		fmt.Println("ffmpeg Error:", err)
		return
	}
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
