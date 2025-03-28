EXAMPLE_FOLDERS := $(shell find . -mindepth 1 -maxdepth 1 -type d -name "go-lessons-*" | sort)
EXAMPLE_FILES := $(shell find $(EXAMPLE_FOLDERS) -type f | sort)

.PHONY: install
install:
	for example_folder in $(EXAMPLE_FOLDERS); do \
		(cd $$example_folder && go install) \
	done

INDEX.md: $(EXAMPLE_FILES)
	echo "# Index of examples\n" > $@
	for example_folder in $(EXAMPLE_FOLDERS); do \
		echo "- **$${example_folder#./}**: `go doc $$example_folder`" >> $@; \
	done