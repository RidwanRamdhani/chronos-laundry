// main.js
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';

// Base API
export const API_BASE = "http://localhost:8080/api";

// Ambil token dari localStorage
export function getToken() {
    return localStorage.getItem("token");
}

// Redirect ke login jika belum login
export function requireAuth() {
    const token = localStorage.getItem("token");
    if (!token) {
        window.location.href = "/pages/login.html";
    }
}

// Logout & hapus semua data auth
export function logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("admin_username");
    localStorage.removeItem("admin_fullname");
    window.location.href = "/pages/login.html";
}
