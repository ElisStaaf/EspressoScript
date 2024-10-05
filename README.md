# EspressoScript
EspressoScript, a programming language for... What was our target audience again? Oh, alright: No one!
No, just kidding. Altough i still don't think a lot of people aren't gonna see this since i'm not too
popular on Github. Oh wait! I have to continue!

Requirements
------------
- Git or Gh (optional)
- Go (optional)

Install
-------
To install, you can:
```bash
# Git
git clone https://ElisStaaf/EspressoScript

# Gh
gh repo clone ElisStaaf/EspressoScript

# cURL
echo "Fix it yourself."
```
Or i *guess* you could build from zip? Building from go is also easy, simply:
```bash
# In EspressoScript directory:
go build src/EspressoScript.go

```
Yeah, you've officially installed it now. Do something cool!

Tutorial
--------
Usage:
```bash
./EspressoScript <file>
````
This language is simple. *Too* simple. This is basically all you can write with it:
```javascript
fun main()
   printf("Hello World!") 
end
```
Interesting, i know. No, it has some other features. These being:  
**Variables**
```javascript
let luckyNumber int = 777
let githubUserName string = "ElisStaaf"
```
**if-else-statements**
```javascript
let num = 5
if num > 7
    printf("Your number is >7")
else
    printf("Your number is <7")
end
```
**functions**
```javascript
fun greet(name string, age int)
    printf("Hello ", name, "! You are", age, " years old!")
end
greet("Joe", 34)
greet("Maria", 45)
greet("Bob", 27)
```
They are pretty cool! But most of them also don't... Really work... Right now. This is a WIP, alright? So, yeah, contribute
please, we need your help. I'm not a good programmer. I'm going to sleep. Enjoy!
