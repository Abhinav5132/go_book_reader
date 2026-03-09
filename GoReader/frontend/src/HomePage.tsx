import { createContext, useContext, useEffect, useRef, useState } from "react"
import { GetBookPath, GetBookFromPath, OpenFolderAndCreateALibrary } from "../wailsjs/go/main/App"
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

    const setNewTxtPath = async (path: string) => {
        const data = await GetBookFromPath(path)

        if (!data || data == "") return
        ctx?.setBookData(data)
        navigator("/open-txt")

    }

    const addNewLibraryFromFolder = async() => {
        const response = await OpenFolderAndCreateALibrary()

        if (!response) return
        if (response.Response == "Sucessfully created new library") {
            // this should then redirect to the library page 
            navigator(`/libraryPage/${response.LibraryId}`)
        } else {
            alert(response.Response)
        }
    }

    return (
        <div>
            <button onClick={txtPath}>
                Open Book
            </button>
            <button onClick={addNewLibraryFromFolder}>Add Library from Folder</button>
            <div>
                <h3>Recents</h3>
                <ul>
                    {
                        recentBooks?.map((book) => (
                            <li key={book.ID}>
                                <div onClick={() => setNewTxtPath(book.Path)}>
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