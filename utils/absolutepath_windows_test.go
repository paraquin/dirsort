//go:build windows

package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	currentDir, _ := os.Getwd()
	currentUser, _ := user.Current()
	testCases := []struct {
		desc   string
		expect string
		got    string
	}{
		{
			desc:   "already absolute path",
			expect: `C:\users\user\testdir`,
			got:    AbsolutePath(`C:\users\user\testdir`),
		},
		{
			desc:   "user home directory symbol",
			expect: filepath.Join(currentUser.HomeDir, "Documents"),
			got:    AbsolutePath(`~\Documents`),
		},
		{
			desc:   "relative path",
			expect: filepath.Join(currentDir, "testdir"),
			got:    AbsolutePath("testdir"),
		},
		{
			desc:   "with double dot",
			expect: filepath.Join(currentDir, "testdir"),
			got:    AbsolutePath(`dir\..\testdir`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.expect != tC.got {
				t.Errorf("test %q: epected %q, but got %q", tC.desc, tC.expect, tC.got)
			}
		})
	}
}
