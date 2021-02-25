# go-findmyip

Find my public IP from external DNS reference.

We basically implemented the following `host` inquiry from *opendns*:
```
host myip.opendns.com resolver1.opendns.com
```

Large portion of the code is borrowed from the post in this
[thread]: https://www.reddit.com/r/golang/comments/esus2j/specifying_dns_server_for_lookup/
by g-a-c.

