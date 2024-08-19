build:
	go build

install:
	make build
	mkdir -p ${HOME}/.dwm/bin
	mv main ${HOME}/.dwm/bin/dwmbar