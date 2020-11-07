# STTG

![gif](https://raw.githubusercontent.com/wille1101/gifs/master/animation.gif)

[![Go Report Card](https://goreportcard.com/badge/github.com/wille1101/sttg)](https://goreportcard.com/report/github.com/wille1101/sttg)

SVT Text-Tv Go (STTG) är en TUI-klient för SVTs text-tv skriven i Go.

## Installation

Ladda ner under Releases, gör filen exekverbar och lägg i din bin-mapp

(Exempel, din bin-mapp kan vara någon annanstans)

```
wget https://github.com/wille1101/sttg/releases/download/1.0.0/sttg_v1.0.0_linux
chmod +x sttg_v1.0.0_linux  && mv sttg_v1.0.0_linux ~/.local/bin/sttg
```

eller

Kompilera från källkoden själv (Behöver ha Go installerat)

```
git clone https://github.com/wille1101/sttg.git
cd sttg && go install
```

## Användning

### UI
Uppe till vänster skriver man in sidan man vill gå till efter att man har klickat `:` eller `i`. Standardsidan är 100.

Uppe till höger visas procenten kvar att skrolla på sidan innan du når botten. Om hela sidan syns visas 100%.

### Tangenter
- `h, l` eller  `vänster, höger`: Gå en sida åt vänster/höger
- `j, k` eller `ner, upp`:  Skrolla ner/upp på en sida
- `:` eller `i`:        Gå direkt till en sida
- `q` eller `Ctrl + c`: Stäng programmet
- `1`:            Gå direkt till Nyheter
- `2`:                  --||--    Ekonomi
- `3`:                  --||--    Sport
- `4`:                  --||--    Väder
- `H`:            Visar hjälpsidan

#### Inspiration
  https://github.com/voidcase/txtv
