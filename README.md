# device

## What is it?

Go library for detecting device information from User Agents, for example assuming you have the following User Agent string:

`Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B143 Safari/8536.25`

You will get:

* Device
  * Type: `phone`
  * Version: `5`
  * Details: `Mobile/10B143`
* Browser
  * Name: `safari`
  * Version: `6.0`
* Operating System
  * Name: `ios`
  * Version: `6.1`
