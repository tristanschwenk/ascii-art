# **ASCII-ART by Tristan SCHWENK**

Welcome to my version of the **ASCII-ART**. It contains 3 modules of the [001-edu project](https://github.com/01-edu/public/tree/master/subjects/ascii-art) :

- Filesystem module (fs)
- Color module
- Output module

## **How can you use the ascii-art script ?**

The script uses **GoLang flags**, so you have to put the name of the flags with the value you want it to have

The word you want to print have to by written at the very end of your line. 

## **Why do you have to put your word at the end ?**

With GoLang flags, you if you put a parameter that is not a flag, then it goes as a "flag tail" and everything else goes in there.
In order to have the flags plus your word, you have to put it at the end

## Module Usage :

### **Filesystem (fs)**
In order to use the filesystem flag, just write -fs *[name_of_font] [word]*

You can choose 3 fonts :
- Standard
- Shadow
- Thinkertoy

If you don't put the fs flag, the standard font is used... as standard

Example :
```
ascii-art % ./main -fs thinkertoy "Hello you"
                                     
o  o     o o                         
|  |     | |                         
O--O o-o | | o-o       o  o o-o o  o 
|  | |-' | | | |       |  | | | |  | 
o  o o-o o o o-o       o--O o-o o--o 
                          |          
                       o--o          
```

### **Color flag**
In order to use the color flag, write -color *[XXX,XXX,XXX] [word]*

Your color have to be in RGB format and separated by commas.

Example
```
 ascii-art % ./main -color 35,150,200 'Hello there'
 _    _          _   _                 _     _                           
| |  | |        | | | |               | |   | |                          
| |__| |   ___  | | | |   ___         | |_  | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \        | __| |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |       \ |_  | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/         \__| |_| |_|  \___| |_|     \___| 
                                                                                                                                                 
```
(the color doesn't print here but will be printed on your terminal or by [clicking here](nothing_to_see/readme-img.png))

### **Output flag**
In order to use the output flag, write -output *[file_name] [word]*

If your file exist, it will write over it 

If it doesn't exist, it will create it and write on it 

Example
```
ascii-art % ./main -output newfile.txt 'Hello there'
ascii-art % cat newfile.txt 
 _    _          _   _                 _     _                           
| |  | |        | | | |               | |   | |                          
| |__| |   ___  | | | |   ___         | |_  | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \        | __| |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |       \ |_  | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/         \__| |_| |_|  \___| |_|     \___| 
                                                                         
 ```
                                                                        


