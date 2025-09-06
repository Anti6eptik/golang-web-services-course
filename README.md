# directory visualizer

![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## 📌 Описание

Этот проект — простой визулизатор строения папки на Go. Он был сделан как домашнее задание первой недели курса https://stepik.org/course/187490/syllabus

## 🚀 Установка

1. **Скачайте и установите Golang** (если у вас его нет) с [официального сайта](https://go.dev/).
2. **Клонируйте репозиторий:**  
   ```sh
   git clone https://github.com/Anti6eptik/golang-web-services-course
   cd HTTP-server-on-GO
   ```
3. **Установите зависимости:**  
   ```sh
   go mod download  # или go mod tidy
   ```
4. **Запустите программу:**  
   ```sh
   go run .\main.go
   ```
5. **Введите в консоль название папки, желаемую глубину и нужна ли вам визуализация файлов**
   
## Примеры
Терминал после запуска:
```sh
Введите название папки
testdata
Введите максимальную глубину
10
Введите 0 если хотите визулизировать только папки или 1 если хотите визуализировать папки и файлы
0
```
Результат:
```sh
testdata
        ├───project
        ├───static
        │       ├───a_lorem
        │       │       └───ipsum
        │       ├───css
        │       ├───html
        │       ├───js
        │       └───z_lorem
        │               └───ipsum
        └───zline
                └───lorem
                        └───ipsum
```

Терминал после запуска:
```sh
Введите название папки
testdata
Введите максимальную глубину
10
Введите 0 если хотите визулизировать только папки или 1 если хотите визуализировать папки и файлы
1
```
Результат:
```sh
testdata
        ├───project
        │       ├───file.txt (19b)
        │       └───gopher.png (70372b)
        ├───static
        │       ├───a_lorem
        │       │       ├───dolor.txt (empty)
        │       │       ├───gopher.png (70372b)
        │       │       └───ipsum
        │       │               └───gopher.png (70372b)
        │       ├───css
        │       │       └───body.css (28b)
        │       ├───empty.txt (empty)
        │       ├───html
        │       │       └───index.html (57b)
        │       ├───js
        │       │       └───site.js (10b)
        │       └───z_lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        ├───zline
        │       ├───empty.txt (empty)
        │       └───lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        └───zzfile.txt (empty)
```
