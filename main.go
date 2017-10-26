package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unidoc/pdf/creator"
	"github.com/unidoc/unidoc/pdf/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], `"path to pdf"`)

		return
	}
	path := os.Args[1]

	fp, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer fp.Close()
	reader, err := model.NewPdfReader(fp)
	dir := filepath.Dir(path)
	_, pathFileName := filepath.Split(path)
	base := strings.TrimRight(pathFileName, filepath.Ext(path))

	if err != nil {
		panic(err)
	}

	evenCreator := creator.New()

	for i := 0; i < len(reader.PageList); i += 2 {
		page, err := reader.GetPage(i + 1)

		if err != nil {
			panic(err)
		}

		if err := evenCreator.AddPage(page); err != nil {
			panic(err)
		}
	}

	if err := evenCreator.WriteToFile(filepath.Join(dir, base+"_even.pdf")); err != nil {
		panic(err)
	}

	oddCreator := creator.New()

	for i := 1; i < len(reader.PageList); i += 2 {
		page, err := reader.GetPage(i + 1)

		if err != nil {
			panic(err)
		}

		if err := oddCreator.AddPage(page); err != nil {
			panic(err)
		}
	}

	if len(reader.PageList)&1 != 0 {
		oddCreator.NewPage()
	}
	if err := oddCreator.WriteToFile(filepath.Join(dir, base+"_odd.pdf")); err != nil {
		panic(err)
	}

	fmt.Println("OK")
}
