-include .config/session.cfg
DIR=days/day${day}
INPUT_URL=https://adventofcode.com/2024/day/${day}/input

create:
	mkdir -p ${DIR}
	cp templates/day.go ${DIR}/day${day}.go
	cp templates/day_test.go ${DIR}/day${day}_test.go
	touch ${DIR}/test.txt
	if [ -f .config/session.cfg ]; then \
		curl --cookie "session=${session}" ${INPUT_URL} > ${DIR}/input.txt; \
	else \
		touch ${DIR}/input.txt; \
	fi
	sed -i "s/DAYNUM/${day}/g" ${DIR}/*.go
