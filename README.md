# gokmp

String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP).

## Disclaimer

This library was written as part of my Master's Thesis and should be used as a helpful implementation reference for people interested in the Knuth-Morris-Pratt algorithm than as a performance string searching library.

I believe the compiler has since caught up to most of the gains that this library bought me back in the day.

See [Documentation](http://godoc.org/github.com/paddie/gokmp) on [GoDoc](http://godoc.org/).

Example:

```go
package main

import (
	"fmt"
	"github.com/paddie/gokmp"
)

const str = "aabaabaaaabbaabaabaaabbaabaabb"
//          "        _          _      _   "
//                   8          19     26
const pattern = "aabb"

func main() {
	kmp, _ := gokmp.NewKMP(pattern)
	ints := kmp.FindAllStringIndex(str)

	fmt.Println(ints)
}
```

Output:

```text
[8 19 26]
```

## Tests and Benchmarks

```bash
go test -v -bench=. -benchtime=15s
```

Output:

```text
=== RUN TestFindAllStringIndex
--- PASS: TestFindAllStringIndex (0.00 seconds)
=== RUN TestFindStringIndex
--- PASS: TestFindStringIndex (0.00 seconds)
=== RUN TestContainedIn
--- PASS: TestContainedIn (0.00 seconds)
=== RUN TestOccurrences
--- PASS: TestOccurrences (0.00 seconds)
=== RUN TestOccurrencesFail
--- PASS: TestOccurrencesFail (0.00 seconds)
PASS
goos: linux
goarch: amd64
BenchmarkKMPIndexComparison-8                   910319210               19.5 ns/op
BenchmarkStringsIndexComparison-8               1000000000               6.04 ns/op
BenchmarkKMPIndexComparisonDanteKO-8               10000           1586468 ns/op
BenchmarkStringsIndexComparisonDanteKO-8           40083            457143 ns/op
BenchmarkKMPIndexComparisonDanteOK-8             5035858              3596 ns/op
BenchmarkStringsIndexComparisonDanteOK-8        44114173               404 ns/op
ok  	github.com/paddie/gokmp	5.854s
```

## Comparison

```bash
gokmp.FindStringIndex / strings.Index = 178/359 = 0.4958
```

**Almost a 2x improvement over the naive built-in method.**
