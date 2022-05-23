package hw1

import "testing"

func TestReadExistFile(t *testing.T) {
	err, res := Read(".env")
	if err != nil {
		t.Errorf("Read(.env) failed: %v", err)
	}
	t.Logf("Read(.env) = %v", res)
}

func TestNotExistFile(t *testing.T) {
	err, res := Read("not_exist_file")
	if err == nil {
		t.Errorf("Read(not_exist_file) failed: %v", err)
	}
	t.Logf("Read(not_exist_file) = %v", res)
}
