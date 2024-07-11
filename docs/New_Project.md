# Starting a New Porject in Go:

Before starting a new project we need to understand the below 2 definitions.

## Packages in Go

A package is just a folder that contains the collection of go files.

## Modules in Go

And the collection of these packages is known as module.

And when we're initializing a project we are really initializing a new module.

## New Project Setup

Create a new project `go_tutorials` by running the below command:

```bash
mkdir -p ./go_tutorials/
```

Go into the project root folder, with the below command:

```bash
cd go_tutorials
```

Initialize the new module named `go_tutorials` or with full github name `github.com/revanthvermareddy/go_tutorials` by running the below command:

```bash
go mod init go_tutorials
```

You should see a new `go.mod` file created in the root directory. This is similar to `package.json` file in node or `requirements.txt` file in python.

