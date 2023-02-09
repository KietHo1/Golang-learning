package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SNode struct {
	Info int //interface{}: dùng cho mọi kiểu dữ liệu
	Next *SNode
}

type SList struct {
	Head *SNode
	Tail *SNode
}

func initList(sl *SList) {
	//Khoi tao danh sach rong
	sl.Head = nil //nil == NULL
	sl.Tail = nil
}

func IsEmpty(sl SList) int {
	if sl.Head == nil {
		return 1
	} else {
		return 0
	}
}

func CreateSNode(x int) *SNode {
	var p *SNode = &SNode{Info: x, Next: nil};
	if (p == nil){
		fmt.Printf("Khong the cap nhat nut moi!");
		return nil;
	}
	return p;
}

func ShowSNode(p *SNode){
	fmt.Printf("%4d", p.Info);
}


func ShowSList(sl SList){
	if (IsEmpty(sl) == 1){
		fmt.Printf("\nDanh sach rong!");
	}
	fmt.Printf("\nNoi dung cua danh sach la: ");
	for p := sl.Head; p != nil; p = p.Next {
		ShowSNode(p);
	}
}

func FindSNode(sl SList, x int) *SNode{
	for p := sl.Head; p != nil; p = p.Next{
		if (p.Info == x){
			return p;
		}
	}
	return nil;
}

func InsertAfter(sl SList, q *SNode, p *SNode) int{
	if (q == nil || p == nil){
		return 0;
	}else
	{
		p.Next = q.Next;
		q.Next = p;
		if (sl.Tail == q){
			sl.Tail = p;
		}
	}
	return 1;
}

func InsertTail(sl SList, p *SNode) int{
	if (p == nil){
		return 0;
	}
	if (IsEmpty(sl) == 1){
		sl.Head = p;
		sl.Tail = p;
	}else
	{
		sl.Tail.Next = p;
		sl.Tail = p;
	}
	return 1;
}

func InsertHead(sl SList, p *SNode) int{
	if (p == nil){
		return 0;
	}
		
	if (IsEmpty(sl) == 1){
		sl.Head = p;
		sl.Tail = p;
	}else
	{
		p.Next = sl.Head;
		sl.Head = p;
	}
	return 1;
}

func DeleteHead(sl SList, x int) int{
	if (IsEmpty(sl) == 1){
		return 0; //Khong thuc hien duoc
	}
	var p *SNode = sl.Head;
	sl.Head = sl.Head.Next;
	if (sl.Head == nil){
		sl.Tail = nil;
	}
	x = p.Info; //Lay thong tin cua nut bi huy
	return 1; //Thuc hien thanh cong
}

func DeleteTail(sl SList, x int) int{
	if (IsEmpty(sl) == 1){
		return 0; //Khong thuc hien duoc
	}
	var q *SNode= sl.Head;
	var p *SNode = sl.Tail;
	if (p == q){
		sl.Head = nil;
		sl.Tail = nil;
	}else
	{
		for (q.Next != p){
			//Tim nut ke truoc nut Tail
			q = q.Next;
		} 
		sl.Tail = q;
		sl.Tail.Next = nil;
	}
	x = p.Info; //Lay thong tin cua nut bi huy
	return 1; //Thuc hien thanh cong
}

func DeleteAfter(sl SList, q *SNode, x int) int{
	if (q == nil || q.Next == nil){
		return 0; //Khong thuc hien duoc
	}
	var p *SNode = q.Next;
	q.Next = p.Next;
	if (p == sl.Tail){
		sl.Tail = q;
	}
	x = p.Info; //Lay thong tin cua nut bi huy
	return 1; //Thuc hien thanh cong
}

func DeleteSNodeX(sl SList, x int) int{
	if (IsEmpty(sl) == 1){
		return 0; //Khong thuc hien duoc
	}
	var p *SNode = sl.Head;
	var q *SNode = nil; //se tro den nut ke truoc p
	for ((p != nil) && (p.Info != x)){//Vong lap tim nut p chua x, q la nut ke truoc p
		q = p;
		p = p.Next;
	}
	if (p == nil){
		//khong tim thay phan tu co khoa bang x
		return 0; //Khong thuc hien duoc
	} 
		
	if (p == sl.Head){
		DeleteHead(sl, x);
		//p co khoa bang x la nut dau danh sach
	}else // xoa nut p co khoa x nam ke sau nut q
	{
		DeleteAfter(sl, q, x);
	}
	return 1; //Thuc hien thanh cong
}

func DeleteSList(sl SList) int{
	if (IsEmpty(sl) == 1){
		return 0; //Khong thuc hien duoc
	}
	for (sl.Head != nil){
		sl.Head = sl.Head.Next;
	}
	sl.Tail = nil;
	return 1; //Thuc hien thanh cong
}

func CreateAutoSList(sl SList){
	var n, kq int;
	for {
			fmt.Printf("\nCho biet so phan tu cua danh sach: ");
			fmt.Scanf("%d", &n);
			if (n <= 0) {
				break
			}
		}
	rand.Seed(time.Now().UnixNano());
	i := 1;
	for (i <= n){
		var x = rand.Intn(99 - 1 + 1) + 1;
		var p *SNode = CreateSNode(x);
		kq = InsertHead(sl, p);
		if (kq == 0){
			fmt.Printf("\nKhong them duoc nut moi!");
		}
		i++;
	}
}