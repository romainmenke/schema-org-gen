package typemap

import "os"

func NewTypeMapCacheFile() (*os.File, error) {
	err := os.RemoveAll("./tm-cache")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll("./tm-cache", 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.Create("./tm-cache/tm-cache.json")
	if err != nil {
		return nil, err
	}

	return f, nil
}

func NewOrExistingTypeMapCacheFile() (*os.File, error) {
	err := os.MkdirAll("./tm-cache", 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.Open("./tm-cache/tm-cache.json")
	if os.IsNotExist(err) {
		f, err = os.Create("./tm-cache/tm-cache.json")
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}
