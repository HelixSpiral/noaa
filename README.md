NOAA - National Oceanic and Atmospheric Administration
---

This is a simple Golang wrapper for parsing data from various APIs provided by NOAA.

There weren't any packages that I could find that performed the functions I needed, so I wrote this one. Woefully incomplete in it's current state but PRs are welcome.

I do intend to fully flesh this out at some point in the future, but for now it solves the immediate problem I had.

# Usage - NDBC - National Data Buoy Center

Craete a new wrapper
```go
n := noaa.New()
```

Call any of the exported functions
```go
resp, err := n.NDBC.GetPictureFromBuoy(44027)
if err != nil {
    panic(err)
}

log.Println(resp)

resp2, err := n.NDBC.GetLatestDataFromBuoy(44027)
if err != nil {
    panic(err)
}

log.Println(resp2)
```
