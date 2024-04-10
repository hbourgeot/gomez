# Gomez, your loyal Tool Manager

This project is a command-line tool written in Go that helps to install different programming tools. It supports at April 9, 2024 the following tools:

* Node.js Version Manager (nvm)
* Fast Node Manager (fnm)
* Python Version Manager (pyenv)
* SDKMAN! (for all Java-related ecosystems)
* Cargo (Rust package manager)

I hope very soon to add more tools to the list. Any help is welcome in the issues or pull requests sections.

## Getting Started

### Prerequisites

- None! Download the binary on the [releases page](https://github.com/hbourgeot/gomez).
- Linux (MacOS and Windows are not supported yet).

### Installing

If you want, you can add the binary to your PATH. For example, you can move the binary to `/usr/local/bin` and then, add the permissions:

```bash
chmod +x /usr/local/bin/gmz
```

With this, you are ready to **go**!

**Note 1**: You don't need to use `sudo` command.

**Note 2**: The Gomez CLI is gmz for short.

## Usage

Gomez has a simple interface. You can see the available commands by running `gmz --help` or also `gmz -h`. This is the output:

```text
Gomez, your loyal helper for install some programming tools, like: Fast Node Manager (fnm), Pyenv,
Node Version Manager (nvm), Cargo and SDKMan, for now.

Why I'm developing this CLI? For automatize the installations of the most programming tools needed
for programmers. With a few keys press. Don't go to the github page of any of these, Gomez can help
you with that. Only write the tool that you need and Gomez will:

1. Install the tool.
2. Configure the tool for the shell of your choice.
3. Install the latest stable version of the programming language (or provide a version).

"Gomez, I'm to serve him, my lord."

Usage:
  gmz [command]

Available Commands:
  cargo       Cargo, the package manager for Rust
  completion  Generate the autocompletion script for the specified shell
  fnm         Install Fast Node Manager (fnm)
  help        Help about any command
  nvm         A brief description of your command
  pyenv       Install Python Version Management
  sdkm        A brief description of your command

Flags:
  -b, --bash   Install for Bash (default true)
  -f, --fish   Install for fish
  -h, --help   help for gmz
  -z, --zsh    Install for ZSH

Use "gmz [command] --help" for more information about a command.
```

For example, if you want to install fnm, run:

```bash
gmz fnm -l -z
```

This command will:

1. Install fnm.
2. Configure fnm for ZSH.
3. Install the latest LTS version of Node.js.


## License

This project is licensed under the MIT License - see the `LICENSE` file for details.