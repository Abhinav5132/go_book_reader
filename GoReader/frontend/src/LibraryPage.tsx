import { useContext, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { models } from "../wailsjs/go/models";
import { GetBookFromPath, GetLibrary } from "../wailsjs/go/main/App";
import { BookDataContext } from "./BookDataContext";


export default function LibraryPage( ) {
    const { id } = useParams()
    const [library, setLibrary] = useState<models.Library>()
    const navigator = useNavigate()
    const ctx = useContext(BookDataContext)
    
    useEffect(() => {
        let cancelled = false;
        const GetLibraryDetails = async () =>{
            const data = await GetLibrary(Number(id)); //TODO: this type casting might fail handel later
            if (!cancelled && data){
                setLibrary(data)
            }
        }
        GetLibraryDetails()
        return () => {
            cancelled = true
        }  
    }, [id])

    const setNewTxtPath = async (path: string) => {
        const data = await GetBookFromPath(path)

        if (!data || data == "") return
        ctx?.setBookData(data)
        navigator("/open-txt")

    }

    if (!library) {
        return <div>Loading...</div>
    }
    
    return (
        <div>
            <h3> { library?.Name } </h3>
            
            <div>
                {
                    library.Books?.map((book) =>(
                        <li key={book.ID}>
                            <div onClick={() => setNewTxtPath(book.Path)}>
                                <div>{book.Name}</div>
                                <div>{book.Path}</div>
                                <div>{book.LastAccessed}</div>    
                            </div>
                        </li>
                    ))
                }
            </div>        
        </div>
    )
}