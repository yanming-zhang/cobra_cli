package gitClient

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"golang.org/x/mod/semver"
)

func Clone(gitRef, repoLocalPath, repoURL string) (*git.Repository, error) {
	var refName plumbing.ReferenceName

	isSemVer := semver.IsValid(gitRef)
	if isSemVer {
		refName = plumbing.NewTagReferenceName(gitRef)
	} else {
		refName = plumbing.NewBranchReferenceName(gitRef)
	}

	repo, err := git.PlainClone(repoLocalPath, false, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: refName,
		SingleBranch:  true,
		Progress:      os.Stdout,
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}
