# fits
Fits is a **fi**le **t**ransfer **s**ervice written in go.

## Install

### Docker

Build docker image:
```
docker build . -t fits
```

Run docker image with your `from` and `to` folders:
```
docker run -v $HOME/fits/from:/from -v $HOME/fits/to:/to -it fits
```
