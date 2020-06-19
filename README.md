# pscan

Simple and fast TCP port scanner written in Golang.

## Precompiled Binaries For Linux

[Latest release](https://github.com/s3rgeym/pscan/releases/latest).

## Manual Install

```zsh
$ make build
$ sudo make install
```

## Usage

```zsh
$ pscan 103.1.172.0/22 -p 22 -c 10000 -l 10000
103.1.174.197:22
103.1.174.193:22
103.1.174.3:22
...
```

IP Ranges: `x.x.x.x/x`, `x.x.x.x-x.x.x.x`.

Help:

```zsh
$ pscan --help
```
