import { API_BASE, getToken, requireAuth, logout } from "./src/main.js";

requireAuth();  // pastikan sudah login

const adminName = document.getElementById("adminName");
const logoutBtn = document.getElementById("logoutBtn");

adminName.textContent = localStorage.getItem("admin_fullname") || "";

logoutBtn.addEventListener("click", () => logout());

// === Fetch Dashboard Data ===
async function loadDashboard() {
    try {
        const response = await fetch(`${API_BASE}/transactions/dashboard`, {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + getToken()
            }
        });

        const result = await response.json();

        if (!response.ok) {
            alert("Gagal mengambil data dashboard");
            return;
        }

        const data = result.data;

        document.getElementById("total").textContent = data.total || 0;
        document.getElementById("antrian").textContent = data.antrian || 0;
        document.getElementById("mencuci").textContent = data.mencuci || 0;
        document.getElementById("menyetrika").textContent = data.menyetrika || 0;
        document.getElementById("siap").textContent = data.siap_diambil || 0;
        document.getElementById("selesai").textContent = data.selesai || 0;

    } catch (err) {
        console.error(err);
        alert("Server error: tidak bisa mengambil data");
    }
}

loadDashboard();
