.PHONY: all clean build

all: clean build

clean: ## Remove previous build
	rm -rf ./_book ./docs

build: 
	gitbook build
	mv ./_book ./docs