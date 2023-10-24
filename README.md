# gophersay [![Go Reference](https://pkg.go.dev/badge/github.com/FraserChapman/gophersay.svg)](https://pkg.go.dev/github.com/FraserChapman/gophersay)

gophersay is a tiny command-line tool written in Go that displays a cute ASCII art of the go gopher along with a customizable message.

It takes any text as input and displays the message inside a speech bubble.

Inspired by the classic [cowsay](https://en.wikipedia.org/wiki/Cowsay)

## Usage

### Prerequisites
Make sure you have Go installed on your system.

### Installation
Clone the repository and navigate to the project directory:

```
git clone https://github.com/FraserChapman/gophersay
cd gophersay
```

Build the project:

```
go build -o gophersay .
```

or on windows...

```
go build -o gophersay.exe .
```

### Running the Program

```
./gophersay <your-message>
```

or on windows...

```
gophersay.exe <your-message>
```

Replace `<your-message>` with the text you want the gopher to say.

If no message is provided, the gopher will say `Don't panic!` by default.

### Examples

```
./gophersay Hello, Gopher!
```

This will display the following output:

```
+---------+
| Hello,  |
| Gopher! |
+---------+
  \
   \
    \    ,_---~~~~~----._
  _,,_,*^____      _____\˴ᐠ*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \ᐠ/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
```

To pass arguments containing spaces, ensuring they are treated as full lines of text, wrap the arguments in quotes. For example:
""

```
./gophersay "Six by nine. Forty two." "That's all there is."
```

This will display the following output:

```
+-------------------------+
| Six by nine. Forty two. |
| That's all there is.    |
+-------------------------+
  \
   \
    \    ,_---~~~~~----._
  _,,_,*^____      _____\˴ᐠ*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \ᐠ/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
```

## Contributing

This project is just a bit of fun...but if you do want to contribute and make it even better, feel free to fork the repository and submit a pull request. Your contributions are always welcome!

### Notes for Windows users...

[Why does Windows Defender think the compiled binary is infected?](https://go.dev/doc/faq#virus)