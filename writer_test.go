package ico

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
)

func readPng(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func TestWriter(t *testing.T) {
	fn := "icondata.png"
	m0, err := readPng(fn)
	if err != nil {
		t.Error(fn, err)
	}
	icoimg, _ := os.Create("new.ico")
	defer icoimg.Close()
	icoerror := Encode(icoimg, m0)
	fmt.Printf("icoerror: %v\n", icoerror)
}

func TestWriterWithMultipleImage(t *testing.T) {
	fn1 := "icondata-small.png"
	m0, err := readPng(fn1)
	if err != nil {
		t.Error(fn1, err)
	}
	fn2 := "icondata.png"
	m1, err := readPng(fn2)
	if err != nil {
		t.Error(fn2, err)
	}
	icoimg, _ := os.Create("new2.ico")
	defer icoimg.Close()
	icoerror := Encode(icoimg, m0, m1)
	fmt.Printf("icoerror: %v\n", icoerror)

}
