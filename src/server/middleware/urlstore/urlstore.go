package urlstore

import (
	"sync"
	"time"
)

type URLSchema struct {
	LongURL string
	Timeout int64
}

type URLStore struct {
	sync.RWMutex
	urls map[string]*URLSchema
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]*URLSchema),
	}
}

func (store *URLStore) Save(us *URLSchema, shortURL string) {
	store.Lock()
	defer store.Unlock()
	store.urls[shortURL] = us
}

func (store *URLStore) Get(shortURL string) (string, bool) {
	store.RLock()
	urlschema, found := store.urls[shortURL]
	if !found {
		return "", false
	}
	store.RUnlock()
	if urlschema.Timeout < time.Now().Unix() {
		store.Delete(shortURL)
		return "", false
	}
	return urlschema.LongURL, found
}

func (store *URLStore) Delete(shortURL string) {
	store.Lock()
	defer store.Lock()
	delete(store.urls, shortURL)
}
