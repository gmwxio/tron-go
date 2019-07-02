package lsp

import (
	"errors"
	"sync"
)

type filecache struct {
	mutex sync.RWMutex
	txt   map[string]string
}

func newCache() filecache {
	return filecache{
		txt: make(map[string]string),
	}
}

func (fc *filecache) put(uri string, txt string) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	fc.txt[uri] = txt
}

func (fc *filecache) get(uri string) (string, error) {
	fc.mutex.RLock()
	defer fc.mutex.RUnlock()
	txt, ex := fc.txt[uri]
	if !ex {
		return "", errors.New("no txt for " + uri)
	}
	return txt, nil
}
