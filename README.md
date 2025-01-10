# Pallet Sorter

Sorting your things on a pallet.
Old Python school project remade in Go + continuation.

## Planned new features

- User interface --> Configure amount, sizes etc etc in browser
- Sorting parameters

## Usage

Currently: <br/>
In ´palletSorter/cmd/main.go´. Configure 'cubes' and the variables 'width, height, depth' according to your needs. <br/>
Can add cubes with POST requests. Example: <br/>
```curl -X POST -H "Content-Type: application/json" -d '{width":1,"height":1,"depth":1}' http://localhost:8080/api/cubes/add```

Future:
All settings changeable while running in the browser.

## Run dev

```go run cmd/main.go```

## Build:

```go build -o palletviewer ./cmd```


## Changelog

### v0.1.0
- Initial preview .exe

### v0.2.0 (In development)
- Add cubes with POST requests
