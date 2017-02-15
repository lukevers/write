package main

var (
	// Version is the version number that this package was built on. This
	// variable should be set by the linker.
	Version = "0.0"

	// Commit is the commit that this package was built on. This variable should
	// be set by the linker.
	Commit = "00000000"
)

func init() {
	Stdout.Printf(
		"Starting up Version %s (%s)",
		Version,
		Commit[:8],
	)
}
