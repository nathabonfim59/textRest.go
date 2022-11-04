# textRest.go
A rest API for your text files

## Why?

I had a huge list of passwords and I needed to go trough every entry and get this data from an API (the filesystem was not available). So I created this program.

## Usage
textRest [text file] [port]

## Web API
localhost:port/?idx=123
> line 123 contents
