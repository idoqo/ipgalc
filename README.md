## idoqo/ipgalc
Given an IPv4 address and a network prefix, ipgalc calculates the corresponding
- Broadcast IP
- Network ID
- First Host
- Last Host
- Group Size and
- Hosts Size

## Usage
```
go run ./main.go --ip $IP --prefix-bit $PREFIX 
e.g
go run ./main.go --ip 127.0.0.1 --prefix-bit 24
```

## How it works
ipgalc works by first splitting up the given IP address into the individual
octets, and uses the prefix to resolve the correct subnet (as well as the subnet
octets).
Both octets are represented as integer arrays (of length 4), and the calculation
is performed thus:

### Network ID
Each element (called octet) in the target IP is _bitwise ANDed_ with the element at same
index in the subnet mask to give the element at that index in the network ID
e.g., say the IP address is "127.0.0.1" and prefix is 24 (i.e `/24`), the subnet mask is "255.255.255.0",
then:
```
networkID[0] = 127 & 255 // 127
networkID[1] = 0 & 255   // 0
networkID[2] = 0 & 255   // 0
networkID[3] = 1 & 0   // 0
```
which gives the network id as 127.0.0.0
### Broadcast IP
Each octet in the target IP is _bitwise ORed_ with the _bitwise complement_ of
the element at same index in the subnet mask. The (negative) result is added to
256 to wrap-around the value, giving the correct value (PS: I can't figure out why
it *needs* the extra 256). Using the previous example, we get:
```
broadcastIP[0] = 256 + (127 | (^255))  // 127
broadcastIP[1] = 256 + (0 | (^255))  // 0
broadcastIP[2] = 256 + (0 | (^255)) // 0
broadcastIP[3] = 256 + (1 | (^255)) // 255
```
giving broadcast IP as 127.0.0.255.

### First Host
// todo :/
### Last Host
// todo :/
### Group Size
// todo :/
### Host Size
// todo :/

## Resources
- [ipcalc](http://jodies.de/ipcalc): an online IP address calculator.
- [Practical Networking on
  YouTube](https://www.youtube.com/watch?v=SM0kdVfhxZ0): Teaches
  a human-friendly way to do subnet magic.

## License
[Do What the Fuck you Want to](http://www.wtfpl.net/).

