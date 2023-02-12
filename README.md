# Optional
Golang library providing generic optional type. Provides zero allocation alternative for using pointer to optional structs.

## Example
Use of default value
```go
func generateURL(host string, port optional.Optional[int]) string {
	return fmt.Sprintf("postgres://%s:%d", host, port.GetOrElse(5432))
}

func main() {
	fmt.Println(generateURL("postgres", optional.New(5533)))
	// Prints
	// postgres://postgres:5533
	fmt.Printf(generateURL("postgres", optional.None[int]()))
	// Prints
	// postgres://postgres:5432
}
```
