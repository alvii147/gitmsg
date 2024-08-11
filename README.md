[![Free Palestine](https://hinds-banner.vercel.app/free-palestine?variant=crimson)](https://www.pcrf.net/)

[![Genocide Watch](https://hinds-banner.vercel.app/genocide-watch?variant=crimson)](https://www.pcrf.net/)

<p align="center">
    <img alt="gitmsg logo" src="img/gitmsg_logo.png" width=600 />
</p>

<p align="center">
<strong><i>gitmsg</i></strong> is an application for writing git commit messages according to the <a href="https://www.midori-global.com/blog/2018/04/02/git-50-72-rule">git 50/72 rule</a>.
</p>

<div align="center">

[![GitHub CI](https://img.shields.io/github/actions/workflow/status/alvii147/gitmsg/github-ci.yml?branch=master&label=Github%20CI&logo=github)](https://github.com/alvii147/gitmsg/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/alvii147/gitmsg/cli)](https://goreportcard.com/report/github.com/alvii147/gitmsg/cli) [![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

</div>

## Installation

Clone the repository:

```bash
git clone https://github.com/alvii147/gitmsg.git
```

Setup *gitmsg*:

```bash
cd gitmsg/
./setup.sh
```

Set up alias to script by adding the following to `~/.bashrc` or `~/.bash_profile`:

```bash
alias gitmsg="/absolute/path/to/gitmsg/gitmsg.sh"
```

## Usage

```
gitmsg helps write git commit messages based on the git 50/72 rule.

Usage:
        gitmsg [OPTIONS]

Options:
        -g, run in GUI mode
        -c, run in CLI mode
        -h, show help
```

### GUI Mode

```bash
gitmsg -g
```

![GUI screenshot](img/gui_screenshot.png)

### CLI mode

```bash
gitmsg -c
```

![CLI screenshot](img/cli_screenshot.png)
