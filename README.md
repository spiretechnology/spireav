# spireav

**spireav** is a wrapper library to FFmpeg that makes it simple to build complex filter graphs, execute transcoding tasks using them, and monitor progress of the tasks.

## Installation

To install this library:

```bash
go get github.com/spiretechnology/spireav
```

**spireav** depends on `ffmpeg` / `ffprobe` being installed locally at runtime. Either install it on the system PATH or install it on the machine and tell the library where to find it.

```go
spireav.FfmpegPath = "/path/to/ffmpeg"
spireav.FfprobePath = "/path/to/ffprobe"
```

## Examples

Several examples of how to use this library are in the `examples` folder.
