import React, { useState } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import ProductList from "./components/ProductList";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || null);

  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={<Login setToken={setToken} />}
        />
        <Route path="/register" element={<Register />} />
        <Route
          path="/products"
          element={
            token ? <ProductList token={token} /> : <Navigate to="/" />
          }
        />
      </Routes>
    </Router>
  );
}

export default App;
