# fits
Fits is a **fi**le **t**ransfer **s**ervice written in go.

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
docker run -v $HOME/fits/from:/from -v $HOME/fits/to:/to -it fits
```

## Config
To use fits put a `from` and `to` destination in your `config.toml`

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
