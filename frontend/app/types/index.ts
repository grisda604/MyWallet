export interface Wallet {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    name: string;
    type: string;
    balance: number;
    color: string;
    icon: string;
    currency?: string;
}

export interface Transaction {
  ID: number;
  CreatedAt: string; // GORM ส่งมาเป็น CreatedAt (หรือ date ตามที่คุณตั้งใน model)
  date: string;      // ใน Model Go เรามี field "date" แยกต่างหาก
  amount: number;
  type: string;      // 'income', 'expense', 'transfer'
  note: string;
  // Relations
  wallet_id: number;
  wallet?: Wallet;   // เพราะ Backend Preload มาให้
  category_id?: number;
  category?: Category;
  target_wallet_id?: number;
  target_wallet?: Wallet;
}
export interface Category {
    ID : number;
    CreateAt : string;
    UpdateAt : string;
    name : string;
    type : string;
    color : string;
    icon : string;
}