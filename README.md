# Pallet Sorter

Sorting your things on a pallet.
Old Python school project remade in Go + continuation.

### Future features

- User interface --> Configure amount, sizes etc etc in browser
- 3D Sorting --> DONE
- How it's filled parameters

### Usage

Currently:
In ```palletSorter/cmd/main.go```. Configure ```cubes``` and the variables ```width, height, depth``` according to your needs.

Future:
All settings changeable while running in the browser.

### Run dev

```go run cmd/main.go```

### Build:

```go build -o palletviewer ./cmd```
