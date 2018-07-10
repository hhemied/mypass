##MyPass Program :

Environment : Linux, OSX and Windows

Ideas:

Function:
  -> FUNC to check if the DB file "~/.mypass/.db" is exist or not
     is not exist, go to another FUNC to create the DB structure.

  -> FUNC to collect UserData
     [Better to use type struct] for the collected data.
     .. Return the UserData as a slice.

  -> FUNC to read from config file the length of PW
     And Location for .mypass tree 
     USE toml

  -> FUNC to generate a strong Password and return it as string.
     Use ASCI

  ->  Func to print out the UserData and the last password on the terminal.

  ->  Func to open the DB and write UserData to it
      Use "defer" to ensure that DB is closed.
