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

## best

- **01310-structMethods** - how to attach methods on structs
- **01500-for** - various ways to loop through slices
- **01600-api** - the simplest API, serves an array of strings
- **01610-api-text-fprintf** - simplest webs server, serves a HTML string
- **01620-api-handler** - simple API with routes that call functions
- **01700-methods** - attaching methods to structs, pseudo OOP in Go
- **01800-functions** - basics of functions, e.g. return values, multiple parameters
- **02000-fetch-and-create-files** - fetchs an array of items via API and creates a HTML file that displays them
- **02100-url** - function that takes a URL and returns info from it, base url, url variables
- **02150-pointers** - simple example of pointers 
- **02160-errorHandling-structs-pointers** - good example of error handling and using pointers instead of values
- **02180-anonymousFunctions** - simple example of a function which accepts either an integer or a float
- **02180-higher-order-functions** - good example of a useful pattern of a function that accepts a function as a parameter
- **02190-defer** - uses defer to postpone closing file until function finishes
- **02200-interfaces** - interface to make sure that structs circle and rectangle have the method getArea()
- **02220-empty-interface** - example of empty interface which functions as TypeScript's "any"
- **02250-concurrency** - simple example of using concurrency to fetch data from 4 API's without having to wait for each one to finish before starting the next
- **02290-web-templates** - good example of using EJS-like templates for websites
