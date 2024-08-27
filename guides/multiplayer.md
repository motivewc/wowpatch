# Detailed Guide

## Background
The Warcraft client uses encryption in multiple places to ensure that communications with the battle.net service and game server can't be intercepted. This is one of the reasons why the game client requires patching.

TrinityCore ships with an encryption key, but this guide will have you generate your own. This configuration is more secure and suitable for use with servers that have multiple players connecting to them.

## Steps
### 1. Decide on a valid fully-qualified domain name
You need to own a domain name through a registrar such as [Cloudflare](https://www.cloudflare.com). We'll use them in this guide but any place you can register a domain and control its DNS will work.

You'll also need to point the domain name to your server's IP address. If you're using Cloudflare, do NOT use proxy mode, use a [gray cloud DNS record](https://developers.cloudflare.com/dns/manage-dns-records/reference/proxied-dns-records/#dns-only-records).

For the remainder of this guide, we'll use the DNS name `server.wow` as an example. Replace this with your own!

### 2. Obtain a valid certificate for your server hostname
Obtaining a valid certificate can be done a multitude of ways, but is outside the purview of this guide. 

I recommend using LetsEncrypt with a DNS challenge to make things simple. A guide for that is available [here](https://www.digitalocean.com/community/tutorials/how-to-acquire-a-let-s-encrypt-certificate-using-dns-validation-with-acme-dns-certbot-on-ubuntu-18-04).

If you prefer to not use LetsEncrypt and do things the old fashioned way, you can generate a certificate signing request and then purchase a certificate online by sharing the certificate signing request with a provider like [Comodo](https://comodosslstore.com/positivessl.aspx). You'll still need to validate your DNS ownership, and this method costs money, so it's not recommended.

```bash
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 \
  -nodes -keyout server.key -out server.crt -subj "/CN=server.wow" \
  -addext "subjectAltName=DNS:server.wow"
```

At the end of this step, you should have two new files, one certificate, and one private key. We'll call these `server.crt` and `server.key`. They should be in PEM format, which is normally the default.

### 3. Update bnetserver.conf
```
CertificatesFile = "server.crt"
PrivateKeyFile = "server.key"
```
We edit these files from the default name so that the keys shipped by default don't accidentally ever overwrite your own certificates.

### 4. Update realmlist
Open a MySQL terminal and run the following query on your auth database.

```mysql
UPDATE realmlist SET address = 'server.wow', localAddress = 'server.wow';
```


### 5. Client
In your World of Warcraft installation, find the `_retail_` folder and edit the file `WTF/Config.wtf`. Ensure that there is a line which says `SET portal "server.wow"`.

### 6. Patch
Run the wowpatch application on your WoW client. It will spit out a file named "Arctium" or "Arctium.exe" depending on your operating system, and you will use this to play instead of the normal exe.

### 7. Launch
Start your server, then start your client and play.