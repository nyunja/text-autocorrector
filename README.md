# Goreloaded

This is a simple text completion/editing/auto-correction tool. Have fun auditing.

## How It Works

The tool takes arguments as the name of a file (input) which contains a text to be modified and the names of the file which will contain the modified tex (output).

The **[main](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/main.go)** file calls the functions to edit the file and writes the output files.

## Commands

- **[FinalizedText](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/functions/FinalizedText.go)**: Calls all the functions to edit the text and pass the output to the main file.

* **[CommandChecker](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/functions/CommandChecker.go)**: Check for command patterns [(cap), (bin), (hex), (up), (low)] or the command with a number [(cap, number)] in the input and calls the appropriate function and passes the returned slice of strings to the FinalizedText function.

* **[TextTransformer](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/functions/TextTransformer.go)**: Contains the logic to execute the tranformation. It:

        - Replaces the word before (hex) with the decimal version of the word. 

                Ex: "1E (hex) files were added" -> "30 files were added"

        - Replaces the word before (bin) with the decimal version of the word. 

                Ex: "It has been 10 (bin) years" -> "It has been 2 years"

        - Converts the word before (up) with the Uppercase version of it. 
        
                Ex: "Ready, set, go (up) !" -> "Ready, set, GO !"

        - Converts the word before (low) with the Lowercase version of it. 

                Ex: "I should stop SHOUTING (low)" -> "I should stop shouting"

        - Converts the word before (cap) with the capitalized version of it.

                Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"
    
        - For **(low), (up), (cap)** and a number like so: **(low, number)**, it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.

                Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING"

* **[Punctuations](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/functions/Punctuations.go)** corrects every instance of the punctuations ., ,, !, ?, : and ; to be close to the previous word and with space apart from the next one. 
        
                Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"
    
        - The tool leaves groups of punctuations together like: ... or !?. 
        
                Ex: "I was thinking ... You were right" -> "I was thinking... You were right".
    
        - The opening and closing single quotes will be placed to the right and left of the word or words in the middle of them, without any spaces. 
        
                Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"
        
                Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"

+ **[Vowels](https://learn.zone01kisumu.ke/git/johnotieno0/go-reloaded/src/branch/master/functions/Vowels.go)** turns a to an if the next word begins with a vowel (a, e, i, o, u) or a h. 

        Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!"
