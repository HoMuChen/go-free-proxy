# go-free-proxy

Periodically fetch free proxies in the background and maintain time to live for each ip port pair.

## Usage 

```
go get github.com/HoMuChen/go-free-proxy
```
```go
package main

import (
    "fmt"
    "time"
    
    "github.com/HoMuChen/go-free-proxy"
)

func main() {
    //proxy.Options{TTL, Period} (seconds)
    proxyService := proxy.New(proxy.Options{10, 20})
    
    //Calling Run() will start to fetch proxies immediately and every {Period} seconds later.
    proxyService.Run()

    for {
        ip, err := proxyService.Random()
        //If no proxies is available, err will be errors.New("Empty list")
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(ip)

        time.Sleep(time.Second)
    }
}
```
