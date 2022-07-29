package utils

var version string = "v0.0.3"

// SetVersion sets the version number for use later in the version command and for request headers.
func SetVersion(v string) {
	version = v
}

// GetVersion retrieves the version number for use later in the version command and for request headers.
func GetVersion() string {
	return version
}
