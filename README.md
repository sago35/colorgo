# colorgo
colorize stdout by regex

## Demo

    $ git help | colorgo  diff Red  lo\w* Grean

![](https://raw.githubusercontent.com/wiki/sago35/colorgo/images/demo.png)

Japanese is ok.

    $ subst /? | colorgo -i cp932 "ドライブ|パス" Red

![](https://raw.githubusercontent.com/wiki/sago35/colorgo/images/demo2.png)

## Usage

```
USAGE:
    gocolor.go [OPTIONS] [ REGEX  COLOR ]*

OPTIONS:
    --input, -i "utf8"  input encoding
    --output, -o "utf8" output encoding
    --help, -h          show help
    --version, -v       print the version

OTHER:
    Default encoding is utf8 for input and output.
    Supported encodings are below.

        cp932 shijtjis eucjp utf8
        and encodings supported by mahonia

    REGEX : Regular expression

    COLOR : Color name

        None
        Black
        Red
        Green
        Yellow
        Blue
        Magenta
        Cyan
        White
```

## Install

    go get github.com/sago35/colorgo

## Contribution

TBD.

## Licence

[MIT](http://opensource.org/licenses/mit-license.php)

## Author

[sago35](https://github.com/sago35)
