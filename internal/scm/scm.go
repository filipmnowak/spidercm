package scm

import (
	"os/exec"
	"sync"
)

const (
	git       = "/usr/bin/git"
	mercurial = "/usr/bin/hg"
	fossil    = "/opt/fossil-2.23/bin/fossil"
)

type SCMAction struct {
	Command string
	Args    []string
}

type SCMResult struct {
	Logs []byte
	Err  error
}

type SCM struct {
	Action  string
	ToRun   []SCMAction
	paths   []string
	message string
	Results []SCMResult
}

func (scm *SCM) setCommands(action string) {
	actions := map[string][]SCMAction{
		"init": {
			{Command: git, Args: []string{"init", "."}},
			{Command: mercurial, Args: []string{"init", "."}},
			//{Command: fossil, Args: []string{"init", ".fossil.sqlite3"}},
			{Command: "/usr/bin/sh", Args: []string{"-c", "fossil init .fossil.sqlite3 && fossil open --force .fossil.sqlite3"}},
		},
		"add": {
			{Command: git, Args: append([]string{"add"}, scm.paths...)},
			{Command: mercurial, Args: append([]string{"add"}, scm.paths...)},
			{Command: fossil, Args: append([]string{"add"}, scm.paths...)},
		},
		"commit": {
			{Command: git, Args: []string{"commit", "-m", scm.message}},
			{Command: mercurial, Args: []string{"commit", "-m", scm.message}},
			{Command: fossil, Args: []string{"commit", "-m", scm.message}},
		},
	}
	scm.ToRun = actions[action]
}

func (scm *SCM) runAction(action string) error {
	scm.setCommands(action)
	var channels []<-chan SCMResult
	for _, a := range scm.ToRun {
		channels = append(channels, gorountineWrapper(a))
	}
	for ch := range merge(channels) {
		scm.Results = append(scm.Results, ch)
	}
	return nil
}

func gorountineWrapper(a SCMAction) <-chan SCMResult {
	ch := make(chan SCMResult)
	go func() {
		cmd := exec.Command(a.Command, a.Args...)
		output, err := cmd.CombinedOutput()
		ch <- SCMResult{
			Logs: output,
			Err:  err,
		}
		close(ch)
	}()
	return ch
}

func (scm SCM) Init() error {
	return scm.runAction("init")
}

func (scm *SCM) Add(paths []string) error {
	scm.paths = paths
	return scm.runAction("add")
}

func (scm *SCM) Commit(message string) error {
	scm.message = message
	return scm.runAction("commit")
}

// https://go.dev/blog/pipelines
func merge(chs []<-chan SCMResult) <-chan SCMResult {
	var wg sync.WaitGroup
	out := make(chan SCMResult)

	output := func(ch <-chan SCMResult) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
