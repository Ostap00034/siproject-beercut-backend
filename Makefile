.PHONY: run stop

run:
	@echo "Запуск всех микросервисов..."
	# Запускаем каждый микросервис в фоне
	go run auth-service/cmd/main.go &
	go run author-service/cmd/main.go &
	go run user-service/cmd/main.go &
	go run exhibition-service/cmd/main.go &
	go run genre-service/cmd/main.go &
	go run picture-service/cmd/main.go &
	go run api-gateway/cmd/main.go &
	@echo "Все микросервисы запущены."
	@wait

stop:
	@echo "Остановка микросервисов по портам 50051–50056..."
	@for port in 50051 50052 50053 50054 50055 50056; do \
	    pid=$$(lsof -ti tcp:$$port); \
	    if [ -n "$$pid" ]; then \
	        echo "Убиваем процесс $$pid, прослушивающий порт $$port"; \
	        kill $$pid; \
	    else \
	        echo "На порту $$port не найдено запущенных процессов"; \
	    fi; \
	done
