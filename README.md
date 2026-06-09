# Invisible

Embed noise or hidden message to text.

![Coverage](https://raw.githubusercontent.com/kitsuyui/octocov-central/main/badges/kitsuyui/invisible/coverage.svg)
[![Github All Releases](https://img.shields.io/github/downloads/kitsuyui/invisible/total.svg)](https://github.com/kitsuyui/invisible/releases/latest)


## Install

Download prebuilt binaries from the
[latest GitHub release](https://github.com/kitsuyui/invisible/releases/latest).

To install from source with Go:

```console
$ go install github.com/kitsuyui/invisible@latest
```

For local development builds, run `go build -o invisible .` from the repository
root. The `bin/build.sh` script is available for release-style cross-platform
builds.

## Add noise

```console
$ invisible add-noise < example/plain.txt > example/noised.txt
```

It looks normal. But the size is grown from 446 bytes to 1070 bytes by noise.

> **Note**: `add-noise` output is non-deterministic — the invisible characters inserted vary on each run.
> Use `--seed <N>` to get reproducible output (e.g. `--seed 42`). `encode` and `decode` are always deterministic.

`add-noise` accepts these options:

- `--frequency <rate>` / `-f <rate>`: noise insertion probability. The default
  is `0.5`. Negative values are rejected.
- `--noise-size <count>` / `-s <count>`: maximum number of invisible characters
  inserted at once. The default is `1`.

```
L⁡o​rem⁣ ⁢ip​sum ⁡d​o﻿lor⁡ ⁠sit⁡ a​met⁣,﻿ ⁠c‌ons⁠ect﻿et⁢ur⁠ ⁣a⁤di⁡p⁣isci‌n‌g e‌li⁢t,‌ ⁢se​d d⁡o‌ e​i⁡u⁠smo⁢d⁡ ​t⁠em⁡p﻿o​r i﻿n⁠cididu‌n⁡t ⁢u⁡t labo⁠re ​e⁡t​ ‌dolo⁣re​ mag‌n⁡a⁠ al⁣i​qua﻿.⁡ ⁤Ut en​im ad⁤ ‌mi﻿nim ⁡v⁢e⁣ni‌am, qui⁤s ⁢n⁡os⁣trud‌ ⁢e⁣xe‌rc⁠it﻿ati⁣on ‌u⁢l‌l⁢a​mc⁠o⁠ laboris n​i⁤si⁡ ​ut⁠ ⁢a⁠liq‌u​i⁣p e‌x e⁠a⁡ c‌o﻿mmo⁠do⁤ c﻿on⁤s⁣e⁠q⁡uat⁠. Duis aute iru⁠r⁤e ⁤d‌o﻿lor ⁤i⁣n⁢ ⁡r⁠e​pr​eh⁡e﻿nd⁤e⁠rit⁣ i⁣n⁣ vo⁡l⁣uptate⁠ v‌el⁣it ​e⁢s⁡s⁡e⁠ ⁠cil⁢l‌u​m dol​ore⁢ ​e⁢u ⁠f​u⁤g﻿i​a⁢t‌ ⁠n⁠ulla par⁣i﻿a﻿t​u⁠r​. ‌Except‌e​ur⁢ ⁤sin​t⁠ ﻿o﻿ccaecat cu​pid⁢a⁠tat ⁤n⁡o⁣n⁤ ⁠proide‌nt, ​sunt⁤ ⁡i‌n﻿ c﻿u⁡lp⁤a ⁢qui‌ off​i﻿c⁢i‌a dese​ru﻿n​t ‌mo‌l⁢li⁡t⁣ ​a⁢ni⁤m⁣ ​id ⁡est ﻿l⁤abo⁡r​u⁢m.

```


## Encode

```console
$ invisible encode -m 'Hello, World!' < example/plain.txt > example/embedded.txt
```

## Encoded text

It looks normal. But the hidden message is in.

```
L⁠o⁠r​e⁤m⁠ ⁣i⁣p⁢s⁡u⁡m​ ⁤d﻿o⁢l⁣o⁢r‌ ​s​i⁣t⁡ ⁣a⁣m﻿e⁡t⁢,⁢ ⁤c⁤o‌n⁢s⁢e‌c​t⁠e​t​u​r​ ​a⁠d⁠i​p⁤i⁠s⁣c⁣i⁢n⁡g⁡ ​e⁤l﻿i⁢t⁣,⁢ ‌s​e​d⁣ ⁡d⁣o⁣ ﻿e⁡i⁢u⁢s⁤m⁤o‌d⁢ ⁢t‌e​m⁠p​o​r​ ​i​n⁠c⁠i​d⁤i⁠d⁣u⁣n⁢t⁡ ⁡u​t⁤ ﻿l⁢a⁣b⁢o‌r​e​ ⁣e⁡t⁣ ⁣d﻿o⁡l⁢o⁢r⁤e⁤ ‌m⁢a⁢g‌n​a⁠ ​a​l​i​q​u⁠a⁠.​ ⁤U⁠t⁣ ⁣e⁢n⁡i⁡m​ ⁤a﻿d⁢ ⁣m⁢i‌n​i​m⁣ ⁡v⁣e⁣n﻿i⁡a⁢m⁢,⁤ ⁤q‌u⁢i⁢s‌ ​n⁠o​s​t​r​u​d⁠ ⁠e​x⁤e⁠r⁣c⁣i⁢t⁡a⁡t​i⁤o﻿n⁢ ⁣u⁢l‌l​a​m⁣c⁡o⁣ ⁣l﻿a⁡b⁢o⁢r⁤i⁤s‌ ⁢n⁢i‌s​i⁠ ​u​t​ ​a​l⁠i⁠q​u⁤i⁠p⁣ ⁣e⁢x⁡ ⁡e​a⁤ ﻿c⁢o⁣m⁢m‌o​d​o⁣ ⁡c⁣o⁣n﻿s⁡e⁢q⁢u⁤a⁤t‌.⁢ ⁢D‌u​i⁠s​ ​a​u​t​e⁠ ⁠i​r⁤u⁠r⁣e⁣ ⁢d⁡o⁡l​o⁤r﻿ ⁢i⁣n⁢ ‌r​e​p⁣r⁡e⁣h⁣e﻿n⁡d⁢e⁢r⁤i⁤t‌ ⁢i⁢n‌ ​v⁠o​l​u​p​t​a⁠t⁠e​ ⁤v⁠e⁣l⁣i⁢t⁡ ⁡e​s⁤s﻿e⁢ ⁣c⁢i‌l​l​u⁣m⁡ ⁣d⁣o﻿l⁡o⁢r⁢e⁤ ⁤e‌u⁢ ⁢f‌u​g⁠i​a​t​ ​n​u⁠l⁠l​a⁤ ⁠p⁣a⁣r⁢i⁡a⁡t​u⁤r﻿.⁢ ⁣E⁢x‌c​e​p⁣t⁡e⁣u⁣r﻿ ⁡s⁢i⁢n⁤t⁤ ‌o⁢c⁢c‌a​e⁠c​a​t​ ​c​u⁠p⁠i​d⁤a⁠t⁣a⁣t⁢ ⁡n⁡o​n⁤ ﻿p⁢r⁣o⁢i‌d​e​n⁣t⁡,⁣ ⁣s﻿u⁡n⁢t⁢ ⁤i⁤n‌ ⁢c⁢u‌l​p⁠a​ ​q​u​i​ ⁠o⁠f​f⁤i⁠c⁣i⁣a⁢ ⁡d⁡e​s⁤e﻿r⁢u⁣n⁢t‌ ​m​o⁣l⁡l⁣i⁣t﻿ ⁡a⁢n⁢i⁤m⁤ ‌i⁢d⁢ ‌e​s⁠t​ ​l​a​b​o⁠r⁠u​m⁤.⁠
⁣⁣⁢⁡⁡​⁤﻿⁢⁣⁢‌​​⁣⁡⁣⁣﻿⁡⁢⁢⁤⁤‌⁢⁢‌​⁠​​​​​
```

## Decode

```console
$ invisible decode < example/embedded.txt
Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!
```

## Encoding format

Encoded messages start with an invisible `v1` format marker. The decoder still
accepts older markerless payloads for compatibility, and rejects payloads that
use the reserved marker prefix with an unsupported version.

## Version

```console
$ invisible version
```

Prints the installed version. Returns `(devel)` when built from source and `(unknown)` when build information is unavailable.

## Development

Install [lefthook](https://github.com/evilmartians/lefthook) and run the following command once after cloning:

```console
$ lefthook install
```

This sets up the following local Git hooks that mirror the CI checks:

- **pre-commit**: runs `go vet ./...` to catch common issues before each commit.
- **pre-push**: runs `go vet ./...` and `go test ./...` to verify correctness before pushing.

CI still runs the full test suite on every pull request and push to `main` — the hooks only bring that feedback earlier, during local development.

## LICENSE

### Source

The 3-Clause BSD License. See also LICENSE file.
