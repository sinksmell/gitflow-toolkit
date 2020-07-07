package git

import (
	"fmt"
	"strings"

	"github.com/sinksmell/gitflow-toolkit/utils"
)

type CommitType string

type RepoType string

const Cmd = "git"

const (
	FEAT     CommitType = "feat"
	FIX      CommitType = "fix"
	DOCS     CommitType = "docs"
	STYLE    CommitType = "style"
	REFACTOR CommitType = "refactor"
	TEST     CommitType = "test"
	CHORE    CommitType = "chore"
	PERF     CommitType = "perf"
	HOTFIX   CommitType = "hotfix"
	EXIT     CommitType = "exit"
)

const CommitTpl = `{{ .Type }}({{ .Scope }}): {{ .Subject }}

{{ .Body }}

{{ .Footer }}

{{ .Sob }}
`

const CommitMessagePattern = `^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (.*)|^Merge\ branch(.*)|^Merge\ remote-tracking\ branch(.*)`

func CheckGitProject() {
	utils.MustExecNoOut(Cmd, "rev-parse", "--show-toplevel")
}

func CheckStagedFiles() bool {
	output := utils.MustExecRtOut(Cmd, "diff", "--cached", "--name-only")
	return strings.TrimSpace(output) != ""
}

func GetLastCommitInfo() *[]string {
	title := utils.MustExecRtOut(Cmd, "log", "-1", "--pretty=format:%s")
	desc := utils.MustExecRtOut(Cmd, "log", "-1", "--pretty=format:%b")
	return &[]string{title, desc}
}

func GetCurrentBranch() string {
	// 1.8+ git symbolic-ref --short HEAD
	return strings.TrimSpace(utils.MustExecRtOut(Cmd, "rev-parse", "--abbrev-ref", "HEAD"))
}

func Rebase(sourceBranch string, targetBranch string) {
	fmt.Println("checkout branch:", targetBranch)
	utils.MustExec(Cmd, "checkout", targetBranch)
	fmt.Println("pull origin branch:", targetBranch)
	utils.MustExec(Cmd, "pull", "origin", targetBranch)
	fmt.Println("checkout branch:", sourceBranch)
	utils.MustExec(Cmd, "checkout", sourceBranch)
	fmt.Println("exec git rebase")
	utils.MustExec(Cmd, "rebase", targetBranch)
	fmt.Println("push", sourceBranch, "to origin")
	utils.MustExec(Cmd, "push", "origin", sourceBranch)
}

func Checkout(prefix CommitType, branch string, remote string) {
	if prefix == FEAT {
		prefix = "feature" // 分支名使用完整单词
	}
	if len(remote) != 0 {
		utils.MustExec(Cmd, "checkout", "-b", string(prefix)+"/"+branch, "-t", remote)
		return
	}
	utils.MustExec(Cmd, "checkout", "-b", string(prefix)+"/"+branch)
}

func Push() {
	utils.MustExec(Cmd, "push", "origin", strings.TrimSpace(GetCurrentBranch()))
}
