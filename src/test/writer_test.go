package main_test

import (
	"image"
	"image/png"
	"os"
	"png2ico/pkg/image/ico"
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
	fn := "data/icondata.png"
	m0, err := readPng(fn)
	if err != nil {
		t.Error(fn, err)
	}
	outIco, _ := os.Create("output.ico")
	defer func() {
		outIco.Close()
		os.Remove("output.ico")
	}()

	if err := ico.Encode(outIco, m0); err != nil {
		t.Fatalf("ico.Encode error: %v\n", err)
	}

	// Check output successful
	if _, err := os.Stat("output.ico"); os.IsNotExist(err) {
		t.Fatalf("output ico not exists error: %v\n", err)
	}
}

func TestWriterWithMultipleImage(t *testing.T) {
	fn1 := "data/gopher.png"
	m0, err := readPng(fn1)
	if err != nil {
		t.Error(fn1, err)
	}
	fn2 := "data/icondata.png"
	m1, err := readPng(fn2)
	if err != nil {
		t.Error(fn2, err)
	}
	outIco, _ := os.Create("output.ico")
	defer func() {
		outIco.Close()
		os.Remove("output.ico")
	}()

	if err := ico.Encode(outIco, m0, m1); err != nil {
		t.Fatalf("ico.Encode error: %v\n", err)
	}
	if _, err := os.Stat("output.ico"); os.IsNotExist(err) {
		t.Fatalf("output ico not exists error: %v\n", err)
	}
}
