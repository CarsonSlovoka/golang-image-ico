package main_test

import (
	"os"
	"png2ico/pkg/image/ico"
	png2 "png2ico/pkg/image/png"
	"testing"
)

func TestWriter(t *testing.T) {
	fn := "data/icondata.png"
	m0, err := png2.ReadImg(fn)
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
	m0, err := png2.ReadImg(fn1)
	if err != nil {
		t.Error(fn1, err)
	}
	fn2 := "data/icondata.png"
	m1, err := png2.ReadImg(fn2)
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
