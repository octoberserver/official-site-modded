build:
	cd "${0%/*}"; \
	make tailwind-build; \
	cp -r ./src/ ./dist/src/; \
	./tmpl-preprocessor/tmpl-pre; \
	make docker-build; \

tailwind-build:
	cd "${0%/*}"; \
	npx tailwindcss -i ./src/common/main-before.css -o ./src/common/main.css

docker-build:
	cd "${0%/*}"; \
	cd "dist"; \
	docker stop modsrv-site; \
	docker rm modsrv-site; \
	docker image rm octmodsrv/official-site:latest; \
	docker build -t octmodsrv/official-site:latest .
	# docker run -p80:80 -p443:443 --name modsrv-site octmodsrv/official-site:latest
	docker run -p8081:80 --name modsrv-site octmodsrv/official-site:latest

tailwind-dev:
	cd "${0%/*}"; \
	npx tailwindcss -i ./src/common/main-before.css -o ./src/common/main.css --watch

