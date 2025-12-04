const API_BASE = "http://localhost:8080/api/auth/login";

const loginForm = document.getElementById("loginForm");
const loginBtn = document.getElementById("loginBtn");
const errorMsg = document.getElementById("errorMsg");

loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    errorMsg.textContent = "";
    loginBtn.disabled = true;
    loginBtn.textContent = "Loading...";

    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();

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
            errorMsg.textContent = result.error || "Login gagal";
            loginBtn.disabled = false;
            loginBtn.textContent = "Login";
            return;
        }

        // Save token
        const token = result.data.token;
        localStorage.setItem("token", token);

        // Save user info (optional)
        localStorage.setItem("admin_username", result.data.username);
        localStorage.setItem("admin_fullname", result.data.full_name);

        // Redirect
        window.location.href = "/pages/dashboard.html";

    } catch (err) {
        console.error(err);
        errorMsg.textContent = "Tidak bisa terhubung ke server";
    }

    loginBtn.disabled = false;
    loginBtn.textContent = "Login";
});
