package main

import (
	"log";
	"fmt";
	"path";
	"os";
	"strings";
	)


func WalkFunc(file_path string, file os.FileInfo, err error) error {
	if (IsFileToMove(file)) {
		dest, err := GetDest(file, *dest_root)

		if (err != nil) {
			log.Fatal(err)
		}

		new_path := path.Join(dest, file.Name())

		err = os.Rename(file_path, new_path)

		if (err != nil) {
			log.Fatal(err)
		}
	}

	return err
}

func GetDest(file os.FileInfo, dest_root string) (string, error) {
	year, month, day := file.ModTime().Date()
	dest := fmt.Sprintf("%s/%d/%s/%02d", dest_root, year, GetMonthsName(month), day)

	return dest, os.MkdirAll(dest, 0775)
}

func IsFileToMove(file os.FileInfo) bool {
	if (file.IsDir()) {
		return false
	}

	fname := strings.ToLower((file.Name()))

	// Todo: multiple suffixes, move suffix to config
	return strings.HasSuffix(fname, ".nef")
}
