# Eluvio_Challenge_Core
## The Core Engineering challenge from Eluvio. An application to concurrently retrieve information using HTTP headers

This is a challenge project from Eluvio for the application to the core engineer summer internship 2021. 
The program is used to retrieve information about items using their item ID, with limitation that only one item per query, and maximum 5 items request concurrently

The code consists of two parts. The [main.go](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/main.go) performs most of the functionality, and [IDStore.go](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/IDstore.go) serves as a metastore for lookups and information storage.

To run the program, first build them by running "go build *.go" (Remember to set your GOPATH)

Input of item IDs should be in a form a .txt file (like [input.txt](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/input.txt), separated by "\n" (new line).

The input should contain three arguments: baseURL, ID text file, and authorization

For example, to retrive items from "https://eluv.io/items/ID" whose IDs are "A","B","C","D","E", stored in "input.txt" at the current directory,  with authorization "Y1JGMmR2RFpRc211MzdXR2dLNk1UY0w3WGpI", one should run the following command: 
```
go build *.go
go run . https://eluv.io/items/ input.txt Y1JGMmR2RFpRc211MzdXR2dLNk1UY0w3WGpI
```
The output should be stored in a txt file named [result.txt](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/result.txt)

Currently [input.txt](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/input.txt) and [result.txt](https://github.com/JWRickyWan/Eluvio_Challenge_Core/blob/main/main/result.txt) contain pseudo IDs and their responses. 

For any questions/comments/suggestions, feel free to send an email to me : jianghongwan1996@gmail.com
