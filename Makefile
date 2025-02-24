BINARY = dwmbar
LOCAL = $(HOME)/.dwm/bin
INSTALL_DIR = /usr/local/bin

.PHONY: all build install clean local_install

all: build

build:
	@echo "Building... $(BINARY)" && go build -o $(BINARY) && strip $(BINARY)

install_local: build
	@echo "Local Installing... $(BINARY)" && \
	mkdir -p $(LOCAL) && \
	cp $(BINARY) $(LOCAL)/$(BINARY)

install: build
	@echo "Installing... $(BINARY)" && \
	mkdir -p $(INSTALL_DIR) && \
	mv $(BINARY) $(INSTALL_DIR)/$(BINARY)

clean:
	rm -f $(BINARY) ./main
