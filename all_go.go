package main
import "fmt"
func zero() int {
	return 0
}
var test = func() func() int {
	return zero
}
func main() {
	fmt.Println(test()()) // 0
}



package main
import "fmt"
func main() {
	a := 1
	p := &a
	*p = 2
	a = 3
	fmt.Println(*p) // 3
}


package main
import "fmt"
func po(i *int) **int {
	return &i
}
func main() {
	i := 1
	p := po(&i)
	fmt.Println(**p) //1
}

package main
import "fmt"
var glob struct {
	name string
}
func main() {
	glob = struct {
		name string
	}{
		name: "a",
	}
	fmt.Println(glob.name)  //a
}


package main
import "fmt"
type glob struct {
	name string
}
func main() {
	type local glob
	fmt.Println(local(glob{name: "local"})) //{local}  // просто приведение типа
}
func (g glob) String() string {
	return g.name
}


package main
import (
	"encoding/json"
	"fmt"
)
const js = `[1, 2, 3, 4]`
func main() {
	a := [3]int{}
	json.Unmarshal([]byte(js), &a)
	fmt.Println(a[0]) // 1 , тк Unmarshal проверяем получателя!
}

package main
import (
	"fmt"
)
type Record struct {
	text string
}
func main() {
	a := []Record{{"text"}}
	a[0].text = ""
	fmt.Println(a)  // [{}] срез из пустой структуры
}


package main
import "fmt"
func main() {
	var x float32
	if x == 0 {
		fmt.Println("0") // 0 
	} else if x == 0.0 {
		fmt.Println("0.0")
	}
}

switch x {
case 1:
   fmt.Println("x = 1")
   fallthrough
case 2:
   fmt.Println("x = 2")
}

switch {
case x == 1:
	fmt.Println("x = 1")
case x == 2:
	fmt.Println("x = 2")
default:
}

type House struct {
    Number 
    BuildDate string
}


type Worker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	workerInfo := Worker{Name: "Ваня", Age: 14, Job: "Go-разработчик"}
	jsonInfo, err := json.Marshal(workerInfo)
	if err != nil {
		fmt.Println("Ошибка записи данных:", err)
	}
	fmt.Println(jsonInfo)
}

type Point struct {
    X int
    Y int
}
func main() {
    in:=Point{}
    fmt.Scan(&(in.X), &(in.Y))
    fmt.Println(in)
}


type Struct struct {
	Name string
}
func (s *Struct) ClearName() { передача по ссылке
	s.Name = "1"
}
func main() {
	var s Struct
	s.Name = "Struct"
	s.ClearName() //  s.ClearName() => (&s).ClearName() авт разименовывание
	fmt.Print(s.Name) // 1
	(&s).ClearName()
	fmt.Print(s.Name) // 1
}

type Struct struct {
	Name string
}
func (s Struct) ClearName() {
	s.Name = "1" // передача по значению
}
func main() {
	var s Struct
	s.Name = "Struct"
	s.ClearName()  // =>  (&s).ClearName() - это автоматического разыменования, встроенному в язык Go.
	fmt.Print(s.Name) // Struct
	(&s).ClearName()  // (&s) явно передаёт указатель на структуру методу.
	fmt.Print(s.Name) // Struct
}


var Phone struct {
	Model string
	Year  int
}

//Определение и инициализация переменной P:
//Здесь определяется и сразу же вызывается анонимная функция
var P = func() interface{} {
	Phone.Model = "IPhone"  //Устанавливает поле Model структуры Phone в значение "IPhone".
	return interface{}(nil) // Возвращает тип interface{}), но это значение не используется.
}()
//Можно упростить
var P = func() struct{} {
	Phone.Model = "IPhone"
	return struct{}{}  //struct{} определяет тип пустой структуры.
                       //{} создаёт значение пустой структуры.
}()

func main() {
	fmt.Println(Phone)
}



type glob struct {
     glob struct {
        Name string
     }
}
func main() {
	fmt.Println(glob{struct{ Name string }{Name: "main"}}) // {{main}} ХЗ почему
}


type Record struct {
	text string
}
func main() {
	m := make(map[int]Record)
	r := m[0]
	r.text = "hello"
	fmt.Println(m[0].text) // пустая строка
}

type Record struct {
	text string
}
func main() {
	m := make(map[int]Record)
	r := Record{}
	m[0] = r
	r.text = "hello"
	fmt.Println(m[0].text) //путая строка
}

func main() {
	m := make(map[interface{}]interface{})
	m["a"] = 1
	m[2] = "b"
	for k, v := range m {
		fmt.Print(k, v) // a1a2 or a2a1
	}
}

import (
	"encoding/json"
	"fmt"
)
const js = `{"A": 1, "b": 2}`
func main() {
	m := make(map[string]int)
	_ = json.Unmarshal([]byte(js), &m)
	fmt.Print(m["A"]) // 1
	fmt.Print(m["b"]) // 2
}

func main() {
	var a any
	v, ok := a.(int)
	fmt.Println(v, ok) // 0 false
}

func main() {
	var a any
	var b any
	a = 1
	b = "1"
	c := 1
	fmt.Println(a == b, a == c, b == c) // false true false
}


import "fmt"
fmt.Print(name)
    var i int 
    fmt.Scan(&i)
    fmt.Print(i)

var emptyInterface interface{}


import (
	"fmt"
	"sync"
)

func main() {
	// создаем группу ожидания
	// фактически это счетчик, который мы можем ожидать пока он не станет равен 0
	var wg sync.WaitGroup
	wg.Add(2)
	// запуск горутины 1
	go func() {
		fmt.Println("run goroutine 1")
		wg.Done()
	}()
	// запуск горутины 2
	go func() {
		fmt.Println("run goroutine 2")
		wg.Done()
	}()
	// порядок вывода сообщений предсказать точно невозможно
	fmt.Println("start")
	// ожидаем когда обе горутины завершатся
	wg.Wait()
	fmt.Println("end")
}


func main() {
	a := 1
	go func() {
		fmt.Println(a)
	}()
	a = 2
	fmt.Println(a)
	time.Sleep(time.Second)
}


func main() {
    var count int
    fmt.Scan(&count)
    //arr:=make([]int, 0)
	arr:=[]int{}
    for i:=1; i<=count; i++ {
        var x int
        fmt.Scan(&x)
        if i%3 == 0 {
            arr = append(arr, x)
        }
    }

    for _, value := range arr {
    	fmt.Println(value)
	}
}


func main() {
    var count int
    fmt.Scan(&count)
    h := map[string]int{}

    var keys []string
    // заполните срез ключей и карту
    for i := 0; i < count; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d", &k,&v)
        h[k] = v
        keys = append(keys, k)
    }
    for _, key := range keys {
        // выведите пары ключ-значение
        fmt.Println(key,"->", h[key])
    }
}

    m := map[int]string{}
    for k, v := range h {
        m[v] = k
        fmt.Println(v,k)
    }    


func main() {
    var count int
    fmt.Scan(&count)

    h := map[string]int{}
    
    keys := []string{}
    
    for i := 0; i < count; i++ {
        var key string
        var value int
        fmt.Scan(&key, &value)
        h[key] = value
        keys = append(keys, key)
    }
    //m := map[int]string{}
    for i := 0; i < len(keys); i++ {
        fmt.Println(h[keys[i]],keys[i])
    }    
}


    sl :=[]Entry{}
    
    var E Entry
    for k, v := range h {
        E.key = k
        E.value = v
        sl = append(sl, E)
    }
    fmt.Println(sl)


type Entry struct {
    key string
    value int
}

func main() {
    var count int
    fmt.Scan(&count)

    h := map[string]int{}
    sl_key :=[]string{}
    
    for i := 0; i < count; i++ {
        var key string
        var value int
        fmt.Scan(&key, &value)
        h[key] = value
        sl_key=append(sl_key, key)
    }
    sl :=[]Entry{}
    
    var E Entry
    for i := 0; i < len(sl_key); i++ {
        E.key = sl_key[i] 
        E.value = h[sl_key[i]]
        sl = append(sl, E)
    }
    fmt.Println(sl)


}

type Register interface {
    Set(int)
    Get() int
}

type Storage struct {
}

func (s *Storage) Get () int {
    return 1
}

func (s *Storage) Set (i int)  {    
}


import (
	"fmt"
	"math"
)
type Register interface {
	Get() float64
	Set(int)
}
type Storage struct {
	val int
}
func (s *Storage) Get() float64 {
	return math.Sqrt(float64(s.val))
}
func (s *Storage) Set(i int) {
	s.val = i
}
func main() {
	s := Storage{}
	s.Set(9)
	fmt.Println(s.Get())
}


package main
import "fmt"
func main() {
    var a,b,x int
    fmt.Scanf("%d %d", &a,&b)
    fmt.Scanf("%d", &x)
    if x>=a && x <= b {
        fmt.Println("В диапазоне")
    } else {
        fmt.Println("Не в диапазоне")
    }
}

func main() {
    arr:= [3]int{}
    fmt.Println(arr)
}


func odd(d int) string{
    if d%2 == 0 {
        return "Чет"
    } else {
        return "Нечет"
    }    
}
func main() {
    var a,b int
    fmt.Scanf("%d %d", &a,&b)
    fmt.Println(odd(a), odd(b))
}

var a,cnt int
fmt.Scan(&a)

	for a > 0 {
		a = a / 10
		cnt++
	}
fmt.Println(cnt)


import (
	"fmt"
	"strconv"
)
func main() {
	var a int64 = 5
	fmt.Println(strconv.FormatInt(a, 2)) //вывод в двочном коде
	a = a >> 1
	fmt.Println(strconv.FormatInt(a, 2))
}

func main() {
    var x,y int64
    fmt.Scanf("%d %d",&x,&y)

    if (x == y) {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}


ch := make(chan int, 3)

fmt.Println(len(ch)) // 0
ch <- 1
ch <- 2
fmt.Println(cap(ch)) // 3
<-ch
close(ch)
fmt.Println(<-ch) // 2


func main() {
	// объявление небуферизованного канала
	ch := make(chan int)
	go func() {
		// запись в канал
		// операция заблокируется пока не будет готов читатель из канала
		ch <- 1
	}()
	// чтение из канала
	// операция заблокируется пока не будет готов писатель в канал
	x := <-ch
	fmt.Println(x) // 1
	go func() {
		ch <- 2
		close(ch)
	}()
	// чтение из закрытого непустого канала
	one := <-ch // 2
	two := <-ch // 0
	fmt.Println(one, two) // 2 0
}

func main() {
	ch := make(chan int)
	out := make(chan struct{})
	// Первая горутина: записывает значение в канал ch
	go func() {
		ch <- 42
	}()
	// Вторая горутина: читает значение из канала ch и выводит его на экран
	go func() {
		value := <-ch
		fmt.Println(value)
		close(out) // Закрытие сигнального канала out после чтения значения
	}()
	// Ожидание закрытия сигнального канала out
	<-out
}

//последовательный вызов Горутин 1 и 2
func main() {
    var x, y int
    fmt.Scan(&x, &y)
    ch := make(chan int, 2)
    done := make(chan struct{})
    go func() {
        ch <- x  //пишем 1 число
    }()
    go func() {        
        x1, _ := <-ch //ждем первое число
        ch <- x1      // отправляем 1 
        ch <- y       // отправляем 2 число
        close(done)// Закрытие сигнального канала out после вывода всех значений
    }()        
    <-done    
    fmt.Println(<-ch)
    fmt.Println(<-ch)    
}



x, y := 10, 11
	ch := make(chan int, 2)
	sck := make(chan struct{})

	go func() {
		ch <- x // пишем 1 число
		sck <- struct{}{}
	}()

	go func() {
		<-sck      // ждем первое число
		ch <- y    // отправляем 2 число
		close(sck) // Закрытие сигнального канала sck
	}()

	<-sck // Ожидание отправки первого числа
	fmt.Println(<-ch)
	fmt.Println(<-ch)


func main() {
	x, y := 10, 11
	ch := make(chan int, 2)
	sck := make(chan struct{})
	done := make(chan struct{})

	go func() {

		ch <- x //пишем 1 число
		sck <- struct{}{}

	}()

	go func() {
		<-sck       //ждем первое число
		ch <- y     // отправляем 2 число
		close(done) // Закрытие сигнального канала out
	}()

	<-done
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


На вход программы подается N целых чисел и записываются в буферизованный канал.
Необходимо прочитать все значения из канала с помощью функции range.

func main() {
    var num int
    fmt.Scan(&num)
    ch := make(chan int, num)
    for i := 0; i < num; i++ {
        var x int
        fmt.Scan(&x)
        ch <- x
    }
    close (ch)
    for x := range ch {
        fmt.Println(x)
    }
}


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup
	// Функция для записи чисел от 1 до 10 в канал
	writeNumbers := func(ch chan int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}
	// Запускаем две горутины для записи данных в ch1 и ch2
	go writeNumbers(ch1)
	go writeNumbers(ch2)
	// Запускаем горутину для перенаправления данных из ch1 в out
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch1 {
			out <- num
		}
	}()
	// Запускаем горутину для перенаправления данных из ch2 в out
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch2 {
			out <- num
		}
	}()
	// Запускаем горутину для закрытия канала out после записи всех данных
	go func() {
		wg.Wait()
		close(out)
	}()
	// Читаем и выводим все значения из канала out до его закрытия
	for num := range out {
		fmt.Println(num)
	}
}


import (
	"fmt"
	"sync"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        for i := 0; i < 10; i+=2 {
            ch1 <- i
        }
        close(ch1)
    }()

    go func() {
        for i := 1; i < 10; i+=2 {
            ch2 <- i
        }
        close(ch2)
    }()

    out := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)
	go func() {
		defer wg.Done()
        for i := 1; i < 10; i+=2 {
            x1:= <- ch1
            x2:= <- ch2
            out <- x1
            out <- x2
        }
	}()
	// Запускаем горутину для закрытия канала out после записи всех данных
	go func() {
		wg.Wait()
		close(out)
	}()

    for v := range out {
        fmt.Println(v)
    }
}


// Запускаем горутину для перенаправления данных из ch1 и ch2 в out через select
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case x1, ok := <-ch1:
				if ok {
					out <- x1
				} else {
					ch1 = nil
				}
			case x2, ok := <-ch2:
				if ok {
					out <- x2
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()


// Запускаем горутину для перенаправления данных из ch1 и ch2 в out
	go func() {
		defer wg.Done()
		for {
				x1, is_open:= <-ch1
		  		if is_open {
		    		out <- x1
			    } else {
			   		ch1 = nil
				}

		  		x2, is_open:= <-ch2
		  		if is_open {
		    		out <- x2
		   		} else {
		   			ch2 = nil
		   		}

				if ch1 == nil && ch2 == nil {
					break
				}
			}
	}()



--------------------------------------------------------------------
Карты
func main() {
    m := map[string]int{"a":1,"b":2,"c":3}
    vals :=[]string{"a","b","c"}
    for _, x:= range (vals ){
        fmt.Println(x, m[x])
    }    
}

func main() {
    var num int
    fmt.Scan(&num)
    vals := []string{}
    mapp := map[string]int{}
    var k string
    var v int
    for i:=0;i<num;i++ {
        fmt.Scanf("%s %d", &k, &v)
        mapp[k]=v*v
        vals=append(vals,k)
    }
    for _, k :=range vals {
        fmt.Println(k, mapp[k])
    }
}


// Функция для поиска пересекающихся ключей с одинаковыми значениями
func intersect(map1, map2 map[string]int) map[string]int {
	intersection := make(map[string]int)
	for key, value := range map1 {
//		if v, ok := map2[key]; ok && v == value {
		if _, ok := map2[key]; ok {
			intersection[key] = value
		}
	}
	return intersection
}

//пересечение двух срезов
func main() {
	map1 := make(map[string]int)
    map2 := make(map[string]int)
    var cnt1, cnt2 int
    fmt.Scan(&cnt1)
    var k string
    var v int    
    for i:=0; i<cnt1; i++ {
        fmt.Scanf("%s %d", &k, &v)
        map1[k]=v
    }
    fmt.Scan(&cnt2)
    for i:=0; i<cnt2; i++ {
        fmt.Scanf("%s %d", &k, &v)
        map2[k]=v
    }    
    for k,v := range intersect(map1,map2){
        fmt.Println(k,v)
    }
}



func pow(s, p int) int { 
    if p== 0 {
        return 1
    }
    res :=s
    for i:=2;i<=p;i++{
        res*=s
    }
    return res
}


// recurSum вычисляет сумму элементов в срезе a рекурсивно
func recurSum(a []int) int {
	// Базовый случай: если длина среза равна 0, возвращаем 0
	if len(a) == 0 {
		return 0
	}
	// Рекурсивный случай: сумма первого элемента и рекурсивная сумма оставшихся элементов
	return a[0] + recurSum(a[1:])
}

func div(a, b int) (int, int) {
	return a / b, a % b
}


func drawCross(a, b int) {
	for i := 0; i < a; i++ { //строка
		for j := 0; j < b; j++ {
			if i == a / 2 {
				fmt.Print("*")
			} else if j == b / 2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func getElement(a string, i int) string {
	res := ""
	if i >= 0 {
		res = string(a[i])
	} else {
		res = string(a[len(a)+i])
	}
	return res
}

// split разделяет строку a на две строки по индексу i
func split(a string, i int) (string, string) {
	if i < 0 || i > len(a) {
		// Если индекс выходит за границы строки, возвращаем пустые строки
		return "", ""
	}
	return a[:i], a[i:]
}

func filter() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i < 5; i++ {
		fmt.Scan(&b)
		if b > a {
			fmt.Println(b)
		}
	}
}

func div() {
	var x, N, res int
	fmt.Scan(&x, &N)
	res = x
	for i := 1; i <= N; i++ {
		res = res / 2
	}
	fmt.Println(res)
}

func isEnd() {
	var cur, end string
	for {
		fmt.Scan(&cur)
		if cur == "end" {
			fmt.Println(end)
			break
		}
		end = cur
	}
}

AZa

func printABC() {
	for i := 65; i <= 90; i++ {
		fmt.Println(string(i), string(i+32))
	}
}

type Money struct {
	Name  string
	Value int
}
    var m Money
    fmt.Scanf("%s %d", &m.Name, &m.Value)
    fmt.Println(m.Value, m.Name)


type Human struct {
	Name string
}
// Определение структуры Alien, встраивающей структуру Human
type Alien struct {
	Human
}


//----------------------------------------------------
type Sender interface {
	Send()
}
type EmailSender struct {
	name string
}
func (s EmailSender) Send() {
	fmt.Println("email: Hello", s.name)
}
type PushSender struct {
	name string
}
func (s PushSender) Send() {
	fmt.Println("push: Hello", s.name)
}

func sendAB(str, name string) {
	var sender Sender
	if str == "email" {
		es := EmailSender{}
		es.name = name
		sender = es
	}
	if str == "push" {
		ps := PushSender{}
		ps.name = name
		sender = ps
	}
	sender.Send()
}

//-----------------------------------------
// Matrix представляет собой тип матрицы
type Matrix struct {
	rows int
	cols int
	data [][]int
}

// NewMatrix создает новую матрицу размером m x n и заполняет ее нулями
func NewMatrix(m, n int) *Matrix {
	data := make([][]int, m)
	for i := range data {
		data[i] = make([]int, n)
	}
	return &Matrix{
		rows: m,
		cols: n,
		data: data,
	}
}

// String возвращает строковое представление матрицы
func (matrix *Matrix) String() string {
	var sb strings.Builder
	for _, row := range matrix.data {
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%d ", val))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
func shoMatrix() {
	m := 3 // Количество строк
	n := 3 // Количество столбцов
	matrix := NewMatrix(m, n)
	fmt.Print(matrix)
}

//------------------------------------------



// Matrix представляет собой тип матрицы
type Matrix struct {
	rows int
	cols int
}

func matrix(m, n int) *Matrix {
	return &Matrix{
		rows: m,
		cols: n,
	}
}


func (matrix *Matrix) String() string {
	var sb strings.Builder
	for i := 0; i < matrix.rows; i++ {
		for j := 0; j < matrix.cols; j++ {
			sb.WriteString(fmt.Sprint("0"))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}


func showMatrix() {
	m := 3 // Количество строк
	n := 4 // Количество столбцов
	_matrix := matrix(m, n)
	fmt.Print(_matrix)
}

294

//пропущенное число
func findGap() {
	var cnt, cur int
	arr := []int{}
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&cur)
		arr = append(arr, cur)
	}
	cur = arr[0]
	for i := 1; i < cnt; i++ {
		if arr[i] != cur+1 {
			fmt.Println(i)
			return
		}
		cur = arr[i]
	}
	fmt.Println(-1)
}



func findZeros(nums []int) []int {
	arrZero := []int{}
	for i, v := range nums {
		if v == 0 {
			arrZero = append(arrZero, i)
		}
	}
	return arrZero
}

func main() {  
    var cnt, cur int
	arr := []int{}
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&cur)
		arr = append(arr, cur)
	}

    fmt.Println(findZeros(arr))
}


//копирование среза
func main() {
    var N int
    fmt.Scan(&N)

    var input []int
    for i := 0; i < N; i++ {
        var num int
        fmt.Scan(&num)
        input = append(input, num)
    }

    var start, end int
    fmt.Scanf("%d %d", &start, &end)
    output:=make([]int, end-start)
    copy(output, input[start:end])
    fmt.Println(output)
}

Сортировка срезов:
sort.Ints(numbers)
sort.Float64s(floats)
sort.Strings(strings)



import (
    "fmt"
    "sort"
)    

func isIncSeq(arr []int) {
	cur := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] != cur+1 {
			fmt.Println("false")
			return
		}
		cur = arr[i]
	}
	fmt.Println("true")
}

func main() {
    var N int
    fmt.Scan(&N)

    var input []int
    for i := 0; i < N; i++ {
        var num int
        fmt.Scan(&num)
        input = append(input, num)
    }
    sort.Ints(input)
    isIncSeq(input)
}


func splitByPairs(str string) []string {
	res := []string{}
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str)-1; i += 2 {
		cur := ""
		cur = string(str[i]) + string(str[i+1])
		res = append(res, cur)
	}
	fmt.Println(res)
	return res
}

//304

//считать последовательность единиц
func findMaxConsecutiveOnes(nums []int) int {
	var cnt, maxCnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			maxCnt = cnt
			cnt = 0
			continue
		}
		cnt++
	}
	if cnt > maxCnt {
		maxCnt = cnt
	}
	return maxCnt
}


func notInMap(m map[string]int, _keys []string) {
	for _, v := range _keys {
		if _, ok := m[v]; ok == false {
			fmt.Println(v)
		}
	}
}
func main() {
    var N int
    fmt.Scan(&N)
    m := map[string]int{}
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        m[k]=v
    }   
    var input []string
    for i := 0; i < 5; i++ {
        var k string
        fmt.Scan(&k)
        input = append(input, k)
    }
    notInMap(m, input)

}


func mapCompare(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func main() {
    
    var N int
    fmt.Scan(&N)
    a := map[string]int{}
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        a[k]=v
    }       
    b := map[string]int{}
    fmt.Scan(&N)
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        b[k]=v
    }    
    fmt.Println(mapCompare(a, b))
}

func main() {    
    var N int
    fmt.Scan(&N)
    a := map[string]int{}
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        a[k]=v
    }       
    //fmt.Println(a)
    sr:=[]string{}
    for k,_ := range a {
        sr=append(sr, k)
    }
    sort.Strings(sr)
    for _, v := range sr {
        fmt.Println(v)
    }
}

// 308


func diffMaps(m1, m2 map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			res[k] = v
		}
	}
	return res
}

func main() {
    var N int
    fmt.Scan(&N)
    a := map[string]int{}
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        a[k]=v
    }       
    b := map[string]int{}
    fmt.Scan(&N)
    for i := 0; i < N; i++ {
        var k string
        var v int
        fmt.Scanf("%s %d",&k,&v)
        b[k]=v
    }
    res := diffMaps(a,b)
    //fmt.Println(res)
    for k, v := range res {
        fmt.Println(k,v)
    }
}

func isAnagram(s string, t string) bool {
	sr1 := map[string]int{}
	sr2 := map[string]int{}

	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		sr1[string(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		sr2[string(t[i])]++
	}

	for k, v := range sr1 {
		if sr2[k] != v {
			return false
		}
	}	
	return true
}


func convertDateTime(str string) {
	const layout = "2006-01-02 15:04:05"
	// Парсинг строки в объект time.Time
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Ошибка парсинга времени:", err)
		return
	}
	// Вывод в требуемом формате 00 hours 01 minutes 02 seconds
	fmt.Printf("%02d hours %02d minutes %02d seconds\n", t.Hour(), t.Minute(), t.Second())
}

func main() {
    var x string
    fmt.Scan(&x)
    runes := []rune(x)
	//first := runes[0]
	//last := runes[len(runes)-1]
    fmt.Printf("%c %c\n",runes[0], runes[len(runes)-1])
}

//без сортировки через карту
//Числа располагаются в произвольном порядке. Требуется найти пропущенное число.
func missingNumber(arr []int) int {
	max, min := arr[0], arr[0]
	mp := make(map[int]int, len(arr))
	mp[arr[0]] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
		mp[arr[i]] = 1
	}
	for i := min; i <= max; i++ {
		if _, ok := mp[i]; ok == false {
			return i
		}

	}
	return max + 1
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != i+min {
			res = i + min
			break
		}
	}
	if res == 0 {
		res = nums[len(nums)-1] + 1
	}
	return res
}

func isPowerOfTwo(n int) bool {
	var a, b int
	a = n
	for {
		cur := a
		a = cur / 2
		b = cur % 2
		if b != 0 {
			return false
		}
		if a <= 2 {
			break
		}
	}
	return true
}
//318 - 28.07.2024


// допишите определение пустого интерфейса
type Empty interface {
}

func main() {
    // Шаг 1 - инициализируйте переменную i типа Empty
    var i Empty
    // Шаг 2 - сравните переменную с nil
    fmt.Println(i == nil)
    // Шаг 3 - присвойте переменной i значение в виде пустой структуры
    i = struct{}{}
    // Шаг 4 - сравните переменную с пустым интерфейсом interface{}
    var emptyInterface interface{}
    fmt.Println(i == emptyInterface)
}


// Определение интерфейса Caller
type Caller interface {
	Nested() Caller
}
// Структура, которая реализует интерфейс Caller
type MyCaller struct {
	name string
}
// Реализация метода Nested для MyCaller
func (mc *MyCaller) Nested() Caller {
	// Для демонстрации возвращаем новый экземпляр MyCaller
	return &MyCaller{name: mc.name + " (nested)"}
}



func main() {
	// Создаем экземпляр MyCaller
	caller := &MyCaller{name: "Initial"}
	// Используем метод Nested
	nestedCaller := caller.Nested()
	// Приводим nestedCaller к типу *MyCaller для доступа к полю name
	if nc, ok := nestedCaller.(*MyCaller); ok {
		fmt.Println(nc.name) // Output: Initial (nested)
	}
}

//----------------------------- 
type Doer interface {
    Do() error
}
	var d Doer //Переменная d имеет тип Doer, но не содержит конкретного значения, которое реализует интерфейс Doer.
	d.Do() // паника - invalid memory address or nil pointer dereference 
    //Попытка вызвать метод Do на nil интерфейсе приводит к панике выполнения (runtime panic), так как d не указывает ни на какой конкретный объект.

//----------------------------------------------------
// Определение интерфейса Doer
type Doer interface {
	Do() error
}

// Реализация интерфейса Doer в структуре MyDoer
type MyDoer struct{}

// Реализация метода Do для MyDoer
func (md MyDoer) Do() error {
	fmt.Println("Doing something...")
	return nil
}

func main() {
	// Инициализация переменной d конкретным значением, реализующим интерфейс Doer
	var d Doer = MyDoer{}

	// Вызов метода Do
	err := d.Do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}


func main() {
	var empty interface{}
	empty = 1
	_ = int64(empty.(int))
	fmt.Println(empty) // 1 ошибок нет
}

func main() {
	ch1 := make(chan int)
	ch2 := ch1
	close(ch1)
	if _, ok := <-ch1; !ok {
		fmt.Print("1")
	}
	if _, ok := <-ch2; ok {
		fmt.Print("2")
	}
} //вывод только 1. так как ch2 закроется как и ch1


func main() {
     ch := make(chan int)
     ch <- 1
     for i := range ch {
         fmt.Print(i)
     }
}//паника - дедлоск


func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Print(<-ch)
	fmt.Print(<-ch)
	go func() {
		fmt.Print(<-ch)
	}()
}// паника и зависнет

func go_1() {
	ch := make(chan int, 2)

	go func() {
		fmt.Print(<-ch)
	}()

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Print(<-ch)
	fmt.Print(<-ch)

} // OK - 123


func main() {
	ch := make(chan int)
	stopCh := make(chan struct{})

	go func() {
		ch <- 1
		fmt.Print(<-ch)
		close(stopCh)
	}()
	go func() {
		fmt.Print(<-ch)
		ch <- 2
	}()
	<-stopCh
} // ok - 12

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	for i := 1; i <= 2; i++ {
		ch1 <- i
		ch2 <- i * i

		select {
		case v := <-ch1:
			fmt.Print(v)
		case v := <-ch2:
			fmt.Print(v)
		}
	}
}// ok любая пара комбинаций из цифр 1,2,4

ch := make(chan int)
var ch = make(chan int, 0)
ch := (chan int)(nil) объявляет переменную ch типа chan int и инициализирует ее значением nil
var ch chan any

var wg sync.WaitGroup
wg.Add(2)
wg.Done()
wg.Wait()


import (
    "fmt"
    "sync"
)

func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    var res [2]int
    var wg sync.WaitGroup
    wg.Add(2)    
    
    go func() {
        res[0] = n1
        wg.Done()
    }()
    go func() {
        res[1] = n2
        wg.Done()
    }()

    wg.Wait()
    if res[0] > res[1] {
        fmt.Println(res[1])
    } else {
        fmt.Println(res[0])
    }
}


func main() {
    var n int
    fmt.Scan(&n)

    var wg sync.WaitGroup
    wg.Add(1)

    sum := 0
    go func() {
        for i := 1; i <= n; i++ {
            sum += i
        }
        wg.Done()
    }()

    wg.Wait()
    fmt.Println(sum)
}


type Record struct {
	text string
}

func main() {
	m := make(map[struct{}]Record)
	m[struct{}{}] = Record{"test"}
	fmt.Println(m[struct{}{}].text) // test
}

347


func main() {
	var a = 1

	func(a int) {
		a += 1 // локальная переменная
	}(a)

	func() int {
		a += 1
		return a
	}()

	fmt.Println(a) //2
}

	e := fmt.Errorf("1 : %w", errors.New("2"))
	fmt.Println(e.Error()) // 1 : 2
    Спецификатор %w позволяет обернуть одну ошибку в другую, сохраняя при этом возможность извлечения оригинальной ошибки с помощью errors.Unwrap.


	e := fmt.Errorf("1 : %w", errors.New("2"))
	e = fmt.Errorf("3 : %w", e)
	e = errors.Unwrap(e)  // откручивает 3 ошибку
	//Разворачивание: errors.Unwrap(e) извлекает внутреннюю ошибку "1 : 2".
	fmt.Println(e.Error()) // 1 : 2

import (
    "encoding/json"
    "fmt"
)
type Person struct {
    Name string
    Age  int
}
func main() {
    person := Person{
        Name: "Bill",
        Age:  25,
    }
    pers, _ := json.Marshal(person)
    fmt.Println(string(pers))
}

{"x": 2}

func isJson(str string) bool {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(str), &js)
	return err == nil
}

// hello -> Hello
func capitalize(s string) string {
    return string(s[0]-32)+s[1:]
}
363 - 28.07.2024

Здравствуйте. Попробуйте применить bufio.Scanner


import (
    "encoding/json"
    "fmt"
    "bufio"
    "os"
)
func isJson(str string) bool {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(str), &js)
	return err == nil
}
func main() {
    var count int
    fmt.Scan(&count)
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        str := scanner.Text()
        //fmt.Print(str, ":") // Println will add back the final '\n'
        fmt.Println(isJson(str))
	}
}

package main

import (
	"fmt"
	"unicode"
    "bufio"
    "os"    
)

func hasNonAlphanumeric(str string) bool {
	res := false
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			res = true
			break
		}
	}
	return res
}

func main() {
    var N int
    fmt.Scan(&N)
    scanner := bufio.NewScanner(os.Stdin)
    i:=0
    for scanner.Scan() {
        str := scanner.Text()        
        //fmt.Println(str,hasNonAlphanumeric(str))
        if hasNonAlphanumeric(str) {
            fmt.Println(i)
        }
        i++
    }       
}

import (
	"fmt"
    "bufio"
    "os"    
)


func letterCounter(str string) {
	a, b := 0, 0
	for i := 0; i < len(str); i++ {
		cur := str[i]
		if cur < 65 {
			continue
		}
		if cur >= 65 && cur < 97 {
			a++
		}
		if cur >= 97 && cur < 97+32 {
			b++
		}
	}
	fmt.Println(a, b)
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text()
    letterCounter(str)
}

import (
	"fmt"
    "bufio"
    "os"    
    "strings"
)


import (
	"fmt"
    "strings"
)

func main() {
    cnt :=0
    fmt.Scan(&cnt)
    var sb strings.Builder
    for i:=0; i<cnt; i++ {
        str:=""
        fmt.Scan(&str)
        sb.WriteString(str)        
    }
    fmt.Println(sb.String()) //объединить символы в строку    
}


import (
	"fmt"
    "bufio"
    "os"    
    "strings"
    "regexp"
)
// Функция для фильтрации строки
func filterString(input string) string {
	// Регулярное выражение для удаления недопустимых символов
	re := regexp.MustCompile(`[^a-zA-Zа-яА-Я ,]+`)
	// Удаляем все недопустимые символы
	filtered := re.ReplaceAllString(input, "")
	// Удаляем лишние пробелы
	filtered = strings.TrimSpace(filtered) //для удаления начальных и конечных пробелов.
	filtered = strings.Join(strings.Fields(filtered), " ") //собираем строку из елементов через разделитель ' '
	fmt.Println(filtered)
	return filtered
}

371 - 29.07.2024

// Функция для нахождения индекса слова с заданным префиксом
func findWordIndexWithPrefix(text, prefix string) int {
	words := strings.Fields(text) // Разбиваем текст на слова

	// Проходим по словам и ищем слово с заданным префиксом
	for i, word := range words {
		if strings.HasPrefix(word, prefix) {
			return i
		}
	}
	return -1
}

import (
    "fmt"
    "strings"
    "bufio"
    "os"        
)

func main() {
    var searchWord string
    fmt.Scan(&searchWord)
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    fmt.Println(findWordIndexWithPrefix(text, searchWord))    
}


import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func truncateSentence(str string, k int) string {
	words := strings.Fields(str) // Разбиваем текст на слова
	var sb strings.Builder
	for i := len(words) - k; i < len(words); i++ {
		sb.WriteString(words[i])
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()    
    scanner.Scan()
    nn := scanner.Text()    
    var n int
    n , err = strconv.Atoi(nn)   
	if err != nil {
		fmt.Println("Ошибка: введено не целое число.")
		return
	}
    fmt.Println(truncateSentence(text, n))
}


func compareStr(a, b string) bool {
	arr_a := map[string]int{}
	arr_b := map[string]int{}
	for _, v := range a {
		arr_a[string(v)]++
	}
	for _, v := range b {
		arr_b[string(v)]++
	}
	if len(arr_a) != len(arr_b) {
		return false
	}
	for k, _ := range arr_a {
		if _, ok := arr_b[k]; !ok {
			return false
		}
	}
	return true
}



// Функция для приведения каждой второй буквы в каждом слове к верхнему регистру
func convertEverySecondLetterToUpper(text string) string {
	words := strings.Fields(text)
	var result []string

	for _, word := range words {
		tmp := []rune{}
		for i, v := range word {
			if i%2 == 1 {
				tmp = append(tmp, unicode.ToUpper((v)))
			} else {
				tmp = append(tmp, (v))
			}
		}
		result = append(result, string(tmp))
	}
	fmt.Println(strings.Join(result, " "))
	return strings.Join(result, " ")
}


func main() {
	//var ch chan int // канал не инициализирован, заблокирован - res = 2
	ch := make(chan int) // вот так верно res = 1 если есть time.Sleep
	go func() {
		ch <- 1
	}()

	// time.Sleep(time.Second) // иначе не успевает попасть 1
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println(2) // канал не инициализирован, заблокирован
	}
} // result 2 



func main() {
	ch1 := make(chan int)

	ch2 := ch1
	go func() {
		ch1 <- 1
	}()

	select {
	case v := <-ch2:
		fmt.Println(v)
	case <-time.After(time.Second): //ждем 1 сек, если от канала ch2 не будет получено, то на экран будет выведено "2".
		fmt.Println("2")
	}
} // result = 1 все норм

func main() {
    var str string
    fmt.Scan(&str)	

    mp := make(map[string]string)    

	err = json.Unmarshal([]byte(str), &mp) //строку json в карту
	if err != nil {
    	fmt.Println("Ошибка десериализации:", err)
	    return
	}
    //fmt.Println(m)
    for k, v :=range mp {
        fmt.Println(k,v)
    }    
}


func heavyOperation(k string) string {
    return db[k]
}

var m = map[string]string{}

func set(k, v string) {
    db[k] = v
}

var db = map[string]string{}

func setDB() {
    for i := 0; i < 8; i++ {
        var k, v string
        fmt.Scan(&k, &v)
        db[k] = v
    }
}

func get(k string) string {
   // допишите код данной функции
    if v, ok := db[k]; ok {
        return v
    } else {
        return heavyOperation(k)
    }
}

func main() {
    setDB()

    for i := 0; i < 5; i++ {
        var s string
        fmt.Scan(&s)
        fmt.Println(get(s))
        //set(s, "x")
    }
}

type Record struct {
    id int
    name string
}

func main() {
    var cnt,idx int
    fmt.Scanf("%d %d", &cnt, &idx)
    
    //fmt.Println(cnt,b)
    mp := map[int]string{}
    var b int
    for i:=0; i<cnt; i++ {
        str:=""
        fmt.Scanf("%d %s", &b, &str)
        mp[b] = str
    }
    fmt.Println(mp[idx])    
}


func main() {
   var cnt,idx int
    fmt.Scanf("%d %d", &cnt, &idx)   
    mp := map[int]Record{}
    r:=Record{}
    for i:=0; i<cnt; i++ {
        fmt.Scanf("%d %s", &r.id, &r.name)
        mp[r.id] = r
    }
    fmt.Println(mp[idx].name)
}


    var input = make([]int, 8)
    var scanInput = make([]any, 8)
        
    for i := range input {
        scanInput[i] = &input[i] // Затем заполняем срез указателями на элементы среза input. 
								// Это необходимо, потому что fmt.Scan требует, чтобы аргументы были указателями.
    }
    fmt.Scan(scanInput...)
    fmt.Print(input)

400 31-07-2024

func moveZeroes(nums []int) {
	res := make([]int, 0)
	for _, v := range nums {
		if v > 0 {
			res = append(res, v)
		}
	}
	for _, v := range nums {
		if v == 0 {
			res = append(res, v)
		}
	}
	fmt.Println(res)
}


func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isPartArr(arr [5]int, prt []int) {
	a := prt[0]
	idx := 0
	for i, v := range arr {
		if v == a {
			idx = i
			break
		}
	}
	end := len(prt) + idx
	arr2 := arr[idx:end]
	fmt.Println(compareSlices(arr2, prt))
}

func main() {
    var arr [5] int
    for i:=0; i<5;i++ {
        fmt.Scan(&arr[i])
    }
    var cnt int
    fmt.Scan(&cnt)
    sr:=[] int{}
    for i:=0; i<cnt;i++ {
        var cur int
        fmt.Scan(&cur)
        sr=append(sr,cur)
    }
    isPartArr(arr,sr)
}


func numberOfOnes(str string) int {
    cnt := 0
    for _, v := range str {
        if string(v) == "1" {
            cnt++
        }
    }
    return cnt
}

//
func getSignal(house [][]int) {
	for i, flats := range house {
		for j, sig := range flats {
			if sig == 1 {
                fmt.Printf("Этаж: %d, Квартира: %d\n", i+1, (i)*len(flats)+j+1)
			}
		}
	}
}


func getOperation(a string) {
	switch a {
	case "+":
		fmt.Println("Сумма")
	case "-":
		fmt.Println("Разность")
	case "*":
		fmt.Println("Умножение")
	case "/":
		fmt.Println("Деление")
	default:
		fmt.Println("Неизвестно")
	}
}

func calcOp(a, b int, op string) {
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("Неизвестно")
	}
}

func main() {
    var a,b int
    var op string
    fmt.Scanf("%d %d %s",&a,&b,&op)
    calcOp(a,b,op)
}

func getTempDesc(a int) {
	if a == -100 {
		fmt.Println("Слишком холодно")
	} else if a >= -80 && a <= -50 {
		fmt.Println("Холодно")
	} else if a >= -20 && a <= 20 {
		fmt.Println("Комфорт")
	} else if a >= 21 && a <= 80 {
		fmt.Println("Жарко")
	} else if a == 100 {
		fmt.Println("Слишком жарко")
	}
}

func getType() {
	var input string
	fmt.Scanln(&input)

	// Попытка преобразования в целое число
	if intValue, err := strconv.Atoi(input); err == nil {
		//fmt.Printf("Это целое число: %d\n", intValue)
		fmt.Println("int")
		return
	}

	// Попытка преобразования в число с плавающей точкой
	if floatValue, err := strconv.ParseFloat(input, 32); err == nil {
		//fmt.Printf("Это число с плавающей точкой: %f\n", floatValue)
		fmt.Println("float32")
		return
	}

	// Попытка преобразования в булево значение
	if boolValue, err := strconv.ParseBool(input); err == nil {
		//fmt.Printf("Это булево значение: %t\n", boolValue)
		fmt.Println("bool")
		return
	}

	// Если ни одно из преобразований не удалось, это строка
	//fmt.Printf("Это строка: %s\n", input)
	fmt.Println("string")


func decomposeNumber(val int) {
	for {
		a := val / 10
		b := val % 10
		fmt.Println(b)
		if a == 0 {
			break
		}
		val = a
	}
}

func splitThousand(val int) string {
	res := []int{}
	for {
		a := val / 1000
		b := val % 1000
		res = append(res, b)
		if a == 0 {
			break
		}
		val = a
	}
	str := ""
	for i := len(res) - 1; i >= 0; i-- {
		tmp := fmt.Sprint(res[i])
		str += tmp
		if i != 0 {
			str += "_"
		}
	}
    return str  // 1_234
}

419 - вечер 31-07-2024

import (
    "fmt"
    "math"
)    

func main() {
    var a1,a2,b1,b2 int
    fmt.Scan(&a1,&a2)
    fmt.Scan(&b1,&b2)
    length:= math.Sqrt(float64((a1-b1)*(a1-b1) + (a2-b2)*(a2-b2)))
    fmt.Printf("%.2f", length)
}

type House struct {
    Number int
}

func main() {
    street := []House{
        {1},
        {2},
        {3},
        {4},
        {5},
    }
    var old_, new_ int
    for i := 0; i < 2; i++ {
        fmt.Scan(&old_, &new_)
        street[old_ -1].Number = new_
    }

    for _, h := range street {
        fmt.Println(h)
    }
}

type House struct {
    Number int
    BuildDate string
}
func main() {
    street := []House{}
    for i:=0;i<5;i++ {
        var h House
        fmt.Scan(&h.Number, &h.BuildDate)
        street=append(street, h)
    }
    for _,v := range street {
        //fmt.Println(v.Number, v.BuildDate)
        fmt.Println(v)
    }
}

func fib(a, b, cnt int) bool {
    fmt.Println(a)
    return (cnt<0) || fib(b, a+b, cnt-1)
}

//-----------------------------------
type Product struct {
    Name string
    Price int
}

type ByPrice  []Product

func (a ByPrice ) Len() int { 
    return len(a) 
}
func (a ByPrice ) Swap(i, j int) { 
    a[i], a[j] = a[j], a[i] 
}
func (a ByPrice ) Less(i, j int) bool { 
    return a[i].Price < a[j].Price
}

func main() {
    var pr Product
    products := []Product{}
    for i:=0;i<3;i++{
        fmt.Scan(&pr.Name, &pr.Price)
        products=append(products, pr)
    }
    sort.Sort(ByPrice(products))
	//sort.Sort(sort.Reverse(ByName(products)))
    fmt.Println(products)
}


type SortByPrice  []Product

func (a SortByPrice ) Len() int { 
    return len(a) 
}
func (a SortByPrice ) Swap(i, j int) { 
    a[i], a[j] = a[j], a[i] 
}
func (a SortByPrice ) Less(i, j int) bool { 
    return a[i].Price < a[j].Price
}

452 вечер 01-08-2024


import (
    "fmt"
    "sort"
)    

const (
    Junior = "junior"
    Middle = "middle"
    Senior = "senior"
)

func getAchievment(score int) []string{
    res:=[]string{}
    if score >= 10 {
        res=append(res, Junior)
    }
    if score >= 50 {
        res=append(res, Middle)
    }
    if score >= 100 {
        res=append(res, Senior)
    }
    return res
}

func main() {
    balls:=map[string]int{}
    var cnt int
    fmt.Scan(&cnt)
    for i:=0;i<cnt;i++{
        var name string
        var score int
        fmt.Scan(&name, &score)
        balls[name]+=score
    }
    //fmt.Println(balls)    
    //ach :=[]string{}
    //fmt.Println(balls)
    //names:=[]string{}
    names:=make([]string, 0, len(balls))
    for key := range balls {
		names = append(names, key)
	}
    sort.Strings(names)
    
    for _,k:=range names {
        fmt.Println(k,getAchievment(balls[k]))
    }    
}


import (
    "time"
)    

type Task struct {
    ID int64
    Title string
    Description string
    CreatedAt time.Time
    UpdatedAt time.Time
    DueDate time.Time
    IsCompleted bool
    IsOvertime bool
}


func getDuplicate(arr []int) {
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if cur == arr[j] {
				fmt.Println(cur)
				return
			}
		}
	}
}


func exitsInArr(a int, arr []int) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func filterDuplicate(arr []int) {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if !exitsInArr(val, res) {
			res = append(res, arr[i])
		}
	}
	fmt.Println(res)
}



func exitsInArr(a int, arr []int) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func filterDuplicate(arr []int) {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if !exitsInArr(val, res) {
			res = append(res, arr[i])
		}
	}
	fmt.Println(res)
}


func main() {
    var input = make([]int, 8)
    var scanInput = make([]any, 8)
       
    for i := range input {
        scanInput[i] = &input[i]
    }
    fmt.Scan(scanInput...)
    //fmt.Print(input)
    filterDuplicate(input)
}


type Student struct {
    Name   string
    Scores []int
    Avg float64
}

func main() {
	var cnt_stud, cnt_score int
	fmt.Scan(&cnt_stud, &cnt_score)
    students := []Student{}
    
	for i := 0; i < cnt_stud; i++ {
        var student Student
        fmt.Scan(&student.Name)
        
        var scanInput = make([]any, cnt_score)
        var input = make([]int, cnt_score)

        for j := range input {
            scanInput[j] = &input[j]
        }
        fmt.Scan(scanInput...)
        student.Scores = input
        
        students=append(students, student)        
    }
        
    for i,stud :=range students {
        sum := 0
        for _, score := range stud.Scores {
            sum+=score
            
        }
        //fmt.Println(float64(sum))
        //fmt.Println(float64(sum) /float64(cnt_score))
        students[i].Avg = float64(sum) / float64(cnt_score)        
    }
    
    min := students[0].Avg
    idx_min := 0
    for i:=1;i<cnt_stud;i++ {
        if students[i].Avg < min {
            idx_min = i
        }
    }
    fmt.Println(students[idx_min].Name)               
    //fmt.Println(students)    
}



func main() {
    var need_avg int
    fmt.Scan(&need_avg)
    names := []string{}
    balls := []int{}
    for i:=0;i<4; i++ {
        name:=""
        fmt.Scan(&name)
        names=append(names, name)
        sum:=0
        for j:=0;j<5; j++ {
            var ball int
            fmt.Scan(&ball)
            sum+=ball
        }
        balls=append(balls, sum/5)        
    }
    //fmt.Println(names)
    //fmt.Println(balls)
    result :=[]string{}
    for i:=0;i<4; i++ {
        if balls[i] == need_avg {
            result=append(result, names[i])
        }
    }    
    fmt.Println(result)
}


package main

import (
    "fmt"
)

type OrderedMap struct {
    keys []string
    data map[string]interface{}
}

func NewOrderedMap() *OrderedMap {
    return &OrderedMap{
        keys: []string{},
        data: make(map[string]interface{}),
    }
}

func (om *OrderedMap) Set(key string, value interface{}) {
    if _, exists := om.data[key]; !exists {
        om.keys = append(om.keys, key)
    }
    om.data[key] = value
}

func (om *OrderedMap) Get(key string) (interface{}, bool) {
    value, exists := om.data[key]
    return value, exists
}

func (om *OrderedMap) Keys() []string {
    return om.keys
}

func main() {
    om := NewOrderedMap()
    om.Set("first", 1)
    om.Set("second", 2)
    om.Set("third", 3)
    
    for _, key := range om.Keys() {
        value, _ := om.Get(key)
        fmt.Println(key, value)
    }
}



import "fmt"

var keys1 []string
var keys2 []string

//соединить карты с сохранением порядка ключей
func joinMap(a,b map[string]int) map[string]int{
    for _,k := range keys2 {
        a[k]=b[k]
        keys1=append(keys1, k)
    }
    return a
}

func printMap(m map[string]int){
    for _,k := range keys1 {
        fmt.Println(k,m[k])           
    }
}

func main() {
    cnt:=0
    fmt.Scan(&cnt)
    m1:=map[string]int{}
    m2:=map[string]int{}
    
    for i:=0;i<cnt;i++ {
        k:="";v:=0
        fmt.Scan(&k,&v)
        m1[k]=v        
        keys1=append(keys1, k)
    }
    
    fmt.Scan(&cnt)
    for i:=0;i<cnt;i++ {
        k:="";v:=0
        fmt.Scan(&k,&v)
        m2[k]=v        
        keys2=append(keys2, k)
    }
    
    printMap(joinMap(m1,m2))

}

func main() {
	var wg sync.WaitGroup
	i := int64(0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	wg.Wait()
	fmt.Println(i) // 2
}


func main() {
	a := 0
	var o1 sync.Once
	o1.Do(func() {
		a = 1
	})
	o1.Do(func() {
		a = 2
	})
	fmt.Println(a) // 1
}


func main() {
	a := 0
	var o sync.Once
	oc := o
	o.Do(func() {
		a = 1
	})
	oc.Do(func() {
		a = 2
	})
	fmt.Println(a) //2
}


import (
	"fmt"
	"sync"
)
func main() {
	var mut sync.Mutex
	a := 0
	go func() {
		mut.Lock()
		a += 1
		mut.Unlock()
	}()
	go func() {
		mut.Lock()
		a += 1
		mut.Unlock()
	}()
	mut.Lock()
	a = -1
	mut.Unlock()
	mut.Lock()
	fmt.Println(a)
}


var n1, n2 int
fmt.Scan(&n1, &n2)
var mut sync.Mutex
var wg sync.WaitGroup
	wg.Add(2)

sum := 0

go func() {
    mut.Lock()
    sum += n1    
    mut.Unlock()
    defer wg.Done()
}()

go func() {
    mut.Lock()
    sum += n2
    mut.Unlock()
    defer wg.Done()
}()

wg.Wait()
fmt.Println(sum)



func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover in mayPanic:", r)
        }
    }()
    panic("Проблема!")
}

func main() {
    mayPanic()
    fmt.Println("Продолжаем выполнение после вызова mayPanic")
}

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Восстановились после паники:", r)
        }
    }()
    
    fmt.Println("Начало программы")
    panic("Что-то пошло не так!")  // Вызов паники
    fmt.Println("Этот код никогда не будет выполнен")
}

Функция recover работает только в отложенных функциях. Это значит, что её нужно вызывать внутри функции, 
которая была отложена с помощью defer. 
Если recover вызывается вне отложенной функции, она всегда вернёт nil и не сможет перехватить панику.


package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Горутина для отправки данных в ch1
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "сообщение из ch1"
    }()

    // Горутина для отправки данных в ch2
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "сообщение из ch2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Получено:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Получено:", msg2)
        }
    }
}



func go_rutin(cnt int) {

	sumCh := make(chan int)
	mulCh := make(chan int)
	powCh := make(chan int)

	// Горутина для отправки данных в ch1
	go func() {
		if cnt != 1 {
			time.Sleep(1 * time.Second)
		}
		sumCh <- 1
	}()

	go func() {
		if cnt != 2 {
			time.Sleep(1 * time.Second)
		}
		mulCh <- 2
	}()

	go func() {
		if cnt != 3 {
			time.Sleep(1 * time.Second)
		}
		powCh <- 3
	}()

	select {
	case msg1 := <-sumCh:
		fmt.Println("Получено:", msg1)
	case msg2 := <-mulCh:
		fmt.Println("Получено:", msg2)
	case msg3 := <-powCh:
		fmt.Println("Получено:", msg3)
	}
}




// Горутина, которая записывает значение в канал после задержки
func writeToChannel(value int, delay time.Duration, ch chan int) {
	time.Sleep(delay) // Засыпаем на указанный период времени
	ch <- value       // Записываем значение в канал
}

func main() {
	// Входные данные
	value := 42             // Значение, которое нужно записать в канал
	delay := 150 * time.Millisecond // Задержка в миллисекундах

	ch := make(chan int) // Создаем канал

	// Запускаем горутину
	go writeToChannel(value, delay, ch)

	select {
	case v := <-ch:
		// Чтение из канала успешно
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		// Если чтение из канала не произошло за 100 миллисекунд
		fmt.Println(0)
	}
}


// Горутина, которая записывает значение в канал после задержки
func writeToChannel(value int, delay time.Duration, ch chan int) {
	time.Sleep(delay) // Засыпаем на указанный период времени
	ch <- value       // Записываем значение в канал
}

func main() {
    var value int
    var _time time.Duration
    
    fmt.Scan(&value, &_time)
    delay := _time * time.Millisecond // Задержка в миллисекундах
    
    ch := make(chan int) // Создаем канал
    go writeToChannel(value, delay, ch)
    
	select {
	case v := <-ch:
		// Чтение из канала успешно
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		// Если чтение из канала не произошло за 100 миллисекунд
		fmt.Println(0)
	}    
}

func do() {    
	// Создаем буферизованный ридер для стандартного ввода
	bufReader := bufio.NewReader(os.Stdin)
	// Используем io.TeeReader для копирования данных из bufReader в os.Stdout
	reader := io.TeeReader(bufReader, os.Stdout)
	// Копируем данные из reader в io.Discard, чтобы они были прочитаны и выведены
	io.Copy(io.Discard, reader) 
}



package main

import (
	"errors"
	"fmt"
)

// Node представляет узел двусвязного списка
type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

// DoublyLinkedList представляет двусвязный список
type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewDoublyLinkedList создает новый пустой двусвязный список
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Add добавляет новый элемент в конец списка
func (dll *DoublyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

// Delete удаляет первый найденный элемент, удовлетворяющий условию
func (dll *DoublyLinkedList[T]) Delete(cond func(item T) bool) bool {
	current := dll.head
	for current != nil {
		if cond(current.value) {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dll.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dll.tail = current.prev
			}

			dll.size--
			return true
		}
		current = current.next
	}
	return false
}

// Reverse разворачивает двусвязный список
func (dll *DoublyLinkedList[T]) Reverse() {
	current := dll.head
	var prev *Node[T] = nil
	dll.tail = dll.head

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}

	dll.head = prev
}

// Swap меняет местами элементы с индексами i и j
func (dll *DoublyLinkedList[T]) Swap(i, j int) error {
	if i >= dll.size || j >= dll.size {
		return errors.New("индекс вне границ списка")
	}

	if i == j {
		return nil // Ничего менять не нужно
	}

	var nodeI, nodeJ *Node[T]
	current := dll.head

	for idx := 0; current != nil && (nodeI == nil || nodeJ == nil); idx++ {
		if idx == i {
			nodeI = current
		}
		if idx == j {
			nodeJ = current
		}
		current = current.next
	}

	if nodeI != nil && nodeJ != nil {
		nodeI.value, nodeJ.value = nodeJ.value, nodeI.value
		return nil
	}

	return errors.New("не удалось найти указанные индексы")
}

// Print печатает элементы списка
func (dll *DoublyLinkedList[T]) Print() {
	current := dll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// Интерфейс с требуемыми методами
type MyInterface[T any] interface {
	Delete(func(item T) bool) bool
	Reverse()
	Swap(i, j int) error
}

func main() {
	// Пример использования двусвязного списка
	list := NewDoublyLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Println("Исходный список:")
	list.Print()

	// Удаляем элемент
	list.Delete(func(item int) bool {
		return item == 3
	})
	fmt.Println("После удаления 3:")
	list.Print()

	// Разворачиваем список
	list.Reverse()
	fmt.Println("После разворота:")
	list.Print()

	// Меняем местами элементы
	err := list.Swap(1, 3)
	if err != nil {
		fmt.Println("Ошибка при обмене:", err)
	} else {
		fmt.Println("После Swap(1, 3):")
		list.Print()
	}
}
