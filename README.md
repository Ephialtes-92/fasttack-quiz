# Simple Quiz App

This is a simple quiz application. The questions are served by the quiz-api-server which also stores the users and theis scores.
The quiz is access by the quiz-cli-client where the user can get the questions, submit his answers and see how he compares to other people that have played the game.

## How to start the server
After cloning the repo, to start the server in a terminal navigate to quiz-cli-client and run:

```
go run main.go
```
> The server will run locally
 
## How to play
After that navigate to the cli-client folder so you can play the quiz

To get the questions:
```
go run main.go get-questions
```
The questions are retrieved from quiz-api-server/data/questions.json

To answer run:
```
go run main.go submit-answers --answers 1=3,2=2,3=2,4=1,5=2,6=3 --username Sotiris
```
>It is required that you provide the answers and your username.
>The answers have to be in the format question_id=option separated by commas with no spaces in between

Example output:
```
You got 6 out of 6 correct!
```
Then you can use :
```
go run main.go get-comparison --username Sotiris
```
to see how you compare to the other players.
If you are the only one that submitted their questions you will get
```
You are the only player so far
```
So make sure to submit questions with at least 2 different usernames before trying to get a comparison.
