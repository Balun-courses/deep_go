package main

type BaseStorageImpl struct{}

func (s BaseStorageImpl) Close() {}

type SyncStorageImpl struct{}

func (s SyncStorageImpl) Close() {}
func (s SyncStorageImpl) Sync()  {}

type BaseStorage interface {
	Close()
}

type SyncStorage interface {
	Close()
	Sync()
}

func main() {
	var baseStorage BaseStorage = BaseStorageImpl{}
	var syncStorage SyncStorage = SyncStorageImpl{}

	println("baseStorage:", baseStorage)
	baseStorage = syncStorage
	println("baseStorage:", baseStorage)

	println("syncStorage:", syncStorage)
	syncStorageCasted := syncStorage.(interface{ Close() })
	println("syncStorageCasted:", syncStorageCasted)

	//syncStorage = baseStorage
}
