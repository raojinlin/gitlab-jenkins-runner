package main

import (
	"os"
	"testing"
)

func TestGitlabGetMergeRequest(t *testing.T) {
	mr, err := getMergeRequest(os.Getenv("GITLAB_BASEURL"), os.Getenv("GITLAB_ACCESS_TOKEN"), "qingtian/web-app", 1257)
	if err != nil {
		panic(err)
	}

	if mr.IID != 1257 {
		t.Fatal("merge request iid error")
	}

	// mr.RebaseInProgress
}

func TestParseFromDescription(t *testing.T) {
	desc := "#title\n```env\nX=1\nY=2\nC=c=2\n```"
	params := parseParamsFromDesc(desc)
	if params["X"] != "1" {
		t.Fatal("X != 1")
	}

	if params["Y"] != "2" {
		t.Fatal("Y != 2")
	}

	if params["C"] != "c=2" {
		t.Fatal("C != c=2")
	}
}

func TestParseFromDescriptionEmpty(t *testing.T) {
	desc := "#title\n```bash\nX=1\nY=2\nC=c=2\n```"
	params := parseParamsFromDesc(desc)
	if len(params) != 0 {
		t.Fatal("params should empty")
	}
}
