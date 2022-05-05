package util

import (
	"context"
	"errors"
	"os"
)

func CreateDirByTeacherDomain(ctx context.Context, path, teacherDomain string) error {
	var err error
	dirs := []string{"/homework", "/slide"}
	_path := path + "/" + teacherDomain

	if err = os.Mkdir(_path, 0700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}
	for _, v := range dirs {
		if err = os.Mkdir(_path+v, 0700); err != nil {
			if !errors.Is(err, os.ErrExist) {
				return err
			}
		}
	}
	return nil
}
