# Go Convert Layout

Portition of [ai/convert-layout](https://github.com/ai/convert-layout) to GoLang.

Simple go lib to convert from one keyboard layout to other. 

## Usage

```golang

var ru = converter.Create("ru")

fmt.Printf(ru.ToEng("руддщ")) // => hello
fmt.Printf(ru.FromEn("ghbdtn")) //=> привет
```

## Layouts

Currently supported keyboard layouts:

* Belarusian
* English (QWERTY, Dvorak, Colemak)
* German
* Kazakh
* Korean
* Russian
* Spanish
* Ukrainian
* Hebrew
* Persian

## Credits

[Andrey Sitnik](https://github.com/ai) the author of [ai/convert-layout](https://github.com/ai/convert-layout) library.