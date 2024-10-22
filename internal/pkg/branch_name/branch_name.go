package branchname

import (
	"strings"
)

func PrepareBranchName(branchName string) string {
	return prepareForSubdomain(prepareForDir(branchName))
}

func prepareForSubdomain(s string) string {
	// _~:/?#[]@!$&'()*+,;= are allowed url characters (RFC3986), but are not suitable for subdomain
	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-."

	var result strings.Builder
	for _, c := range s {
		if strings.ContainsRune(allowed, c) {
			result.WriteRune(c)
		}
	}

	return result.String()
}

func prepareForDir(s string) string {
	// Slashes are not allowed in Unix filenames
	return strings.ReplaceAll(s, "/", "-")
}
