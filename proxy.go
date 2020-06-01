package proxy

import (
    "time"
    "io/ioutil"
    "net/http"
    "regexp"
    "errors"
    "fmt"
)

type Proxy struct {
    options     Options
    proxies     map[string]int
    jobs        chan bool
    socksUrl    string
}

type Options struct {
    TTL     int
    Period  int
}

func New(o Options) *Proxy {
    return &Proxy{
        options:     o,
        proxies:     make(map[string]int),
        jobs:        make(chan bool),
        socksUrl:    "https://www.socks-proxy.net",
    }
}

func (proxy *Proxy) GetTTL() int {
    return proxy.options.TTL
}

func (proxy *Proxy) GetPeriod() int {
    return proxy.options.Period
}

func (proxy *Proxy) GetAll() (proxies []string, err error) {
    now := int(time.Now().Unix())

    for ip, exp := range proxy.proxies {
        if exp > now {
            proxies = append(proxies, ip)
        }
    }

    return
}

func (proxy *Proxy) Random() (string, error) {
    now := int(time.Now().Unix())

    for ip, exp := range proxy.proxies {
        if exp > now {
            return ip, nil
        }
    }

    return "", errors.New("Empty list")
}

func (proxy *Proxy) Insert(ip string) error {
    proxy.proxies[ip] = int(time.Now().Unix()) + proxy.options.TTL

    return nil
}

func (proxy *Proxy) worker(jobs <-chan bool) {
    for range jobs {
        fmt.Println("receiving a job...")

        ips, err := proxy.FetchProxies()

        if err != nil {
            continue
        }

        for _, ip := range ips {
            proxy.Insert(ip)
        }
    }

    return
}

func (proxy *Proxy) Run() error {
    go proxy.worker(proxy.jobs)

    go func() {
        proxy.jobs <- true

        ticker := time.NewTicker(time.Duration(proxy.options.Period * int(time.Second)))
        for {
            <-ticker.C
            proxy.jobs <- true
        }
    }()

    return nil
}


func (proxy *Proxy) FetchProxies() (ips []string, err error) {
    res, err := http.Get(proxy.socksUrl)
    defer res.Body.Close()
    if err != nil {
        return ips, err
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return ips, err
    }

    ips = proxy.ParseIPs(string(body))

    return ips, nil
}

func (proxy *Proxy) ParseIPs(html string) (ips []string) {
    re, _ := regexp.Compile(`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+`)
    ips = re.FindAllString(html, -1)

    return ips
}
