all: clean build install

build:
	mkdir build/
	zip -j build/passwordstore.alfredworkflow src/*.*
	
install:
	open build/passwordstore.alfredworkflow

clean:
	rm -rf build/
