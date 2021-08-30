build_all_apis: build_server1 build_server2 build_server3

build_server1:
	go build cmd\api1\api1.go
build_server2:
	go build cmd\api2\api2.go
build_server3:
	go build cmd\api2xml\api2xml.go

build_checker:
	go build cmd\checker\checker.go

run_checker:
	checker

brchecker: build_checker run_checker


brapisw: build_all_apis run_apisw
brapisl: build_all_apis run_apisl
	
#windows
run_apisw:
	powershell saps "cmd '/c api1'"
	powershell saps "cmd '/c api2'"
	powershell saps "cmd '/c api2xml'"

#linux
run_apisl:
	terminal -e "/c api1"
	terminal -e "/c api2"
	terminal -e "/c api2"
