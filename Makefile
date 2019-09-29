help:
        @echo ""
        @echo "***********************************************************************"
        @echo "* Available commands:                                                 *"
        @echo "*   ► fetch - fetches all required dependencies                       *"
        @echo "*   ► build-door-opener - compiles source code into runnable binaries *"
        @echo "***********************************************************************"
        @echo ""

fetch:
        @echo ""
        @echo "*************************************************************"
        @echo "* Fetching necessary dependencies. This may take a while... *"
        @echo "*************************************************************"
        @echo ""
        go get

build-door-opener:
        @echo ""
        @echo "********************************************************"
        @echo "* Building door opener source code into binary file... *"
        @echo "********************************************************"
        @echo ""
        env GOOS=linux GOARCH=arm GOARM=6 go build -o blarden-raspberry main.go open-door.go
