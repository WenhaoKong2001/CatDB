# Design of CatDB

## API
```go
//open a CatDb,if filePath doesn't exist it will creat one CatDB with the given filePath
func Open(filePath string) (*DB, error)

//
func (db *DB) Put(key string, value string) error

//
func (db *DB) Get(key string) (string, bool, error)

//
func (db *DB) Delete(key string)(bool , error)
```

## Block
Block represent a disk page which usually is 4096 bytes.
It needs to be converted into Node in order to do CRUD operation.

## Node
Node represents a memory node which provides Btree operations.It
needs to be converted into Block when store to disk.

## ConvertServer
ConvertServer provide the ability to convert a memory node into
disk block or vice versa,and some useful utils to operate Block
and Node.

