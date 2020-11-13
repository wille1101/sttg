# STTG

![gif](https://raw.githubusercontent.com/wille1101/gifs/master/animation.gif)

[![Go Report Card](https://goreportcard.com/badge/github.com/wille1101/sttg)](https://goreportcard.com/report/github.com/wille1101/sttg)

SVT Text-Tv Go (STTG) är en TUI-klient för SVTs text-tv skriven i Go.

## Installation

### Linux

#### Ladda ner under Releases, gör filen exekverbar och lägg i din bin-mapp

(Exempel, din bin-mapp kan vara någon annanstans)

```
wget https://github.com/wille1101/sttg/releases/download/1.0.1/sttg_v1.0.1_linux
chmod +x sttg_v1.0.1_linux && mv sttg_v1.0.1_linux ~/.local/bin/sttg
```

eller

#### Kompilera från källkoden själv (Behöver ha Go installerat)

```
git clone https://github.com/wille1101/sttg.git
cd sttg && go install
```

### Windows

#### Ladda ner under Releases och starta .exe-filen

eller

#### Kompilera från källkoden själv (Behöver ha Go installerat)

Ladda ner repon som en .zip, packa upp innehållet och navigera till den uppackade mappen i cmd. Skriv sen:
```
go install
```
Nu kan du starta programmet igenom att skriva `sttg` i cmd.

Om du vill ha en .exe-fil, navigera till den uppackade mappen och skriv:
```
go build
```

## Användning

### UI
Uppe till vänster skriver man in sidan man vill gå till efter att man har klickat `:` eller `i`. Standardsidan är 100.

Uppe till höger visas procenten kvar att skrolla på sidan innan du når botten. Om hela sidan syns visas 100%.

### Tangenter
- `h, l` eller  `vänster, höger`: Gå en sida åt vänster/höger
- `j, k` eller `ner, upp`:  Skrolla ner/upp på en sida
- `g, G`: Gå till toppen/botten av hela sidan
- `d, u`: Gå ner/upp en halvsida (Halva höjden av fönstret, inte SVT-sidan)
- `:` eller `i`:        Gå direkt till en sida
- `q` eller `Ctrl + c`: Stäng programmet
- `1-9`:            Gå direkt till en sida, med siffran du skriver in som första sidosiffra. 
- `H`:            Visar hjälpsidan

## Konfiguration
För att ändra standardtangenterna, samt andra värden som används, kan man skapa en `config.yml`-fil i programmets mapp, eller i:
#### Linux  
`~/.config/sttg/`

#### Windows
`C:\Users\ANVÄNDARE\AppData\Roaming\sttg\`

Standardtangenterna som nämns ovanför används när en `config.yml`-fil inte definerar en tangent, eller om den  inte finns överhuvudtaget.

### Exempelfil med alla värden satta till standard

```
---
#Definera tangenterna för STTG
#Varje handling kan tilldelas två tangenter samtidigt
Keys:
  #Skrolla upp på en sida
  Up:
  - k
  - up

  #Skrolla ner på en sida
  Down:
  - j
  - down

  #Gå en sida åt vänster
  Left:
  - h
  - left

  #Gå en sida åt höger
  Right:
  - l
  - right

  #Gå direkt till en sida
  SetPage:
  - i
  - ":"

  #Visa hjälpsidan
  GetHelp:
  - H
  - ""

  #Stäng programmet
  Quit:
  - q
  - "ctrl+c"

Page:
  #Standardsidan som visas när STTG startas
  DefPageNr: 100

```

### Exempelfil som ändrar navigeringen med Vim-tangenterna h, j, k, l till w, a, s, d

```
---
Keys:
  Up:
  - w
  - up
  Down:
  - s
  - down
  Left:
  - a
  - left
  Right:
  - d
  - right

```

## Inspiration
  https://github.com/voidcase/txtv
