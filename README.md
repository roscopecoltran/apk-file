# apk-file

[![Travis CI](https://travis-ci.org/jessfraz/apk-file.svg?branch=master)](https://travis-ci.org/jessfraz/apk-file)

Search apk packages via the command line.

> Basically `apt-file` but for alpine.

## Installation

#### Binaries

- **darwin** [386](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-darwin-386) / [amd64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-freebsd-386) / [amd64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-linux-386) / [amd64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-linux-amd64) / [arm](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-linux-arm) / [arm64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-windows-386) / [amd64](https://github.com/jessfraz/apk-file/releases/download/v0.1.1/apk-file-windows-amd64)

#### Via Go

```bash
$ go get github.com/jessfraz/apk-file
```

## Usage

```console
$ apk-file posix.so
FILE                                                       PACKAGE             BRANCH              REPOSITORY          ARCHITECTURE
/usr/lib/lua/5.1/posix.so                                  lua5.1-posix        edge                main                armhf
/usr/lib/lua/5.3/posix.so                                  lua5.3-posix        edge                main                armhf
/usr/lib/lua/5.2/posix.so                                  lua5.2-posix        edge                main                armhf
/usr/lib/lua/5.1/posix.so                                  lua5.1-posix        edge                main                x86_64
/usr/lib/lua/5.2/posix.so                                  lua5.2-posix        edge                main                x86_64
/usr/lib/lua/5.3/posix.so                                  lua5.3-posix        edge                main                x86_64
/usr/lib/lua/5.3/posix.so                                  lua5.3-posix        edge                main                x86
/usr/lib/lua/5.1/posix.so                                  lua5.1-posix        edge                main                x86
/usr/lib/lua/5.2/posix.so                                  lua5.2-posix        edge                main                x86
/usr/lib/python2.7/site-packages/uptime/_posix.so          py-uptime           edge                testing             armhf
/usr/lib/python2.7/site-packages/uptime/_posix.so          py-uptime           edge                testing             x86_64
/usr/lib/python2.7/site-packages/uptime/_posix.so          py-uptime           edge                testing             x86
/usr/lib/python2.7/site-packages/psutil/_psutil_posix.so   py-psutil           edge                community           armhf
/usr/lib/python2.7/site-packages/psutil/_psutil_posix.so   py-psutil           edge                community           x86_64
/usr/lib/python2.7/site-packages/psutil/_psutil_posix.so   py-psutil           edge                community           x86
/usr/lib/lua/5.3/rex_posix.so.2.8                          lua5.3-rex-posix    edge                main                armhf
/usr/lib/lua/5.3/rex_posix.so                              lua5.3-rex-posix    edge                main                armhf
/usr/lib/lua/5.2/rex_posix.so.2.8                          lua5.2-rex-posix    edge                main                armhf
/usr/lib/lua/5.2/rex_posix.so                              lua5.2-rex-posix    edge                main                armhf
/usr/lib/lua/5.1/rex_posix.so.2.8                          lua5.1-rex-posix    edge                main                armhf
/usr/lib/lua/5.1/rex_posix.so                              lua5.1-rex-posix    edge                main                armhf
/usr/lib/lua/5.2/rex_posix.so                              lua5.2-rex-posix    edge                main                x86_64
/usr/lib/lua/5.2/rex_posix.so.2.8                          lua5.2-rex-posix    edge                main                x86_64
/usr/lib/lua/5.3/rex_posix.so                              lua5.3-rex-posix    edge                main                x86_64
/usr/lib/lua/5.3/rex_posix.so.2.8                          lua5.3-rex-posix    edge                main                x86_64
/usr/lib/lua/5.1/rex_posix.so                              lua5.1-rex-posix    edge                main                x86_64
/usr/lib/lua/5.1/rex_posix.so.2.8                          lua5.1-rex-posix    edge                main                x86_64
/usr/lib/lua/5.2/rex_posix.so                              lua5.2-rex-posix    edge                main                x86
/usr/lib/lua/5.2/rex_posix.so.2.8                          lua5.2-rex-posix    edge                main                x86
/usr/lib/lua/5.3/rex_posix.so                              lua5.3-rex-posix    edge                main                x86
/usr/lib/lua/5.3/rex_posix.so.2.8                          lua5.3-rex-posix    edge                main                x86
/usr/lib/lua/5.1/rex_posix.so                              lua5.1-rex-posix    edge                main                x86
/usr/lib/lua/5.1/rex_posix.so.2.8                          lua5.1-rex-posix    edge                main                x86
/usr/lib/python2.7/site-packages/_psutil_posix.so          py-psutil           v3.3                community           armhf
/usr/lib/python2.7/site-packages/_psutil_posix.so          py-psutil           v3.3                community           x86_64
/usr/lib/python2.7/site-packages/_psutil_posix.so          py-psutil           v3.3                community           x86
/usr/lib/lua/5.3/rex_posix.so                              lua5.3-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.3/rex_posix.so.2.8                          lua5.3-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.2/rex_posix.so                              lua5.2-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.2/rex_posix.so.2.8                          lua5.2-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.1/rex_posix.so                              lua5.1-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.1/rex_posix.so.2.8                          lua5.1-rex-posix    v3.3                main                armhf
/usr/lib/lua/5.1/posix.so                                  lua5.1-posix        v3.3                main                armhf
/usr/lib/lua/5.2/posix.so                                  lua5.2-posix        v3.3                main                armhf
/usr/lib/lua/5.3/posix.so                                  lua5.3-posix        v3.3                main                armhf
/usr/lib/lua/5.2/rex_posix.so                              lua5.2-rex-posix    v3.3                main                x86_64
/usr/lib/lua/5.2/rex_posix.so.2.8                          lua5.2-rex-posix    v3.3                main                x86_64
/usr/lib/lua/5.3/posix.so                                  lua5.3-posix        v3.3                main                x86_64
/usr/lib/lua/5.3/rex_posix.so                              lua5.3-rex-posix    v3.3                main                x86_64
/usr/lib/lua/5.3/rex_posix.so.2.8                          lua5.3-rex-posix    v3.3                main                x86_64
```

```console
$ apk-file bin/file
FILE                                                                  PACKAGE             BRANCH              REPOSITORY          ARCHITECTURE
/usr/lib/erlang/lib/gs-1.6/examples/ebin/file_dialog.beam             erlang-gs           edge                community           armhf
/usr/lib/erlang/lib/kernel-4.2/ebin/file.beam                         erlang-kernel       edge                community           armhf
/usr/lib/erlang/lib/kernel-4.2/ebin/file_io_server.beam               erlang-kernel       edge                community           armhf
/usr/lib/erlang/lib/kernel-4.2/ebin/file_server.beam                  erlang-kernel       edge                community           armhf
/usr/lib/erlang/lib/stdlib-2.8/ebin/filelib.beam                      erlang-stdlib       edge                community           armhf
/usr/lib/erlang/lib/stdlib-2.8/ebin/file_sorter.beam                  erlang-stdlib       edge                community           armhf
/usr/lib/erlang/lib/stdlib-2.8/ebin/filename.beam                     erlang-stdlib       edge                community           armhf
/usr/lib/erlang/lib/gs-1.6/examples/ebin/file_dialog.beam             erlang-gs           edge                community           x86
/usr/lib/erlang/lib/kernel-4.2/ebin/file.beam                         erlang-kernel       edge                community           x86
/usr/lib/erlang/lib/kernel-4.2/ebin/file_server.beam                  erlang-kernel       edge                community           x86
/usr/lib/erlang/lib/kernel-4.2/ebin/file_io_server.beam               erlang-kernel       edge                community           x86
/usr/lib/erlang/lib/stdlib-2.8/ebin/file_sorter.beam                  erlang-stdlib       edge                community           x86
/usr/lib/erlang/lib/stdlib-2.8/ebin/filename.beam                     erlang-stdlib       edge                community           x86
/usr/lib/erlang/lib/stdlib-2.8/ebin/filelib.beam                      erlang-stdlib       edge                community           x86
/usr/lib/erlang/lib/gs-1.6/examples/ebin/file_dialog.beam             erlang-gs           edge                community           x86_64
/usr/lib/erlang/lib/kernel-4.2/ebin/file.beam                         erlang-kernel       edge                community           x86_64
/usr/lib/erlang/lib/kernel-4.2/ebin/file_io_server.beam               erlang-kernel       edge                community           x86_64
/usr/lib/erlang/lib/kernel-4.2/ebin/file_server.beam                  erlang-kernel       edge                community           x86_64
/usr/lib/erlang/lib/stdlib-2.8/ebin/filename.beam                     erlang-stdlib       edge                community           x86_64
/usr/lib/erlang/lib/stdlib-2.8/ebin/filelib.beam                      erlang-stdlib       edge                community           x86_64
/usr/lib/erlang/lib/stdlib-2.8/ebin/file_sorter.beam                  erlang-stdlib       edge                community           x86_64
/usr/bin/file                                                         file                edge                main                armhf
/usr/bin/file                                                         file                edge                main                x86_64
/usr/bin/file                                                         file                edge                main                x86
/usr/bin/fileinfo                                                     leptonica           edge                testing             armhf
/usr/lib/erlang/lib/gs-1.5.16/examples/ebin/file_dialog.beam          erlang17-gs         edge                testing             armhf
/usr/lib/ruby/gems/2.2.0/gems/hoe-3.13.1/template/bin/file_name.erb   ruby-hoe            edge                testing             armhf
/usr/lib/erlang/lib/stdlib-2.4/ebin/filename.beam                     erlang17-stdlib     edge                testing             armhf
/usr/lib/erlang/lib/stdlib-2.4/ebin/file_sorter.beam                  erlang17-stdlib     edge                testing             armhf
/usr/lib/erlang/lib/stdlib-2.4/ebin/filelib.beam                      erlang17-stdlib     edge                testing             armhf
/usr/lib/erlang/lib/kernel-3.2/ebin/file.beam                         erlang17-kernel     edge                testing             armhf
/usr/lib/erlang/lib/kernel-3.2/ebin/file_server.beam                  erlang17-kernel     edge                testing             armhf
/usr/lib/erlang/lib/kernel-3.2/ebin/file_io_server.beam               erlang17-kernel     edge                testing             armhf
/usr/lib/erlang/lib/kernel-3.2/ebin/file.beam                         erlang17-kernel     edge                testing             x86_64
/usr/lib/erlang/lib/kernel-3.2/ebin/file_io_server.beam               erlang17-kernel     edge                testing             x86_64
/usr/lib/erlang/lib/kernel-3.2/ebin/file_server.beam                  erlang17-kernel     edge                testing             x86_64
/usr/bin/fileinfo                                                     leptonica           edge                testing             x86_64
/usr/lib/erlang/lib/gs-1.5.16/examples/ebin/file_dialog.beam          erlang17-gs         edge                testing             x86_64
/usr/lib/erlang/lib/stdlib-2.4/ebin/filename.beam                     erlang17-stdlib     edge                testing             x86_64
/usr/lib/erlang/lib/stdlib-2.4/ebin/filelib.beam                      erlang17-stdlib     edge                testing             x86_64
/usr/lib/erlang/lib/stdlib-2.4/ebin/file_sorter.beam                  erlang17-stdlib     edge                testing             x86_64
/usr/lib/erlang/lib/kernel-3.2/ebin/file.beam                         erlang17-kernel     edge                testing             x86
/usr/lib/erlang/lib/kernel-3.2/ebin/file_server.beam                  erlang17-kernel     edge                testing             x86
/usr/lib/erlang/lib/kernel-3.2/ebin/file_io_server.beam               erlang17-kernel     edge                testing             x86
/usr/bin/fileinfo                                                     leptonica           edge                testing             x86
/usr/lib/erlang/lib/gs-1.5.16/examples/ebin/file_dialog.beam          erlang17-gs         edge                testing             x86
/usr/lib/erlang/lib/stdlib-2.4/ebin/file_sorter.beam                  erlang17-stdlib     edge                testing             x86
/usr/lib/erlang/lib/stdlib-2.4/ebin/filename.beam                     erlang17-stdlib     edge                testing             x86
/usr/lib/erlang/lib/stdlib-2.4/ebin/filelib.beam                      erlang17-stdlib     edge                testing             x86
/usr/share/doc/swish-e/examples/prog-bin/file.pl                      swish-e-doc         edge                main                armhf
```

```
$ ./apk-file --output=yaml -arch=x86_64 version.h
```
