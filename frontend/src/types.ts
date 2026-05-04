export interface ProductItem {
  id: string;
  name: string;
  price: number;
  categories: string[];
  calories: number;
  protein: number;
  fiber: number;
  shops: ShopName[];
}

export enum ShopName {
  Maxima = "Maxima",
  Iki = "Iki",
  Lidl = "Lidl",
  Rimi = "Rimi",
}