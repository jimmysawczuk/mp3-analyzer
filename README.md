# mp3-analyzer

Right now, this program takes an MP3 file piped via STDIN, analyzes it for ID3 information as well as length and size, and dumps them to STDOUT. **This program is still heavily untested.**


## Installation
```bash
$ go get github.com/jimmysawczuk/mp3-analyzer
$ go install mp3-analyzer
$ cat my-test-file.mp3 | mp3-analyzer
```
