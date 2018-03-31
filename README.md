# go-polymorphic-shapes
Just a simple example on how to achieve polymorphism in Go using Interfaces


## Polymorphism using Go

Polymorphism in GO can be achieved using [Interfaces](https://golang.org/ref/spec#Interface_types). The idea behind the concept is similar to some Ojbect Oriented languages like JAVA or C++, but done in a different way.

Suppose you have a type `Rectangle` and `Circle` that implements a `Shape` type behaviour. If you have an object `shape` of type `Shape` and you do `shape.area()` you want to get the right value whereas `shape` is a `Rectangle` or a `Circle`. Each of them will have a different implementation, but you don't care about that; you just want to get the right area calculation for that shape.

First you need to define some method signatures within an `interface` that will be called `Shape`:

```go
type Shape interface {
   Area() float32
   Perimeter() float32
   Dimensions() string
}
```

Then you define `Circle` and `Rectagle` types:

```go
type Rectangle struct {
   length float32
   width  float32
}

type Circle struct {
   radius float32
}
```

And you make them implement `Shape` interface. In Go, you don't need to explicitly set that a type implements an interface like in JAVA or C++; you just have to implement the methods defined within the interface:

```go
// Rectangle Shape implementation:
func (r Rectangle) Area() float32 {
   return r.length * r.width
}

func (r Rectangle) Perimeter() float32 {
   return 2 * (r.length + r.width)
}

func (r Rectangle) Dimensions() string {
   return fmt.Sprintf("length %.2f cm and width %.2f cm", r.length, r.width)
}
```

```go
// Circle Shape implementation:
func (c Circle) Area() float32 {
   return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
   return 2 * math.Pi * c.radius
}

func (c Circle) Dimensions() string {
   return fmt.Sprintf("radius %.2f cm", c.radius)
}
```

Then you create, for example, a rectangle of 2x3 and and a circle of radio 1:

```go
s := Rectangle{2, 3}
c := Circle{1}
```

And finally you create a var of type `Shape` and set it to one of the shapes defined above:

```go
var myShape Shape
myShape = s

fmt.Printf("\nArea: %f", myShape.Area())
```

That should return the rectangle area. You could also put both shapes within a collection of type `Shape` and iterate the collection like this:

```go
shapesCollection := []Shape{s, c}

for _, shape := range shapesCollection {
      fmt.Printf(
         "A %s of %s has an area of %.2f cm2 and a perimeter of %.2f cm.\n",
         reflect.TypeOf(shape).Name(),
         shape.Dimensions(),
         shape.Area(),
         shape.Perimeter(),
      )
   }
```

That should produce this output:

```bash
A Rectangle of length 2.00 cm and width 3.00 cm has an area of 6.00 cm2 and a perimeter of 10.00 cm.
A Circle of radius 1.00 cm has an area of 3.14 cm2 and a perimeter of 6.28 cm.
```

## Installation and Running

If you want to try out the above example, just get the package using `go get` tool:

```bash
go get github.com/charlysan/go-polymorphic-shapes
```

And then you can build it and run it like this:

```bash
cd $GOPATH/src/github.com/charlysan/go-polymorphic-shapes
go run main.go
```
