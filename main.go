package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bndr/gojenkins"
)

var (
	triggerJobBuild bool
	baseUrl         string
	params          string
	jobName         string
	jenkinsUser     string
	jenkinsToken    string
)

func waitJobStart(job *gojenkins.Job) {
	ctx := context.Background()
	for {
		isQueued, err := job.IsQueued(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if isQueued {
			fmt.Println("waiting... job", jobName, "queued")
		} else {
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func ParseParams(params string) map[string]string {
	labelArr := strings.Split(params, ",")
	result := make(map[string]string, 0)

	for _, label := range labelArr {
		param := strings.SplitN(label, "=", 2)
		if len(param) < 2 {
			continue
		}

		result[param[0]] = param[1]
	}

	return result
}

func main() {
	flag.StringVar(&jobName, "job", jobName, "Jenkins job name")
	flag.StringVar(&baseUrl, "base", baseUrl, "Jenkins base url")
	flag.StringVar(&params, "params", params, "Gitlab labels variable value")
	flag.StringVar(&jenkinsUser, "user", jenkinsUser, "Jenkins user")
	flag.StringVar(&jenkinsToken, "token", jenkinsUser, "Jenkins user token")
	flag.BoolVar(&triggerJobBuild, "build", triggerJobBuild, "Trigger jenkins job build")
	flag.Parse()

	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, baseUrl, jenkinsUser, jenkinsToken)
	_, err := jenkins.Init(ctx)
	if err != nil {
		fmt.Println("init jenkins error:", err.Error())
		os.Exit(1)
	}

	job, err := jenkins.GetJob(ctx, jobName)
	if err != nil {
		fmt.Printf("get job %s error:%s\n", jobName, err.Error())
		os.Exit(1)
	}

	waitJobStart(job)
	nextBuildNumber := job.GetDetails().NextBuildNumber
	var build *gojenkins.Build
	for {
		build, err = job.GetBuild(ctx, nextBuildNumber)
		if build != nil {
			break
		}

		if err != nil {
			if triggerJobBuild && strings.Contains(err.Error(), "404") {
				buildParams := ParseParams(params)
				fmt.Printf("build job %s with params: %+v\n", jobName, buildParams)
				_, err = job.InvokeSimple(ctx, buildParams)
				if err != nil {
					fmt.Println("invoke error", err.Error())
					os.Exit(2)
				}

				waitJobStart(job)
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		time.Sleep(1 * time.Second)
	}

	consoleOutput, err := build.GetConsoleOutputFromIndex(ctx, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for consoleOutput.HasMoreText {
		fmt.Print(consoleOutput.Content)
		consoleOutput, err = build.GetConsoleOutputFromIndex(ctx, consoleOutput.Offset)
		if err != nil {
			fmt.Println("get console output from index", consoleOutput.Offset, err.Error())
			return
		}
	}

	build, err = job.GetBuild(ctx, nextBuildNumber)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("build.Info().Result: %v\n", build.Info().Result)
	if build.Info().Result != "SUCCESS" {
		os.Exit(127)
	}
}
