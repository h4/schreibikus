# Schreibikus

Simple tool for moving files to separate folders based on file creation time

## Build

```
go get github.com/stathat/jconfig
go build
```

## Usage 

```
./schreibikus [args]
```

### Args

* -source_root — Source root path
* -dest_root — Dest root path
* -config_path — Path to config.json file

source_root and dest_root or config_path must be defined.

## Config 

See `config_sample.json` for available params
