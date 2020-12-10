# Road to Go Pro — Packages and Modules

Please read the tutorial on [Medium]().

## Run the example code

```bash
# cd to this directory and then run
go run cmd/main.go
# check the outputs in the console
```

## Project structure

### cmd

We used to only have one package (`main`) in the project. Since we're introducing packages and modules, it's a great time to place stuff in more suitable places now.  
`main.go` is now placed under `cmd` folder. `cmd` folder indicates the stuff placed inside are executable.  

### pkg

`pkg` folder is a place where we store all the packages we build. In this example, we created a package called `stringset`.

## Go Module

This project is developed using go module feature. If you are using go version 1.13 or later, this feature comes out of the box. You can check out the `go.mod` file for the details of the dependency used in this project.  
Vendor mode is used in this project. Therefore, you don't need to download the dependencies from the Internet. Go will use the copy in the `vendor` folder.

Here are some useful commands:

```bash
# clean up the go modules
go mod tidy
# refresh contents in the vendor folder
go mod vendor
```

## Go Doc

```bash
# cd to this directory and then run
godoc
```

`godoc` tool starts a server at port [6060](http://localhost:6060). You can check the documentation it generates.  
For instance, to check the documentation of `stringset` package, you can go to [this page](http://localhost:6060/pkg/github.com/songx23/RoadToGoPro/part05/pkg/stringset/).