package endtoend

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	wasmPath := filepath.Join("..", "..", "bin", "sqlc-gen-hello.wasm")
	if _, err := os.Stat(wasmPath); err != nil {
		t.Fatalf("sqlc-gen-hello.wasm not found: %s", err)
	}

	wasmModule, err := os.ReadFile(wasmPath)
	if err != nil {
		t.Fatal(err)
	}

	shaSum := sha256.Sum256(wasmModule)
	sha256 := fmt.Sprintf("%x", shaSum)

	sqlcPath, err := exec.LookPath("sqlc")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range FindTests(t, "testdata") {
		dir := dir
		t.Run(dir, func(t *testing.T) {
			yamlPath := filepath.Join(dir, "sqlc.yaml")
			yaml, err := os.ReadFile(yamlPath)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Contains(yaml, []byte(sha256)) {
				yaml = bytes.Replace(yaml, []byte(`sha256: "`+sha256+`"`), []byte(fmt.Sprintf(`sha256: "%s"`, sha256)), 1)
				if err := os.WriteFile(yamlPath, yaml, 0644); err != nil {
					t.Fatal(err)
				}
			}

			want := expectedOutput(t, dir)
			cmd := exec.Command(sqlcPath, "diff")
			cmd.Dir = dir
			got, err := cmd.CombinedOutput()
			if diff := cmp.Diff(string(want), string(got)); diff != "" {
				t.Errorf("sqlc diff mismatch (-want +got):\n%s", diff)
			}
			if len(want) == 0 && err != nil {
				t.Error(err)
				t.Log(want)
				t.Log(got)
			}
		})
	}
}

func FindTests(t *testing.T, root string) []string {
	t.Helper()
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "sqlc.yaml" {
			dirs = append(dirs, filepath.Dir(path))
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return dirs
}

func expectedOutput(t *testing.T, dir string) []byte {
	t.Helper()
	stderrPath := filepath.Join(dir, "stderr.txt")
	if _, err := os.Stat(stderrPath); err != nil {
		if os.IsNotExist(err) {
			return []byte{}
		}
		t.Fatal(err)
	}
	output, err := os.ReadFile(stderrPath)
	if err != nil {
		t.Fatal(err)
	}
	return output
}
