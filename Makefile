BIN_NAME = dwmbar
LOCAL_DIR = $(HOME)/.dwm/bin
INSTALL_DIR = /usr/local/bin

.PHONY: all build install clean local_install

all: build

build:
	go build -o $(BIN_NAME)
	strip $(BIN_NAME)

install_local: build
	mkdir -p $(BUILD_DIR)
	cp $(BIN_NAME) $(LOCAL_DIR)/$(BIN_NAME)

install: build
	cp $(BIN_NAME) $(INSTALL_DIR)/$(BIN_NAME)

clean:
	rm -f $(BIN_NAME)
	rm -f main
