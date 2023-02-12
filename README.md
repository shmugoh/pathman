# `pathman`

This project should not be confused with [/therootcompany/pathman](https://github.com/therootcompany/pathman).

#
`pathman` is a simple CLI Manager for the Windows Environment Variables that  
aims to simplify the process of managing environment variables in a quick and 
efficient manner. 

`pathman` allows the user to add and remove keys of PATH environment variables in 
a convenient and user-friendly manner. The user can specify a specific folder as 
the value to add, otherwise it will be the folder they are currently in. 
System-Level and User-Level PATHs can be edited.

The program is written in Go and makes use of the Windows API to manipulate PATH 
variables
by Registry.

## Installation Guide

1. Download the latest version of `pathman` from [here]
(https://github.com/juanpisss/`pathman`/releases/latest/download/`pathman`.exe)

2. Place the executable file in a folder of your choice. 
Do keep in mind this folder will be included in your system's PATH.

3. Run the `pathman` executable from the same folder. `pathman` will 
automatically add the folder to your system's PATH for universal usage. 

That's it! You are now ready to use `pathman` on your system.

## Usage

`pathman` makes it easy to add folders/values to your system's environment variable. Here's how to use `pathman` to add a folder to PATH:

- To add the current folder to PATH, simply run `pathman` in the terminal. 
`pathman` will automatically add the current folder to PATH.

- If you want to add a different folder to PATH, use the `--folder` flag followed 
by the path to the desired folder. For example, to add `/example/folder` to PATH, 
run the following command: 
    ```bash
    pathman --folder /example/folder
    ```

- To add a folder to a different environment variable, use the `--path` flag 
followed by the name of the variable If the variable does not exist, `pathman` 
will create it for you. For example, to add a folder to a variable named 
`MY_VAR`, run the following command:
    ```bash
    pathman --folder /example/folder --path MY_VAR
    ```
    
## FAQ

### Why only for Windows?
It's really a personal thing; I only use Windows as my operating system and
created `pathman` as a fun Golang side project to solve this one issue I keep 
having whenever I wanna manually install CLI tools (that are not available on 
chocolately, or are but are broken), and I also wanted to mess with the Windows 
API directly.