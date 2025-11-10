package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Pow(s, p int) int {
	if p == 0 {
		return 1
	}
	res := s
	for i := 2; i <= p; i++ {
		res *= s
	}
	return res
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func drawCross(a, b int) {
	for i := 0; i < a; i++ { //строка
		for j := 0; j < b; j++ {
			if i == a/2 {
				fmt.Print("*")
			} else if j == b/2 {
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
	if i > 0 {
		res = string(a[i-1])
	} else {
		res = string(a[len(a)+i])
	}
	return res
}

func hasPrefix(s string, suf string) bool {
	for i := 0; i < len(suf); i++ {
		if s[i] != suf[i] {
			return false
		}
	}
	return true
}

func sum() {
	var a, res int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		res += a
	}
	fmt.Println(res)
}
func down() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	for i := b; i >= a; i-- {
		fmt.Println(i)
	}
}

func div2() {
	var x, N, res int
	fmt.Scan(&x, &N)
	res = x
	for i := 1; i <= N; i++ {
		res = res / 2
	}
	fmt.Println(res)
}

func isSum() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b == c {
		fmt.Printf("%d + %d = %d", a, b, c)
	} else if a+c == b {
		fmt.Printf("%d + %d = %d", a, c, b)
	} else if b+c == a {
		fmt.Printf("%d + %d = %d", b, c, a)
	} else {
		fmt.Println("нет")
	}
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

// часть строки
func partString(str string, a, b int) {
	//	var a, b int
	//	var str string
	//	fmt.Scan(&str)
	//	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(string(str[a:b]))
}

func printN() {
	var b int
	var str string
	fmt.Scan(&b)
	fmt.Scan(&str)
	for i := 0; i < b; i++ {
		fmt.Println(str)
	}
}

func printABC() {
	for i := 65; i <= 90; i++ {
		fmt.Println(string(i), string(i+32))
	}
}

type Money struct {
	Name  string
	Value int
}

type Person struct {
	Name string
	Age  int
}

type Human struct {
	Name string
}

// Определение структуры Alien, встраивающей структуру Human
type Alien struct {
	Human
}

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
	//	var str, name string
	//	fmt.Scan(&str)
	//	fmt.Scan(&name)

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

// Matrix представляет собой тип матрицы
type Matrix struct {
	rows int
	cols int
}

// NewMatrix создает новую матрицу размером m x n и заполняет ее нулями
func NewMatrix(m, n int) *Matrix {
	return &Matrix{
		rows: m,
		cols: n,
	}
}

func matrix(m, n int) *Matrix {
	return &Matrix{
		rows: m,
		cols: n,
	}
}

// String возвращает строковое представление матрицы
/*
func (matrix *Matrix) String2() string {
	var sb strings.Builder
	for _, row := range matrix.data {
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%d ", val))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
*/

func (matrix *Matrix) String() string {
	var sb strings.Builder
	for i := 0; i < matrix.rows; i++ {
		for j := 0; j < matrix.cols; j++ {
			//fmt.Print("0")
			sb.WriteString(fmt.Sprint("0"))
		}
		//fmt.Println("")
		sb.WriteString("\n")
	}
	return sb.String()
}

func showMatrix() {
	m := 3 // Количество строк
	n := 4 // Количество столбцов
	//_matrix := matrix(m, n)
	//fmt.Print(_matrix)
	fmt.Println(matrix(m, n))
}

func printN1() {
	var cnt, cur int
	arr := [10]int{}
	fmt.Scan(&cnt)
	for i := 0; i < 10; i++ {
		fmt.Scan(&cur)
		arr[i] = cur
	}
	for i := 0; i < cnt; i++ {
		fmt.Print(arr[i], " ")
	}
}

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

func findZeros_() {
	var cnt, cur int
	arr := []int{}
	//arrZero := []int{}
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&cur)
		arr = append(arr, cur)
	}
	fmt.Println(findZeros(arr))
}

func filterNegative(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if v < 0 {
			res = append(res, v)
		}
	}
	return res
}

func rotateArr() {
	input := [5]int{1, 2, 3, 4, 5}
	var output [5]int
	for i := 0; i < len(input); i++ {
		output[5-1-i] = input[i]
	}
	fmt.Println(output)
}

func sliceCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isInc(arr []int) {
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
	//	for k, v := range m {
	//		fmt.Println(k, v)
	//	}
	for _, v := range _keys {
		if _, ok := m[v]; ok == false {
			fmt.Println(v)
		}
	}
}
func testMap() {
	m := map[string]int{"a": 1, "b": 2, "d": 3, "e": 4}
	ks := []string{"a", "f"}
	notInMap(m, ks)
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

func testMapCompare() {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := map[string]int{"a": 1, "b": 2, "d": 4}
	fmt.Println(mapCompare(a, b))
}

func diffMaps(m1, m2 map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			res[k] = v
		}
	}
	return res
}

func testDiffMaps() {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(diffMaps(a, b))
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
	// Вывод в требуемом формате
	fmt.Printf("%d hours %0d minutes %0d seconds\n", t.Hour(), t.Minute(), t.Second())
}

func hasNonAlphanumeric(str string) bool {
	res := false
	// Проверка каждого символа в строке
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			res = true
			break
		}
	}
	return res
}

func _sort(arr []int) []int {
	min, idx := arr[0], 0
	res := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			idx = i
		}
	}
	res = append(res, idx)
	return res
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

func missingNumber2(arr []int) int {
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
	//fmt.Println(mp)
	for i := min; i <= max; i++ {
		if _, ok := mp[i]; ok == false {
			return i
		}

	}
	return max + 1
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

func testCaller() {
	// Создаем экземпляр MyCaller
	caller := &MyCaller{name: "Initial"}
	// Используем метод Nested
	nestedCaller := caller.Nested()
	// Приводим nestedCaller к типу *MyCaller для доступа к полю name
	if nc, ok := nestedCaller.(*MyCaller); ok {
		fmt.Println(nc.name) // Output: Initial (nested)
	}
}

type Doer interface {
	Do() error
}

type NotNil interface {
}

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

}

func go_2() {
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
}

func go_3() {
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
}

func q1() {
	var a = 1
	func(a int) {
		a += 1
	}(a)

	func() int {
		a += 1
		return a
	}()

	fmt.Println(a)
}

func err1() {
	e := fmt.Errorf("1 : %w", errors.New("2"))
	e = fmt.Errorf("3 : %w", e)
	e = errors.Unwrap(e)
	fmt.Println(e.Error())
}

func isJson(str string) bool {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(str), &js)
	return err == nil
}

var letters = "abcdefghijklmnopqrstuvwxyz0123456789"

func check(password string) bool {
	for i := 0; i < len(password); i++ {
		if !strings.Contains(letters, string(password[i])) {
			return false
		}
	}
	return true
}

// draw Rectangle
func drawRectangle(a, b int) {
	for i := 0; i < a; i++ { //строка
		for j := 0; j < b; j++ {
			if i == 0 || i == a-1 {
				fmt.Print("*")
			} else if j == 0 || j == b-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func trimPrefix(str, suf string) string {
	x := []rune(str)
	return string(x[len(suf):])
}

type FileInfo struct {
	IsExists bool  // признак того, что файл существует
	Size     int64 // размер файла
	IsDir    bool  // является ли объект директорией
}

// IsExists функция для проверки существования файла и получения информации о нем
func IsExists(filePath string) (*FileInfo, error) {
	file_Info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return &FileInfo{
			IsExists: false,
			Size:     0,
			IsDir:    false,
		}, err
	}

	return &FileInfo{
		IsExists: true,
		Size:     file_Info.Size(),
		IsDir:    file_Info.IsDir(),
	}, nil

}

func IsExists1(filePath string) (bool, error) {

	if !fs.ValidPath(filePath) {
		return false, nil
	}

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false, err
	}

	return false, nil
}

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

func upper(str string) {
}

// Функция для фильтрации строки
func filterString(input string) string {
	// Регулярное выражение для удаления недопустимых символов
	re := regexp.MustCompile(`[^a-zA-Zа-яА-Я ,]+`)
	// Удаляем все недопустимые символы
	filtered := re.ReplaceAllString(input, "")
	// Удаляем лишние пробелы
	filtered = strings.TrimSpace(filtered)
	filtered = strings.Join(strings.Fields(filtered), " ")
	fmt.Println(filtered)
	return ""
}

func getIndex(str, word string) int {
	return strings.Index(str, word)
}

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

func truncateSentence(str string, k int) string {
	words := strings.Fields(str) // Разбиваем текст на слова
	var sb strings.Builder
	for i := len(words) - k; i < len(words); i++ {
		sb.WriteString(words[i])
		sb.WriteString(" ")
	}
	return sb.String()
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

func numberOfSubstrings(word string, patterns []string) int {
	cnt := 0
	for _, v := range patterns {
		if strings.Contains(word, v) {
			cnt++
		}
	}
	return cnt
}

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

func getSignal(house [][]int) {
	for i, flats := range house {
		for j, sig := range flats {
			if sig == 1 {
				fmt.Println("Этаж:", i+1, "Квартира:", (i+1)*len(flats)+j)
			}
		}
	}
}

func drawClock(a int) {
	for i := 0; i < a/2; i++ {
		for j := 0; j < a-i; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	fmt.Print("------")

	for i := a / 2; i >= 0; i-- {
		for j := 0; j < a-i; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func max3(a, b, c int) int {
	return max(max(a, b), max(b, c))
}

func insidePeriod(a1, a2, b1, b2 int) bool {
	if b1 < a1 || a1 < b2 { // начало диапазона внутири b1 - b2
		return true
	}
	if b1 < a2 || a2 < b2 { // конец диапазона внутири b1 - b2
		return true
	}
	if a1 < b1 || b1 < a2 { // начало диапазона внутири b1 - b2
		return true
	}
	if a1 < b2 || b2 < a2 { // конец диапазона внутири b1 - b2
		return true
	}
	return false
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

func getTemp(a int) {
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

func convertTo2(val int) {
	res := ""
	for {
		a := val / 2
		b := val % 2
		if b == 0 {
			res += "0"
		} else {
			res += "1"
		}
		if a == 0 {
			break
		}
		val = a
	}
	res2 := ""
	for i := len(res) - 1; i >= 0; i-- {
		res2 += string(res[i])
	}
	//fmt.Println(res)
	fmt.Println(res2)
}

func decomposeNumber(val int) {
	//res:=""
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

func decomposeThousand(val int) {
	res := []int{}
	for {
		a := val / 1000
		b := val % 1000
		//fmt.Println(b)
		res = append(res, b)
		if a == 0 {
			break
		}
		val = a
	}
	//fmt.Println(res)
	str := ""
	for i := len(res) - 1; i >= 0; i-- {
		//var x int
		//x = res[i]
		tmp := strconv.Itoa(res[i])
		str += tmp
		if i != 0 {
			str += "_"
		}
		//fmt.Println(res[i])
	}
	fmt.Println(str)
}

func getSumMulNumber(val int) (int, int) {
	sum := 0
	mul := 1
	for {
		a := val / 10
		b := val % 10
		sum += b
		mul *= b
		if a == 0 {
			break
		}
		val = a
	}
	return sum, mul
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

type Student struct {
	Name   string
	Scores []int
}

/*
func getMinScore() {
	var cnt_stud, cnt_score int
	fmt.Scan(&cnt_stud, &cnt_score)
	for i := 0; i < cnt_stud; i++ {
		fmt
	}
}
*/

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

func main() {
	/*
		fmt.Println("Hello world")
		fmt.Println("222")
		arr := []int{1, 2, 3, 4}
		arr = append(arr, 5, 6, 7)
		fmt.Println(Pow(2, 3))
	*/
	//a, b := 3, 2
	//fmt.Println(div(a, b))
	//drawCross(5, 5)
	//fmt.Println(getElement("12345", -2))
	//fmt.Println(hasPrefix("abba", "abr"))
	//partString("12345", 1, 3)
	//printABC()
	//sendAB("email", "Andrew")
	//showMatrix()
	//rotateArr()
	//splitByPairs("1122334")
	//n := []int{0, 1, 1, 0}
	//fmt.Println(findMaxConsecutiveOnes(n))
	//testMap()
	//testMapCompare()
	//testDiffMaps()
	//convertDateTime("2000-01-01 00:01:02")
	//fmt.Println(hasNonAlphanumeric("hello5"))
	//arr := []int{5, 3, 4, 1, 2}
	//fmt.Println(missingNumber(arr))
	//fmt.Println(missingNumber2(arr))
	//fmt.Println(isPowerOfTwo(18))
	//testCaller()
	//var d Doer ; d.Do()

	//var null NotNil = new(interface{})
	//fmt.Println(null)

	//go_3()
	//q1()
	//e := fmt.Errorf("1 : %w", errors.New("2"))
	//fmt.Println(e.Error())
	//fmt.Println(isJson("{}}"))
	//fmt.Println(check("abcDD"))
	//drawRectangle(3, 4)
	//fmt.Println(trimPrefix("abcde", "ab"))
	//letterCounter("Abb")
	//filterString("some1   bad")
	//fmt.Println(findWordIndexWithPrefix("short text", "text"))
	//fmt.Println(truncateSentence("hello world is a basic hello world program", 3))
	//fmt.Println(compareStr("abc", "aaabbcb"))
	//convertEverySecondLetterToUpper("hello")
	//arr:=[]int{1,2,3,4}
	//moveZeroes([]int{1, 0, 0, 2, 3, 4})
	//isPartArr([5]int{1, 2, 3, 4, 5}, []int{2, 3, 4})
	//houses := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}}
	//getSignal(houses)
	//drawClock(8)
	//fmt.Println(max3(4, 5, 3))
	//fmt.Println(insidePeriod(20, 50, 1, 100))
	//getOperation("/")
	//getTemp(10)
	//convertTo2(10)
	//decomposeNumber(123)
	//decomposeThousand(1234)
	//str := ""
	//str = fmt.Sprint(100)
	//fmt.Println(getSumNumber(234))
	//sum, mul := getSumMulNumber(234)
	//fmt.Println(sum, mul)
	//getDuplicate([]int{4, 2, 3, 4})
	//filterDuplicate([]int{4, 2, 3, 4})

	go_rutin(2)

}
