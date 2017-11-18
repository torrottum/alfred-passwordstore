all: clean build install

build:
	go build -o src/alfred-pass src/alfred-pass.go
	mkdir build/
	zip -j build/passwordstore.alfredworkflow src/*
	
install:
	open build/passwordstore.alfredworkflow

clean:
	rm -rf src/alfred-pass
	rm -rf build/
