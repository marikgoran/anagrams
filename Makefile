help:
	@printf 'Usage:\n Run "make container" once to build the code and the docker containter\n Run "make deploy" to stsrt the docker container build in previous step\n Run "make test" to run few tests against the running service\n Run "make clean" to remove the build artifacts (code and docker image)\n'
	
container:
	@go build
	@docker build -t anagrams .
	
deploy:
	@docker run -d -p 8080:8080 anagrams -t anagrams

clean:
	@docker ps --filter "ancestor=anagrams" -q | xargs -i docker stop {} 
	@docker rmi anagrams --force 2> /dev/null || true
	@rm -f anagrams 
	
tests:
	curl http://localhost:8080/anagrams/A
	@echo
	curl http://localhost:8080/anagrams/abb
	@echo
	curl http://localhost:8080/anagrams/1234567890
	
