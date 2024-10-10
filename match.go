package snapshot

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iancoleman/strcase"
)

func Match(t *testing.T, renderer Renderer, value any) {
	t.Helper()

	current := &bytes.Buffer{}
	if err := renderer.Render(current, value); err != nil {
		t.Fatalf("snapshot cannot render: %v", err)
	}

	snapF := filename(t, snapshotExt+renderer.Ext())
	if !exists(t, snapF) {
		write(t, snapF, current)
		t.Fatalf("snapshot: created new snapshot: %s", snapF)
	}

	existing := &bytes.Buffer{}
	read(t, snapF, existing)

	failedF := filename(t, failedExt+renderer.Ext())
	if !bytes.Equal(existing.Bytes(), current.Bytes()) {
		write(t, failedF, existing)
		t.Fatalf("snapshot: does not match: %s", failedF)
	}

	if exists(t, failedF) {
		if err := os.Remove(failedF); err != nil {
			t.Fatalf("snapshot: cannot remove failed %s: %v", failedF, err)
		}
	}
}

type Renderer interface {
	Render(w io.Writer, value any) error
	Ext() string
}

const (
	snapshotExt = ".snap"
	failedExt   = ".failed"
)

func exists(t *testing.T, name string) bool {
	t.Helper()

	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		t.Fatalf("snapshot: cannot stat %s: %v", name, err)
	}
	return true
}

func read(t *testing.T, name string, w io.Writer) {
	t.Helper()

	f, err := os.Open(name)
	if err != nil {
		t.Fatalf("snapshot: cannot open file %s: %v", name, err)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		t.Fatalf("snapshot: cannot read file %s: %v", name, err)
	}
}

func write(t *testing.T, name string, r io.Reader) {
	t.Helper()

	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("snapshot: cannot create directory %s: %v", dir, err)
	}

	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("snapshot: cannot create file %s: %v", name, err)
	}
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		t.Fatalf("snapshot: cannot write file %s: %v", name, err)
	}
}

func filename(t *testing.T, ext string) string {
	t.Helper()

	name, _ := strings.CutPrefix(t.Name(), "Test")
	return filepath.Join(
		"testdata",
		strcase.ToSnake(name)+ext,
	)
}
