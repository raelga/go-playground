# Introduction to Go

## Features

- Compiles to Machine Code
- Design for concurrence
- Statically typed
- Garbage collected

## Language

### Variables

https://blog.golang.org/gos-declaration-syntax

#### Definition

```go
var foo string;
foo = "bar"     // OK, new value for foo
foo = 5         // compile error, expected string, found int
newfoo = "bar"  // compile error, foo is not defined
bar := "foo"    // implicit var
```

#### Basic Types

https://golang.org/pkg/go/types/ or `go doc types`

```go
    Invalid: {Invalid, 0, "invalid type"},

        Bool:          {Bool, IsBoolean, "bool"},
        Int:           {Int, IsInteger, "int"},
        Int8:          {Int8, IsInteger, "int8"},
        Int16:         {Int16, IsInteger, "int16"},
        Int32:         {Int32, IsInteger, "int32"},
        Int64:         {Int64, IsInteger, "int64"},
        Uint:          {Uint, IsInteger | IsUnsigned, "uint"},
        Uint8:         {Uint8, IsInteger | IsUnsigned, "uint8"},
        Uint16:        {Uint16, IsInteger | IsUnsigned, "uint16"},
        Uint32:        {Uint32, IsInteger | IsUnsigned, "uint32"},
        Uint64:        {Uint64, IsInteger | IsUnsigned, "uint64"},
        Uintptr:       {Uintptr, IsInteger | IsUnsigned, "uintptr"},
        Float32:       {Float32, IsFloat, "float32"},
        Float64:       {Float64, IsFloat, "float64"},
        Complex64:     {Complex64, IsComplex, "complex64"},
        Complex128:    {Complex128, IsComplex, "complex128"},
        String:        {String, IsString, "string"},
        UnsafePointer: {UnsafePointer, 0, "Pointer"},

        UntypedBool:    {UntypedBool, IsBoolean | IsUntyped, "untyped bool"},
        UntypedInt:     {UntypedInt, IsInteger | IsUntyped, "untyped int"},
        UntypedRune:    {UntypedRune, IsInteger | IsUntyped, "untyped rune"},
        UntypedFloat:   {UntypedFloat, IsFloat | IsUntyped, "untyped float"},
        UntypedComplex: {UntypedComplex, IsComplex | IsUntyped, "untyped complex"},
        UntypedString:  {UntypedString, IsString | IsUntyped, "untyped string"},
        UntypedNil:     {UntypedNil, IsUntyped, "untyped nil"},
```

```go
var x byte = 255    // byte alias for uint8
var y byte = x + 5  // overflow to 4
x++                 // overflow to 0

var y byte = 0
y = y - 1           // underflow to 255
```

```go
var foo byte = 2
var foo int32 = 260
sum = foo + bar             // compilation error different types
foo = bar                   // compilation error different types
var conv uint8 = uint8(260) // 4
```

### Arrays

```go
var a [2]int
a[0] = 1
a[1] = "foo"    // compile error, wrong type
a[1] = 2        // OK
a[2] = 3        // compile error, out of bounds
index := 2
a[index] = 3    // runtime error, out of bounds
n := 5
var b [n]int    // compile error, array size must be constant
var c [2]int    // OK
c = a           // OK, copy all values
var d [4]int
d = a           // compile error, wrong type [2]int vs [4]int
```

Go will put the array values contiguous in memory adding padding if needed to align with 4/8 bytes.

```go
var a [2][2][2]int

0x00 a[0][0][0]
0x01 a[0][0][1]
0x02 a[0][1][0]
0x03 a[0][1][1]
0x04 a[1][0][0]
0x05 a[1][0][1]
0x06 a[1][1][0]
0x07 a[1][1][1]
```

### Slices

`TODO`

### Structs

```go
type Person struct {
    Name string
    Age string
    Gender string
    Children []Person   // OK, can be empty, is a slice
    Mother Person       // compile error, invalid recurse type
    Father Person       // Type Person contains a Person, infinite references
    Parents []Person    // OK, can be empty, is a slice
}
```

### Pointer

```go
var x int = 42
var y *int  // pointer to integrer
y = &x      // &x returns the pointer to x (address of x)
z := *y     // *y derreferences y, gets 42, the value of x
x = 24      // updates x value
z = *y      // *y returns the new value of x, 24
```

### Maps

`TODO`

### Definition scope

```go
x := 0
if condition {
    x := 42
}
y := x      // y is 0
if condition {
    y = 21
}
z := y      // z is 21
```

### Functions

```go
func foo(x, y int, z bool) int {
    if z {
        return x + y
    } else {
        return "bar"    // compile error, wrong return type
    }
}

func bar() (string, int) {
    return "The answer is", 42
}
func main() {
    var x int = foo(1, 2, true)     // OK
    var y int = foo(1, 2)           // compile error, wrong  number of arguments
    var z int = foo(1, 2, "bar")    // compile error, last argument not a bool
    var v bool = foo(1, 2, true)    // compile error, wrong return type
    a := 1
    b := 2
    s := foo(a, b, true) + a        // OK, compiler knows that a, b and foo are int
    str, num := bar()               // OK
}
```

### Syntax

- Semicolons are inserted by the compiler

```go
func foo(x int) int {
    x++
    return x
}
foo(1)
```

The compiler will put the `;`:

```go
func foo(x int) int {
    x++;
    return x;
};

foo(1);
```

### Loops

```go
abc := [3]string{"a", "b", "c"}
for i, v := rang abc {
    // i is the index or key if abc is a map
    // v is the value
}
```

#### Range

## Resources

- https://github.com/ardanlabs/gotraining
- https://github.com/dariubs/GoBooks
- http://openmymind.net/The-Little-Go-Book/
- https://www.youtube.com/watch?v=UvpSDtbdGDM
- https://www.youtube.com/watch?v=G3PvTWRIhZA
- https://www.youtube.com/watch?v=ytEkHepK08c
- https://codegangsta.gitbooks.io/building-web-apps-with-go
- https://golang.org/doc/code.html#Organization