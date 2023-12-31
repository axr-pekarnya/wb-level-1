# Уровень L1

## Теоретические вопросы

### 1. Какой самый эффективный способ конкатенации строк?

Приведём все возможные способы конкатенации строк в `Go`.

#### 1. Оператор "+"

```go
result := str1 + str2
```

#### 2. Форматирование строк

```go
result := fmt.Sprintf("%s%s", str1, str2)
```

#### 3. Метод `Join()` модуля `strings`

Наиболее локаничным и применимы этот метод будет по отношению к большому
количеству слов.

```go
result := string.Join([]string{str1, str2}, "")
```

#### 4. Тип `strings.Builder` и метод `WriteString()`

Наиболее эффективный метод, по своей сути схожий с `bytes.Buffer`, но при вызове
метода `String()` не происходит повторной аллокации памяти. Предварительной памяти, 
как и оптимизации для малого буфера нет, поэтому для достижения большей эффективности 
применяется метод `Grow(n int)`, где n - увеличиваемый размер буффера.

```go
var builder strings.Builder
builder.Grow(len(str1) + len(str2))
builder.WriteString(str1)
builder.WriteString(str2)
result := builder.String()
```

### 2. Что такое интерфейсы, как они применяются в `Go`?

Интерфейс - средство, предопределяющее или регламентирующее взаимодействие 
объектов. Справедливо замечание, что интерфейсы свойственны языкам программирования, содержащим
в себе парадигму объектно-ориентированного программирования. В отличие от главных представителей 
ООП - `C++` и `Java`, в которых используются классы с виртуальными функциями с последующими 
наследованием и перегрузкой и настоящий `interface` с последующим использованием ключевого 
слова `implements` соответственно, `Go`, как всегда, впереди планеты всей.

Интерфейс определяется вполне логично, разберёмся на примере.

```go
type Reader interface {
    Read() string
}
```

В отличие от вышеупомянутых способов определения структуры, удовлетворяющей интерфейсу, 
в `Go` всё локанично - cтруктура удовлетворяет интерфейсу, если для неё определены все методы 
интерфейса.

Пусть существует струтура

```go
type Scanner struct {
    buffer []byte
}

func (s *Scanner) Scan() {
    ...
}
```

тогда для удовлетворения интерфейсу `Reader` в структуру `Scanner` достаточно добавить метод 
`Read()`

```go
func (s *Scanner) Read() {
    ...
}
```

Напрашивается очевидная логическая цепочка: тип удовлетворяет интерфейсу, если содержит методы,
определённые в этом интерфейсы, следовательно, пустому интерфейсу удовлетворяет любой тип.

### 3. Чем отличаются `RWMutex` от `Mutex`?

В `Go` `sync.Mutex` и `sync.RWMutex` - это два примитива синхронизации, которые предоставляют механизмы синхронизации в Go, 
    но они различаются по своей функциональности
    и способу использования:

#### `Mutex` используется для обеспечения эксклюзивного доступа к ресурсу. 
Только одна горутина может удерживать мьютекс в определенный момент времени, 
что означает, что другие горутины должны ждать, пока мьютекс будет разблокирован.

    
```go
var mutex sync.Mutex
// Блокировка мьютекса
mutex.Lock()
// Выполнение критической секции
mutex.Unlock()
```

#### `RWMutex` предоставляет два уровня блокировки: чтение и запись. 
По сути, несколько горутин могут одновременно читать данные, если мьютекс находится в режиме чтения. 
Однако только одна горутина может владеть мьютексом в режиме записи.

```go
var rwMutex sync.RWMutex
// Блокировка для чтения
rwMutex.RLock()
// Выполнение операций чтения
rwMutex.RUnlock()
//
// Блокировка для записи
rwMutex.Lock()
// Выполнение операций записи
rwMutex.Unlock()
```

Таким образом, `RWMutex` полезен, когда у вас есть данные, которые часто считываются, но редко записываются, поскольку он позволяет параллельное чтение, 
но требует эксклюзивной блокировки для записи. `Mutex`, с другой стороны, обеспечивает исключительный доступ к ресурсу в любое время, 
что может быть более дорого в терминах производительности, если есть много операций чтения.

### 4. Чем отличаются буферизированные и не буферизированные каналы?

#### Не буферизированные каналы (unbuffered channels)

Каналы, которые не имеют буфера, называются не буферизированными или синхронными каналами. В не буферизированных каналах отправка и получение 
данных происходит синхронно. Отправка данных в не буферизированный канал блокируется до тех пор, пока другая горутина не готова принять данные 
из канала, и наоборот. Это гарантирует безопасность передачи данных между горутинами, но может вызвать блокировку, если отправка и получение не 
происходят одновременно.

```go
ch := make(chan int) // Создание не буферизированного канала
go func() {
    value := <-ch // Получение данных из канала
    //...
}()
ch <- 42 // Отправка данных в канал
```

#### Буферизированные каналы (buffered channels)

Буферизированные каналы имеют фиксированное количество слотов для данных (буфер). Отправка в буферизированный канал блокируется только в том случае, 
если буфер полностью заполнен. Получение данных из буферизированного канала блокируется только в том случае, если буфер полностью пуст. Это позволяет более 
гибко управлять асинхронной передачей данных между горутинами.

Пример буферизированного канала:
```go
ch := make(chan int, 3) // Создание буферизированного канала с буфером размером 3
go func() {
    value := <-ch // Получение данных из канала
    //...
}()
ch <- 42 // Отправка данных в канал (не блокируется)
```

### 5. Какой размер у структуры `struct{}{}`?

Структура `struct{}` не содержит полей и, следовательно, не занимает памяти для хранения данных. Ее размер равен нулю. Эту структуру можно использовать в Go,
например, как тип-индикатор или маркер для каналов или мапов, когда вам важно только наличие или отсутствие значения, но не его содержание.

Пример использования пустой структуры для создания мапа:
```go
myMap := make(map[string]struct{})
// Добавление элемента в мап (здесь значение фактически не имеет значения)
myMap["ключ"] = struct{}{}
// Проверка наличия ключа
if _, exists := myMap["ключ"]; exists {
    fmt.Println("Ключ существует")
}
```

### 6. Есть ли в `Go` перегрузка методов или операторов?

В Go нет поддержки перегрузки методов или операторов, как это может быть в некоторых других языках программирования, таких как C++ или Java. 
Вместо этого Go придерживается принципа "простоты" и "явности", предоставляя однострочное имя метода или функции для каждой операции. 
Это делает код более читаемым и предсказуемым.

### 7. В какой последовательности будут выведены элементы `map[int]int`?

Пример:
```go
m[0]=1
m[1]=124
m[2]=281
```

В Go, порядок итерации по элементам `map` является случайным и не гарантирован. Это связано с тем, что внутреннее представление карты оптимизировано 
для быстрого доступа, а не для определенного порядка элементов.

### 8. В чем разница между `make` и `new`?

В Go `make` и `new` используются для создания объектов разных типов, и у них разные назначения:

#### `make` 

Используется для создания `slice`, `map` и `channel`. Эти типы данных требуют инициализации, и `make` выделяет и инициализирует память для них, 
возвращая готовую к использованию переменную.

   ```go
   mySlice := make([]int, 10)       // Создание среза
   myMap := make(map[string]int)    // Создание карты
   myChannel := make(chan int)     // Создание канала
   ```

#### `new` 

Используется для создания указателей на нулевое значение типа данных. Он выделяет память и возвращает указатель на новый объект, инициализируя его 
нулевым значением (`0` для чисел, `nil` для указателей и ссылок на объекты).

   ```go
   var myIntPointer *int
   myIntPointer = new(int) // Создание указателя на int
   ```

### 9. Сколько существует способов задать переменную типа `slice` или `map`?

В Go существует несколько способов задать переменную типа среза (`slice`) и карты (`map`). Вот некоторые из них:

#### `slice`

1. Инициализация через литералы
   ```go
   mySlice := []int{1, 2, 3}
   ```

2. Использование функции `make` для создания среза с определенной длиной и возможно буфером
   ```go
   mySlice := make([]int, 5)
   ```

3. Создание среза из существующего массива
   ```go
   myArray := [3]int{1, 2, 3}
   mySlice := myArray[:]
   ```

#### `map`

1. Инициализация через литералы
   ```go
   myMap := map[string]int{
       "один": 1,
       "два": 2,
       "три": 3,
   }
   ```

2. Создание пустой `map` с использованием функции `make`
   ```go
   myMap := make(map[string]int)
   ```

3. Создание карты с указанием начальной ёмкости при использовании функции `make`
   ```go
   myMap := make(map[string]int, 10)
   ```

4. Создание карты с использованием литерала без инициализации значений
   ```go
   myMap := map[string]int{}
   ```

5. Создание карты с присвоением значений по одному элементу
   ```go
   myMap := make(map[string]int)
   myMap["один"] = 1
   myMap["два"] = 2
   ```

## Я вам комплиятор?

### 10. Что выведет данная программа и почему?

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```

Вывод программы: 1 1

При инициализации переменной `p` типа *int, её определили как ссылку на
 переменную `a` типа `int`, поэтому при первом выводе `p = 1`. Далее вызывается 
 функция `update(p *int)`, изменяющая указатель, но это не оказывается влияние на 
 указатель `p`, определённый в `main()`.

Примечание: внутри функции `update()` `p` ссылается на `b`.

### 11. Что выведет данная программа и почему?

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

Программа выведет числа от 0 до 4 в произвольном порядке и сообщит о возникшем
`deadlock`. Дело в том, что нужно передавать указатель на `wg`, чтобы объект
изменялся, в нашем случае `wg.Done` не срабатывает, горутины завершают работу,
а главня горутина блокируется вызовом `wg.Wait()`. 

Исправленный вариант вызова
анонимной функции.

```go
go func(wg *sync.WaitGroup, i int){
    //...
}(&wg, i)
```

### 12. Что выведет данная программа и почему?

```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```

Вывод программы: 0.

В первой строке инициализируется перменная `n := 0` в области видимости 
функции `main()`. Далее инициализируется другая переменная `n := 1` в области
видимости `if`.

### 13. Что выведет данная программа и почему?


```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```

Вывод: [100 2 3 4 5]

В `Go` тип данных `slice` ссылается на элементы. Всё становится совсем очевидно, 
если вывести адрес памяти нулевого элемента и самого `slice`. При вызове функции
`someAction(v []int8, b int8)` создаётся копия передаваемого `slice`, которая всё ещё
ссылается на те же ячейки памяти, на которые ссылается `a` в `main()`.

### 14. Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```

Вывод:[b b a][a a]

При вызове анонимной функции мы не передаём указатель на исходный `slice`, поэтому 
внутри этой функции работаем с копией, соответсвенно, операции не оказывают эффекта
на исходные данные.
