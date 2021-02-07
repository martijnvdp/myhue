# myhue
Test go to interact with the api of the philips hue 

and learning go 
source: 
https://github.com/amimof/huego

environment vars 

```
"HUEUSER":"someusername"
"HUEIP":"1.1.1.1" (optional wil search for bridge)
"HUETOKEN":"token after first user creation" (after first run it will give you a token to set with this var)
```

after first run a user token wil be created and displayed for future use \
(before first run press the button on the hue bridge) \

# examples 

```
myhue list
```
myhue set:
```
Usage:
  myhue set [flags]

Flags:
  -b, --bright int   brightness level (default 60)
  -h, --help         help for set
  -l, --light int    light id
  -o, --on           Light on or off

  myhue set light 1 --on -b 49
```

# build
build exe: 
go build
# refs
https://github.com/amimof/huego
\
https://golang.org/
