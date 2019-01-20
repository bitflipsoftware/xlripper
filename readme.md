XLSX
====

#### Quickly Extract Data from XLSX Sheets in Go

## Background

We have encountered (once in Node.js and once in Golang) libraries that extract data from XLSX files, but that do so much, much too slowly (or with too much RAM) for our use case. Our use case involves extracting a spreadsheet that has about 65 columns, and about 380,000 rows. The total (compressed) file size is somewhere around 80 MB of data.

Writing a C++ Node.js plugin, which uses xlsxio (https://github.com/brechtsanders/xlsxio) we acheived a benchmark of approximately 30 seconds using about 4 GB of RAM (though it was hard to tell how much the plugin would have used without the JavaScript engine). We later ported our main application to Go and attempted to use https://github.com/tealeg/xlsx. It's not clean how long this library would take, but it's more than 5 minutes and at least 7GB of RAM. We believe there are two bottlenecks; 1) Go's built-in XML parsing is too slow, and 2) tealeg's library provides many features that we do not need which require parsing and memory overhead.

## Solution

Our xlsx library does one thing only. Its purpose is to take in an xlsx file and to return arrays of strings represengting the data found in the xlsx file's sheets.

## Priorities

1) Native Go implementation, no cgo.
2) As Fast as Possible
3) Optionally optimize for lower memory overhead once #2 has been acheived.
