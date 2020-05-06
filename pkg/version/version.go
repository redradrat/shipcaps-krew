package version

var (
	// gitCommit contains the git commit identifier.
	gitCommit string

	// gitTag contains the git tag or describe output.
	gitTag string
)

// GitCommit returns the value stamped into the binary at compile-time or a
// default "unknown" value.
func GitCommit() string {
	if gitCommit == "" {
		return "unknown"
	}
	return gitCommit
}

// GitTag returns the value stamped into the binary at compile-time or a
// default "unknown" value.
func GitTag() string {
	if gitTag == "" {
		return "unknown"
	}
	return gitTag
}
