# myhue
Test go to interact with the api of the philips hue 

and learning go 
source: 
https://github.com/amimof/huego

environment vars 

```
"HUEUSER":"someusername"
"HUEIP":"1.1.1.1"
"HUETOKEN":"token after first user creation"
```

after first run a user token wil be created and displayed for future use \
(before first run press the button on the hue bridge) \

# examples 

```
myhue list
myhue list --width
myhue list -w --reachable --filter "name"

myhue set --light 1 --on
myhue set -l 1 --on -br 60
myhue set -l 1 
```
# build
build exe: 
go build
# refs
https://github.com/amimof/huego
\
https://golang.org/
