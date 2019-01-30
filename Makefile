SCRIPTS_PATH      := scripts

.PHONY: all
all: build

.PHONY: build
build: 
	$(SCRIPTS_PATH)/docs.sh