<p align='center'><img width="128" src='./vendor/noto-emoji-tangerine.png' alt="A tangerine emoji"/></p>
 <h1 align='center'>gitsrc</h1>

<p align="center">Conveniently pull from remote Git repositories you use often.</p>

<br /><br />

If you work with projects which depend on multiple independent services, you probably have to update them all regularly and know that a little help can come handy!

Usage
-----

Write the list of repositories you want to pull from in a file, one per line (e.g. `example.txt`):

```txt
github.com/gonzalo-bulnes/dice
github.com/gonzalo-bulnes/gitsrc
```

```bash
gitsrc update example.txt # that's all!
```

Installation
------------

[![Go Report Card](https://goreportcard.com/badge/github.com/gonzalo-bulnes/gitsrc)](https://goreportcard.com/report/github.com/gonzalo-bulnes/gitsrc)
[![Build Status](https://travis-ci.org/gonzalo-bulnes/gitsrc.svg?branch=master)](https://travis-ci.org/gonzalo-bulnes/gitsrc)

Binaries for official releases may be downloaded from the [releases page on GitHub](https://github.com/gonzalo-bulnes/gitsrc/releases).

If you want to compile it from source, try:

```bash
go get github.com/gonzalo-bulnes/gitsrc/...
```

For Unix/Linux users, you can install `gitsrc` using the following command. You may want to change the version number in the command below from `v0.1.0` to whichever version you want:

```bash
curl -sL -o /usr/local/bin/gitsrc \
    https://github.com/gonzalo-bulnes/gitsrc/releases/download/v0.1.0/gitsrc-linux-amd64 \
 && chmod +x /usr/local/bin/gitsrc
```

Contributing
------------

[![GoDoc](https://godoc.org/github.com/gonzalo-bulnes/gitsrc?status.svg)](https://godoc.org/github.com/gonzalo-bulnes/gitsrc)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-e7359e.svg?style=popout)](http://makeapullrequest.com)

Whether it is your first pull request or your 100th, the [contributing guidelines][contributing] are here to help you get started!

Please note that by participating in this project, you agree to abide by its [code of conduct]. That is true for pull requests, and also when participating in issues.

  [contributing]: ./CONTRIBUTING.md
  [code of conduct]: ./CODE_OF_CONDUCT.md


Credits
-------

The tangerine emoji in the header was rendered from an SVG that belongs to Google and [was published under the Apache License v2.0 as part of Noto Emoji](https://github.com/googlei18n/noto-emoji).

License
-------

    gitsrc
    Copyright (C) 2019 Gonzalo Bulnes Guilpain

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
