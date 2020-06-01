package proxy_test

import (
    "testing"
    //"time"
    "errors"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-free-proxy"
)

func TestGetTTL(t *testing.T) {
    proxy := proxy.New(proxy.Options{})
    ttl := proxy.GetTTL()

    assert.Equal(t, ttl, 0, "default ttl should 0")

    t.Log(ttl)
}

func TestGetPeriod(t *testing.T) {
    proxy := proxy.New(proxy.Options{})
    period := proxy.GetPeriod()

    assert.Equal(t, period, 0, "default period should 0")

    t.Log(period)
}

func TestGetAll(t *testing.T) {
    proxy := proxy.New(proxy.Options{})
    ips, err := proxy.GetAll()

    assert.Equal(t, err, errors.New("Empty list"))
    assert.Len(t, ips, 0)

    t.Log(ips)
}

func TestRandom(t *testing.T) {
    proxy := proxy.New(proxy.Options{})
    ip, err := proxy.Random()

    assert.Equal(t, err, errors.New("Empty list"))
    assert.Equal(t, ip, "")

    t.Log(ip)
}

func TestInsertExpImmedidately(t *testing.T) {
    proxy := proxy.New(proxy.Options{})
    proxy.Insert("127.0.0.1")

    ips, err := proxy.GetAll()
    assert.Equal(t, err, errors.New("Empty list"))
    assert.Len(t, ips, 0)

    t.Log(ips)
}

func TestInsert(t *testing.T) {
    proxy := proxy.New(proxy.Options{30, 5})
    proxy.Insert("127.0.0.1")

    ips, err := proxy.GetAll()
    assert.NoError(t, err)
    assert.Len(t, ips, 1)

    t.Log(ips)
}

func TestExpire(t *testing.T) {
    proxy := proxy.New(proxy.Options{30, 5})
    proxy.Insert("127.0.0.1")
    proxy.Expire("127.0.0.1")

    ips, err := proxy.GetAll()
    assert.Equal(t, err, errors.New("Empty list"))
    assert.Len(t, ips, 0)

    t.Log(ips)
}

func TestParseIPs(t *testing.T) {
    proxy := proxy.New(proxy.Options{30, 5})

    ips := proxy.ParseIPs("i am a paragrapgh containging some ip address, like 127.0.0.1:80 or 10.0.0.1:3421")

    assert.Len(t, ips, 2)

    t.Log(ips)
}

func TestFetch(t *testing.T) {
    proxy := proxy.New(proxy.Options{30, 5})

    ips, err := proxy.FetchProxies()

    assert.NoError(t, err)

    t.Log(ips)
}

//func TestRun(t *testing.T) {
//    proxy := proxy.New(proxy.Options{30, 5})
//
//    proxy.Run()
//
//    time.Sleep(15 * time.Second)
//}
