# fits
Fits is a **fi**le **t**ransfer **s**ervice written in go.

## Install

### Docker

Build the container:
```
docker build . -t fits
```

Run the container with your from and to folders:
```
docker run -v $HOME/fits/from:/from -v $HOME/fits/to:/to -it fits
```
