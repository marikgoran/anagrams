# Anagrams

Simple golang web services to find anagrams of a given word


### Requirments

This service was developed and tested on Debian 9. Requirements are make, docker 1.12+ and golang 1.6+

### Usage

The process of building and deploying the anagram service is controlled via Makefile. These are the steps to build and deploy the service on a local Linux system that has already has docker installed. 

	git clone https://github.com/marikgoran/anagrams
	cd anagrams/
	make 			# to show the usage
	make container		# builds the code and makes the docker image
	make deploy		# run the container based on the build image and exposes port 8080
	make tests		# runs few tests to check the functionality of the service
	make clean		# stops the container and deletes the image and the binary
	
