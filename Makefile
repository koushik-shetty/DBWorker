all : clean fmt build test 

OUTDIR = out/
OUT = dbworker.exe
TARGET = $(OUTDIR)$(OUT)
MAIN = main.go
DIRS = ./app ./lib ./utils

build: $(DIRS)
	go build -o $(TARGET) $(MAIN)

clean:
	rm -rf out/

test: 
	go test $(DIRS)
	
fmt:
	go fmt $(DIRS)