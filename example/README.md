Демонстрация работы protoc plugin gen-go-tarantool


Запуск Tarantool при локальной установки в директории tnt ```tarantool init.lua```

В директории proto размещается proto файлы
Для make proto необходимо установить: 
* https://github.com/yoheimuta/protolint
* https://github.com/pseudomuto/protoc-gen-doc

Демо приложение вызовет функцию push_messages Tarantool приложения,
передаст в нее структуру go сгенерированную из proto,
вернет ее как ответ на запрос.