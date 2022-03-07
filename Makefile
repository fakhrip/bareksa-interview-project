serve:
	docker-compose up --build --remove-orphans -d

test:
	docker-compose -f docker-compose.test.yaml -p bareksa-inteview-project_test up --abort-on-container-exit --remove-orphans --build unit_test