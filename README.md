# GoSurf

CLI Client for Surfline

## Usage

`gosurf` can read forecasts and tides for all places Surfline services from the command line (for Windows, Linux and Mac).

### Forecasts

To get a forecast (North Los Angeles, CA, USA):

```sh
$ gosurf f -s 58581a836630e24c44878fd5
+-----------+-------------+--------------+-----------+------------------------+
|   DATE    | TIME OF DAY |    RATING    |   RANGE   |        FORECAST        |
+-----------+-------------+--------------+-----------+------------------------+
| 5/24/2021 | AM          | Fair         | 3.0-5.0ft | Waist to head high     |
| 5/24/2021 | PM          | Poor to Fair | 3.0-5.0ft | Waist to shoulder high |
| 5/25/2021 | AM          | Fair         | 3.0-4.0ft | Waist to chest high    |
| 5/25/2021 | PM          | Poor to Fair | 3.0-4.0ft | Waist to stomach high  |
| 5/26/2021 | AM          | Poor to Fair | 3.0-4.0ft | Waist to chest high    |
| 5/26/2021 | PM          | Poor to Fair | 3.0-4.0ft | Waist to chest high    |
+-----------+-------------+--------------+-----------+------------------------+
```

Or to get a forecast for a different subregion (specifically Ventura, CA, USA) and day range:

```sh
$ gosurf f -s 58581a836630e24c4487900c -d 5
+-----------+-------------+--------------+-----------+------------------------+
|   DATE    | TIME OF DAY |    RATING    |   RANGE   |        FORECAST        |
+-----------+-------------+--------------+-----------+------------------------+
| 5/24/2021 | AM          | Fair         | 3.0-4.0ft | Waist to shoulder high |
| 5/24/2021 | PM          | Poor to Fair | 3.0-4.0ft | Waist to chest high    |
| 5/25/2021 | AM          | Fair         | 2.0-3.0ft | Waist to stomach high  |
| 5/25/2021 | PM          | Poor to Fair | 2.0-3.0ft | Waist to stomach high  |
| 5/26/2021 | AM          | Fair         | 2.0-3.0ft | Waist to stomach high  |
| 5/26/2021 | PM          | Poor to Fair | 2.0-3.0ft | Waist to stomach high  |
| 5/27/2021 | AM          | Fair         | 2.0-3.0ft | Waist to stomach high  |
| 5/27/2021 | PM          | Poor to Fair | 2.0-3.0ft | Waist to stomach high  |
| 5/28/2021 | AM          | Fair         | 2.0-3.0ft | Waist to stomach high  |
| 5/28/2021 | PM          | Poor to Fair | 2.0-3.0ft | Waist to chest high    |
+-----------+-------------+--------------+-----------+------------------------+
```

Be sure to use the `subregion` ID (if coming from Surfline or their API directly) or the `TYPEID` (if coming from the CLI).

### Tide

To get the tides for Solimar Beach, CA, USA:

```sh
$ gosurf t -s 5842041f4e65fad6a770895f -d 3
+------------+-------+-------------+--------+
|    DATE    | TIME  | DESCRIPTION | HEIGHT |
+------------+-------+-------------+--------+
| 12/31/2020 | 03:09 | LOW         |   2.39 |
|            | 09:18 | HIGH        |   6.04 |
|            | 16:51 | LOW         |  -0.82 |
|            | 23:25 | HIGH        |   3.58 |
| 1/1/2021   | 03:45 | LOW         |   2.53 |
|            | 09:53 | HIGH        |   5.94 |
|            | 17:28 | LOW         |  -0.69 |
| 1/2/2021   | 00:08 | HIGH        |   3.61 |
|            | 04:32 | LOW         |   2.59 |
|            | 10:34 | HIGH        |   5.64 |
|            | 18:08 | LOW         |  -0.46 |
+------------+-------+-------------+--------+
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
$ gosurf s -t 58f7ed58dadb30820bb38f8b -md 1
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

The last two choices in the tree (4 and 5) are subregions (used for forecasts) and spots (used for tides).

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

Download the right binary (for example, `gosurf_darwin_amd64` for Mac) and move it to somewhere in your `$PATH` (so that it can be loaded into your command line environment).

For Mac, you might do something like this to download `gosurf` 2.0.0:

```sh
curl -L https://github.com/mhelmetag/gosurf/releases/download/2.0.0/gosurf_darwin_amd64 -o /usr/local/bin/gosurf
chmod a+x /usr/local/bin/gosurf
```
