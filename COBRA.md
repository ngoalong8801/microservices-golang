# Corba Introduction

## How to set up ?
```shell
go install github.com/spf13/cobra-cli@latest
which -a cobra
```
## What is cobra ?
Cobra is a powerful library and tool in Golang that is used to create CLI (Command-Line Interface) applications.

## Why use cobra ?
- Automatic help generation for commands and flags
- Increased productivity because of commands such as cobra init appname & cobra add cmdname.
## Installing
```shell
    go get github.com/spf13/cobra/cobra
```
After that you're ready to use this powerful tool in your own project:
```shell
  import "github.com/spf13/cobra"
```

## Using cobra
1. <strong>cobra init --pkg-name app </strong> : to initialize cobra app
2. <strong>cobra add  </strong>add any cobra service needed for your project
Your structure directory will look like this: 

![image](https://user-images.githubusercontent.com/76799846/184332024-e7d227dc-7d80-43a9-beb9-62a71adca6dd.png)

## Implementing
- cmd/root.go is the file cover all of your service. U have to access and modify function Run in it as long as compatible with your purpose.
- After that u need to <strong>go install</strong> to up-to-date CLI (or just simply ignore it if u want to use <strong>go run main.go</strong>)
- <strong>cobra init </strong> command that u used before will create CLI with your name project
- Each of service you have will be run by using: 
```
 <project-name> <name-service> 
```
- If u want to run project using:
```
<project-name>
 ```
Make sure that u have performed go install before running this command

## Conclusion
Yay well done bro !! \
Hope u have interesting moments with the brief pieces of introduction that I've written.

## References
1. https://pkg.go.dev/github.com/spf13/cobra
2. https://github.com/spf13/cobra
3. https://www.educative.io/answers/how-to-use-cobra-in-golang