import axios from 'axios';
import { BrowserRouter, Link, Routes, Route } from "react-router-dom";
import { Button } from 'react-bootstrap';
import { useState } from 'react';
import { Home } from "./Home";
import { Threads } from "./Threads";




export default function App() {
    const [respStr, setRespStr] = useState()

    // call api
    const setCookie = () => {
        axios.get("http://localhost:8080/debug-set-cookie", { withCredentials: true })
            .then((result) => {
                setRespStr("クッキーをセットしました。")
            })
    }

    const readCookie = () => {
        axios.get("http://localhost:8080/debug-read-cookie", { withCredentials: true })
            .then(() => {
                setRespStr("クッキーをサーバ側に送信。サーバのlogを確認してください。")
            })
    }

    return (
        <BrowserRouter>
            <div>
                <Link to="/home">Home</Link><br />
                <Link to="/threads">Threads</Link>
            </div>

            <Routes>
                <Route path="/home" element={<Home />} />
                <Route path="/threads" element={<Threads />} />
            </Routes>

            <Button onClick={() => setCookie()}>set cookie</Button>
            <Button onClick={() => readCookie()}>read cookie</Button>

            <p>{respStr}</p>

        </BrowserRouter>
    );
}