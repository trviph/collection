test:
	go test -coverprofile=./test_profile ./... 

cover: test
	go tool cover -html=./test_profile && unlink ./test_profile