## Часть 1. Sort.

Нужно реализовать unix утилиту sort с набором параметров:

`-f` - игнорировать регистр букв

`-u` - выводить только первое среди нескольких равных

`-r` - сортировка по убыванию

`-o <файл>` - выводить в файл, без этой опции выводить в stdout

`-n` - сортировка чисел

`-k <номер столбца>` - сортировать по столбцу (разделитель столбцов по умолчанию можно оставить пробел)

Также нужно написать тесты на эту функциональность. Тесты должны быть как для успешных случаев, так и для неуспешных. Примеры с тестами мы будем показывать ещё на следующих лекциях (можно посмотреть [тут](https://github.com/go-park-mail-ru/lectures/tree/master/1/6_uniq)). 

В `1/readings_1.md` есть список книг по го, а так же по всем частым и нужным операциям, там вы можете найти многие примеры кода, которые вам пригодятся.

### Пример работы:
```bash
    $ cat data.txt
    Napkin
    Apple
    January
    BOOK
    January
    Hauptbahnhof
    Book
    Go

    $ go run sort.go data.txt
    Apple
    BOOK
    Book
    Go
    Hauptbahnhof
    January
    January
    Napkin

    $ go run sort.go -r data.txt
    Napkin
    January
    January
    Hauptbahnhof
    Go
    Book
    BOOK
    Apple

    $ go run sort.go -u data.txt
    Apple
    BOOK
    Book
    Go
    Hauptbahnhof
    January
    Napkin

    $ go run sort.go -u -f data.txt
    Apple
    BOOK
    Go
    Hauptbahnhof
    January
    Napkin
```
