# Browser_History
## _Get Browser Histories for Chrome, Firefox, Safari, and Edge on MacOS (Intel and M1/2)_
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

The `browser_history` Golang binary produces the last 100 entries of any detected local users and respective Chrome, Firefox, Safari, and Edge browser histories and prints them to standard out

## Usage

```sh
./browser_history
```

## Caveats

The count of results per browser can be adjusted by adjusting the `main.go` file for the `sql_limit` variable and recompiling.
