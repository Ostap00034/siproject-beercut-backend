# Порты, на которых слушают микросервисы
PORTS := 50051 50052 50053 50054 50055 50056

.PHONY: run stop stop-unix stop-windows

run:
	@echo "Запуск всех микросервисов..."
	@go run auth-service/cmd/main.go &
	@go run author-service/cmd/main.go &
	@go run user-service/cmd/main.go &
	@go run exhibition-service/cmd/main.go &
	@go run genre-service/cmd/main.go &
	@go run picture-service/cmd/main.go &
	@go run api-gateway/cmd/main.go &
	@echo "Все микросервисы запущены."
	@wait

stop:
	@echo "Остановка микросервисов на портах $(PORTS)..."
ifeq ($(OS),Windows_NT)
	@$(MAKE) stop-windows
else
	@$(MAKE) stop-unix
endif

stop-unix:
	@echo "(Unix) Ищем и убиваем процессы..."
	@for port in $(PORTS); do \
	  pid=$$(lsof -t -i tcp:$$port); \
	  if [ -n "$$pid" ]; then \
	    echo "Убиваем процесс $$pid на порту $$port"; \
	    kill $$pid; \
	  else \
	    echo "На порту $$port нет процессов"; \
	  fi; \
	done

stop-windows:
	@for %%P in ($(PORTS)) do @for /f "tokens=5" %%A in ('netstat -aon ^| find "%%P" ^| find "LISTENING"') do taskkill /F /PID %%A
