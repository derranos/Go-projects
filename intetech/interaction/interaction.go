package interaction

import (
	"os"
	"os/exec"
	"path/filepath"
)

func SetLibs() error {
	err := exec.Command("pip", "install", "pandas").Run()
	if err != nil {
		return err
	}
	err = exec.Command("pip", "install", "pyarrow").Run()
	if err != nil {
		return err
	}
	return nil
}

func FindPath() (*string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(currentDir, "../python.py")
	return &path, err
}

func PyRun(path string, data1 []byte, data2 []byte) ([]byte, error) {
	cmd := exec.Command("python3", path, string(data1), string(data2))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}
