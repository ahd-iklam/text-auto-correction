# Text Completion and Editing Tool

## Introduction

This project is a simple text completion, editing, and auto-correction tool written in Go. It leverages functions from previous repositories to implement a variety of text modifications. The goal is to provide a utility that processes a given text file and outputs a modified version based on specified rules.

## Project Requirements

- The project is written in Go.
- Code adheres to good practices.
- Test files for unit testing are recommended.

## Features

The tool processes text files and applies the following modifications:

1. **Hexadecimal to Decimal Conversion:**
   - Every instance of `(hex)` replaces the preceding word (a hexadecimal number) with its decimal equivalent.
   - Example: `"1E (hex) files were added"` becomes `"30 files were added"`.

2. **Binary to Decimal Conversion:**
   - Every instance of `(bin)` replaces the preceding word (a binary number) with its decimal equivalent.
   - Example: `"It has been 10 (bin) years"` becomes `"It has been 2 years"`.

3. **Text Case Transformations:**
   - **Uppercase:** `(up)` converts the preceding word to uppercase.
     - Example: `"Ready, set, go (up)!"` becomes `"Ready, set, GO!"`.
   - **Lowercase:** `(low)` converts the preceding word to lowercase.
     - Example: `"I should stop SHOUTING (low)"` becomes `"I should stop shouting"`.
   - **Capitalized:** `(cap)` converts the preceding word to capitalized case.
     - Example: `"Welcome to the Brooklyn bridge (cap)"` becomes `"Welcome to the Brooklyn Bridge"`.

4. **Text Case Transformations with Word Count:**
   - If a number is specified, it transforms the specified number of preceding words.
     - **Uppercase:** `(up, <number>)`
       - Example: `"This is so exciting (up, 2)"` becomes `"This is SO EXCITING"`.
     - **Lowercase:** `(low, <number>)`
       - Example: `"I AM SHOUTING (low, 2)"` becomes `"I am shouting"`.
     - **Capitalized:** `(cap, <number>)`
       - Example: `"the brooklyn bridge (cap, 2)"` becomes `"The Brooklyn Bridge"`.

5. **Punctuation Adjustment:**
   - Standard punctuation marks (., ,, !, ?, :, ;) should be placed close to the preceding word and separated by a space from the following word.
     - Example: `"I was sitting over there ,and then BAMM !!"` becomes `"I was sitting over there, and then BAMM!!"`.
   - Special handling for groups of punctuation (e.g., `...`, `!?.`):
     - Example: `"I was thinking ... You were right"` becomes `"I was thinking... You were right"`.

6. **Quotation Marks:**
   - Single quotes should enclose the word(s) directly, with no spaces.
     - Example: `"I am exactly how they describe me: ' awesome '"` becomes `"I am exactly how they describe me: 'awesome'"`.
     - Example: `"As Elton John said: ' I am the most well-known homosexual in the world '"` becomes `"As Elton John said: 'I am the most well-known homosexual in the world'"`.

7. **Article Adjustment:**
   - Change "a" to "an" if the following word begins with a vowel or an 'h'.
     - Example: `"There it was. A amazing rock!"` becomes `"There it was. An amazing rock!"`.

## Usage

To use the tool, run the following commands:

1. **Prepare a text file with the content to be modified.**

   ```sh
   $ cat sample.txt
   it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
Run the Go program with the input and output file names as arguments.

2 Run the Go program with the input and output file names as arguments.
$ go run . sample.txt result.txt
3 Check the output file for the modified text.
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
Allowed Packages
The project only uses standard Go packages.

Contribution
Feel free to fork this repository and contribute by submitting pull requests. For major changes, please open an issue first to discuss what you would like to change.

