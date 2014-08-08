# Control iTerm2 from cli

`iterm2` is a command line tool to control iTerm2. You can change tab colors and titles with this command.


## Installation

Install via Homebrew

```
brew tap suin/suin
brew install iterm2-cli
```

## Usage

### Change tab color

You can change tab colors by just run `tab color $colorname` command like this:

```console
$ iterm2 tab color light-blue
```

140+ colors are available. Type `tab color` to show all available color names.


### Change tab title

Run `tab title $title` to change tab title.

```console
$ iterm2 tab title my project
```
