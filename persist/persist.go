package persist

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

func Save(path string, v interface{}) error {

	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var Marshall = func(v interface{}) (io.Reader, error) {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), nil
	}
	r, err := Marshall(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var Unmarshall = func(r io.Reader, v interface{}) error {
		return json.NewDecoder(r).Decode(v)
	}
	return Unmarshall(f, v)

}
