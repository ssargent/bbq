# BBQ Manager
If you think that name is lousy - you're not alone, I haven't thought of a better one yet.  If this is not your first time here you may be wondering where did *all* the code go?  I decided to reboot and shift the focus.  The legacy version is still around in a [branch](https://github.com/ssargent/bbq/tree/legacy-bbq).  The legacy version taught me about go and about kubernetes.  It did that job fairly well in introducing me to them through a discovery process.  I've a lot more experience with both of those technologies now.  Enough to know that I want to reboot and start somewhat fresh, somewhat simpler.  Also, I want to focus on different things.  Specifically Rust & Networking.  There will still be go, but not only go.

In this repo you'll see a small number of top level directories.

| Directory | Purpose |
|--|---|
|apis | Protobuf API Definitions for gRPC services |
|baste | Baste is the codename for the golang implementation of bbq services.  It will be the initial implementation for some things, and the only implementation for others.. |
|saute | Saute is the codename for the rust implementation of some select bbq services.  Specifically the data collector will be written in rust and will send data via quic/udp to a daemon (go) which will record the data |
|scripts| Scripts will house repo level scripts and useful things| 

## Architecture

At a high level the system will look like this.

![bbq architecture](/bbq-architecture.png "bbq architecture")