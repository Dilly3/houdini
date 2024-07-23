package github

import (
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/pkg/github"
	"strconv"
	"strings"
)

func splitID(url string) string {
	split := strings.Split(url, "commits/")
	return split[len(split)-1]
}
func mapToCommitsInfo(commits []github.CommitResponse, repoName string) []model.CommitInfo {
	var commitsInfo []model.CommitInfo
	for i := 0; i < len(commits); i++ {
		commit := commits[i]
		commitInfo := mapToCommitInfo(&commit, repoName)
		commitInfo.ID = splitID(commit.URL)
		commitsInfo = append(commitsInfo, commitInfo)
	}
	return commitsInfo
}
func mapToCommitInfo(commit *github.CommitResponse, repoName string) model.CommitInfo {
	id := splitID(commit.URL)
	return model.CommitInfo{
		ID:          id,
		Message:     commit.Message,
		AuthorName:  commit.Author.Name,
		AuthorEmail: commit.Author.Email,
		Date:        commit.Author.Date,
		URL:         commit.URL,
		RepoName:    repoName,
	}
}

func mapRepoResponse(repo *github.RepoResponse) model.RepoInfo {
	return model.RepoInfo{
		ID:          repo.ID,
		Name:        repo.Name,
		CreatedAt:   repo.CreatedAt,
		UpdatedAt:   repo.UpdatedAt,
		URL:         repo.URL,
		Description: repo.Description,
		Language:    repo.Language,
		Forks:       repo.Forks,
		Stars:       repo.Stars,
		OpenIssues:  repo.OpenIssues,
	}
}
func stringToInteger(s string) int {
	if s == "" {
		return 0
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}
