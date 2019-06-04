package safefile

import (
	"io"
	"os"
)

func WriteReader(filename string, reader io.Reader, perm os.FileMode) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, perm)
	if err != nil {
		return
	}

	// Handle resource panic as well
	defer func() {
		if p := recover(); p != nil {
			os.Remove(f.Name())
			panic(p)
		}

		if err != nil {
			os.Remove(f.Name())
		}
	}()

	_, err = io.Copy(f, reader)
	if err != nil {
		return
	}

	err = f.Sync()

	// Close the file regardless
	if e := f.Close(); e == nil {
		err = e
	}

	return
}
