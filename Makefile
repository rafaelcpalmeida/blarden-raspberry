help:
        @echo ""
        @echo "**************************************************************************"
        @echo "* Available commands:                                                    *"
        @echo "*   ► fetch - fetches all required dependencies                          *"
        @echo "*   ► build - compiles source code into runnable binaries                *"
        @echo "*   ► install - creates new service and moves binary to proper location  *"
        @echo "*   ► update - moves binary to proper location                           *"
        @echo "**************************************************************************"
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
        @echo "********************************************************"
        @echo "* Building door opener source code into binary file... *"
        @echo "********************************************************"
        @echo ""
        env GOOS=linux GOARCH=arm GOARM=6 go build -o blarden-raspberry main.go open-door.go

install:
        @echo ""
        @echo "*************************************************"
        @echo "* Placing binary file on /usr/local/bin path... *"
        @echo "*************************************************"
        @echo ""
        sudo mv blarden-raspberry /usr/local/bin
        @echo ""
        @echo "*************************************"
        @echo "* Registering new system service... *"
        @echo "*************************************"
        @echo ""
        sudo cp blarden.service /etc/systemd/system/blarden.service
        sudo chmod 644 /etc/systemd/system/blarden.service
        sudo systemctl enable blarden

update:
        @echo ""
        @echo "*************************************************"
        @echo "* Placing binary file on /usr/local/bin path... *"
        @echo "*************************************************"
        @echo ""
        sudo mv blarden-raspberry /usr/local/bin
