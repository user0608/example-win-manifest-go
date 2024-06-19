# Example of Creating an Administrator-Privileged Executable in Windows With Golang

This is a simple example demonstrating how to create an executable in Go that requires administrator privileges on Windows systems.

## Required Tools

For this example, the following tools are used:

- **rsrc**: Tool for embedding .ico and manifest resources in Go programs for Windows. More information and the tool can be found at [akavel/rsrc](https://github.com/akavel/rsrc).
- **make**: Used for automating the build process. You can download it directly from [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm) or install it using Chocolatey with `choco install make`.

## Steps to Compile and Build

1. **Generate Manifest File (.syso)**:
   ```bash
   make manifest
   make buidl # go build should link automatically all *.syso file
