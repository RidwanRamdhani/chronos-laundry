import { API_BASE, requireAuth, getToken } from "/src/main.js";

requireAuth();

// Form elements
const serviceTypeSelect = document.getElementById("serviceType");
const itemNameInput = document.getElementById("itemName");
const priceInput = document.getElementById("price");
const priceForm = document.getElementById("priceForm");

// Ambil ID dari URL
const params = new URLSearchParams(window.location.search);
const priceId = params.get("id");

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

// ================================
// Load Price Detail
// ================================
async function loadPriceDetail() {
    try {
        const res = await fetch(`${API_BASE}/service-prices/${priceId}`, {
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const data = await res.json();

        if (!res.ok) {
            alert("Gagal memuat data harga");
            return;
        }

        const p = data.data;

        // Isi otomatis form
        serviceTypeSelect.value = p.service_type;
        itemNameInput.value = p.item_name;
        priceInput.value = p.price;

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
}

// Load data awal
await loadServiceTypes();
await loadPriceDetail();

// ================================
// Submit Update Form
// ================================
priceForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const payload = {
        service_type: serviceTypeSelect.value,
        item_name: itemNameInput.value.trim(),
        price: parseInt(priceInput.value),
    };

    try {
        const res = await fetch(`${API_BASE}/service-prices/${priceId}`, {
            method: "PUT",
            headers: {
                "Authorization": "Bearer " + getToken(),
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });

        const result = await res.json();

        if (!res.ok) {
            alert(result.message || "Gagal memperbarui harga");
            return;
        }

        alert("Harga berhasil diperbarui!");
        window.location.href = "./service-prices.html";

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});
