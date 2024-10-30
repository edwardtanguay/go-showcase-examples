# go-showcase-examples

This is my main site for creating and storing working examples of Go concepts as I learn.

Also see my Tech Site [Skills page](https://tanguay-eu.vercel.app/skills)

## how to use

- execute example with Go
  - `cd 00100-hello`
  - `go run main.go`

- create executable file
  - Windows
    - `cd 00100-hello`
    - `go build -o hello.exe main.go`
    - in Windows terminal, `hello`
      - also in Git bash: `./hello`
  - Mac/Linux and Windows Bash
    - `cd 00100-hello`
    - `go build -o hello main.go`
    - in Windows terminal, `hello`

- run main.go with additional files
  - `cd 00210-foriloop`
  - `go run main.go tools.go`

- create executable with multiple files
  - `cd 00600-math`
  - `go build -o shownums main.go tools.go`
  - `sudo mv shownums /usr/local/bin` (for Windows, copy to folder that is included in the environment paths)

