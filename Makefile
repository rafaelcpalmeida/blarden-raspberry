help:
	@echo ""
	@echo "***********************************************************"
	@echo "* Available commands:                                     *"
	@echo "*   ► fetch - fetches all required dependencies           *"
	@echo "*   ► build - compiles source code into runnable binaries *"
	@echo "***********************************************************"
	@echo ""

fetch:
	@echo ""
	@echo "*************************************************************"
	@echo "* Fetching necessary dependencies. This may take a while... *"
	@echo "*************************************************************"
	@echo ""
	go get

build:
	@echo ""
	@echo "********************************************"
	@echo "* Building source code into binary file... *"
	@echo "********************************************"
	@echo ""
	env GOOS=linux GOARCH=arm GOARM=6 go build -o open-door open-door.go
