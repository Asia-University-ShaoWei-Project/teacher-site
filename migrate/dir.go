package migrate

import (
	"errors"
	"os"
)

func SetupDir() error {
	var err error
	dirs := []string{"homework", "slide"}
	if err = os.Mkdir("static/doc", 0700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}
	if err = os.Mkdir("static/doc/rikki", 0700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}
	for _, v := range dirs {
		if err = os.Mkdir("static/doc/rikki/"+v, 0700); err != nil {
			if !errors.Is(err, os.ErrExist) {
				return err
			}
		}
	}
	return nil
}
