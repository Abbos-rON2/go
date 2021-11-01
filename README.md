# bigint
Implement a package named "bigint" in golang. That has the following functions and methods.

- func NewInt(num string) (Bigint, error)
- func (z *Bigint) Set(num string) error
- func Add(a, b Bigint) Bigint
- func Sub(a, b Bigint) Bigint
- func Multiply(a, b Bigint) Bigint
- func Mod(a, b Bigint) Bigint
- func (x *Bigint) Abs() Bigint

examples:

```go
a:=bigint.NewInt("988847123412385995937737458959")
b:=bigint.NewInt("21231231231231231231231231233")
b.Set("1") // b = "1"
c:=bigint.Add(a, b) // c = "988847123412385995937737458960"
d:=bigint.Sub(a, b) // d = "988847123412385995937737458958"
e:=bigint.Multiply(a, b) // e = "988847123412385995937737458959"
f:=bigint.Mod(a, b) // f = "0"
```
