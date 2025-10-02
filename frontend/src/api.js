const BASE_URL = "http://localhost:8090/api/v1";

export async function registerUser(data) {
  const res = await fetch(`${BASE_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  return res.json();
}

export async function loginUser(data) {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  return res.json();
}

export async function getProducts() {
  const res = await fetch(`${BASE_URL}/product`);
  return res.json();
}

export async function createProduct(data, token) {
  const res = await fetch(`${BASE_URL}/product`, {
    method: "POST",
    headers: { 
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`
    },
    body: JSON.stringify(data),
  });
  return res.json();
}

export async function updateProduct(productID, data, token) {
  const res = await fetch(`${BASE_URL}/product/${productID}`, {
    method: "PUT",
    headers: { 
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`
    },
    body: JSON.stringify(data),
  });
  return res.json();
}

export async function deleteProduct(productID, token) {
  const res = await fetch(`${BASE_URL}/product/${productID}`, {
    method: "DELETE",
    headers: { "Authorization": `Bearer ${token}` },
  });
  return res.json();
}
