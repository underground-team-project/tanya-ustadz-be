start:
	docker-compose up tanyaustadz

stop:
	docker-compose down

purge:
	docker-compose -f docker-compose.yaml down --volumes

build:
	docker-compose -f docker-compose.yaml up -d --build tanyaustadz
