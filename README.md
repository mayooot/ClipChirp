# ClipChirp

Copy file to your clipboard in macOS

# Usage

~~~bash
$ chmod +x clipchirp
$ sudo mv clipchirp /usr/local/bin/clipchirp
$ clipchirp bird.png

🦤 Successfully copied /Users/ming/ClipChirp/bird.png to clipboard
~~~

# Build and Install

~~~bash
$ chmod +x setup.sh
$ ./setup.sh
ClipChirp setup complete!
~~~

# Benchmark

~~~bash
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/mayooot/ClipChirp
cpu: Apple M3
Benchmark_resolveAbsolutePathUsingRealpath-8         997           1140647 ns/op           47162 B/op        104 allocs/op
Benchmark_resolveAbsolutePathInGo-8               749814              1509 ns/op             538 B/op          7 allocs/op
Benchmark_copyToClipboardUsingFinder-8                10         118017212 ns/op           12492 B/op         84 allocs/op
Benchmark_copyToClipboardInObjC-8                   6285            181118 ns/op             538 B/op          7 allocs/op
PASS
ok      github.com/mayooot/ClipChirp    5.022s
~~~
