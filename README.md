# GoPing
Ping program written in Go with the option of flooding a host with pings (purely for educational purposes)

## Ping a host
```console
john@john-MBP goping % sudo ./goping -d example.com   
```

## Ping a host n times
```console
john@john-MBP goping % sudo ./goping -d example.com -c 4
```

## Flood a host 
Do not use this outside of a test environment or without permission of the target
```console
john@john-MBP goping % sudo ./goping -d example.com -f 
```
