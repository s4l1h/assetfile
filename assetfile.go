package assetfile

import (
	"fmt"
	"sync"
)

//New assetfile
func New() *AssetFile {
	return &AssetFile{
		Files:   make(map[string]int),
		Counter: 20,
	}
}

// AssetFile manage files
type AssetFile struct {
	Files   map[string]int
	Counter int
	sync.RWMutex
}

// Add  File
func (manager *AssetFile) Add(file string, param ...interface{}) {
	manager.Lock()
	defer manager.Unlock()
	if param != nil {
		file = fmt.Sprintf(file, param...)
	}
	if _, ok := manager.Files[file]; !ok {
		manager.Counter++
		manager.Files[file] = manager.Counter
	}
}

// AddIndex  File
func (manager *AssetFile) AddIndex(file string, index int, param ...interface{}) {
	manager.Lock()
	defer manager.Unlock()
	if param != nil {
		file = fmt.Sprintf(file, param...)
	}
	manager.Files[file] = index

}

// List return asset files
func (manager *AssetFile) List() []string {
	return SortedKeys(manager.Files)
}
