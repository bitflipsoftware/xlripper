xlripper
========

#### Quickly Extract Data from XLSX Sheets in Go

## Background

We have encountered (once in Node.js and once in Go) libraries that extract data from XLSX files, but that do so much, much too slowly (or with too much RAM) for our use case. Our use case involves extracting a spreadsheet that has about 65 columns, and about 380,000 rows. The total (compressed) file size is somewhere around 80 MB of data.

Writing a C++ Node.js plugin, which uses `xlsxio` (https://github.com/brechtsanders/xlsxio) we achieved a benchmark of approximately 30 seconds using about 4 GB of RAM (though it was hard to tell how the plugin would have performed without the JavaScript engine). We later ported our main application to Go and attempted to use https://github.com/tealeg/xlsx. It's not clear how long this library would take, but it's more than 5 minutes and at least 7GB of RAM. We believe there are two bottlenecks; 1) (primarily) Go's built-in XML parsing is too slow, and 2) (secondarily) tealeg's library provides features that we do not need which may slow it down slightly. However, profiling reveals that the largest bottleneck is Go's very sad XML parser.

## Solution

The `xlripper` library does one thing only. Its purpose is to take in an xlsx file and to return arrays of strings representing the data found in the xlsx file's sheets.

## Priorities

1. Native Go implementation, no cgo.
2. Go as fast as possible
3. Optimize for lower memory overhead if it can be done without making it too slow.

## Installation

This is a Go library, there is no main function. To use this library in your own application:

`go get -u github.com/bitflip-software/xlsx`
