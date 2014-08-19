package binarydist

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"
)

var diffT = []struct {
	old *os.File
	new *os.File
}{
	{
		old: mustWriteRandFile("test.old", 1e3),
		new: mustWriteRandFile("test.new", 1e3),
	},
	{
		old: mustOpen("testdata/sample.old"),
		new: mustOpen("testdata/sample.new"),
	},
}

func TestDiff(t *testing.T) {
	for _, s := range diffT {
		got, err := ioutil.TempFile("/tmp", "bspatch.")
		if err != nil {
			panic(err)
		}
		os.Remove(got.Name())

		exp, err := ioutil.TempFile("/tmp", "bspatch.")
		if err != nil {
			panic(err)
		}

		start := time.Now()
		cmd := exec.Command("bsdiff", s.old.Name(), s.new.Name(), exp.Name())
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		os.Remove(exp.Name())
		if err != nil {
			panic(err)
		}
		fmt.Println("bsdiff cost: %s", time.Now().Sub(start))

		start = time.Now()
		err = Diff(s.old, s.new, got)
		if err != nil {
			t.Fatal("err", err)
		}
		fmt.Println("binarydist cost: %s", time.Now().Sub(start))

		_, err = got.Seek(0, 0)
		if err != nil {
			panic(err)
		}
		gotBuf := mustReadAll(got)
		expBuf := mustReadAll(exp)

		if !bytes.Equal(gotBuf, expBuf) {
			t.Fail()
			t.Logf("diff %s %s", s.old.Name(), s.new.Name())
			t.Logf("%s: len(got) = %d", got.Name(), len(gotBuf))
			t.Logf("%s: len(exp) = %d", exp.Name(), len(expBuf))
			i := matchlen(gotBuf, expBuf)
			t.Logf("produced different output at pos %d; %d != %d", i, gotBuf[i], expBuf[i])
		}
	}
}
