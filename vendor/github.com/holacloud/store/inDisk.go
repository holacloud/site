package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type StoreDisk[T Identifier] struct {
	dataDir string
}

func NewStoreDiskCached[T Identifier](dataDir string) (*StoreCached[T], error) {
	// This is a helper to create a cached store with a disk backend
	// It is not required to use the store, but it is a convenience function to
	// create a cached store with a disk backend
	disk, err := NewStoreDisk[T](dataDir)
	if err != nil {
		return nil, err
	}
	return NewStoreCached(disk, NewStoreMemory[T]())
}

func NewStoreDisk[T Identifier](dataDir string) (*StoreDisk[T], error) {

	// ensure dir
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, fmt.Errorf("ERROR: ensure data dir '%s': %s\n", dataDir, err.Error())
	}

	return &StoreDisk[T]{
		dataDir: dataDir,
	}, nil
}

func (f *StoreDisk[T]) List(ctx context.Context) ([]*T, error) {
	// Read directory directly
	entries, err := os.ReadDir(f.dataDir)
	if err != nil {
		return nil, fmt.Errorf("reading directory: %s", err.Error())
	}

	var result []*T
	for _, entry := range entries {
		if entry.IsDir() || !strings.EqualFold(".json", path.Ext(entry.Name())) {
			continue
		}

		filename := path.Join(f.dataDir, entry.Name())
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("error opening '%s': %s\n", filename, err.Error())
			continue
		}

		var item *T
		if err := json.NewDecoder(file).Decode(&item); err != nil {
			log.Printf("error decoding '%s': %s\n", filename, err.Error())
			file.Close()
			continue
		}
		file.Close()
		result = append(result, item)
	}

	return result, nil
}

func (f *StoreDisk[T]) Put(ctx context.Context, item *T) error {
	id := (*item).GetId()
	targetFilename := path.Join(f.dataDir, id+".json")

	// 1. Create temp file in the same directory (ensures same filesystem for atomic rename)
	tmpFile, err := os.CreateTemp(f.dataDir, fmt.Sprintf("tmp-%s-*.json", id))
	if err != nil {
		return fmt.Errorf("creating temp file: %s", err.Error())
	}
	tmpName := tmpFile.Name()

	// Cleanup temp file in case of failure
	defer func() {
		tmpFile.Close()
		if err != nil {
			os.Remove(tmpName)
		}
	}()

	// 2. Write JSON
	e := json.NewEncoder(tmpFile)
	e.SetIndent("", "    ")
	if err = e.Encode(item); err != nil {
		return fmt.Errorf("interim persistence %s: %s\n", tmpName, err.Error())
	}

	// 3. Sync to disk (optional but safer)
	if err = tmpFile.Sync(); err != nil {
		return fmt.Errorf("syncing temp file: %s", err.Error())
	}

	// Close explicitly before rename (defer Close is too late for Windows, good practice generally)
	tmpFile.Close()

	// 4. Atomic Rename
	if err = os.Rename(tmpName, targetFilename); err != nil {
		return fmt.Errorf("renaming %s to %s: %s", tmpName, targetFilename, err.Error())
	}

	return nil
}

func (f *StoreDisk[T]) Get(ctx context.Context, id string) (*T, error) {
	filename := path.Join(f.dataDir, id+".json")

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, nil // Not found
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var item *T
	if err := json.NewDecoder(file).Decode(&item); err != nil {
		return nil, fmt.Errorf("decoding item '%s': %s", id, err.Error())
	}

	return item, nil
}

func (f *StoreDisk[T]) Delete(ctx context.Context, id string) error {
	filename := path.Join(f.dataDir, id+".json")
	err := os.Remove(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Already deleted or never existed, benign
		}
		return fmt.Errorf("item '%s' persistence error: %s", id, err.Error())
	}

	return nil
}
