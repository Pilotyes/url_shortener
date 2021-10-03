package version

const (
	programName = "url_shortener"
)

var (
	gitTag, gitCommit, gitBranch string
)

// Version returns the name and version information of this program.
func Version() (string, string, string, string) { // nolint
	return programName, gitTag, gitCommit, gitBranch
}
