import { createContext, useContext, useRef, useState } from "react"
import { GetBookPath } from "../wailsjs/go/main/App"
import { useNavigate } from "react-router-dom"
import { BookDataContext } from "./BookDataContext";


export default function HomePage() {
    const ctx = useContext(BookDataContext)
    const navigator = useNavigate();
    const txtPath = async () => {
        const data = await GetBookPath()

        if (!data || data == "") return 
        ctx?.setBookData(data)
        navigator("/open-txt")
    }
    return (
            <button onClick={txtPath}>
                Open Book
            </button>
    )
}