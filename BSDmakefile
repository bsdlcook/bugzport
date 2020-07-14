PREFIX?=	/usr/local

GO_CMD=		${PREFIX}/bin/go
GO_BIN?=	bp
GO_TARGET=	cmd/bp
GO_FLAGS?=

default:
	${GO_CMD} build ${GO_FLAGS} -o ${GO_BIN} ${GO_TARGET}/main.go

clean:
	rm -f ${GO_BIN}
	${GO_CMD} clean
.PHONY: clean

mod:
	${GO_CMD} mod tidy -v
	${GO_CMD} mod verify
.PHONY: mod

lint:
	${PREFIX}/bin/golangci-lint run 
.PHONY: lint

install: default
	install -s -m 0755 ${GO_BIN} ${PREFIX}/bin
.PHONY: install

deinstall:
	rm -f ${PREFIX}/bin/${GO_BIN}
.PHONY: deinstall
