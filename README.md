# MyHue
Test go cli app to interact with the api of the philips hue 
using cobra and viper libs

Just to Learn go 
source: 
https://github.com/amimof/huego

## Usage ##
```
myhue list

myhue set -l 1 -b 50

```
After first run a user token wil be created and and stored in the config file \
before the first run you have to press the button on the hue bridge 

config file is stored in ###home\config.myhue.yaml###

## Examples ## 

```
myhue list
```
myhue set:
```
Usage:
  myhue set [flags]

Flags:
  -b, --bright int   brightness level
  -h, --help         help for set
  -l, --light int    light id
  -o, --on           Light on or off

  myhue set light 1 --on -b 49var a bool
		
```
## environment vars ##

```
"HUE_USER":"someusername"
"HUE_IP":"1.1.1.1" (optional wil search for bridge)
"HUE_TOKEN":"token after first user creation" (after first run it will give you a token to set with this var)
```

# build
build exe: 
go build

# refs
```
https://github.com/amimof/huego
https://golang.org
https://github.com/spf13/cobra
https://github.com/spf13/viper
```
