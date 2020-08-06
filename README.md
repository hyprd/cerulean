# cerulean
Go wrapper CLI application for youtube-dl.

Since YouTube videos lack reliable metadata accessibility, the ID3 tags of downloads may be wrong or malformed.

There are many ID3 tag editors available online (such as [mp3tag](https://www.mp3tag.de/en/)).

## Requirements
[Go](https://golang.org/)

[youtube-dl](https://youtube-dl.org/)

[ffmpeg](https://ffmpeg.org/)

Add the youtube-dl executable to your PATH.

## Installation

`go install cerulean` and fire away.

File downloads are saved in the repository directory.

## Usage

`cerulean [URL]`

The URL to add is the character sequence after ?v=
