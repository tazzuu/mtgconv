# mtgconv

Magic the Gathering Deck list converter.

`mtgconv` is a CLI tool & Go library to convert deck lists between different input sources and output formats.

## Features

- Fetch a deck list from supported online platforms
- Convert a deck list into multiple output formats
- Search for decks on supported  online platforms and batch-convert the results
- Core package built in Go, designed for re-use as a library, and designed with a modular system for adding new input sources and output formats

## Supported Sources

- [moxfield.com/](https://moxfield.com/)

### Coming Soon

- [archidekt.com](https://archidekt.com)
- [Shiny card tracker .csv export](https://play.google.com/store/apps/details?id=io.getshiny.shiny&hl=en-US&pli=1)

## Supported Outputs

- `.dck` decklist format
  - includes "Compatibility Mode" support to output reduce card information to ensure support for various programs
- `.txt` plain text format
- `.json` serialized format

# Usage

See all the cli flags and options available;

```bash
mtgconv --help
mtgconv convert --help
mtgconv search --help
```

## Examples

Convert a Moxfield deck to .txt decklist format

```bash
mtgconv --user-agent "$MOXKEY" convert --output-format txt https://moxfield.com/decks/AZKbE6E6kUWW2zMfHR41sQ
```

Search for the top 100 Commander decks and save them as .dck deck list files

```bash
mtgconv --user-agent "$MOXKEY" search moxfield.com
```

- **NOTE**: querying Moxfield requires an API key

# Installation

## Download

Download a pre-compiled version for your system from the [Releases](https://github.com/tazzuu/mtgconv/releases) page.

## Build Locally

Building locally requires Go version 1.24.5+ to be [installed](https://go.dev/doc/install)

Build with the included Makefile recipe:

```bash
make build
```

# Project Layout

- `cmd/mtgconv`: cli entrypoint
- `pkg/mtgconv2/core`: core pipeline, configs, and types
- `pkg/mtgconv2/sources`: input source adapters
- `pkg/mtgconv2/outputs`: output handlers

# Planned Features

- web interface & hosted app instance
- static resource repository of retrieved & converted deck lists
- add more input sources
