Command-line Quiz Game

# Usage

```
go run quiz

go run quiz --csv=problems2.csv
```

# Requirements

## Part 1

- ✅ Interactive quiz game
- ✅ Ask a series of questions, the questions & answered pairs being rows in a CSV
- ✅ CSV to load is configurable; defaults to `problems.csv`
- ✅ Keep a total of correct and incorrect answers
- ✅ At the end of the quiz, tell the user how many questions they were asked, and how many they got right

## Part 2

- Add a time limit
- Quiz terminates as soon as timer expires (don't wait for input)
- Time limit is configurable; defaults to 30s

## Bonus

- ✅ Trim spaces
- ✅ Case insensitive
- Flag to shuffle quiz order
