package store

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type StoreDiskGzip[T Identifier] struct {
	dataDir string
}

func NewStoreDiskGzipCached[T Identifier](dataDir string) (*StoreCached[T], error) {
	disk, err := NewStoreDiskGzip[T](dataDir)
	if err != nil {
		return nil, err
	}
	return NewStoreCached(disk, NewStoreMemory[T]())
}

func NewStoreDiskGzip[T Identifier](dataDir string) (*StoreDiskGzip[T], error) {
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, fmt.Errorf("ERROR: ensure data dir '%s': %s\n", dataDir, err.Error())
	}

	return &StoreDiskGzip[T]{
		dataDir: dataDir,
	}, nil
}

func (f *StoreDiskGzip[T]) List(ctx context.Context) ([]*T, error) {
	entries, err := os.ReadDir(f.dataDir)
	if err != nil {
		return nil, fmt.Errorf("reading directory: %s", err.Error())
	}

	var result []*T
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".json.gz") {
			continue
		}

		filename := path.Join(f.dataDir, entry.Name())
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("error opening '%s': %s\n", filename, err.Error())
			continue
		}

		gr, err := gzip.NewReader(file)
		if err != nil {
			log.Printf("error creating gzip reader '%s': %s\n", filename, err.Error())
			file.Close()
			continue
		}

		var item *T
		if err := json.NewDecoder(gr).Decode(&item); err != nil {
			log.Printf("error decoding '%s': %s\n", filename, err.Error())
			gr.Close()
			file.Close()
			continue
		}
		gr.Close()
		file.Close()
		result = append(result, item)
	}

	return result, nil
}

func (f *StoreDiskGzip[T]) Put(ctx context.Context, item *T) error {
	id := (*item).GetId()
	targetFilename := path.Join(f.dataDir, id+".json.gz")

	tmpFile, err := os.CreateTemp(f.dataDir, fmt.Sprintf("tmp-%s-*.json.gz", id))
	if err != nil {
		return fmt.Errorf("creating temp file: %s", err.Error())
	}
	tmpName := tmpFile.Name()

	defer func() {
		tmpFile.Close()
		if err != nil {
			os.Remove(tmpName)
		}
	}()

	gw := gzip.NewWriter(tmpFile)
	e := json.NewEncoder(gw)
	e.SetIndent("", "    ")
	if err = e.Encode(item); err != nil {
		gw.Close()
		return fmt.Errorf("interim persistence %s: %s\n", tmpName, err.Error())
	}
	if err = gw.Close(); err != nil {
		return fmt.Errorf("closing gzip writer: %s", err.Error())
	}

	if err = tmpFile.Sync(); err != nil {
		return fmt.Errorf("syncing temp file: %s", err.Error())
	}

	tmpFile.Close()

	if err = os.Rename(tmpName, targetFilename); err != nil {
		return fmt.Errorf("renaming %s to %s: %s", tmpName, targetFilename, err.Error())
	}

	return nil
}

func (f *StoreDiskGzip[T]) Get(ctx context.Context, id string) (*T, error) {
	filename := path.Join(f.dataDir, id+".json.gz")

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("creating gzip reader for item '%s': %s", id, err.Error())
	}
	defer gr.Close()

	var item *T
	if err := json.NewDecoder(gr).Decode(&item); err != nil {
		return nil, fmt.Errorf("decoding item '%s': %s", id, err.Error())
	}

	return item, nil
}

func (f *StoreDiskGzip[T]) Delete(ctx context.Context, id string) error {
	filename := path.Join(f.dataDir, id+".json.gz")
	err := os.Remove(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("item '%s' persistence error: %s", id, err.Error())
	}

	return nil
}
