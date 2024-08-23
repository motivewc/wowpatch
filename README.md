# wowpatch
It's yet another WoW client patcher, but this time written in Go and without any in-client 
memory modifications. 

This approach will ONLY work if you:
1. are connecting to a server with a valid TLS certificate that chains to a trusted root CA
in your system trust store.
1. are using a hostname and not an IP address for your portal cvar setting.
1. are connecting to a server that uses the same [gamecrypto key](https://github.com/TrinityCore/TrinityCore/blob/343f637435cc97ddedd725945cbad417c8f14391/src/server/game/Server/Packets/AuthenticationPackets.cpp#L221) as what is hardcoded (so basically, TrinityCore)

The upsides of this approach are that this patcher should work on Windows and macOS, and if you're a private server
owner, you can simply redistribute the generated Wow.exe file. I have only tested this on macOS, as that is where
I develop and is the reason for this tool's existence.

The code is super shitty, but I'm releasing this so that there's documentation for other mac users out there.

An absolutely **enormous** amount of thanks to [Fabian](https://github.com/Fabi) from [Arctium](https://arctium.io/) for basically
all of the knowledge that went into this.

## Usage
1. Open a Terminal
1. Install Go (`brew install go`)
1. `go install github.com/motivewc/wowpatch`
1. `cd /Applications/World of Warcraft/_retail_/World of Warcraft.app/Contents/MacOS/`
1. `wowpatch "World of Warcraft"`
1. A file named "Arctium" should be generated in the current working directory. Keep the exe named as such to be nice to Blizzard and use this exe to launch WoW from now on.