# Jxbscare

## Overview
Jxbscare <small>(I censored the slur for you)</small> is a CV-to-PDF generator consisting of a mediocre DSL made with ANTLR and implemented in Go. These so called *jxb* files are then turned into PDFs with maroto.

## Installation
I don't believe in distributing executables, you will have to build it from source.
```
git clone https://github.com/wsand02/jxbscare
cd jxbscare
go build
```

## Usage
```
./jxbscare -p a4 [jxb file here]
```

## Example jxb
<small>Any resemblance to actual events, locales, or persons is purely coincidental.</small>
```
cv Skitjobb
name Anders Andersson
selfie IMG_1573.jpeg

begin contact
address Gatan 200F, 66633, Atlanta
phone 77777777777
email example@example.com
website example.com
end contact

row insert contact
insert selfie

row insert name
```

For more info you can read the [grammar](parser/JXB.g4).

### Example output
coming soon...

## License
Jxbscare is licensed under the [MIT License](LICENSE).