test:
	docker build -f Dockerfile.test -t jjliggett/jjversion-test . && docker run --rm jjliggett/jjversion-test
