package main

import "bytes"

import "bufio"

import "fmt"

var str = "链表（Linked list）是一种常见的基础数据结构,是一种线性表,但是并不会按线性的顺序存储数据,而是在每一个节点里存到下一个节点的指针(Pointer).由于不必须按顺序存储,链表在插入的时候可以达到O(1)的复杂度,比另一种线性表顺序表快得多,但是查找一个节点或者访问特定编号的节点则需要O(n)的时间,而顺序表相应的时间复杂度分别是O(logn)和O(1)."

func read() {
	fmt.Println("------read------")
	data := []byte(str)
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}

func readByte() {
	fmt.Println("------readByte------")
	data := []byte(str)
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}

func readBytes() {
	fmt.Println("------readBytes------")
	data := []byte(str)
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}

func readLine() {
	fmt.Println("------readLine------")
	data := []byte("Hello\nLibre")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}

func readRune() {
	fmt.Println("------readRune------")
	data := []byte(str)
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}

func readSlice() {
	fmt.Println("------readSlice------")
	data := []byte("Hello,Libre")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}

func readString() {
	fmt.Println("------readString------")
	data := []byte("Hello,Libre")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println(line, err)
}

func buffered() {
	fmt.Println("------buffered------")
	data := []byte("Hello,Libre")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [8]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println(rn)
	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println(rn)
}

func peek() {
	fmt.Println("------peek------")
	data := []byte("Hello,Libre")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	bl, err := r.Peek(2)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(8)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
}

func main() {
	read()
	readByte()
	readBytes()
	readLine()
	readRune()
	readSlice()
	readString()
	buffered()
	peek()
}
