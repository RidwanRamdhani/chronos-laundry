const API_BASE = "http://localhost:8080/api/auth/login";

const loginForm = document.getElementById("loginForm");
const loginBtn = document.getElementById("loginBtn");
const errorMsg = document.getElementById("errorMsg");
const errorText = document.getElementById("errorText");
const btnText = document.getElementById("btnText");
const btnSpinner = document.getElementById("btnSpinner");

// Show error message with animation
function showError(message) {
    errorText.textContent = message;
    errorMsg.classList.add("show");
}

// Hide error message
function hideError() {
    errorMsg.classList.remove("show");
    errorText.textContent = "";
}

// Set loading state
function setLoading(isLoading) {
    loginBtn.disabled = isLoading;
    if (isLoading) {
        btnText.classList.add("d-none");
        btnSpinner.classList.remove("d-none");
    } else {
        btnText.classList.remove("d-none");
        btnSpinner.classList.add("d-none");
    }
}

loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    hideError();
    setLoading(true);

    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();

    // Basic validation
    if (!username || !password) {
        showError("Username dan password harus diisi");
        setLoading(false);
        return;
    }

    try {
        const response = await fetch(API_BASE, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        const result = await response.json();

        if (!response.ok) {
            showError(result.error || "Login gagal. Periksa username dan password Anda.");
            setLoading(false);
            return;
        }

        // Save token
        const token = result.data.token;
        localStorage.setItem("token", token);

        // Save user info (optional)
        localStorage.setItem("admin_username", result.data.username);
        localStorage.setItem("admin_fullname", result.data.full_name);

        // Show success message briefly before redirect
        btnText.innerHTML = '<i class="fas fa-check-circle me-2"></i>Login Berhasil!';
        btnText.classList.remove("d-none");
        btnSpinner.classList.add("d-none");

        // Redirect after short delay
        setTimeout(() => {
            window.location.href = "/pages/dashboard.html";
        }, 800);

    } catch (err) {
        console.error("Login error:", err);
        showError("Tidak dapat terhubung ke server. Pastikan server backend berjalan.");
        setLoading(false);
    }
});

// Clear error message when user starts typing
document.getElementById("username").addEventListener("input", hideError);
document.getElementById("password").addEventListener("input", hideError);

// Check if already logged in
window.addEventListener("DOMContentLoaded", () => {
    const token = localStorage.getItem("token");
    if (token) {
        // Optionally verify token validity here
        // For now, just redirect if token exists
        window.location.href = "/pages/dashboard.html";
    }
});
