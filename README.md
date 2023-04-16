## Quake game log analytics


## Build
First, we must build the progam using the follow command line:

```
go build -o bin/quake-game-log-analytics main.go
```

## Run
After build, you can run the program passing as argument the log file path to be analysed:

```
./bin/quake-game-log-analytics resources/qgames.log
```

You can optionally pass an argument to output the game report in a file, using **-o=<output_file>**:

```
./bin/quake-game-log-analytics --o=report.out resources/qgames.log
```

## Build and Run
You can build and run the program without create a binary file:

```
go run main.go resources/qgames.log 
```