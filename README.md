# calcy-lib

The library/module part of a study in polymorphic deployment of a package across multiple UIs.

See github.com/mdwhatcott/calcy-apps for the other part.

-----

## Task 0

Define a struct called `Addition` which implements the following interface: `type Calculator { Calculate(a, b int) int }`

## Task 1

Build a command-line program that 1) receives two integers as arguments, 2) passes them to your `Addition` calculator, and 3) prints the resulting integer to the screen. Call `panic` on any errors encountered in the program.

## Task 2

Ensure that the `Addition` calculator is separated from the `main` function by putting them in different packages/folders (`package main` and `package calc`).

## Task 3

You've already done some manual testing of the `Addition` calculator via your Command Line Interface (CLI). Now, provide a quick, sure, and repeatable proof (ie. a few automated unit tests in `package calc`) which also prove the calculator works correctly

## Task 4

Extract a `Handler` from the code in the `main` function. Upon instantiation this `Handler` will receive 1) a reference to `os.Stdout` as an `io.Writer`, and 2) an instance of a `Calculator`, which will both be assigned as struct fields. It will provide a `Handle` method to receive the command line arguments as a `[]string`. The `Handle` method will then 1) parse the arguments as integers, 2) pass them to the calculator, and 3) print the result to the output writer. Any error encountered in that process should be returned back to the caller (`main`) to be logged, also resulting in a non-zero exit code.

## Task 5

Prove that the handler works as advertised with automated unit tests (this might be a bit tricky..)

## Task 6

Create a new repository/module called `github.com/jcrob2/calc-apps`. This will now become the new home of your `main` function and the `Handler`. Create two folders at the root of the repository, one called `main`, and the other called `handlers`. Move the handler and its tests (created during Tasks 2 & 3) into `handlers/cli.go` and `handlers/cli_test.go`. Move the `main` function into `main/calc-cli/main.go`. Add `github.com/jcrob2/calc-lib` as a dependency of `github.com/jcrob2/calc-apps`. Get everything compiling and working again.

## Task 7

Let's code together to release a formal version of the calc-lib module.

## Task 8

Create several more calculator implementations of your own choosing. Provide proof that they work as intended. Release a new version of the library to the world.

## Task 9

Incorporate the newly released version of `calc-lib` in `calc-apps` and extend the CLI program in `calc-apps` to incorporate these new operations using a command-line flag. Here's what you're working towards:

```
$ calc-cli -op '+' 3 4
7

$ calc-cli -op '-' 7 5
2

go run main.go -op 'asdf' 1 2
10:11:07 main.go:20: unsupported operand: asdf
exit status 1
```

## Task 10

Create a new command-line program in `calc-apps` that will accept a CSV text file via STDIN:

```
1,+,2
2,-,1
NaN,+,2
1,+,NaN
1,nop,2
3,*,4
20,/,10
```

...and will produce the following CSV text file via STDOUT:

```
1,+,2,3
2,-,1,1
3,*,4,12
20,/,10,2

```

...and prove that it works with automated tests.

## Task 11

Create a new command-line program in `calc-apps` that will function as an HTTP server, handling the following request/response scenarios:

Add:

```
$ curl -v "http://localhost:8080/add?a=3&b=4"
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /add?a=3&b=4 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Mon, 27 Feb 2023 17:34:10 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8
< 
7
```

Multiply:
```
$ curl -v "http://localhost:8080/mul?a=3&b=4"
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /mul?a=3&b=4 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Mon, 27 Feb 2023 17:34:37 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
< 
12
```

## Task 12

Start over and implement all of the above from scratch, keeping everything well-tested and as clean as possible.


---

Other possibilities:

- Make the Calculator interface variadic: `type Calculator interface { Calculate(...int) int }` (Is this a breaking change?)
- Convert the CSV handler into a client of the HTTP handler
- Make the CSV handler concurrent, sending up to N requests at a time.
- Rewrite the tests using gunit.
- Write integration tests that compile and invoke each executable.
- Rewrite a few components via TDD.