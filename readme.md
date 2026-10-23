<h1 align="center"> Number Guessing Game CLI </h1>

<p align="center">
A quick command line guessing game written in Golang.
</p>

<p align="center">Project URL: https://roadmap.sh/projects/number-guessing-game</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/palette/macchiato.png">
</p>

## What it does

The game picks a random number and you try to guess it. You choose a difficulty (Easy: 10
chances, Medium: 5, Hard: 3) and after each guess it tells you whether the number is greater or
less than your guess. At the end of the round it shows how long you took and starts a new round.

## Why it exists

It is the [Number Guessing Game](https://roadmap.sh/projects/number-guessing-game) project from the
roadmap.sh back-end roadmap, built in Go following the project's specification.

## Requirements

- Go 1.24.8 or higher
- git

## Installation

Linux:

```sh
git clone https://github.com/rillmind/numberGuessingGameCli.git
cd numberGuessingGameCli
go build -o numberGuessing ./cmd
sudo cp numberGuessing /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd numberGuessingGameCli
go build -o numberGuessing .\cmd\
```

The Windows block above clones `taskManagerCli` by mistake. The right URL is
`https://github.com/rillmind/number-guessing-game-cli.git` (the old name `numberGuessingGameCli`
redirects to it).

## How to test

There are no automated tests. `go vet ./...` checks the code; to play:

```sh
go build -o numberGuessing ./cmd
./numberGuessing
```

Check that:

- after choosing `3` (Hard) you get exactly 3 guesses, each answered with
  `Incorrent! The number is greater than N` or `... less than N`;
- a right guess prints `Congratulations! You guessed the correct number in N attempts`
  (N counts the wrong guesses before it);
- every round ends with `Time passed: ...` and the menu comes back with the next round number.

There is no option to quit: press `Ctrl+C`. The number is drawn with `rand.IntN(101)`, so it can
be anything from 0 to 100, not 1 to 100 as the menu says.
