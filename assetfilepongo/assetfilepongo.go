s4l1hpackage assetfilepongo

import (
	"github.com/akmyazilim/assetfile"
	"github.com/flosch/pongo2"
)

// AssetFile for asset
type AssetFile struct {
	manager *assetfile.AssetFile
}

// Add  File
func (afp *AssetFile) Add(file string, param ...interface{}) *pongo2.Value {
	afp.manager.Add(file, param...)
	return pongo2.AsValue("")
}

// AddIndex  File with index
func (afp *AssetFile) AddIndex(file string, index int, param ...interface{}) *pongo2.Value {
	afp.manager.AddIndex(file, index, param...)
	return pongo2.AsValue("")
}

// List return asset files
func (afp *AssetFile) List() []string {
	return afp.manager.List()
}

// New pongo2 asset file
func New() *AssetFile {
	return &AssetFile{
		manager: assetfile.New(),
	}
}
