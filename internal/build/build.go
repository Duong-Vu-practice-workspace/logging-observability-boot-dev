package build

// Set at build time via -ldflags "-X boot.dev/linko/internal/build.GitSHA=..."
var (
	GitSHA    = "unknown"
	BuildTime = "unknown"
)
