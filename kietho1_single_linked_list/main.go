package main

import "fmt"			


func ShowMenu(){
	fmt.Printf("\n************************************************");
	fmt.Printf("\n*                 MENU                         *");
	fmt.Printf("\n*-----------------------------------------------");
	fmt.Printf("\n*  1. Tao nut moi va xuat ra noi dung           *");
	fmt.Printf("\n*  2. Tao danh sach tu dong(ngau nhien)         *");
	fmt.Printf("\n*  3. Xuat noi dung danh sach                   *");
	fmt.Printf("\n*  4. Them node vao dau danh sach               *");
	fmt.Printf("\n*  5. Them node vao cuoi danh sach              *");
	fmt.Printf("\n*  6. Them nut p chua x vao sau nut q chua y    *");
	fmt.Printf("\n*  7. Tim mot nut chua gia tri x bat ky         *");
	fmt.Printf("\n*  8. Xoa node dau                              *");
	fmt.Printf("\n*  9. Xoa node cuoi                             *");
	fmt.Printf("\n* 10. Xoa nut p chua x dang sau nut q           *");
	fmt.Printf("\n* 11. Xoa toan bo danh sach                     *");
	fmt.Printf("\n* 0. Thoat chuong trinh                         *");
	fmt.Printf("\n*************************************************");
}
//===============================================================
func Process(){
	var X, Y int;
	var P, Q *SNode;
	var luaChon, kq int;
	var sl SList;
	initList(&sl);
	for {
	ShowMenu();
		for {
			fmt.Printf("\nHay chon mot chuc nang (0->13): ");
			fmt.Scanf("%d", &luaChon);
			if (luaChon < 0 || luaChon > 13) {
				break
			}
		}
		switch (luaChon){
		case 1:
			fmt.Printf("\nNhap mot so nguyen bat ky: ");
			fmt.Scanf("%d", &X);
			P = CreateSNode(X);
			fmt.Printf("\nNoi dung nut moi vua tao la: ");
			ShowSNode(P);
		case 2:
			initList(&sl);
			CreateAutoSList(sl);
			ShowSList(sl);
		case 3:
			ShowSList(sl);
		case 4:
			fmt.Printf("\nCho biet gia tri nut muon them dau danh sach: ");
			fmt.Scanf("%d", &X);
			P = CreateSNode(X);
			kq = InsertHead(sl, P);
			if (kq == 0){
				fmt.Printf("\nKhong thuc hien duoc viec them vao dau danh sach!");
			}else{
				ShowSList(sl);
			}
		case 5:
			fmt.Printf("\nCho biet gia tri nut muon them cuoi danh sach: ");
			fmt.Scanf("%d", &X);
			P = CreateSNode(X);
			kq = InsertTail(sl, P);
			if (kq == 0){
				fmt.Printf("\nKhong thuc hien duoc viec them vao cuoi danh sach!");
			}else{
				ShowSList(sl);
			}
		case 6:
			fmt.Printf("\nNhap mot so nguyen x bat ky muon them: ");
			fmt.Scanf("%d", &X);
			fmt.Printf("Nhap mot gia tri nut ke truoc: ");
			fmt.Scanf("%d", &Y);
			P = CreateSNode(X);
			Q = FindSNode(sl, Y);
			kq = InsertAfter(sl, Q, P);
			if (kq == 0){
				fmt.Printf("Khong thuc hien duoc viec them node q vao sau node p!");
			}else{
				ShowSList(sl);
			}
		case 7:
			fmt.Printf("\nNhap mot so nguyen x bat ky muon tim: ");
			fmt.Scanf("%d", &X);
			P = FindSNode(sl, X);
			if (P == nil){
				fmt.Printf("\nKhong tim thay nut nao co gia tri bang %d", X);
			}else{
				fmt.Printf("\nTon tai nut co gia tri bang: %d", X);
			}
		case 8:
			kq = DeleteHead(sl, X);
			if (kq == 1){
				fmt.Printf("Thong tin cua danh sach la: ");
				ShowSList(sl);
				fmt.Printf("\nGia tri cua nut dau vua bi xoa la: %d", X);
			}else{
				fmt.Printf("Xoa nut dau khong thanh cong!");
			}
		case 9:{
			kq = DeleteTail(sl, X);
			if (kq == 1){
				fmt.Printf("Thong tin cua danh sach la: ");
				ShowSList(sl);
				fmt.Printf("\nGia tri cua nut cuoi vua bi xoa la: %d", X);
			}else{
				fmt.Printf("Xoa nut cuoi khong thanh cong!");
			}
		}
			
		case 10:
			fmt.Printf("\nNhap mot so nguyen x bat ky muon xoa: ");
			fmt.Scanf("%d", &X);
			fmt.Printf("Nhap mot gia tri nut ke truoc: ");
			fmt.Scanf("%d", &Y);
			Q = FindSNode(sl, Y);
			kq = DeleteAfter(sl, Q, X);
			if (kq == 0){
				fmt.Printf("Khong thuc hien duoc viec nut p dang sau nut q!");
			}else{
				ShowSList(sl);
			}
		case 11:{
			kq = DeleteSList(sl);
			if (kq == 1){
				fmt.Printf("\nXoa toan bo danh sach thanh cong!");
				ShowSList(sl);
			}else{
				fmt.Printf("Xoa toan bo danh sach khong thanh cong!");
				ShowSList(sl);
			}
		}
		default:
        	fmt.Println("Chuc nang so 12 tro di chua co san!");
		}
		if (luaChon != 0) {
			break
		}
	}
}

func main() {
	Process()
}