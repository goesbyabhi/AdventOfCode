# My Advent Of Code attempts

This repo contains my attempts at Advent Of Code and my honest advice would be to not look at my attempts because they suck lmao
For AOC 2023, I've decided to try my hands at Go. I don't have any prior experience at Go so please ignore all those weird ugly formattings. Thank you.

### Tip
Please write a small script that would curl and automatically download the current date's inputs. Really helps in saving 2 minutes of your life.
I have written a basic script. It's the `fetch_input.sh` file. If you want, you can just straigt up grok it because that what I did for the curl section.

#### How to setup the script and shit

1) First open any input page i.e. any `adventofcode.com/{year}/day/{day}/input` page. Click on Inspect element and switch to Network section. Now reload the page. In the dev tools, you will see an `input` file. Open it, and click on the `Headers` part and scroll down till you see `Cookie`. Now copy the entire cookie (yea its going to be a big random block of letter along with your sesion keys).

2) Open your `.bashrc` file and go to the bottom and type
```sh
export SESSION="your cookie thing that you've copied (make sure that it has the session part TOO)"
```
3) Source the .bashrc file using 
```sh
source .bashrc
```
4) Once done, copy my `fetch_input.sh` script and save it somewhere. Now just do
```sh
chmod+x fetch_input.sh
```
This allows the execution of the script

5) If you want then you can set an alias for it too just like how I've done
```sh
alias myfetch="/path/to/fetch_input.sh"
```
and again source it just like in the 3rd step.
now you can call it like this to get the input of the current date and year
```sh
myfetch > input.txt
```
Or
```sh
myfetch <whatever day you want example 3> > filename.txt
```
so it will be 
```sh
myfetch 3 > input.txt
```
This takes the day you want as an input.

## Mhmm
Yea that's it. In the end have fun or something lol because I'm learning a lot by manually doing shit like this
