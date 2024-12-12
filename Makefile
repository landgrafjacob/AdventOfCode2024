DIR=days/day${day}

create:
	mkdir -p ${DIR}
	cp templates/day.go ${DIR}/day${day}.go
	cp templates/day_test.go ${DIR}/day${day}_test.go
	touch ${DIR}/input.txt ${DIR}/test.txt
	sed -i "s/DAYNUM/${day}/g" ${DIR}/*.go
