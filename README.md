[![Go Report Card](https://goreportcard.com/badge/github.com/maneac/go-ultralight)](https://goreportcard.com/report/github.com/maneac/go-ultralight)
[![GoDoc](http://godoc.org/github.com/maneac/go-ultralight?status.svg)](http://godoc.org/github.com/maneac/go-ultralight)

### Please read the installation instructions below

Go-Ultralight provides unofficial bindings for the Ultralight UI library, endeavouring to match the original API as closely as possible. This project borrows heavily from the wonderful work of Raff found [here](https://github.com/raff/ultralight-go).

#### What is Ultralight?

Ultralight (https://ultralig.ht) is a HTML UI library, written in C++, that provides a performant and lightweight alternative to Electron, with JavaScript support.

# Installation

## Prerequisites

You must have a working CGo installation and have the GOPATH set.

## Automated - RECOMMENDED

1. Run:<br/><br/> `go get github.com/maneac/go-ultralight/go-ultralight` <br/><br/>to download the project and the installation and setup utility.
2. Navigate to your project's directory and execute:<br/><br/>`go-ultralight [OPTIONS]`<br/><br/>to automatically download the Ultralight SDK, and copy the necessary binaries for running your application. For more information on the available options, please read the utility's help (`go-ultralight --help`).
3. That's it! Now you're ready to Go-Ultralight!

## Manual

1. Run:<br><br>`go get github.com/maneac/go-ultralight`<br><br>to fetch the repository.

2. Download the Ultralight SDK for your system from https://ultralig.ht.

3. Locate your installation of Go-Ultralight (typically in $GOPATH/src/github.com/maneac/go-ultralight or $GOPATH/pkg/mod/github.com/maneac/go-ultralight), and extract the Ultralight SDK into an 'SDK' folder inside. Example structure:

   `github.com
     |-maneac
       |-go-ultralight
         |-go-ultralight
         |-examples
         |-SDK
       	  |-bin
       	    |-linux
       	      |-libAppCore.so
       	      |-libUltralight.so
       	      ...
            |-windows
              |-x64
                |-AppCore.dll
                ...
              |-x86
                |-AppCore.dll
                ...
       	  |-deps
       	  |-include
       	  ...`
   
4. Copy the binary files (*.dylib, *.dll, *.so) for your target system type from the 'SDK/bin' folder into your project directory. Example project directory:

   `exampleProject
     |-AppCore.dll
     |-main.go
     |-Ultralight.dll
     |-UltralightCore.dll
     |-WebCore.dll`

# Use

After installation, use the setup utility `go-ultralight` , or follow step 4 of the manual installation to copy the necessary binary files to your project directory. These files are required to run the compiled program.

Please build your project with `go build` instead of using `go run`, as run has been known to cause issues with the Browser example.

For examples, please see the 'examples' directory, which contains Go implementations of the sample projects provided with the Ultralight SDK.

# Known Issues

+ Automatic fetching of the SDK files on Linux may produce a certificate error.
+ Untested on MacOS

# To Do

- Implement the mouse and keyboard events
- Write tests
- Fix any memory leaks
