## dnsbyo

## Motivation
dnsbyo is a personal learning project to understand how DNS Servers work under the hood.
You will find other repositories suffixed with byo, which references [Build your Own X](https://github.com/danistefanovic/build-your-own-x).

I've created this series of small projects to experiement a new way of learning new things, by building a prototype of everything I wish to learn.

I intend to follow most part of [this](https://roadmap.sh/backend) roadmap. So there will probably be a 'byo' project for many of the boxes presented on that roadmap.

## Try it

Open a terminal window and run:

```
make
```

The server is now running on port 8090.

_amazon.com_ and _google.com_ are the hostnames available for querying.  Try adding other records to **records.json**

Open a second terminal window to run:

```
> nslookup -port=8090 google.com localhost

Server:         localhost
Address:        ::1#8090 
                          
Name:   google.com        
Address: 216.58.196.142   
Name:   google.com        
Address: 216.58.196.142   
```

## TODO
- [x] Implement basic A type queries
- [ ] Handle unregistered names
- [ ] Add [gopacket](https://github.com/google/gopacket) dependency
- [x] Easily add new names
