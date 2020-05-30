# GoSurf

CLI Client for Surfline

## Usage

To get a forecast, you need a subregion ID (can be found in the subregional forecast page URL).

`gosurf` can read forecasts for all places Surfline services from the command line (for Windows, Linux and Mac).

Since Surfline cut over to their v2 API, I've been working to make it work with `gosurf`. Search and tide are coming!

### Forecasts

To get a forecast (the default subregion is Santa Barbara, CA, USA):

```sh
$ gosurf f
+-----------+-------------+--------------+-----------+--------------------+
|   DATE    | TIME OF DAY |    RATING    |   RANGE   |      FORECAST      |
+-----------+-------------+--------------+-----------+--------------------+
| 5/29/2020 | AM          | POOR         | 0.5 - 1.0 | Shin to knee high  |
| 5/29/2020 | PM          | POOR         | 0.5 - 1.0 | Shin to knee high  |
| 5/30/2020 | AM          | POOR_TO_FAIR | 1.0 - 2.0 | Knee to thigh high |
| 5/30/2020 | PM          | POOR         | 1.0 - 2.0 | Knee to thigh high |
| 5/31/2020 | AM          | POOR_TO_FAIR | 1.0 - 2.0 | Knee to thigh high |
| 5/31/2020 | PM          | POOR         | 1.0 - 2.0 | Knee to thigh high |
| 6/1/2020  | AM          | POOR_TO_FAIR | 1.0 - 2.0 | Knee to thigh high |
| 6/1/2020  | PM          | POOR         | 1.0 - 2.0 | Shin to knee high  |
| 6/2/2020  | AM          | POOR         | 0.5 - 1.0 | Shin to knee high  |
| 6/2/2020  | PM          | POOR         | 0.5 - 1.0 | Shin to knee high  |
+-----------+-------------+--------------+-----------+--------------------+
```

Or to get a forecast for a different subregion (specifically Ventura, CA, USA):

```sh
$ gosurf -d 3 -s 58581a836630e24c4487900c f
+-----------+-------------+--------------+-----------+------------------------+
|   DATE    | TIME OF DAY |    RATING    |   RANGE   |        FORECAST        |
+-----------+-------------+--------------+-----------+------------------------+
| 5/29/2020 | AM          | POOR_TO_FAIR | 2.0 - 3.0 | Waist to stomach high  |
| 5/29/2020 | PM          | POOR_TO_FAIR | 2.0 - 3.0 | Waist to chest high    |
| 5/30/2020 | AM          | FAIR         | 3.0 - 4.0 | Waist to shoulder high |
| 5/30/2020 | PM          | POOR_TO_FAIR | 3.0 - 4.0 | Waist to shoulder high |
| 5/31/2020 | AM          | FAIR         | 3.0 - 4.0 | Waist to shoulder high |
| 5/31/2020 | PM          | POOR_TO_FAIR | 3.0 - 4.0 | Waist to chest high    |
+-----------+-------------+--------------+-----------+------------------------+
```

## Configuration File

To more easily configure a common place for a forecast, you can create a `.gosurf.yml` file in your home directory. It should look like the example in the repo (`.gosurf.sample.yml`).

The CLI will read in the values to override the global flags' defaults (by flag name; so `subregion` would be for the `subregion` flag). These loaded values can always be overridden by specifying global flags while running a command.

## Installation

Download the right binary (for example, `gosurf_darwin_amd64` for Mac) and move it to somewhere in your `$PATH` (so that it can be loaded into your command line environment).

For Mac, you might do something like to download `gosurf` 2.0.0:

```sh
curl -L https://github.com/mhelmetag/gosurf/releases/download/2.0.0/gosurf_darwin_amd64 -o /usr/local/bin/gosurf
chmod a+x /usr/local/bin/gosurf
```
