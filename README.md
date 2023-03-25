# autocvd-go

[![CI](https://github.com/shunk031/autocvd-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/shunk031/autocvd-go/actions/workflows/ci.yaml)
[![Release](https://github.com/shunk031/autocvd-go/actions/workflows/release.yaml/badge.svg)](https://github.com/shunk031/autocvd-go/actions/workflows/release.yaml)

A golang cli tool for setting `CUDA_VISIBLE_DEVICES` based on GPU utilization.
This project is heavily inspired by the python version of [`jonasricker/autocvd`](https://github.com/jonasricker/autocvd).

The aim of this tool, as well as [`jonasricker/autocvd`](https://github.com/jonasricker/autocvd), is as follows:

> On a system with multiple NVIDIA GPUs, autocvd eliminates the need for manually specifying the CUDA_VISIBLE_DEVICES environment variable. This comes in especially handy on systems with multiple users, like a shared GPU server. It is dependency-free and requires no code changes.

## Requirements

autocvd-go uses `nvidia-smi` to query GPU utilization. Make sure that it is installed and callable.

## Installation

Download the binary from [GitHub Releases](https://github.com/shunk031/autocvd-go/releases/latest) and drop it in your $PATH.

```shell
wget https://github.com/shunk031/autocvd-go/releases/latest/download/auto-cvd_linux_x86_64.tar.gz \
    && tar -xvzf autocvd_linux_x86_64.tar.gz autocvd \
    && rm autocvd_linux_x86_64.tar.gz

# Then, move `autocvd` binary in your $PATH
```

## Usage

To execute a command on a free GPU as well as [`jonasricker/autocvd`](https://github.com/jonasricker/autocvd), run the following:

```console
$ eval $(autocvd) <command>
```

Possible use cases when training deep learning models that use GPUs include the following:

```console
$ eval $(autocvd) python train.py
```

## Examples

This example is identical to [`jonasricker/autocvd`](https://github.com/jonasricker/autocvd).

```shell
# run command on two free GPUs
$ eval $(autocvd -n 2) <command>

# run command on least-used GPU (i.e., do not wait if no GPU is free)
$ eval $(autocvd -l) <command>

# exclude certain GPUs
$ eval $(autocvd -x 0 2) <command>

# if no free GPU is available immediately, wait for 60 seconds only
$ eval $(autocvd -t 60) <command>

# export environment variables into the current shell
$ . <(autocvd -e)  # alternatively: source <(autocvd -e)
```

## LICENSE

MIT

## Acknowledgment

- jonasricker/autocvd: Tool to automatically set CUDA_VISIBLE_DEVICES based on GPU utilization. Usable from command line and code. https://github.com/jonasricker/autocvd#examples 
