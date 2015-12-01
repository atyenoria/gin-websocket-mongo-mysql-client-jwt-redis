package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func Double(stdin io.Reader, stdout io.Writer) error {
	buf, err := ioutil.ReadAll(stdin)
	if err != nil {
		return err
	}

	stdout.Write(buf)
	stdout.Write(buf)

	return nil
}

func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("foo\n")
	stdout := new(bytes.Buffer)

	err := Double(stdin, stdout)
	if err != nil {
		t.Fatal("failed to call Double(): %s", err)
	}

	expected := []byte("foo\nfoo\n")

	if bytes.Compare(expected, stdout.Bytes()) != 0 {
		t.Fatal("not matched")
	}
}
