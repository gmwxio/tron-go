package lsp

import (
	"errors"
	"sync"

	"github.com/wxio/tron-go/adl"
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
func (fc *filecache) delete(uri string) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	delete(fc.txt, uri)
}

type astcache struct {
	mutex sync.RWMutex
	txt   map[string]*adl.Module
}

func newAstCache() astcache {
	return astcache{
		txt: make(map[string]*adl.Module),
	}
}

func (fc *astcache) put(uri string, txt adl.Module) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	fc.txt[uri] = &txt
}

func (fc *astcache) get(uri string) (*adl.Module, error) {
	fc.mutex.RLock()
	defer fc.mutex.RUnlock()
	txt, ex := fc.txt[uri]
	if !ex {
		return nil, errors.New("no txt for " + uri)
	}
	return txt, nil
}
func (fc *astcache) delete(uri string) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	delete(fc.txt, uri)
}
