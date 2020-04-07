package profile

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf(`filename "%s" is not exist`, filename)
		}

		return "", err
	}

	if data == nil {
		return "", nil
	}

	return string(data), nil
}

// 文件名列表
func ReadDir(directory string) ([]string, error) {
	if err := IsDirectory(directory); err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	if files == nil {
		return nil, nil
	}

	var r []string
	for _, f := range files {
		if f == nil || f.IsDir() {
			continue
		}

		r = append(r, f.Name())
	}

	return r, nil
}

func IsDirectory(directory string) error {
	stat, err := os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(`directory "%s" is not exist`, directory)
		}

		return err
	}

	if stat == nil || !stat.IsDir() {
		return fmt.Errorf(`directory "%s" must be a directory`, directory)
	}

	return nil
}
