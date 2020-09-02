# cerulean
Go wrapper CLI application for youtube-dl.

Since YouTube videos lack reliable metadata accessibility, the ID3 tags of downloads may be wrong or malformed.

There are many ID3 tag editors available online (such as [mp3tag](https://www.mp3tag.de/en/)) which can fix these issues.

## Requirements
[Go](https://golang.org/)

[Cobra](https://github.com/spf13/cobra)

[youtube-dl](https://youtube-dl.org/)

[ffmpeg](https://ffmpeg.org/)

Add the youtube-dl executable to your PATH.

## Installation

Clone repo then `go install cerulean`

File downloads are saved in the current directory.

## Usage

`cerulean [URL]`

The URL to add is the character sequence after ?v=

Works for playlists too. Use the playlist URL instead (after &list=) for this.
