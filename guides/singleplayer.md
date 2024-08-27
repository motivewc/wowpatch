# Single Player Guide

## Background
The Warcraft client uses encryption in multiple places to ensure that communications with the battle.net service and game server can't be intercepted. This is one of the reasons why the game client requires patching.

This configuration is suitable for play on a single computer that runs both the game client and TrinityCore.

## Steps

## 1. Decide on a domain name
This can be anything, but it needs to resolve to your machine. We'll use `server.wow` as our example host in this guide.

## 2. Create a self-signed certificate for `server.wow` 
```bash
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 \
  -nodes -keyout server.key -out server.crt -subj "/CN=server.wow" \
  -addext "subjectAltName=DNS:server.wow"
```

This will output two new files, `server.crt` and `server.key`. Move these to the same place as your `bnetserver` binary.

## 3. Add the certificate to your system trust sture
### Mac Users
```bash
sudo security add-trusted-cert -d -r trustAsRoot -k /Library/Keychains/System.keychain server.crt
```
### Windows
Double click the crt, Install Certificate -> Local Machine -> Place certificate in the following store: Trusted Root Certification Authorities.

## 4. Edit your bnetserver.conf file
```
LoginREST.InternalAddress="server.wow"
LoginREST.ExternalAddress="server.wow"
CertificatesFile = "server.crt"
PrivateKeyFile = "server.key"
```

## 5. Update realmlist
Open a MySQL terminal and run the following query on your auth database.

```mysql
UPDATE realmlist SET address = 'server.wow', localAddress = 'server.wow';
```

## 6. Update client portal
In your World of Warcraft installation, find the `_retail_` folder and edit the file `WTF/Config.wtf`. Ensure that there is a line which says `SET portal "server.wow"`.

## 7. (Optional) Update /etc/hosts
Remember how we picked the domain `server.wow`? If you truly own the domain name, point it to your IP address. Otherwise, you'll need to modify your system hosts file.

### Mac
```bash
sudo bash -c 'echo 127.0.0.1 server.wow >> /etc/hosts'
```

### Windows
Edit `C:\Windows\System32\drivers\etc\hosts` with Notepad.exe in Administrator mode and add `127.0.0.1 server.wow`. 

## 8. Patch
Run the wowpatch application on your WoW client. It will spit out a file named "Arctium" or "Arctium.exe" depending on your operating system, and you will use this to play instead of the normal exe.

## 9. Launch
Start your server, then start your client and play.