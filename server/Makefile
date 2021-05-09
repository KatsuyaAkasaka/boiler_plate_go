.PHONY: genMock
genMock:
	mockgen -source ./pkg/domain/repository/$(param).go -destination ./pkg/usecase/mock/$(param).go

.PHONY: updateAllMock
updateAllMock:
	mockgen -source ./pkg/domain/repository/user.go -destination ./pkg/usecase/mock/user.go

.PHONY: connDB
connDB:
	mysql -uroot -p boiler_plate_go_$(env) -h 127.0.0.1 -proot

.PHONY: initDB
initDB:
	mysql -uroot -h 127.0.0.1 -proot < initdb.d/init_$(env).sql
	mysql -uroot -p boiler_plate_go_$(env) -h 127.0.0.1 -proot < initdb.d/init.sql
ifeq ("$(env)", "local")
	mysql -uroot -p boiler_plate_go_$(env) -h 127.0.0.1 -proot < initdb.d/sample_data.sql
endif

.PHONY: testUC
testUC:
	go test -v ./pkg/usecase/...

.PHONY: testEntry
testEntry:
	go test -v ./pkg/infra/entry/...
