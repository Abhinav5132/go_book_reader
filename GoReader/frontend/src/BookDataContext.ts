import { createContext } from "react";

export type BookDataContextType = {
    bookData: string | null,
    setBookData: React.Dispatch<React.SetStateAction<string | null>>;
}
export const BookDataContext = createContext<BookDataContextType | null>(null);