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
func (manager *AssetFile) Add(file string) {
	manager.Lock()
	defer manager.Unlock()
	if _, ok := manager.Files[file]; !ok {
		manager.Counter++
		manager.Files[file] = manager.Counter
	}
}

// AddWithParam  File
func (manager *AssetFile) AddWithParam(file string, param interface{}) {
	manager.Add(fmt.Sprintf(file, param))
}

// AddIndex  File
func (manager *AssetFile) AddIndex(file string, index int) {
	manager.Lock()
	defer manager.Unlock()
	manager.Files[file] = index

}

// List file
func (manager *AssetFile) List() []string {
	return SortedKeys(manager.Files)
}
