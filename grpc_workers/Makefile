compile:
	protoc -I ./pkg/proto ./pkg/proto/grps_workers_server.proto --go_out=./pkg/proto/ --go_opt=paths=source_relative --go-grpc_out=./pkg/proto/ --go-grpc_opt=paths=source_relative

run:
	go run ./cmd/gws/main.go --config=./config/config.yaml

build:
	go build ./cmd/gws


insert_employee:
	psql -d workers \
	 -U postgres \
	 -p 5432 \
	 -h 127.0.0.1 \
	 -c "\copy employee(title) FROM '$(DATA_PATH)'"
	
insert_department:
	psql -d workers \
	 -U postgres \
	 -p 5432 \
	 -h 127.0.0.1 \
	 -c "\copy department(title) FROM '$(DATA_PATH)'"

insert_jobs:
	psql -d workers \
	 -U postgres \
	 -p 5432 \
	 -h 127.0.0.1 \
	 -c "\copy jobTitle(title, departmentID) FROM '$(DATA_PATH)' WITH DELIMITER ',' CSV HEADER"

insert_emp_job:
	psql -d workers \
	 -U postgres \
	 -p 5432 \
	 -h 127.0.0.1 \
	 -c "\copy employee_job(employeeid,jobid,hoursrate,annualsalary,typefp,typicalhours,yy) FROM '$(DATA_PATH)' WITH DELIMITER ',' CSV HEADER"