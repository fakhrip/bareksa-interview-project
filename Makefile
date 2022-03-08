serve:
	docker volume ls | grep -q bareksa-interview-project_bareksa_project_log && docker volume rm bareksa-interview-project_bareksa_project_log
	docker-compose up --build --remove-orphans -d

test:
	docker-compose -f docker-compose.test.yaml -p bareksa-inteview-project_test up --abort-on-container-exit --remove-orphans --build unit_test