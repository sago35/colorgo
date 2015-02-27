# colorgo
colorize stdout by regex

## Demo

    $ git help | colorgo  diff Red  lo\w* Grean

![](https://raw.githubusercontent.com/wiki/sago35/colorgo/images/demo.png)

## Usage

```
USAGE:
    gocolor.go [OPTIONS] [ REGEX  COLOR ]*

OPTIONS:
    --input, -i "utf8"  入力文字コード
    --output, -o "utf8" 出力文字コード
    --help, -h          show help
    --version, -v       print the version

OTHER:
    入出力文字コードは、デフォルト値は入力はutf8、出力はutf8となっています
    以下を設定可能です

        cp932 shijtjis eucjp utf8

    REGEXは、色づけを行う正規表現
    COLORは、以下の色の名前を指定可能です

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
