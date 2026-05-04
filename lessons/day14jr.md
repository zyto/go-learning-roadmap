# Go | Day 14 — структуры: собираем данные в понятные модели

Привет.

До этого момента ты уже писал код, который работает с отдельными значениями:

* числами
* строками
* `slice`
* `map`
* функциями
* ошибками
* указателями

Но в реальных задачах данные почти никогда не живут по одному.

Представь пользователя интернет-магазина.

У него есть:

* `ID`
* имя
* email
* возраст
* признак активности

Можно хранить всё это отдельными переменными:

```go
id := 1
name := "Anna"
email := "anna@example.com"
age := 30
active := true
```

Но чем больше данных, тем сложнее помнить, что относится к чему.

Поэтому в Go есть `struct`.

`struct` позволяет сказать:

> Вот эти поля вместе описывают одну сущность.

Сегодня ты научишься:

* объявлять структуры
* создавать значения структур
* читать и менять поля
* понимать zero value структуры
* видеть, где структура копируется
* менять структуру через указатель
* вкладывать одну структуру в другую
* работать со списком структур
* искать, фильтровать, группировать и агрегировать данные по полям
* понимать DTO как простую модель передачи данных

В следующем уроке появятся `const`, именованные типы и enum-like constants.

После них методы станут намного понятнее.

Но сегодня методов ещё не будет.

---

## 1. Что такое `struct`

`struct` — это тип, который состоит из полей.

Поле — это имя плюс тип.

Например:

```go
type User struct {
	ID   int
	Name string
}
```

Здесь объявлен новый тип `User`.

У него два поля:

* `ID` типа `int`
* `Name` типа `string`

Теперь можно создать переменную такого типа:

```go
package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	var user User

	fmt.Println(user)
}
```

Результат:

```text
{0 }
```

Почему так:

* `var user User` создаёт значение типа `User`
* поля явно не заполнены
* `ID` получает zero value для `int`, то есть `0`
* `Name` получает zero value для `string`, то есть пустую строку

Запомни:

> Zero value структуры — это структура, у которой каждое поле получило zero value своего типа.

---

## 2. Доступ к полям через точку

Чтобы обратиться к полю структуры, используется точка.

```go
package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	user := User{
		ID:   1,
		Name: "Anna",
	}

	fmt.Println(user.ID)
	fmt.Println(user.Name)
}
```

Результат:

```text
1
Anna
```

`user.ID` читается так:

> Возьми поле `ID` у значения `user`.

`user.Name` читается так:

> Возьми поле `Name` у значения `user`.

Это не `map`.

В `map` ключ может быть разным во время выполнения:

```go
scores["go"]
scores["sql"]
```

В структуре поля известны заранее:

```go
user.ID
user.Name
```

Если написать `user.Email`, но поля `Email` в типе `User` нет, программа не скомпилируется.

Это хорошо.

Компилятор помогает поймать ошибку раньше запуска.

---

## 3. `struct` или `map`

Очень важный выбор:

> Если набор полей заранее известен, обычно нужна структура, а не `map`.

Например, данные пользователя можно записать через `map`:

```go
user := map[string]string{
	"id":   "1",
	"name": "Anna",
}
```

На первый взгляд удобно.

Но есть проблемы:

* `id` пришлось хранить строкой, хотя это число
* если написать `"nmae"` вместо `"name"`, компилятор не заметит
* по коду не видно полный набор полей пользователя
* каждый раз нужно помнить, какие ключи договорились использовать

Структура делает модель явной:

```go
type User struct {
	ID   int
	Name string
}

user := User{
	ID:   1,
	Name: "Anna",
}
```

Теперь:

* `ID` — число
* `Name` — строка
* поля видны в объявлении типа
* компилятор проверяет имена полей
* компилятор проверяет типы значений

Если написать:

```go
user.Nmae = "Ann"
```

код не скомпилируется.

Если написать:

```go
user.ID = "1"
```

код тоже не скомпилируется, потому что `ID` имеет тип `int`.

Правило для этого урока:

* если поля заранее известны — используй `struct`
* если ключи действительно динамические — используй `map`

Пример динамических ключей:

```go
scores := map[string]int{
	"go":  10,
	"sql": 8,
}
```

Здесь `map` уместна: предметы могут быть разными, а значение у всех однотипное — баллы.

---

## 4. Как создать структуру через литерал

Самый удобный способ создать структуру — именованный литерал.

```go
package main

import "fmt"

type Product struct {
	ID    int
	Title string
	Price int
}

func main() {
	product := Product{
		ID:    10,
		Title: "Keyboard",
		Price: 5000,
	}

	fmt.Println(product)
}
```

Результат:

```text
{10 Keyboard 5000}
```

Почему это удобно:

* видно, какое значение в какое поле попадает
* поля можно читать без знания порядка в объявлении
* если позже появится новое поле, старый код будет проще понять

Есть и позиционная форма:

```go
product := Product{10, "Keyboard", 5000}
```

Она работает, но для учебного и backend-кода лучше почти всегда использовать именованную форму:

```go
product := Product{
	ID:    10,
	Title: "Keyboard",
	Price: 5000,
}
```

В Go читаемость важнее экономии пары строк.

---

## 5. Можно заполнить не все поля

В именованном литерале не обязательно указывать все поля.

```go
package main

import "fmt"

type Account struct {
	ID     int
	Email  string
	Active bool
}

func main() {
	account := Account{
		ID:    1,
		Email: "anna@example.com",
	}

	fmt.Println(account.ID)
	fmt.Println(account.Email)
	fmt.Println(account.Active)
}
```

Результат:

```text
1
anna@example.com
false
```

Поле `Active` не заполнено явно.

Значит, оно получило zero value для `bool`, то есть `false`.

Это может быть нормально.

Например:

* новый аккаунт по умолчанию не активен
* заказ по умолчанию ещё не оплачен
* товар по умолчанию скрыт

Но это может быть и ошибкой.

Поэтому всегда задавай себе вопрос:

> Подходит ли zero value для этого поля?

---

## 6. Поля можно менять

Если структура лежит в переменной, её поля можно менять через точку.

```go
package main

import "fmt"

type User struct {
	Name   string
	Active bool
}

func main() {
	user := User{Name: "Anna"}

	user.Active = true
	user.Name = "Ann"

	fmt.Println(user)
}
```

Результат:

```text
{Ann true}
```

Что произошло:

* сначала `Name` был `"Anna"`
* `Active` был `false`
* потом `Active` стал `true`
* потом `Name` стал `"Ann"`

Тип поля важен.

Если поле `Active bool`, туда нельзя записать строку.

Если поле `Price int`, туда нельзя записать `"5000"`.

Структуры делают данные не только удобными, но и более безопасными.

---

## 7. Структура — это значение

Это главный мостик из прошлого урока.

`struct` в Go — значение.

При присваивании структура копируется.

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	first := User{Name: "Anna", Age: 30}
	second := first

	second.Name = "Bob"

	fmt.Println(first)
	fmt.Println(second)
}
```

Результат:

```text
{Anna 30}
{Bob 30}
```

Почему:

* `second := first` создаёт копию структуры
* `second.Name = "Bob"` меняет только копию
* `first.Name` остаётся `"Anna"`

Структура не становится “ссылкой” только потому, что внутри несколько полей.

Это всё ещё значение.

---

## 8. Функция тоже получает копию структуры

Передача аргумента в функцию похожа на присваивание.

```go
package main

import "fmt"

type User struct {
	Name string
}

func rename(user User) {
	user.Name = "Bob"
}

func main() {
	user := User{Name: "Anna"}

	rename(user)

	fmt.Println(user)
}
```

Результат:

```text
{Anna}
```

Почему имя не изменилось?

Потому что функция `rename` получила копию структуры.

Она изменила поле у копии, а не у исходного значения из `main`.

Есть два нормальных способа изменить результат снаружи.

Первый способ — вернуть новое значение:

```go
package main

import "fmt"

type User struct {
	Name string
}

func renamed(user User, name string) User {
	user.Name = name
	return user
}

func main() {
	user := User{Name: "Anna"}

	user = renamed(user, "Bob")

	fmt.Println(user)
}
```

Результат:

```text
{Bob}
```

Здесь всё явно:

* функция вернула изменённую структуру
* `main` сохранил результат обратно в `user`

Это хороший вариант, когда функция делает преобразование.

---

## 9. Указатель на структуру

Второй способ — передать указатель.

```go
package main

import "fmt"

type User struct {
	Name string
}

func rename(user *User, name string) {
	user.Name = name
}

func main() {
	user := User{Name: "Anna"}

	rename(&user, "Bob")

	fmt.Println(user)
}
```

Результат:

```text
{Bob}
```

Что важно:

* `&user` берёт адрес структуры
* `user *User` означает “указатель на `User`”
* функция меняет исходную структуру

Полная форма доступа к полю через указатель:

```go
(*user).Name = name
```

Но Go разрешает писать короче:

```go
user.Name = name
```

Это просто удобная запись.

Указатель от этого не перестаёт быть указателем.

Теперь короткое правило выбора.

Если функция просто строит новое значение из старого, чаще возвращай структуру:

```go
func renamed(user User, name string) User {
	user.Name = name
	return user
}
```

Такой код читается как преобразование:

> взяли пользователя, получили пользователя с новым именем.

Если функция должна изменить уже существующее значение, можно принять указатель:

```go
func activate(user *User) {
	if user == nil {
		return
	}

	user.Active = true
}
```

Такой код читается как действие над существующим объектом:

> активировали именно этого пользователя.

На текущем уровне держи простую модель:

* преобразование данных — вернуть новое значение
* изменение конкретного существующего значения — указатель может быть уместен
* если указатель может быть `nil` — сначала проверка
* не превращай все структуры в указатели автоматически

---

## 10. Если указатель может быть `nil`

Указатель на структуру может быть `nil`.

Если обратиться к полю через `nil`-указатель, будет `panic`.

Поэтому сначала проверка.

```go
package main

import "fmt"

type User struct {
	Name string
}

func printName(user *User) {
	if user == nil {
		fmt.Println("no user")
		return
	}

	fmt.Println(user.Name)
}

func main() {
	printName(nil)

	user := User{Name: "Anna"}
	printName(&user)
}
```

Результат:

```text
no user
Anna
```

Это тот же паттерн, который уже был в уроках про ошибки и указатели:

* сначала обрабатываем плохой или пустой случай
* делаем `return`
* ниже пишем основной путь

Такой стиль называется guard clause.

---

## 11. Структура внутри структуры

Поле структуры может быть другой структурой.

```go
package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Address Address
}

func main() {
	user := User{
		Name: "Anna",
		Address: Address{
			City:   "Moscow",
			Street: "Tverskaya",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.Address.City)
	fmt.Println(user.Address.Street)
}
```

Результат:

```text
Anna
Moscow
Tverskaya
```

Такой приём полезен, когда часть данных логически живёт внутри другой модели.

Например:

* пользователь и адрес
* заказ и доставка
* товар и размеры

На этом уроке используем только обычное вложение через именованное поле.

Встраивание структур без имени поля будет позже, рядом с композицией.

---

## 12. Список структур

Часто структура нужна не одна.

Например, список пользователей:

```go
package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	users := []User{
		{ID: 1, Name: "Anna", Active: true},
		{ID: 2, Name: "Bob", Active: false},
		{ID: 3, Name: "Kate", Active: true},
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Active)
	}
}
```

Результат:

```text
1 Anna true
2 Bob false
3 Kate true
```

`[]User` — это обычный `slice`.

К нему применимы уже знакомые инструменты:

* `len`
* `append`
* `range`
* фильтрация
* поиск
* группировка

Разница только в том, что элемент теперь не `int` и не `string`, а структура.

---

## 13. Важная ловушка `range`

Когда ты пишешь:

```go
for _, user := range users {
	user.Active = true
}
```

переменная `user` получает копию элемента.

Изменение копии не меняет исходный элемент `slice`.

```go
package main

import "fmt"

type User struct {
	Name   string
	Active bool
}

func main() {
	users := []User{
		{Name: "Anna"},
		{Name: "Bob"},
	}

	for _, user := range users {
		user.Active = true
	}

	fmt.Println(users)
}
```

Результат:

```text
[{Anna false} {Bob false}]
```

Чтобы изменить элементы исходного `slice`, иди по индексам:

```go
package main

import "fmt"

type User struct {
	Name   string
	Active bool
}

func main() {
	users := []User{
		{Name: "Anna"},
		{Name: "Bob"},
	}

	for i := range users {
		users[i].Active = true
	}

	fmt.Println(users)
}
```

Результат:

```text
[{Anna true} {Bob true}]
```

Запомни:

> Если нужно изменить элемент `[]Struct`, используй индекс.

---

## 14. Поиск структуры по полю

Допустим, есть список пользователей.

Нужно найти пользователя по `ID`.

```go
package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func findUserByID(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}

	return User{}, false
}

func main() {
	users := []User{
		{ID: 1, Name: "Anna"},
		{ID: 2, Name: "Bob"},
	}

	user, ok := findUserByID(users, 2)
	fmt.Println(ok)
	fmt.Println(user)

	missing, ok := findUserByID(users, 10)
	fmt.Println(ok)
	fmt.Println(missing)
}
```

Результат:

```text
true
{2 Bob}
false
{0 }
```

Почему возвращаем `(User, bool)`?

Потому что `User{}` — это просто zero value структуры.

Сам по себе он не говорит, найден пользователь или нет.

`bool` отвечает на вопрос:

> Нашли значение?

Это тот же паттерн `(value, ok)`, который уже был в `map`.

---

## 15. Фильтрация структур

Фильтрация списка структур почти не отличается от фильтрации чисел.

Просто условие теперь смотрит на поле.

```go
package main

import "fmt"

type User struct {
	Name   string
	Active bool
}

func activeUsers(users []User) []User {
	result := make([]User, 0)

	for _, user := range users {
		if user.Active {
			result = append(result, user)
		}
	}

	return result
}

func main() {
	users := []User{
		{Name: "Anna", Active: true},
		{Name: "Bob", Active: false},
		{Name: "Kate", Active: true},
	}

	fmt.Println(activeUsers(users))
}
```

Результат:

```text
[{Anna true} {Kate true}]
```

Алгоритм знакомый:

* создаём пустой результат
* идём по исходному списку
* если условие подходит, добавляем элемент
* возвращаем результат

---

## 16. Группировка структур

Теперь задача чуть ближе к backend.

Есть товары.

Нужно сгруппировать их по категории.

```go
package main

import "fmt"

type Product struct {
	Title    string
	Category string
}

func groupByCategory(products []Product) map[string][]Product {
	result := make(map[string][]Product)

	for _, product := range products {
		result[product.Category] = append(result[product.Category], product)
	}

	return result
}

func main() {
	products := []Product{
		{Title: "Keyboard", Category: "devices"},
		{Title: "Mouse", Category: "devices"},
		{Title: "Desk", Category: "furniture"},
	}

	groups := groupByCategory(products)

	fmt.Println(groups["devices"])
	fmt.Println(groups["furniture"])
}
```

Результат:

```text
[{Keyboard devices} {Mouse devices}]
[{Desk furniture}]
```

Разберём строку:

```go
result[product.Category] = append(result[product.Category], product)
```

Если категории ещё нет, `result[product.Category]` возвращает nil slice.

`append` умеет работать с nil slice.

Поэтому отдельная проверка ключа здесь не нужна.

---

## 17. Агрегация по структурам

Агрегация — это когда мы из набора данных считаем итог.

Например:

* сумму
* количество
* максимум
* группировку

Посчитаем сумму заказов по статусу.

```go
package main

import "fmt"

type Order struct {
	Status string
	Total  int
}

func sumByStatus(orders []Order) map[string]int {
	result := make(map[string]int)

	for _, order := range orders {
		result[order.Status] += order.Total
	}

	return result
}

func main() {
	orders := []Order{
		{Status: "paid", Total: 100},
		{Status: "new", Total: 50},
		{Status: "paid", Total: 70},
	}

	totals := sumByStatus(orders)

	fmt.Println(totals["paid"])
	fmt.Println(totals["new"])
}
```

Результат:

```text
170
50
```

Почему это работает:

* для нового ключа `map[string]int` возвращает `0`
* `+=` прибавляет сумму заказа
* по каждому статусу накапливается итог

---

## 18. DTO — простая модель передачи данных

DTO расшифровывается как Data Transfer Object.

На практике пока думай так:

> DTO — это простая структура, которая описывает форму данных.

Например, данные для создания пользователя:

```go
type CreateUserDTO struct {
	Name  string
	Email string
	Age   int
}
```

Пока DTO — это просто структура.

Без:

* методов
* интерфейсов
* JSON-тегов
* тестов

Пример проверки DTO:

```go
package main

import "fmt"

type CreateUserDTO struct {
	Name  string
	Email string
	Age   int
}

func isValidUser(dto CreateUserDTO) bool {
	if dto.Name == "" {
		return false
	}

	if dto.Email == "" {
		return false
	}

	if dto.Age <= 0 {
		return false
	}

	return true
}

func main() {
	valid := CreateUserDTO{
		Name:  "Anna",
		Email: "anna@example.com",
		Age:   30,
	}

	invalid := CreateUserDTO{
		Name: "Bob",
	}

	fmt.Println(isValidUser(valid))
	fmt.Println(isValidUser(invalid))
}
```

Результат:

```text
true
false
```

В будущих уроках DTO встретится рядом с JSON, файлами и HTTP.

Сейчас достаточно понять:

* DTO хранит данные
* структура делает форму данных явной
* проверка полей может быть обычной функцией

---

## 19. Экспортируемые поля

В уроке про пакеты уже было правило:

* имя с большой буквы экспортируется
* имя с маленькой буквы не экспортируется

Для полей структуры правило такое же.

```go
type User struct {
	ID   int
	Name string
	age  int
}
```

`ID` и `Name` начинаются с большой буквы.

`age` начинается с маленькой.

Внутри того же пакета доступны все поля.

Но если структура будет использоваться из другого пакета, экспортируемые поля важны.

Для будущего:

* DTO часто делают с экспортируемыми полями
* внутренние детали можно оставлять неэкспортируемыми

Глубоко проектировать пакеты в этом уроке не нужно.

Достаточно помнить:

> Правило большой буквы работает и для полей структур.

---

## 20. Где здесь backend

Структуры в backend-коде встречаются постоянно.

Они описывают:

* пользователя
* товар
* заказ
* запрос
* ответ
* конфигурацию
* промежуточный результат обработки

Например:

```go
type Product struct {
	ID       int
	Title    string
	Category string
	Price    int
	Active   bool
}
```

Такой тип сразу говорит:

* что такое товар
* какие поля у него есть
* какие типы у этих полей

Это лучше, чем держать данные в `map[string]string`, если поля заранее известны.

Структура даёт компилятору возможность помогать тебе.

А это очень недооценённая суперсила Go.

Если бы эти же данные лежали в `map[string]string`, цена стала бы строкой, активность стала бы строкой, а ошибка в имени ключа обнаружилась бы только во время работы программы.

Для backend-моделей это почти всегда хуже.

---

## 21. Что пока не трогаем

Чтобы урок не расползся, сегодня не используем:

* методы
* интерфейсы
* JSON-теги
* `encoding/json`
* тесты
* `any`
* embedding
* generics

Эти темы будут позже.

Сегодня цель проще:

> Научиться уверенно моделировать данные через `struct`.

---

## Типичные ошибки

### 1. Ждать, что структура изменится в функции без возврата или указателя

```go
func rename(user User) {
	user.Name = "Bob"
}
```

Эта функция меняет копию.

Чтобы изменение увидел вызывающий код:

* верни изменённую структуру
* или передай указатель

---

### 2. Менять копию из `range`

```go
for _, user := range users {
	user.Active = true
}
```

Так меняется копия.

Для изменения исходного `slice`:

```go
for i := range users {
	users[i].Active = true
}
```

---

### 3. Писать позиционные литералы там, где легко ошибиться

Плохо читается:

```go
user := User{1, "Anna", true}
```

Лучше:

```go
user := User{ID: 1, Name: "Anna", Active: true}
```

---

### 4. Забывать про zero value

```go
account := Account{Email: "a@example.com"}
```

Все незаполненные поля получат zero value.

Это может быть правильно.

А может быть скрытой ошибкой.

---

### 5. Читать поле через `nil`-указатель

```go
func printName(user *User) {
	fmt.Println(user.Name)
}
```

Если `user == nil`, будет `panic`.

Сначала проверка:

```go
if user == nil {
	return
}
```

---

## Мини-шпаргалка

```go
type User struct {
	ID   int
	Name string
}
```

Объявить структуру.

```go
user := User{ID: 1, Name: "Anna"}
```

Создать значение.

```go
fmt.Println(user.Name)
```

Прочитать поле.

```go
user.Name = "Ann"
```

Изменить поле.

```go
func rename(user User) User {
	user.Name = "Bob"
	return user
}
```

Вернуть изменённую копию.

```go
func activate(user *User) {
	if user == nil {
		return
	}

	user.Active = true
}
```

Изменить исходную структуру через указатель.

```go
for i := range users {
	users[i].Active = true
}
```

Изменить элементы `[]User`.

---

## Задачи

Решай задачи в отдельных маленьких файлах или папках.

После каждой задачи запускай:

```text
go run .
```

Если файл один:

```text
go run main.go
```

И не забывай:

```text
go fmt
```

---

### Задача 1. Первый `struct`

Объяви структуру:

```go
type User struct {
	ID   int
	Name string
}
```

Создай:

```go
var user User
```

Выведи:

* структуру целиком
* `user.ID`
* результат проверки `user.Name == ""`

Проверь, какие zero value получили поля.

---

### Задача 2. Карточка товара

Объяви структуру:

```go
type Product struct {
	ID    int
	Title string
	Price int
}
```

Создай товар через именованный литерал.

Данные:

```go
ID: 10
Title: "Keyboard"
Price: 5000
```

Выведи каждое поле отдельно.

---

### Задача 3. Неполный литерал

Объяви структуру:

```go
type Account struct {
	ID     int
	Email  string
	Active bool
}
```

Создай аккаунт, указав только `ID` и `Email`.

Потом:

* выведи `Active`
* измени `Active` на `true`
* выведи структуру целиком

---

### Задача 4. Копия структуры

Объяви:

```go
type User struct {
	Name string
	Age  int
}
```

Создай:

```go
first := User{Name: "Anna", Age: 30}
second := first
```

Измени `second.Name`.

Выведи `first` и `second`.

Перед запуском попробуй предсказать результат.

---

### Задача 5. Функция получает копию

Напиши функцию:

```go
func rename(user User, name string)
```

Внутри измени `user.Name`.

В `main`:

* создай пользователя
* вызови `rename`
* выведи пользователя

Объясни, почему имя не изменилось.

---

### Задача 6. Вернуть изменённую структуру

Напиши функцию:

```go
func renamed(user User, name string) User
```

Функция должна вернуть структуру с новым именем.

В `main` сохрани результат обратно в переменную.

---

### Задача 7. Изменить через указатель

Объяви:

```go
type User struct {
	Name   string
	Active bool
}
```

Напиши:

```go
func activate(user *User)
```

Требования:

* если `user == nil`, сразу выйти
* иначе установить `Active = true`

Проверь:

* `activate(&user)`
* `activate(nil)`

---

### Задача 8. Вложенная структура

Объяви:

```go
type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Address Address
}
```

Создай пользователя с адресом.

Выведи:

* имя
* город
* улицу

---

### Задача 9. Список пользователей

Объяви:

```go
type User struct {
	ID     int
	Name   string
	Active bool
}
```

Создай `[]User`:

```go
[]User{
	{ID: 1, Name: "Anna", Active: true},
	{ID: 2, Name: "Bob", Active: false},
	{ID: 3, Name: "Kate", Active: true},
}
```

Через `range` выведи поля каждого пользователя.

---

### Задача 10. Ловушка `range`

Создай `[]User` с полем `Active`.

Сначала попробуй сделать всех активными так:

```go
for _, user := range users {
	user.Active = true
}
```

Выведи результат.

Потом сделай правильно:

```go
for i := range users {
	users[i].Active = true
}
```

Сравни результаты.

---

### Задача 11. Найти пользователя по `ID`

Напиши:

```go
func findUserByID(users []User, id int) (User, bool)
```

Требования:

* если пользователь найден, вернуть его и `true`
* если не найден, вернуть `User{}` и `false`

Проверь два вызова:

* поиск `id == 2`
* поиск `id == 10`

Ожидаемое поведение:

* для `id == 2` должно получиться `ok == true`, а имя пользователя должно быть `"Bob"`
* для `id == 10` должно получиться `ok == false`

---

### Задача 12. Активные пользователи

Напиши:

```go
func activeUsers(users []User) []User
```

Функция должна вернуть новый `slice` только с активными пользователями.

Исходный `slice` не изменять.

Проверь на данных из задачи 9.

Ожидаемое поведение:

* в результате должны остаться `Anna` и `Kate`
* `Bob` в результат попасть не должен
* длина результата должна быть `2`

---

### Задача 13. Группировка товаров

Объяви:

```go
type Product struct {
	Title    string
	Category string
}
```

Напиши:

```go
func groupByCategory(products []Product) map[string][]Product
```

Данные:

```go
[]Product{
	{Title: "Keyboard", Category: "devices"},
	{Title: "Mouse", Category: "devices"},
	{Title: "Desk", Category: "furniture"},
}
```

Ожидаемое поведение:

* в группе `devices` должны быть `Keyboard` и `Mouse`
* в группе `furniture` должен быть `Desk`
* длина `groups["devices"]` должна быть `2`
* длина `groups["furniture"]` должна быть `1`

---

### Задача 14. Сумма заказов по статусу

Объяви:

```go
type Order struct {
	Status string
	Total  int
}
```

Напиши:

```go
func sumByStatus(orders []Order) map[string]int
```

Данные:

```go
[]Order{
	{Status: "paid", Total: 100},
	{Status: "new", Total: 50},
	{Status: "paid", Total: 70},
}
```

Ожидаемо:

* `paid` → `170`
* `new` → `50`

---

### Задача 15. DTO и проверка полей

Объяви:

```go
type CreateUserDTO struct {
	Name  string
	Email string
	Age   int
}
```

Напиши:

```go
func isValidUser(dto CreateUserDTO) bool
```

Правила:

* если `Name == ""`, вернуть `false`
* если `Email == ""`, вернуть `false`
* если `Age <= 0`, вернуть `false`
* иначе вернуть `true`

Проверь минимум три примера:

```go
valid := CreateUserDTO{Name: "Anna", Email: "anna@example.com", Age: 30}
emptyEmail := CreateUserDTO{Name: "Bob", Age: 20}
badAge := CreateUserDTO{Name: "Kate", Email: "kate@example.com", Age: 0}
```

Ожидаемое поведение:

* `valid` должен дать `true`
* `emptyEmail` должен дать `false`
* `badAge` должен дать `false`

---

### Задача 16. Мини-практика: каталог товаров

Объяви:

```go
type Product struct {
	ID       int
	Title    string
	Category string
	Price    int
	Active   bool
}
```

Данные:

```go
products := []Product{
	{ID: 1, Title: "Keyboard", Category: "devices", Price: 5000, Active: true},
	{ID: 2, Title: "Mouse", Category: "devices", Price: 2500, Active: true},
	{ID: 3, Title: "Desk", Category: "furniture", Price: 12000, Active: false},
	{ID: 4, Title: "Lamp", Category: "furniture", Price: 3000, Active: true},
}
```

Напиши функции:

```go
func findProductByID(products []Product, id int) (Product, bool)
func activeProducts(products []Product) []Product
func totalPrice(products []Product) int
func groupProductsByCategory(products []Product) map[string][]Product
func deactivateProduct(products []Product, id int) bool
```

Требования:

Для `findProductByID`:

* принять список товаров и `id`
* если товар найден, вернуть товар и `true`
* если товар не найден, вернуть `Product{}` и `false`
* проверить `id == 2` и `id == 99`

Ожидаемое поведение:

* для `id == 2` должен найтись `Mouse`
* для `id == 99` должен получиться `ok == false`

Для `activeProducts`:

* принять список товаров
* вернуть новый `slice` только с товарами, у которых `Active == true`
* исходный `slice` не изменять

Ожидаемое поведение:

* в результате должны быть `Keyboard`, `Mouse`, `Lamp`
* `Desk` в результат попасть не должен
* длина результата должна быть `3`

Для `totalPrice`:

* принять список товаров
* вернуть сумму всех `Price`

Ожидаемое поведение на исходных данных:

* сумма должна быть `22500`

Для `groupProductsByCategory`:

* принять список товаров
* вернуть `map[string][]Product`
* ключ — категория
* значение — товары этой категории

Ожидаемое поведение:

* `devices` содержит `Keyboard` и `Mouse`
* `furniture` содержит `Desk` и `Lamp`
* длина `groups["devices"]` равна `2`
* длина `groups["furniture"]` равна `2`

Для `deactivateProduct`:

* принять список товаров и `id`
* если товар найден, изменить `Active` на `false` у товара в исходном `slice`
* вернуть `true`, если товар найден
* вернуть `false`, если товар не найден
* если товар уже был неактивен, это не ошибка: всё равно можно вернуть `true`, если `id` найден

В `deactivateProduct` меняй элемент через индекс.

Проверь:

* `deactivateProduct(products, 2)` должен вернуть `true`
* после этого у товара `Mouse` поле `Active` должно стать `false`
* `deactivateProduct(products, 99)` должен вернуть `false`

Не используй:

* методы
* интерфейсы
* JSON
* тесты
* `any`

---

## Вопросы для проверки

1. Что такое `struct`?
2. Что такое поле структуры?
3. Что такое zero value структуры?
4. Чем именованный литерал лучше позиционного?
5. Как прочитать поле структуры?
6. Как изменить поле структуры?
7. Почему `second := first` копирует структуру?
8. Почему функция с параметром `User` получает копию?
9. Когда лучше вернуть изменённую структуру?
10. Когда нужен `*User`?
11. Почему перед работой с `*User` иногда нужна проверка `nil`?
12. Как создать вложенную структуру?
13. Что означает тип `[]User`?
14. Почему переменная в `for _, user := range users` является копией?
15. Как правильно менять элементы `[]User`?
16. Зачем при поиске возвращать `(User, bool)`?
17. Как работает группировка через `map[string][]Product`?
18. Что такое агрегация?
19. Что такое DTO на практическом уровне?
20. Почему методы и JSON-теги не нужны в этом уроке?

---

## Итог

После этого урока ты должен уверенно читать и писать такой код:

```go
type User struct {
	ID     int
	Name   string
	Active bool
}
```

И понимать:

* структура объединяет связанные поля
* поля читаются и меняются через точку
* незаполненные поля получают zero value
* структура копируется при присваивании
* структура копируется при передаче в функцию
* указатель на структуру позволяет изменить исходное значение
* `nil`-указатель нужно проверять
* `[]Struct` обрабатывается обычными циклами
* для изменения элемента `[]Struct` нужен индекс
* поиск, фильтрация, группировка и агрегация строятся на уже знакомых `for`, `if`, `slice`, `map` и функциях

Следующий шаг — сделать значения внутри структур ещё выразительнее.

Для этого понадобятся `const`, именованные типы и enum-like constants.
