# Deputy Task

## About the solution
There are there implementation of the task
1. Trivial linear
2. Maps
3. Sorted Array with binary search

Trivial linear does not have overhead, but slows down fast with the data growing. Works well with limited roles, as 5 in the example

Maps have a lot of overhead, which makes them slow on modest data sets, but the fastest solution with the data growing

Binary search lays somewhere in between, using some overhead to sort data but having log(n) lookups

## Things to improve
* handle two missing error
* implement interfaces if suitable for business domain
* adjust the code style to the company standards (or at least run go fmt)

## How to build
`go build main.go`

## How to run
The code works with command line arguments, suhc as
`./main 1`

## How to test
`go test ./...`
