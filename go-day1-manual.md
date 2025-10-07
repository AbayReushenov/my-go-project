# День 1: Старт с Go - Подробный комментированный мануал

## 1. Установка Go

### Windows
1. Переходим на официальный сайт golang.org/dl
2. Скачиваем установщик для Windows (.msi файл)
3. Запускаем установщик и следуем инструкциям
4. По умолчанию Go установится в C:\Go
5. Проверяем установку командой:
   ```bash
   go version
   ```

### macOS
1. Переходим на golang.org/dl
2. Скачиваем .pkg файл для macOS
3. Запускаем установщик (установит в /usr/local/go)
4. Проверяем установку:
   ```bash
   go version
   ```

### Linux (Ubuntu/Debian)
1. Скачиваем архив с golang.org/dl:
   ```bash
   wget https://golang.org/dl/go1.21.x.linux-amd64.tar.gz
   ```
2. Удаляем старую версию (если есть):
   ```bash
   sudo rm -rf /usr/local/go
   ```
3. Распаковываем:
   ```bash
   sudo tar -C /usr/local -xzf go1.21.x.linux-amd64.tar.gz
   ```
4. Добавляем в PATH (в ~/.bashrc или ~/.profile):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   source ~/.bashrc
   ```
5. Проверяем:
   ```bash
   go version
   ```

## 2. Базовая структура Go-проектов

### Структура для небольших проектов
```
/myproject
├── main.go
├── go.mod
└── README.md
```

### Структура для проектов со вспомогательными пакетами
```
/myproject
├── internal/
│   └── helper/
│       └── helper.go
├── main.go
├── go.mod
└── README.md
```

### Структура для больших проектов
```
/myproject
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── handler/
│   └── service/
├── pkg/
│   └── utils/
├── go.mod
├── Makefile
└── README.md
```

**Ключевые принципы:**
- `main.go` - точка входа в программу
- `internal/` - пакеты, недоступные для импорта извне
- `cmd/` - исполняемые файлы
- `pkg/` - публичные библиотеки
- `go.mod` - файл модуля с зависимостями

## 3. Создание первого проекта

### Шаг 1: Создание директории и инициализация модуля
```bash
mkdir hello-go
cd hello-go
go mod init hello-go
```

### Шаг 2: Создание main.go
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Шаг 3: Запуск программы
```bash
# Запуск без компиляции
go run main.go

# Компиляция в бинарный файл
go build main.go
# На Windows создастся main.exe, на Linux/macOS - main

# Запуск скомпилированного файла
./main  # Linux/macOS
main.exe  # Windows
```

## 4. Работа с переменными

### Объявление переменных
```go
package main

import "fmt"

func main() {
    // Способ 1: var + тип
    var name string = "Golang"
    var age int = 15
    
    // Способ 2: var с автоопределением типа
    var language = "Go"
    var version = 1.21
    
    // Способ 3: короткое объявление (:=)
    developer := "Google"
    year := 2009
    
    // Несколько переменных одновременно
    var x, y int = 10, 20
    a, b := "hello", "world"
    
    // Константы
    const pi = 3.14159
    const greeting = "Привет"
    
    fmt.Println(name, age, language, version)
    fmt.Println(developer, year)
    fmt.Println(x, y, a, b)
    fmt.Println(pi, greeting)
}
```

**Сравнение с JavaScript:**
```javascript
// JavaScript
let name = "Golang";        // Go: name := "Golang"
const age = 15;             // Go: const age = 15
var language = "Go";        // Go: var language = "Go"
```

### Основные типы данных
```go
package main

import "fmt"

func main() {
    // Числовые типы
    var integer int = 42
    var float32Number float32 = 3.14
    var float64Number float64 = 2.718281828
    
    // Строки
    var text string = "Hello, Go!"
    
    // Булев тип
    var isActive bool = true
    
    // Массивы
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    
    // Слайсы (динамические массивы)
    var fruits []string = []string{"apple", "banana", "orange"}
    
    fmt.Printf("Целое: %d\n", integer)
    fmt.Printf("Float32: %.2f\n", float32Number)
    fmt.Printf("Float64: %.9f\n", float64Number)
    fmt.Printf("Строка: %s\n", text)
    fmt.Printf("Булево: %t\n", isActive)
    fmt.Printf("Массив: %v\n", numbers)
    fmt.Printf("Слайс: %v\n", fruits)
}
```

## 5. Функции

### Базовый синтаксис функций
```go
package main

import "fmt"

// Простая функция без параметров
func sayHello() {
    fmt.Println("Привет!")
}

// Функция с параметрами
func greet(name string) {
    fmt.Printf("Привет, %s!\n", name)
}

// Функция с возвращаемым значением
func add(a int, b int) int {
    return a + b
}

// Функция с несколькими возвращаемыми значениями
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Функция с именованными возвращаемыми значениями
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // автоматически вернет sum и product
}

func main() {
    sayHello()
    greet("Гофер")
    
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", quotient)
    }
    
    s, p := calculate(4, 7)
    fmt.Printf("4 + 7 = %d, 4 * 7 = %d\n", s, p)
}
```

**Сравнение с JavaScript:**
```javascript
// JavaScript
function add(a, b) {
    return a + b;
}

// Go
func add(a int, b int) int {
    return a + b
}
```

## 6. Структуры

### Определение и использование структур
```go
package main

import "fmt"

// Определение структуры
type Person struct {
    Name    string
    Age     int
    Email   string
    IsAdmin bool
}

// Метод для структуры
func (p Person) Introduce() {
    fmt.Printf("Привет! Меня зовут %s, мне %d лет\n", p.Name, p.Age)
}

// Метод с указателем (может изменять структуру)
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%s теперь %d лет!\n", p.Name, p.Age)
}

func main() {
    // Создание структуры - способ 1
    var person1 Person
    person1.Name = "Алексей"
    person1.Age = 25
    person1.Email = "alexey@example.com"
    person1.IsAdmin = false
    
    // Создание структуры - способ 2
    person2 := Person{
        Name:    "Мария",
        Age:     30,
        Email:   "maria@example.com",
        IsAdmin: true,
    }
    
    // Создание структуры - способ 3
    person3 := Person{"Дмитрий", 28, "dmitry@example.com", false}
    
    fmt.Printf("Пользователь 1: %+v\n", person1)
    fmt.Printf("Пользователь 2: %+v\n", person2)
    fmt.Printf("Пользователь 3: %+v\n", person3)
    
    // Вызов методов
    person1.Introduce()
    person2.Introduce()
    
    person1.HaveBirthday()
}
```

**Сравнение с JavaScript:**
```javascript
// JavaScript (класс)
class Person {
    constructor(name, age, email, isAdmin) {
        this.name = name;
        this.age = age;
        this.email = email;
        this.isAdmin = isAdmin;
    }
    
    introduce() {
        console.log(`Привет! Меня зовут ${this.name}, мне ${this.age} лет`);
    }
}

// Go (структура)
type Person struct {
    Name    string
    Age     int
    Email   string
    IsAdmin bool
}

func (p Person) Introduce() {
    fmt.Printf("Привет! Меня зовут %s, мне %d лет\n", p.Name, p.Age)
}
```

## 7. Импорт пакетов

### Стандартные пакеты
```go
package main

import (
    "fmt"      // Форматированный ввод/вывод
    "time"     // Работа со временем
    "strings"  // Работа со строками
    "math"     // Математические функции
    "os"       // Взаимодействие с ОС
)

func main() {
    // Использование fmt
    fmt.Println("Привет, мир!")
    
    // Использование time
    now := time.Now()
    fmt.Printf("Сейчас: %s\n", now.Format("2006-01-02 15:04:05"))
    
    // Использование strings
    text := "  Hello, Go!  "
    fmt.Printf("Обрезанная строка: '%s'\n", strings.TrimSpace(text))
    fmt.Printf("Верхний регистр: %s\n", strings.ToUpper(text))
    
    // Использование math
    fmt.Printf("Квадратный корень из 16: %.2f\n", math.Sqrt(16))
    fmt.Printf("Число Пи: %.6f\n", math.Pi)
    
    // Использование os
    fmt.Printf("Переменная PATH: %s\n", os.Getenv("PATH"))
}
```

### Создание и импорт своих пакетов
```go
// Файл: internal/calculator/calculator.go
package calculator

// Экспортируемая функция (начинается с заглавной буквы)
func Add(a, b int) int {
    return a + b
}

// Неэкспортируемая функция (начинается со строчной буквы)
func multiply(a, b int) int {
    return a * b
}

// Экспортируемая функция, использующая неэкспортируемую
func Square(n int) int {
    return multiply(n, n)
}
```

```go
// Файл: main.go
package main

import (
    "fmt"
    "hello-go/internal/calculator"  // Импорт локального пакета
)

func main() {
    result := calculator.Add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    square := calculator.Square(4)
    fmt.Printf("4² = %d\n", square)
    
    // Это не сработает - функция не экспортирована:
    // calculator.multiply(2, 3) // Ошибка компиляции
}
```

## 8. Пошаговые команды для первого дня

### Полный процесс создания проекта:

1. **Создание проекта:**
   ```bash
   mkdir my-first-go-app
   cd my-first-go-app
   go mod init my-first-go-app
   ```

2. **Создание main.go:**
   ```bash
   # В любом текстовом редакторе создайте main.go с содержимым:
   ```
   ```go
   package main
   
   import "fmt"
   
   func main() {
       fmt.Println("Мое первое Go приложение!")
   }
   ```

3. **Запуск и компиляция:**
   ```bash
   # Запуск без компиляции
   go run main.go
   
   # Компиляция
   go build
   
   # Запуск скомпилированного файла
   ./my-first-go-app    # Linux/macOS
   my-first-go-app.exe  # Windows
   ```

4. **Расширение функциональности:**
   ```bash
   # Создайте папку для внутренних пакетов
   mkdir -p internal/utils
   ```
   
   ```go
   // internal/utils/helper.go
   package utils
   
   import "fmt"
   
   func Greet(name string) {
       fmt.Printf("Привет, %s! Добро пожаловать в Go!\n", name)
   }
   ```
   
   ```go
   // main.go (обновленная версия)
   package main
   
   import (
       "fmt"
       "my-first-go-app/internal/utils"
   )
   
   func main() {
       fmt.Println("Мое первое Go приложение!")
       utils.Greet("Разработчик")
   }
   ```

## 9. Менторские советы для перехода с JavaScript на Go

### Основные различия:

1. **Типизация:**
   - JavaScript: динамическая типизация
   - Go: статическая типизация
   
2. **Объявление переменных:**
   ```javascript
   // JavaScript
   let name = "John";
   const age = 30;
   ```
   ```go
   // Go
   name := "John"
   const age = 30
   ```

3. **Функции:**
   ```javascript
   // JavaScript
   function add(a, b) {
       return a + b;
   }
   ```
   ```go
   // Go
   func add(a int, b int) int {
       return a + b
   }
   ```

4. **Объекты vs Структуры:**
   ```javascript
   // JavaScript
   const person = {
       name: "John",
       age: 30,
       greet() {
           console.log(`Hi, I'm ${this.name}`);
       }
   };
   ```
   ```go
   // Go
   type Person struct {
       Name string
       Age  int
   }
   
   func (p Person) Greet() {
       fmt.Printf("Hi, I'm %s\n", p.Name)
   }
   ```

5. **Обработка ошибок:**
   ```javascript
   // JavaScript
   try {
       const result = riskyOperation();
   } catch (error) {
       console.error(error);
   }
   ```
   ```go
   // Go
   result, err := riskyOperation()
   if err != nil {
       fmt.Printf("Ошибка: %v\n", err)
       return
   }
   ```

### Советы для первого дня:

1. **Не торопитесь с сложными концепциями** - сосредоточьтесь на синтаксисе
2. **Привыкайте к статической типизации** - она поможет избежать ошибок
3. **Изучите соглашения Go** - имена пакетов, экспорт функций
4. **Практикуйтесь с простыми программами** - hello world, калькулятор
5. **Используйте go fmt** для автоматического форматирования кода

## 10. Полезные команды Go

```bash
go version          # Версия Go
go env              # Переменные окружения Go
go mod init <name>  # Инициализация модуля
go mod tidy         # Очистка зависимостей
go build            # Компиляция
go run <file>       # Запуск без компиляции
go fmt              # Форматирование кода
go vet              # Статический анализ
go test             # Запуск тестов
go get <package>    # Установка пакета
```

## Заключение

Первый день изучения Go должен быть посвящен:
1. Установке и настройке среды
2. Пониманию базовой структуры проектов
3. Изучению синтаксиса переменных, функций и структур
4. Созданию первого работающего приложения
5. Сравнению с привычными концепциями JavaScript

Помните: Go проще JavaScript во многих аспектах, но требует более строгого подхода к типам и структуре кода.