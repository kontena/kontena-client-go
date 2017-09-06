package cli

// Options represents the command options that can be passed to modify the
// output.
type Options struct {
	Debug       bool
	Verbose     bool
	Quiet       bool
	Token       string
	SSLCertPath string
}
