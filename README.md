# svt

A minimal looking customizable  **Sorting Visualizer** for the **Terminal**

![](assets/svt.gif)

## Installation

Install it to your $GOPATH/bin using the `go install` command:

```console
$ go install github.com/aod/svt/cmd/svt
```

## Usage

```console
$ svt --help
Usage of svt:
  -a int
        Array size (default 12)
  -d duration
        Delay between sorts (default 16ms)
  -s value
        Sorting algorithm. Possible values are:
                - bubble
                - cocktail
         (default bubble)
  -t int
        Column thiccness (default 4)
```
