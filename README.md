# Jxbscare

## Overview
Jxbscare <small>(I censored the slur for you)</small> is a CV-to-PDF generator using TOML and Maroto and is implemented in Go.

### Disclaimer
Note this project is pretty experimental at the moment and has only been really tested with A4 papersize. *It may also break if you look at it wrong and requires thorough future testing.*
It also used consist of a custom DSL made with ANTLR but that has since been scrapped in favor of just using TOML and hardcoding the CV layout. You can find the remains of my grandiose failure in the git history :DDD

## Installation
I don't believe in distributing executables, you will have to build it from source.
```
git clone https://github.com/wsand02/jxbscare
cd jxbscare
go build
```

## Usage
```
./jxbscare <input> <output> 
```

## Example
coming soon...


## License
Jxbscare is licensed under the [MIT License](LICENSE).