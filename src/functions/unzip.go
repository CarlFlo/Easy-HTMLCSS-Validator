package functions

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// användning
// functions.Unzip("./Ny mapp/Choco-Latte.zip", "./Ny mapp/")
// https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang

// Unzip unzips a file
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsZip returns true if string ends with .zip
func IsZip(str string) bool {

	if strings.EqualFold(str[len(str)-4:], ".zip") { // Its a zip file
		return true
	} else if strings.EqualFold(str[len(str)-4:], ".rar") { // Its a rar file
		fmt.Println("ERROR: .rar is not supported yet. Use zip or a plain folder instead")
		time.Sleep(time.Second * 10)
		os.Exit(1)
	}

	fmt.Println("String extension is:", str)

	return false
}
