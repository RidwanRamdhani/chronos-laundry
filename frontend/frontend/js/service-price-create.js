import { API_BASE, requireAuth, getToken } from "./src/main.js";

requireAuth();

// Elements
const serviceTypeSelect = document.getElementById("serviceType");
const priceForm = document.getElementById("priceForm");

// ================================
// Load Service Types
// ================================
async function loadServiceTypes() {
    try {
        const res = await fetch(`${API_BASE}/service-types`, {
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const data = await res.json();

        if (!res.ok) {
            alert("Gagal memuat jenis layanan");
            return;
        }

        serviceTypeSelect.innerHTML = `
            <option value="">-- Pilih Jenis Layanan --</option>
        `;

        data.data.forEach((type) => {
            serviceTypeSelect.innerHTML += `
                <option value="${type}">${type}</option>
            `;
        });

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
}

loadServiceTypes();

// ================================
// Submit Form
// ================================
priceForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const payload = {
        service_type: serviceTypeSelect.value,
        item_name: document.getElementById("itemName").value.trim(),
        price: parseInt(document.getElementById("price").value),
    };

    try {
        const res = await fetch(`${API_BASE}/service-prices`, {
            method: "POST",
            headers: {
                "Authorization": "Bearer " + getToken(),
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });

        const result = await res.json();

        if (!res.ok) {
            alert(result.message || "Gagal menambahkan harga");
            return;
        }

        alert("Harga berhasil ditambahkan!");
        window.location.href = "./service-prices.html";

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});
