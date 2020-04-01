package cmd

import (
	"github.com/mritd/gitflow-toolkit/git"
	"github.com/spf13/cobra"
)

// handleCheckOut wrap check cmd
func handleCheckOut(cmd *cobra.Command, ct git.CommitType, args []string) {
	switch len(args) {
	case 1:
		git.Checkout(ct, args[0], "")
	case 2:
		git.Checkout(ct, args[0], args[1])
	default:
		_ = cmd.Help()
	}
}

// Create feature branch
func NewFeatBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "feat BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create feature branch",
		Long: `
Create a branch with a prefix of feature.`,
		Aliases: []string{"git-feat"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.FEAT, args)
		},
	}
}

// Create fix branch
func NewFixBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "fix BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create fix branch",
		Long: `
Create a branch with a prefix of fix.`,
		Aliases: []string{"git-fix"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.FIX, args)
		},
	}
}

// Create hotfix branch
func NewHotFixBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "hotfix BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create hotfix branch",
		Long: `
Create a branch with a prefix of hotfix.`,
		Aliases: []string{"git-hotfix"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.HOTFIX, args)
		},
	}
}

// Create perf branch
func NewPerfBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "perf BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create perf branch",
		Long: `
Create a branch with a prefix of perf.`,
		Aliases: []string{"git-perf"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.PERF, args)
		},
	}
}

// Create refactor branch
func NewRefactorBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "refactor BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create refactor branch",
		Long: `
Create a branch with a prefix of refactor.`,
		Aliases: []string{"git-refactor"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.REFACTOR, args)
		},
	}
}

// Create test branch
func NewTestBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "test BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create test branch",
		Long: `
Create a branch with a prefix of test.`,
		Aliases: []string{"git-test"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.TEST, args)
		},
	}
}

// Create chore branch
func NewChoreBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "chore BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create chore branch",
		Long: `
Create a branch with a prefix of chore.`,
		Aliases: []string{"git-chore"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.CHORE, args)
		},
	}
}

// Create style branch
func NewStyleBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "style BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create style branch",
		Long: `
Create a branch with a prefix of style.`,
		Aliases: []string{"git-style"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.STYLE, args)
		},
	}
}

// Create docs branch
func NewDocsBranch() *cobra.Command {
	return &cobra.Command{
		Use:   "docs BRANCH_NAME [TRACK_BRANCH]",
		Short: "Create docs branch",
		Long: `
Create a branch with a prefix of docs.`,
		Aliases: []string{"git-docs"},
		Run: func(cmd *cobra.Command, args []string) {
			handleCheckOut(cmd, git.DOCS, args)
		},
	}
}
