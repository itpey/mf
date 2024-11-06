<h1 align="center">mf</h1>

<p align="center">
mf (short for Make File) is a command-line tool written in Go that creates a new file with the specified filename if it doesn't already exist.
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/itpey/mf">
    <img src="https://pkg.go.dev/badge/github.com/itpey/mf.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/itpey/mf/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/itpey/mf" alt="license">
  </a>
</p>

# Installation

You can install mf in one of two ways:

## Install from GitHub Releases

You can download a precompiled binary from the [Releases](https://github.com/itpey/mf/releases) page:

1. Go to the [Releases](https://github.com/itpey/mf/releases) page.
2. Download the appropriate binary for your system.
3. Add the binary to your system's `PATH` for easy access.

## Install Using `go install`

Make sure you have [Go](https://go.dev) installed and configured on your system. Use go install to install mf:

```bash
go install github.com/itpey/mf@latest
```

Ensure that your `GOBIN` directory is in your `PATH` for the installed binary to be accessible globally.

# Usage

```bash
mf <filename>
```

# Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/itpey/mf/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

# License

mf is open-source software released under the MIT License. You can find a copy of the license in the [LICENSE](https://github.com/itpey/mf/blob/main/LICENSE) file.

# Author

mf was created by [itpey](https://github.com/itpey).
