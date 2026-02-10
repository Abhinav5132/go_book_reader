import { createContext, useContext, useEffect, useRef, useState } from "react"
import { GetBookPath } from "../wailsjs/go/main/App"
import { GetFirstTenRecentBooks } from "../wailsjs/go/main/App"
import { useNavigate } from "react-router-dom"
import { BookDataContext } from "./BookDataContext";
import { models } from "../wailsjs/go/models";


export default function HomePage() {
    const ctx = useContext(BookDataContext)
    const navigator = useNavigate();
    const [recentBooks, setRecentBooks] = useState<models.Book[]>([])


    useEffect(() => {
            let cancelled = false
            const loadRecents = async () => {
                const recents = await GetFirstTenRecentBooks()
                if (!cancelled && recents && recents.length > 0) {
                    setRecentBooks(recents)
                }
            } 
            loadRecents()

            return () => {
                cancelled = true
            }
        }
    , []) // ensures that the effect is run only once


    const txtPath = async () => {
        const data = await GetBookPath()

        if (!data || data == "") return 
        ctx?.setBookData(data)
        navigator("/open-txt")
    }

    
    return (
        <div>
            <button onClick={txtPath}>
                Open Book
            </button>

            <div>
                <h3>Recents</h3>
                <ul>
                    {
                        recentBooks?.map((book) => (
                            <li key={book.Id}>
                                <div>
                                    <div>{book.Name}</div>
                                    <div>{book.Path}</div>
                                    <div>{book.LastAccessed}</div>    
                                </div>    
                            </li>
                        ))
                    }
                </ul>
            </div>
        </div>
    )
}