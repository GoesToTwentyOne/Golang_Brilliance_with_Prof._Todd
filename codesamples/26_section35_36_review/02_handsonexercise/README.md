# What is BET?
- Benchmarking
- ExampleTest
- Testing
#  Covered This section
```
tests
benchmarks
coverage
coverage shown in web browser
examples shown in documentation in a web browser
Godoc Explanation
gofmt :formats go code
go vet :reports suspicious constructs
golint :reports poor coding style

```
#commands
```
godoc -http=:8080 
go test 
go test -bench . 
don’t forget the “.” in the line above
go test -cover 
go test -coverprofile c.out 
go tool cover -html=c.out

```