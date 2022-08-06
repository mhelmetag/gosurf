# GoSurf

CLI Client for Surfline

## Usage

`gosurf` can read forecasts and tides for all places Surfline services from the command line (for Windows, Linux and Mac).

### Forecasts

To get a forecast (North Los Angeles, CA, USA):

```sh
$ gosurf f -r 58581a836630e24c44878fd5
Fetching 3 day(s) of forecasts for North Los Angeles...
+----------+-------------+--------------+-----------+----------------+
|   DATE   | TIME OF DAY |    RATING    |   RANGE   |    FORECAST    |
+----------+-------------+--------------+-----------+----------------+
| 8/6/2022 | AM          | Fair         | 2.0-3.0ft | Thigh to waist |
|          | PM          | Poor to Fair | 2.0-3.0ft | Thigh to waist |
| 8/7/2022 | AM          | Fair         | 2.0-3.0ft | Thigh to waist |
|          | PM          | Poor to Fair | 2.0-3.0ft | Knee to thigh  |
| 8/8/2022 | AM          | Poor to Fair | 1.0-2.0ft | Knee to thigh  |
|          | PM          | Poor         | 1.0-2.0ft | Knee to thigh  |
+----------+-------------+--------------+-----------+----------------+
```

Or to get a forecast for a different subregion (specifically Ventura, CA, USA) and day range:

```sh
$ gosurf f -r 58581a836630e24c4487900c -d 5
Fetching 5 day(s) of forecasts for Ventura...
+-----------+-------------+--------------+-----------+----------------+
|   DATE    | TIME OF DAY |    RATING    |   RANGE   |    FORECAST    |
+-----------+-------------+--------------+-----------+----------------+
| 8/6/2022  | AM          | Poor to Fair | 2.0-3.0ft | Thigh to waist |
|           | PM          | Poor to Fair | 2.0-3.0ft | Thigh to waist |
| 8/7/2022  | AM          | Poor to Fair | 2.0-3.0ft | Thigh to waist |
|           | PM          | Poor to Fair | 2.0-3.0ft | Thigh to waist |
| 8/8/2022  | AM          | Poor to Fair | 1.0-2.0ft | Knee to thigh  |
|           | PM          | Poor to Fair | 1.0-2.0ft | Knee to thigh  |
| 8/9/2022  | AM          | Poor to Fair | 1.0-2.0ft | Knee to thigh  |
|           | PM          | Poor         | 1.0-2.0ft | Knee to thigh  |
| 8/10/2022 | AM          | Poor         | 0.5-1.0ft | Shin to knee   |
|           | PM          | Poor         | 0.5-1.0ft | Shin to knee   |
+-----------+-------------+--------------+-----------+----------------+
```

Be sure to use the `subregion` ID (if coming from Surfline or their API directly) or the `TYPEID` or ID (if coming from the CLI; either solo search or interactive search).

### Tide

To get the tides for Solimar Beach, CA, USA:

```sh
$ gosurf t -s 5842041f4e65fad6a770895f -d 3
Fetching 3 day(s) of tides for Solimar...
+----------+-------+-------------+---------+
|   DATE   | TIME  | DESCRIPTION | HEIGHT  |
+----------+-------+-------------+---------+
| 8/6/2022 | 05:57 | High        | 2.95ft  |
|          | 09:58 | Low         | 2.53ft  |
|          | 16:59 | High        | 5.48ft  |
| 8/7/2022 | 00:55 | Low         | 0.43ft  |
|          | 07:33 | High        | 3.22ft  |
|          | 11:24 | Low         | 2.76ft  |
|          | 18:02 | High        | 5.91ft  |
| 8/8/2022 | 01:51 | Low         | -0.30ft |
|          | 08:29 | High        | 3.58ft  |
|          | 12:40 | Low         | 2.69ft  |
|          | 19:01 | High        | 6.40ft  |
+----------+-------+-------------+---------+
```

Be sure to use the `spot` ID (if coming from Surfline or their API directly) or the `TYPEID` (if coming from the CLI).

### Search

This is for searching Surfline's taxonomy tree. I recommend only using a maxDepth of 0 (default) or 1.

The default is the top level of the tree, Earth:

```sh
$ gosurf s
+--------------------------+---------+--------+---------------+
|            ID            |  TYPE   | TYPEID |     NAME      |
+--------------------------+---------+--------+---------------+
| 58f7f00ddadb30820bb69bbc | geoname | N/A    | Africa        |
| 58f7ed51dadb30820bb38791 | geoname | N/A    | North America |
| 58f7eef9dadb30820bb5626e | geoname | N/A    | Oceania       |
| 58f7eef1dadb30820bb556be | geoname | N/A    | Asia          |
| 58f7eef8dadb30820bb5601b | geoname | N/A    | Europe        |
| 58f7eef5dadb30820bb55cba | geoname | N/A    | South America |
+--------------------------+---------+--------+---------------+
```

You can then work your way down by passing the next level into the search command like `gosurf s -t 58f7ed51dadb30820bb38791` (for North America) and so on.

To get the records contained in Ventura County (using a max depth of 1 since I'm looking for spots specifically):

```sh
$ gosurf s -t 58f7ed58dadb30820bb38f8b -d 1
+--------------------------+---------+--------------------------+------------------------+
|            ID            |  TYPE   |          TYPEID          |          NAME          |
+--------------------------+---------+--------------------------+------------------------+
| 58f7ed59dadb30820bb39233 | geoname | N/A                      | Casa Conejo            |
| 58f7ed58dadb30820bb38f96 | geoname | N/A                      | Ventura                |
| 58f7edbcdadb30820bb3fd33 | geoname | N/A                      | Oxnard Shores          |
| 58f7edc0dadb30820bb401ff | spot    | 5842041f4e65fad6a770895f | Solimar                |
| 58f80a9ddadb30820bd12fce | spot    | 584204214e65fad6a7709cfd | C St. Overview         |
| 58f80a72dadb30820bd0ff32 | spot    | 584204204e65fad6a77096b1 | Ventura Point          |
| 58f7edbddadb30820bb3fe4b | spot    | 5842041f4e65fad6a7708957 | Pitas Point            |
| 58f7f229dadb30820bb94b7d | spot    | 584204204e65fad6a770904d | Mondos                 |
| 58f7edbddadb30820bb3ff16 | spot    | 5842041f4e65fad6a7708959 | Emma Wood              |
| 59c1970edadb30820b1d5a7f | spot    | 59c1970dbb6f23001cd20dd7 | Ventura Point Overview |
| 58f7edbfdadb30820bb4015d | spot    | 5842041f4e65fad6a770895e | Summer Beach           |
| 58f7ed5fdadb30820bb39978 | spot    | 5842041f4e65fad6a7708828 | C St.                  |
| 58f7ed58dadb30820bb38f9e | spot    | 5842041f4e65fad6a770880d | Gold Coast Beachbreaks |
| 58f7ed58dadb30820bb39099 | spot    | 5842041f4e65fad6a7708811 | Ventura Harbor         |
| 58f7edc3dadb30820bb404f6 | spot    | 5842041f4e65fad6a7708963 | Ventura Overhead       |
| 58f7edbcdadb30820bb3fd40 | spot    | 5842041f4e65fad6a770894c | Oxnard                 |
+--------------------------+---------+--------------------------+------------------------+
```

### Search Interactive

This is for interactively searching Surfline's taxonomy tree a bit more easily (similar to the site). Basically I take care of navigating the tree (correct depth and fetching what's next) and you just select as you go with enter.

There are two choices (default is subregion):

- `gosurf si -t subregion` mimics the 'Forecasts' tree
- `gosurf si -t spot` mimics the 'Cams & Reports' tree

```sh
$ gosurf si
Use the arrow keys to navigate: ↓ ↑ → ←  and / toggles search
? Select Taxonomy:
  ▸ Africa (58f7f00ddadb30820bb69bbc)
    Asia (58f7eef1dadb30820bb556be)
    Europe (58f7eef8dadb30820bb5601b)
    North America (58f7ed51dadb30820bb38791)
↓   Oceania (58f7eef9dadb30820bb5626e)
```

As the hints say: navigation is done using the up and down keys, forward slash (`/`) enters search mode for faster navigation and selecting is done with enter.

## Installation

### Binary

Download the right binary (for example, `gosurf_darwin_arm64` for Mac with M1) and move it to somewhere in your `$PATH` (so that it can be loaded into your command line environment).

For Mac, you might do something like this to download `gosurf` 3.0.0:

```sh
curl -L https://github.com/mhelmetag/gosurf/releases/download/3.0.0/gosurf_darwin_arm64 -o /usr/local/bin/gosurf
chmod a+x /usr/local/bin/gosurf
```

### From Source

```sh
go get https://github.com/mhelmetag/gosurf
```

And then either have your `$GOPATH/bin` in your `$PATH` or move that file into your `$PATH`.

### Homebrew from Source

If you prefer having homebrew do it for you:

```sh
brew tap mhelmetag/tap
brew install gosurf
```
