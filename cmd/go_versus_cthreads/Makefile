PHONY: all test
all: cthreads go_versus_cthreads

$(info Setting ulimits to enable 100k C threads on Linux)

define set_limits
  set -x; ulimit -s  256; ulimit -i  120000
endef


# the GNU time command (not the bash builtin):
TIME:=/usr/local/bin/time -f "time: %es mem: %MKB"

TMP ?=  /tmp

# Have to go through this pantomime to satisfy the go compiler which doesn't
# want any "nasty" C files polluting the go project.
$(TMP)/cthreads.c: cthreads.c.txt
	cp $^ $@

cthreads: $(TMP)/cthreads.c
	gcc -g -o $@ $^ -lpthread

go_versus_cthreads: gothreads.go
	go build 

test: cthreads go_versus_cthreads
	@echo -e  "----------------------------   10"
	$(TIME) ./go_versus_cthreads 10
	@echo -e  ""
	$(set_limits); $(TIME) ./cthreads 10
	@echo -e  "\n--------------------------   1000"
	$(TIME) ./go_versus_cthreads 1000
	@echo -e  ""
	$(set_limits); $(TIME) ./cthreads 1000
	@echo -e  "\n--------------------------  10000"
	$(TIME) ./go_versus_cthreads 10000
	@echo -e  ""
	$(set_limits); $(TIME) ./cthreads 10000
	@echo -e  "\n-------------------------- 100000"
	$(TIME) ./go_versus_cthreads 100000
	@echo -e  ""
	$(set_limits); $(TIME) ./cthreads 100000
	@echo FINISHED
	
