# PANGOLIN
## Introduction
My first approach to IT was in 1982 when my brother brought home a ZX Sinclair Spectrum 16K.
I was fascinated by its games (at the time, I was playing a game called 'Bruce Lee') and decided I wanted to create my own. For that reason, I started reading the ZX Manual.

In that manual, I found a game called 'Pangolin': it was a simple animal guessing game. The PC was asking simple questions, expecting a Yes or No answer and, after a few answers, was trying to guess your animal.

What fascinated me was that the game could learn: when he was wrong, it asked a few questions about your animal so that next time it could guess it.

This repo is a small tribute to that game: I wrote it using the go lang.
I kept the same interaction: everything is from the command line and is very simple.

The only difference with the original is that I added some persistence to make the game remember what it has learned.

## The game
The game starts by asking you to think about an animal.
Then it will ask you questions until it tries to guess it.
If it isn't able to guess it, it will ask you a question that would allow him to distinguish your animal so that next time he will know about it.

Accepted answers are `yes` and `no` and their shorter counterpart `y` and `n`.

## How it works
The game's memory is a binary tree, where each cell has a `yes` and a `no` branch.
Answers are the tree leaves (`yes` and `no` are `nil`), while the questions are the tree nodes.
When a new animal is added, the game will ask you a question about the new animal and its answer (`yes` or `no`) so that the bot will appropriately insert the question into the binary tree.

## Known issues
From time to time, the tree might need to be balanced. No code is provided to perform such an operation.
