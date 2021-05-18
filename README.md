# fits
Fits is a **fi**le **t**ransfer **s**ervice written in go.

## Why?
I just want a service that watch a folder and move added files files from a A to B. No more no less, just make one task and make it good.

Examples:
- Your NAS only offers you bloaty and payed enterprise apps for backing up files. Set up a cronjob to generate a backup on your server on and use fits to transfer them to your NAS via FTP/SMB.

- Your company gets files uploaded to a folder and they need to be moved somewhere for processing. Just set up a fits and let it  transfer the new files.

- You run a kubernetes cluster and want a service that can run as a pod and move files between FTP/SMB/SSH to FTP/SMB/SSH.

- You generate blog posts on your local machine and you want them to automatically be uploaded to your website. Use fits to automatically take files from your system and SSH them to your server.

## Usage
When you start fits it will transfer files between your `from` and `to` from your config.
Make sure to set the paths correctly before running fits. Check the [Config](#config) section how to edit it.

### Go
Build the go binary:
```
go build .
```

Run fits:
```
./fits
```

### Docker

Build docker image:
```
docker build . -t fits
```

Run docker image with your `from` and `to` folders:
```
docker run -v $HOME/fits/from:/from -v $HOME/fits/to:/to -v /$HOME/code/fits/config:/config -it fits
```

If you're using SSH you need specify the mount path for `known_hosts` to docker:
```
docker run -v /$HOME/fits/from:/from -v /$HOME/.ssh:/.ssh -v /$HOME/code/fits/config:/config -it fits
```

## Config
To use fits put a `from` and `to` destination in your `config/config.toml`

### File system
Supports `[from]` and `[to]`

Example File system `config.toml`
```
[from]
path = "$HOME/fits/from"

[to]
path = "$HOME/fits/to"
```

### FTP
Supports `[to]`

Example ftp `config.toml`
```
[from]
path = "$HOME/fits/from"

[to]
path = "ftp://ftp.mydomain.com/to" or "ftp://192.168.1.100/to"
username = "username"
password = "password"
```

### SSH
Supports `[to]`

**Note!**  
SSH will check `known_hosts` at the running users `$HOME`-folder.
Please add the host you are sending to in `~/.ssh/known_hosts`  
This can be done with:  
```
ssh-keyscan -H 192.168.1.100 >> ~/.ssh/known_hosts
```

Example ssh `config.toml`
```
[from]
path = "$HOME/fits/from"

[to]
path = "ssh://mydomain.com/to" or "ssh://192.168.1.100/to"
username = "username"
password = "password"
```

### SMB
Supports `[to]`

Example ftp `config.toml`
```
[from]
path = "$HOME/fits/from"

[to]
path = "smb://192.168.1.100/to" or "//192.168.1.100/to"
username = "username"
password = "password"
```
