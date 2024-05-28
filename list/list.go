package list

import (
	"fmt"
	"sort"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/hashicorp/go-version"
)

// GetLatestVersion Gets the latest version of the specified image.
// This applies to public repositories.
func GetLatestVersion(img string) (string, error) {
	repo, err := name.NewRepository(img)
	if err != nil {
		return "", fmt.Errorf("name.NewRepository error: %w", err)
	}

	list, err := remote.List(repo)
	if err != nil {
		return "", fmt.Errorf("remote.List error: %w", err)
	}

	versions := make([]*version.Version, 0)
	for _, v := range list {
		ver, err := version.NewVersion(v)
		if err != nil {
			continue
		}

		versions = append(versions, ver)
	}

	vc := version.Collection(versions)
	sort.Sort(vc)

	return vc[len(vc)-1].String(), nil
}
