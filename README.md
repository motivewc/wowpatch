# wowpatch
It's yet another WoW client patcher, but this time written in Go in order to maximize cross-platform
compatibility, and without any in-client memory modifications. This means that you can generate
a single patched client and redistribute it to others as long as they use the same operating system and processor
architecture.

## Who can use this?
This approach will ONLY work if you:
1. are connecting to a server with a valid TLS certificate that chains to a trusted root CA
in your system trust store.
1. are using a hostname and not an IP address for your portal cvar setting in `WTF/Config.wtf`.
1. are connecting to a server that uses the same [gamecrypto key](https://github.com/TrinityCore/TrinityCore/blob/343f637435cc97ddedd725945cbad417c8f14391/src/server/game/Server/Packets/AuthenticationPackets.cpp#L221) as what is hardcoded (so basically, TrinityCore)

## Usage
If you are an advanced user, the below will likely make sense to you. In case you need additional guidance, see either our [single player](guides/singleplayer.md) or [multiplayer](guides/multiplayer.md) guide.

```bash
./wowpatch -h
This application takes as input a retail World of Warcraft client and will generate a modified executable
from it by using binary patching. The resulting executable can be run safely and connect to private servers.

Usage:
  wowpatch [flags]

Examples:
wowpatch -l ./your/wow/exe -o ./patched-exe

Flags:
  -h, --help                    help for wowpatch
  -o, --output-file string      where to output a modified client (default "Arctium")
  -s, --strip-binary-codesign   removes macOS codesigning from resulting binary (default true)
  -l, --warcraft-exe string     the location of the WoW executable (default "/Applications/World of Warcraft/_retail_/World of Warcraft.app/Contents/MacOS/World of Warcraft")
```

## FAQ
**Q: Why does this generate an exe with the name `Arctium` by default?**

**A:** In the event your client crashes, this helps Blizzard filter out the private server noise from
their automated client telemetry. 

## Thanks
An absolutely **enormous** amount of thanks to [Fabian](https://github.com/Fabi) from [Arctium](https://arctium.io/) for basically
all of the knowledge that went into this.