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

- to create example directories
  - `cd 00900-cli-crex`
  - `go build -o crex main.go tools.go`
  - copy go to a PATH directory, e.g. `sudo mv crex /usr/local/bin`
  - in root of this project, `crex` to get help, or `crex 00999-advanced-for-loops`

## known issues

- couldn't get these to work: 
  - 00920, 00930

## best

- **01310-structMethods** - how to attach methods on structs
- **01500-for** - various ways to loop through slices
- **01600-api** - the simplest API, serves an array of strings
- **01610-api-text-fprintf** - simplest webs server, serves a HTML string


