# go-externalip

Warning: Quick hack. IPv4 & NAT-PMP only. For use in ```ddclient.conf``` or similar.

Usage:
```
$ go get -u github.com/okket/go-externalip
$ go-externalip
11.22.33.44
````

Details see https://github.com/jackpal/go-nat-pmp

Web service alternative:

```
$ curl http://ipv4.myexternalip.com/raw
11.22.33.44
```
