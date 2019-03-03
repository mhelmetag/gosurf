# GoSurf

CLI Client for Surfline.

## Usage

To get a forecast, you need an area ID, region ID and subregion ID.

`gosurf` can get search the places Surfline services and read the forecasts for those places all from the command line (for Windows, Linux and Mac).

### Searching for Places

To search for all areas::

```
$ gosurf p
+--------+-----------------+
|   ID   |      NAME       |
+--------+-----------------+
|   4716 | North America   |
|   4710 | Central America |
|   4711 | South America   |
|   4707 | Caribbean       |
|   4712 | Europe          |
|  54137 | Middle East     |
|   4715 | Africa          |
|   4709 | Indian Ocean    |
|   4705 | Asia            |
|   4713 | Australia       |
|   4714 | New Zealand     |
|   4706 | Pacific Islands |
| 107392 | WCT Events      |
+--------+-----------------+
```

To search for regions in Central America, specify regions as the place type (`--t regions`) and then provide the area ID (`--a 4710`):

```
$ gosurf --a 4710 p --pt regions
+------+-------------+
|  ID  |    NAME     |
+------+-------------+
| 3250 | Guatemala   |
| 2745 | Panama      |
| 2736 | Costa Rica  |
| 3234 | Nicaragua   |
| 3252 | El Salvador |
+------+-------------+
```

### Forecasts

To get a forecast (the default subregion is Santa Barbara, CA, USA):

```
$ gosurf f
+-----------+--------------+-----------------------------+
|   DATE    |  CONDITION   |           REPORT            |
+-----------+--------------+-----------------------------+
| 2/26/2018 | POOR TO FAIR | 1-2ft. - knee to thigh high |
| 2/27/2018 | POOR         | 2-3ft. - knee to chest high |
| 2/28/2018 | POOR         | 1-2ft. - ankle to knee high |
| 3/1/2018  | POOR         | 1-2ft. - ankle to knee high |
| 3/2/2018  | POOR TO FAIR | 1-2ft. - ankle to knee high |
| 3/3/2018  | POOR TO FAIR | 2-3ft. - knee to waist high |
| 3/4/2018  | POOR TO FAIR | 2-3ft. - knee to waist high |
+-----------+--------------+-----------------------------+
```

Or to get a forecast for a different subregion (specifically South Cost Rica):

```
$ gosurf --a 4710 --r 2736 --sr 3314 f
+-----------+--------------+--------------------------------+
|   DATE    |  CONDITION   |             REPORT             |
+-----------+--------------+--------------------------------+
| 2/26/2018 | FAIR TO GOOD | 3-5ft. - waist to head high    |
| 2/27/2018 | FAIR         | 3-4ft. - waist to shoulder     |
|           |              | high                           |
| 2/28/2018 | FAIR         | 3-4ft. - waist to chest high   |
| 3/1/2018  | FAIR         | 3-4ft. - waist to chest high   |
| 3/2/2018  | FAIR TO GOOD | 3-5ft. - waist to head high    |
| 3/3/2018  | FAIR TO GOOD | 3-4ft. - waist to shoulder     |
|           |              | high                           |
| 3/4/2018  | FAIR         | 2-3ft. - thigh to chest high   |
+-----------+--------------+--------------------------------+
```

## Configuration File

To more easily configure a common place for a forecast, you can create a `.gosurf.yml` file in your home directory. It should look like the example in the repo (`.gosurf.sample.yml`).

The CLI will read in the values to override the global flags' defaults (by flag name; so `area` would be for the `area` flag). These loaded values can always be overridden by specifying global flags while running a command.

## Installation

Download the right binary (for example, `gosurf_darwin_amd64` for Mac) and move it to somewhere in your `$PATH` (so that it can be loaded into your command line environment).

For Mac, you might do something like to download `gosurf` 0.0.5:

```
$ curl -L https://github.com/mhelmetag/gosurf/releases/download/0.0.5/gosurf_darwin_amd64 -o /usr/local/bin/gosurf
$ chmod a+x /usr/local/bin/gosurf
```
