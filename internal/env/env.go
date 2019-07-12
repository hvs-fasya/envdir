package env

import (
	"io/ioutil"
	"os"
)

//SetEnvs read all files from dir and set env values - key: file name value: file content
func SetEnvs(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	files, err := d.Readdir(-1)
	if err != nil {
		return err
	}
	for _, f := range files {
		envValue, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			return err
		}
		err = os.Setenv(f.Name(), string(envValue))
		if err != nil {
			return err
		}
	}
	return nil
}
