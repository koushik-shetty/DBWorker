all : clean fmt build test 

OUTDIR = out/
OUT = dbworker.exe
TARGET = $(OUTDIR)$(OUT)
MAIN = main.go
DIRS = ./app ./lib ./utils
RESOURCE_DIR = resources/

build: $(DIRS) copy_res
	go build -o $(TARGET) $(MAIN)

compile: $(DIRS)
	go build -o $(TARGET) $(MAIN)

clean:
	rm -rf out/

test: 
	go test $(DIRS)
	
fmt:
	go fmt $(DIRS)

copy_res:
	@mkdir -p $(OUTDIR)$(RESOURCE_DIR)
	@cp -r $(RESOURCE_DIR)* $(OUTDIR)$(RESOURCE_DIR)
