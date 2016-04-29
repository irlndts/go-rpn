# go-rpn
reverse polish notation calculator

## Installation
```
go get github.com/irlndts/go-rpn
```

### Parse string
```go
rpn_line := rpn.Parse("3 + 5 + 8 - 1 * ( 1 + 8 )")
```

### Calculate RPN string
```go
result := rpn.Calculate(rpn_line)
```
