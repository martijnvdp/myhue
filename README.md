# myhue
Test go to interact with the api of the philips hue 

and learning go 
source: 
https://github.com/amimof/huego

environment vars \

'''
"HUEUSER":"someusername"
"HUEIP":"1.1.1.1"
"HUETOKEN":"token after first user creation"

'''
after first run a user token wil be created and displayed for future use \
(before first run press the button on the hue bridge) \
\
#
example 

'''
myhue -list
myhue -light 1 -on
'''
#
build exe: \
go build
#ref
https://github.com/amimof/huego
\
https://golang.org/
