.PHONY: genMock
genMock:
	mockgen -source ./pkg/domain/repository/$(param).go -destination ./pkg/usecase/mock/$(param).go

.PHONY: connDB
connDB:
	mysql -uroot -p boiler_plate_go_$(env) -h 127.0.0.1 -proot

.PHONY: initDB
initDB:
	mysql -uroot -h 127.0.0.1 -proot < initdb.d/init_$(env).sql
	mysql -uroot -p boiler_plate_go_$(env) -h 127.0.0.1 -proot < initdb.d/init.sql

.PHONY: testUC
testUC:
	go test -v ./pkg/usecase/...

.PHONY: testEntry
testEntry:
	go test -v ./pkg/infra/entry/...
