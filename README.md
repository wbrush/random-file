# random-file

This program is used to add numbers to the front of a file. These numbers are randomly generated. I use this program to "randomize" my music in players that don't have a random setting. 

#Usage

There are 2 command line parameters. They are:
	path : path to the directory which contains the files (default: current directory)
	write : flag to determine if the changes are written to the files or just displayed on the screen (default: display only)
	
Examples:
	random-file.exe -path=./files
	random-file.exe -path=./path/to/files -write

# Todo
- finish last unit test

# Notes
This was written several years ago. I've since moved to a TDD approach and so the unit tests would be complete in any new programs that I write. I like the concept of libraries or packages. I try to write most of my code in a package so that it can be easily added and/or removed.