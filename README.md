# clearout
#### wrapper of fmt.Print with clear escape sequence

![](./sample.gif)

# usage

## example

```main.go
import "github.com/maito1201/clearout"

func main() {
    cout := clearout.Output{}
    for i := 1; i <= 100; i++ {
        cout.Println("do something")
        cout.Printf("progress %v%%",i)
        cout.Render()
    }
}
```

```
do something
progress 100%
```

## with option

```main.go
cout := clearout.Output{Out: os.Stderr, Prefix: "prefix\n", Suffix: "suffix\n"}
cout.Println("message")
cout.Render()
```

```
prefix
message
suffix
```

## method chain

```main.go
cout := clearout.Output{}
cout.Print(1).Println("2").Printf("%v", 3).WithPrefix("prefix\n").WithSuffix("suffix\n").Render()
```

```
prefix
12
3suffix
```