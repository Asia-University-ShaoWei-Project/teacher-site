package migrate

import "os"

func SetupDir() error {
	var err error
	dirs := []string{"homework", "slide"}
	if err = os.Mkdir("static/doc", 0700); err != nil {
		return err
	}
	if err = os.Mkdir("static/doc/rikki", 0700); err != nil {
		return err
	}
	for _, v := range dirs {
		if err = os.Mkdir("static/doc/rikki/"+v, 0700); err != nil {
			return err
		}
	}
	return nil
}
