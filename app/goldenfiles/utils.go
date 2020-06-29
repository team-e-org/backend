package goldenfiles

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func updateGoldenFile(t *testing.T, path string, bytes []byte) {
	t.Log("update golden file")
	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}
}

func goldenFilePath(t *testing.T) string {
	return filepath.Join("testdata", t.Name()+".golden")
}

func readGoldenFile(t *testing.T) []byte {
	path := goldenFilePath(t)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading .golden: %s", err)
	}
	fmt.Printf("data: %+v\n", data)

	if len(data) == 0 {
		return nil //this is to simplify comparison with *bytes.Buffer in ResponseRecorder
	}
	return data
}

func UpdateAndOrRead(t *testing.T, bytes []byte) []byte {
	path := goldenFilePath(t)

	if *update {
		updateGoldenFile(t, path, bytes)
	}

	return readGoldenFile(t)
}
