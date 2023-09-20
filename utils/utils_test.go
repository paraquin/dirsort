package utils

import (
	"testing"
)

func TestExt(t *testing.T) {
	testCases := []struct {
		desc   string
		expect string
		got    string
	}{
		{
			desc:   "no ext",
			expect: "",
			got:    Ext("file_without_ext"),
		},
		{
			desc:   "has ext",
			expect: "txt",
			got:    Ext("file_with_ext.txt"),
		},
		{
			desc:   "double ext",
			expect: "gz",
			got:    Ext("archive.tar.gz"),
		},
		{
			desc:   "hidden file without ext",
			expect: "",
			got:    Ext(".gitignore"),
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
