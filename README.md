# clearout
#### wrapper of fmt.Print with clear escape sequence

![](./sample.gif)

# usage

## example

```main.go
import "github.com/maito1201/clearout"

func main() {
    r := clearout.Output{}
    for i := 1; i <= 100; i++ {
        r.Println("do something")
        r.Printf("progress %v%%",i)
        r.Render()
    }
}
```

```
do something
progress 100%
```

## with option

```main.go
r := clearout.Output{Out: os.Stderr, Prefix: "prefix\n", Suffix: "suffix\n"}
r.Println("message")
r.Render()
```

```
prefix
message
suffix
```

## method chain

```main.go
r := clearout.Output{}
r.Print(1).Println("2").Printf("%v", 3).WithPrefix("prefix\n").WithSuffix("suffix\n").Render()
```

```
prefix
12
3suffix
```