package semversion

import "golang.org/x/mod/semver"

func CompareSemVersionStrings(version1, version2 string) int {
	return semver.Compare(version1, version2)
}
