package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f := "./from/file"
	t := "./to/file"

	fmt.Print(MoveFile(f, t))
}

func MoveFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(dst)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	si, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Stat error: %s", err)
	}

	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}
