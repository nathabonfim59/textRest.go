# textRest.go
A rest API for your text files

## Why?

I had a huge list of passwords and I needed to go trough every entry and get this data from an API (the filesystem was not available). So I created this program.

## Usage
textRest [text file] [port]

![image](https://user-images.githubusercontent.com/21281852/200024805-10f696f5-498f-43e4-a04f-4c5b96c68f2e.png)

## Web API
localhost:port/?idx=123
> line 123 contents
![image](https://user-images.githubusercontent.com/21281852/200025432-6d1267a1-302b-4cf9-a58a-4a276fa8d2e9.png)
