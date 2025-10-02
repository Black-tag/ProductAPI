import React, { useEffect, useState } from "react";

export default function ProductList() {
  const [products, setProducts] = useState([]);
  const [newProduct, setNewProduct] = useState({ Name: "", Price: "" });
  const [editProduct, setEditProduct] = useState(null);
  const [message, setMessage] = useState("");

  // Get token from localStorage
  const token = localStorage.getItem("token");

  // Fetch products
  const fetchProducts = async () => {
    try {
      const res = await fetch("http://localhost:8090/api/v1/product", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) {
        if (res.status === 204) {
          setProducts([]);
          return;
        }
        const errData = await res.json().catch(() => ({}));
        console.error("Error fetching products:", errData);
        return;
      }

      const data = await res.json();
      setProducts(data);
    } catch (err) {
      console.error("Error fetching products:", err);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  // Create product
  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8090/api/v1/product", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(newProduct),
      });

      if (!res.ok) {
        const errData = await res.json().catch(() => ({}));
        setMessage(`Error: ${errData.error || "Unauthorized"}`);
        return;
      }

      setNewProduct({ Name: "", Price: "" });
      setMessage("Product created successfully!");
      fetchProducts();
    } catch (err) {
      console.error("Error creating product:", err);
    }
  };

  // Update product
  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(
        `http://localhost:8090/api/v1/product/${editProduct.ID}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify(editProduct),
        }
      );

      if (!res.ok) {
        const errData = await res.json().catch(() => ({}));
        setMessage(`Error updating product: ${errData.error || res.status}`);
        return;
      }

      setEditProduct(null);
      fetchProducts();
    } catch (err) {
      console.error("Error updating product:", err);
    }
  };

  // Delete product
  const handleDelete = async (ID) => {
    try {
      const res = await fetch(`http://localhost:8090/api/v1/product/${ID}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) {
        const errData = await res.json().catch(() => ({}));
        setMessage(`Error deleting product: ${errData.error || res.status}`);
        return;
      }

      fetchProducts();
    } catch (err) {
      console.error("Error deleting product:", err);
    }
  };

  return (
    <div>
      <h2>Products</h2>
      {message && <p>{message}</p>}

      {/* Create Product Form */}
      <form onSubmit={handleCreate} style={{ marginBottom: "20px" }}>
        <input
          type="text"
          placeholder="Name"
          value={newProduct.Name}
          onChange={(e) =>
            setNewProduct({ ...newProduct, Name: e.target.value })
          }
          required
        />
        <input
          type="number"
          placeholder="Price"
          value={newProduct.Price}
          onChange={(e) =>
            setNewProduct({ ...newProduct, Price: parseFloat(e.target.value) })
          }
          required
        />
        <button type="submit">Add Product</button>
      </form>

      {/* Product Table */}
      <table border="1" cellPadding="10" style={{ borderCollapse: "collapse" }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Price</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {products.length === 0 ? (
            <tr>
              <td colSpan="4">No products available.</td>
            </tr>
          ) : (
            products.map((p) => (
              <tr key={p.ID}>
                <td>{p.ID}</td>
                <td>
                  {editProduct && editProduct.ID === p.ID ? (
                    <input
                      type="text"
                      value={editProduct.Name}
                      onChange={(e) =>
                        setEditProduct({ ...editProduct, Name: e.target.value })
                      }
                    />
                  ) : (
                    p.Name
                  )}
                </td>
                <td>
                  {editProduct && editProduct.ID === p.ID ? (
                    <input
                      type="number"
                      value={editProduct.Price}
                      onChange={(e) =>
                        setEditProduct({
                          ...editProduct,
                          Price: parseFloat(e.target.value),
                        })
                      }
                    />
                  ) : (
                    p.Price
                  )}
                </td>
                <td>
                  {editProduct && editProduct.ID === p.ID ? (
                    <>
                      <button onClick={handleUpdate}>Save</button>
                      <button onClick={() => setEditProduct(null)}>Cancel</button>
                    </>
                  ) : (
                    <>
                      <button onClick={() => setEditProduct(p)}>Edit</button>
                      <button onClick={() => handleDelete(p.ID)}>Delete</button>
                    </>
                  )}
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}
