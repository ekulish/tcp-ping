#TCP Ping

This tool is built as training example for Golang, but have been used to monitor SSL connection to a hardware device as a part of network configuration debugging process. It allows to dial any IP/port and output the result on screen or into the defined file.

##Usage
```
-f="": File name which is used instead of printing results on the screen
-h="127.0.0.1": Host name or IP address
-p=80: Host port
-r=60: Dial period (in seconds)
-t=10: Connection timeout (in seconds)
```
