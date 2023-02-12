# pathman

`pathman` is a CLI Manager for the Windows PATH environment variable that 
aims to simplify the process of managing PATH variables in a quick and efficient manner. 

`pathman` allows the user to add and remove keys of PATH environment variables in a 
convenient and user-friendly manner. The user can specify a specific folder as the value
to add, otherwise it will be the folder they are currently in. System-Level and User-Level
PATHs can be edited.

The program is written in Go and makes use of the Windows API to manipulate PATH variables
by Registry.