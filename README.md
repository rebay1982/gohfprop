# hfprop
I was looking for a simple way to just display the HF propagation condition on the command line because that's where
I spend most of my time. Everytime I wanted to check conditions I either needed to load up a webpage on my phone or 
fire up a web browser to fetch the conditions I'm interested in.

So I wrote a super simple command line utility that does just that for me.

**Credit to n0nbh (www.hamqsl.com) for providing the calculated HF conditions. 73!**

## Requirements 
The code doesn't have any dependency on anything but the Go runtime. Find the installation instructions
(here)[https://go.dev/doc/install].

## Build or Run
```
# Build:
go build hfprop.go

# Run:
go run hfprop.go
```

## Install
Just copy the binary from the build step described above to a directory in your path.

## Running
```
hfprof

        HF Propagation

 Band            Day     Night
-------------------------------
 80m-40m         Poor    Good
 30m-20m         Good    Good
 17m-15m         Good    Good
 12m-10m         Good    Poor
```

Super simple, no need for a web browser. Enjoy!
