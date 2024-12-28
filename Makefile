build:
	go build

install:
	make build
	strip main
	mkdir -p ${HOME}/.dwm/bin
	mv main ${HOME}/.dwm/bin/dwmbar