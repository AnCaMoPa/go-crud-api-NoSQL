# go-crud-api-NoSQL
A template for a RESTful CRUD API in go with fiber and MongoDB. The API can work with any collection of a NoSQL database regardless the structure of said document.

To TEST the API you cand use Postman or Thunder Client in Visual Studio Code

## Installing

0. Install extra packages: 
    ```go install github.com/cosmtrek/air@latest```

1. Install air in the system (arch)
    yay -Syu air

2. Configure the following enviroment Variables needed, otherway air will not work
    export PATH=$PATH:$GOPATH/bin
    export GOPATH=$HOME/go

3. Create your own .env file using the .envtemplate file.

4. If air is correctly installed you will be able to execute the following command ```make dev```, this will run de App in development mode, any change you make while the App is running will compile and apply.

If you want to make your owns short cut with air, you only need to copy the ```dev``` one in the Makefile file.


   
    