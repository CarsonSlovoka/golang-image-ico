package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"png2ico/pkg/image/ico"
	"png2ico/pkg/image/png"
)

func main() {
	var (
		inputPng       = flag.String("fromFile", "", "Input source PNG")
		outputFilename = flag.String("name", "", "The output filepath of ICO")
		isForceMode    = flag.Bool("f", false, "(Optional) Force write to the output file, regardless of whether the output file already exists.")
	)

	flag.Parse()

	if *inputPng == "" || *outputFilename == "" {
		flag.Usage()
		os.Exit(1)
	}

	outputPath := filepath.Join(filepath.Dir(*inputPng), *outputFilename)
	if _, err := os.Stat(outputPath); !os.IsNotExist(err) && !*isForceMode {
		log.Fatalf("The output file already exists.\n%v", err)
	}

	sourceImg, err := png.ReadImg(*inputPng)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()
	if err = ico.Encode(outFile, sourceImg); err != nil {
		log.Fatal(err)
	}

	// Open the output directory.
	exec.Command("explorer", filepath.Dir(outputPath)).Start()
}
