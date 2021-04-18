# fits
Fits is a **fi**le **t**ransfer **s**ervice written in go.

# Usage
When you start fits it will transfer files between your `from` and `to` from your config.
Make sure to set the paths correctly before running fits. Check the [Config](#config) section how to edit it.

## Go
Build the go binary:
```
go build .
```

Run fits:
```
./fits
```

## Docker

Build docker image:
```
docker build . -t fits
```

Run docker image with your `from` and `to` folders:
```
docker run -v $HOME/fits/from:/from -v $HOME/fits/to:/to -it fits
```

# Config
To use fits put a `from` and `to` destination in your `config.toml`

## File system
Example File system `config.toml`:
```
[from]
path = "$HOME/fits/from"

[to]
path = "$HOME/fits/to"
```

## Ftp
At the moment you can only set ftp in `to` settings.

Example ftp `config.toml`:
```
[from]
path = "$HOME/fits/from"

[to]
path = "ftp://ftp.mydomain.com/to" or "ftp://192.168.1.100/to"
username = "username"
password = "password"
```
