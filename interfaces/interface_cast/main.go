package main

type BaseStorage interface {
	Close()
}

type SyncStorage interface {
	Close()
	Sync()
}

func main() {
	var baseStorage BaseStorage
	var syncStorage SyncStorage

	baseStorage = syncStorage
	syncStorage = baseStorage
}
