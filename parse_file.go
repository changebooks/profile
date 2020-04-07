package profile

import "errors"

func ParseFile(filename string) (map[string]string, error) {
	if filename == "" {
		return nil, errors.New("filename can't be empty")
	}

	s, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(s), nil
}
