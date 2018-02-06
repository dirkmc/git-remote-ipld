package main

import (
	"bytes"
	"github.com/magik6k/git-remote-ipld/util"
	"io"
	"io/ioutil"
//	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCapabilities(t *testing.T) {
	tmpdir := setupTest(t)
	defer os.RemoveAll(tmpdir)

	// git clone ipld::20dae521ef399bcf95d4ddb3cefc0eeb49658d2a
	args := []string{"git-remote-ipld", "origin", "20dae521ef399bcf95d4ddb3cefc0eeb49658d2a"}
	// testCase(t, args, "capabilities", []string{"push", "fetch"})

	listExp := []string{
		"20dae521ef399bcf95d4ddb3cefc0eeb49658d2a refs/heads/master",
		"@refs/heads/master HEAD",
	}
	// listForPushExp := []string{
	// 	"@refs/heads/master HEAD",
	// }
	// testCase(t, args, "list", listExp)
	// testCase(t, args, "list for-push", listForPushExp)
	testCase(t, args, "fetch 20dae521ef399bcf95d4ddb3cefc0eeb49658d2a refs/heads/master\n", listExp)
}

func testCase(t *testing.T, args []string, input string, expected []string) {
	reader := strings.NewReader(input + "\n")
	var writer bytes.Buffer
	//logger := log.New(ioutil.Discard, "", 0)
	//err := Main(args, reader, &writer, logger)
	err := Main(args, reader, &writer, nil)
	if err != io.EOF {
		t.Fatal(err)
	}

	response := writer.String()
	exp := strings.Join(expected, "\n")
	if strings.TrimSpace(response) != exp {
		t.Fatalf("Args: %s\nInput:\n%s\nExpected:\n%s\nActual:\n%s\n", args, input, exp, response)
	}
}

func setupTest(t *testing.T) string {
	wd, _ := os.Getwd()
	src := filepath.Join(wd, "..", "..", "mock", "git-empty")
	// src := filepath.Join(wd, "..", "..", "mock", "git")
	si, err := os.Stat(src)
	if err != nil {
		t.Fatal(err)
	}
	if !si.IsDir() {
		t.Fatal("source is not a directory")
	}

	tmpdir, err := ioutil.TempDir("", "git-test")
	if err != nil {
		t.Fatal(err)
	}

	dst := filepath.Join(tmpdir, ".git")
	err = util.CopyDir(src, dst)
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("GIT_DIR", dst)
	return tmpdir
}
