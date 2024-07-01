# CoreUtils is a project that mimics some linux commands.

### 1. head

- Print the first 10 lines of input by default.
- Add a -n flag to specify the number of lines to print.
- Default 10 lines.

```console
user@user-VirtualBox:~$ go build head.go
user@user-VirtualBox:~$ ./head -n 4 fileName
```

### 2. tail

- Print the last 10 lines of input by default.
- Add a -n flag to specify the number of lines to print.
- Default 10 lines.

```console
user@user-VirtualBox:~$ go build tail.go
user@user-VirtualBox:~$ ./tail -n 4 fileName
```

### 3. wc (word count)

- Count lines, words, and characters in the input.
- Add -l, -w, and -c flags to display only lines, words, or characters respectively.
- Default displays lines, words, and characters.

```console
user@user-VirtualBox:~$ go build wc.go
user@user-VirtualBox:~$ ./wc -w -c  fileName
```

### 4. cat

- Concatenate and print files.
- Add a -n flag to number output lines.

```console
user@user-VirtualBox:~$ go build cat.go
user@user-VirtualBox:~$ ./cat -n fileName1 fileName2 fileName3
```

### 5. echo

- Print arguments to standard output.
- Add a -n flag to omit the trailing newline.

```console
user@user-VirtualBox:~$ go build echo.go
user@user-VirtualBox:~$ ./echo -n hello world
```

### 6. env
- Print enviroment variables to standard output.

```console
user@user-VirtualBox:~$ go build env.go
user@user-VirtualBox:~$ ./env
```

### 7. tree

- Print tree structure of a directroy to standard output.
- Add a -L flag to stop at a certain level.

```console
user@user-VirtualBox:~$ go build tree.go
user@user-VirtualBox:~$ ./tree -L 5 DirectoryName
```

### 8. yes

- Repeatedly output a message to standard output.
- Default message "y".

```console
user@user-VirtualBox:~$ go build yes.go
user@user-VirtualBox:~$ ./yes hello world
```

### 9. true

- Returns a successful exit status code of 0.

```console
user@user-VirtualBox:~$ go build true.go
user@user-VirtualBox:~$ ./true 
```

### 10. false

- Returns an unsuccessful exit status code of 1.

```console
user@user-VirtualBox:~$ go build false.go
user@user-VirtualBox:~$ ./false
```
