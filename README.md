# GoSurf

CLI Client for Surfline.

## Usage

To get a forecast, you need an area ID, region ID and subregion ID.

`gosurf` can get search the places Surfline services and read the forecasts for those places all from the command line (for Windows, Linux and Mac).

### Searching for Places

To search for all areas::

```sh
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

To search for regions in Central America, specify regions as the place type (`--pt regions`) and then provide the area ID (`--a 4710`):

```sh
$ gosurf --a 4710 p --pt regions                                                                      [19:51:57]
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

```sh
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

```sh
gosurf --a 4710 --r 2736 --sr 3314 f
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
