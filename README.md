# gopreload

This repo adds functionality that is similar to `LD_PRELOAD` to Golang.
If you import this package and then run your program with the `GO_PRELOAD` environment variable set to a path to plugin, it will be loaded.
To add this to your program all you have to do is add the following into your main file:

```go
import _ "github.com/Granulate/gopreload"
```
