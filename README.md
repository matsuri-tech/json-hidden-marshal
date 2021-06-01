# json-hidden-marshal

## Example

```go
type User struct {
    Name     string `json:"name"`
    Password string `json:"password" hidden:"mask"` // -> masked like "*****", preserving the string length 
    Hidden   int    `json:"hidden" hidden:"-"`
    Hidden2  int    `json:"hidden" hidden:"true"` // hidden:"-" or hidden:"true" to skip
}
```
