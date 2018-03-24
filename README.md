expfunc
=======
expfunc is simple cli tool for golang  
export function name from go source file


Usage:
------
output the function names:  
`expfunc -file /path/to/src.go`

with Test for testing  
`expfunc -test -file /path/to/src.go`

only exported function  
`expfunc -exported -file /path/to/src.go`

help  
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
`go get -v -u github.com/yaeshimo/expfunc`


