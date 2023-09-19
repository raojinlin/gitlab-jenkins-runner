package main

import (
	"strings"

	"github.com/xanzy/go-gitlab"
)

func parseParamsFromDesc(str string) map[string]string {
	params := make(map[string]string, 0)
	inCodeBlock := false
	for _, v := range strings.Split(str, "\n") {
		if strings.Contains(v, "```env") {
			inCodeBlock = true
			continue
		}

		if !inCodeBlock {
			continue
		}

		for param, val := range ParseParams(v) {
			params[param] = val
		}

		if strings.Contains(v, "```") {
			inCodeBlock = false
		}
	}

	return params
}

func getMergeRequest(baseUrl, token, pid string, mergeRequest int) (*gitlab.MergeRequest, error) {
	gitlabClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return nil, err
	}

	mr, _, err := gitlabClient.MergeRequests.GetMergeRequest(pid, mergeRequest, &gitlab.GetMergeRequestsOptions{})
	return mr, err
}
