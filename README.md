expfunc
=======
expfunc is simple cli tool for golang  
export function name from go source file


Usage:
------
output function name:  
`expfunc -file /path/to/src.go`

with Test for testing  
`expfunc -test -file /path/to/src.go`

export only  
`expfunc -exported -file /path/to/src.go`

show help  
`expfunc --help`


Flags:
------
| Flags | Description | type(default) |
| :---- | :---------- | :------ |
| -f -file  | specify source file | string("") |
| -e -exported  | only exported function | bool(false) |
| -t -test  | with Test for testing package | bool(false) |


Install:
--------
`go get github.com/dorymint/expfunc`


