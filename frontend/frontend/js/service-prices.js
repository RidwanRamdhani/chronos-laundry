import { API_BASE, requireAuth, getToken } from "/src/main.js";

requireAuth();

const priceTable = document.getElementById("priceTable");

loadPrices();

// ================================
// Ambil semua harga layanan
// ================================
async function loadPrices() {
    priceTable.innerHTML = `
        <tr><td colspan="4" class="text-center">Loading...</td></tr>
    `;

    try {
        const res = await fetch(`${API_BASE}/service-prices`, {
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const data = await res.json();

        if (!res.ok) {
            priceTable.innerHTML = `
                <tr><td colspan="4" class="text-center text-danger">Gagal memuat data</td></tr>
            `;
            return;
        }

        renderPrices(data.data);

    } catch (err) {
        console.error(err);
        priceTable.innerHTML = `
            <tr><td colspan="4" class="text-center text-danger">Server error</td></tr>
        `;
    }
}

// ================================
// Render tabel
// ================================
function renderPrices(items) {
    if (items.length === 0) {
        priceTable.innerHTML = `
            <tr><td colspan="4" class="text-center">Tidak ada data harga layanan</td></tr>
        `;
        return;
    }

    priceTable.innerHTML = "";

    items.forEach((p) => {
        const tr = document.createElement("tr");

        tr.innerHTML = `
            <td>${p.service_type}</td>
            <td>${p.item_name}</td>
            <td>Rp ${p.price.toLocaleString()}</td>
            <td>
                <a href="./service-price-update.html?id=${p.id}" class="btn btn-warning btn-sm me-2">
                    Edit
                </a>
                <button class="btn btn-danger btn-sm" data-id="${p.id}">
                    Hapus
                </button>
            </td>
        `;

        // event tombol hapus
        tr.querySelector("button").addEventListener("click", () => {
            deletePrice(p.id);
        });

        priceTable.appendChild(tr);
    });
}

// ================================
// Hapus harga layanan
// ================================
async function deletePrice(id) {
    if (!confirm("Yakin ingin menghapus harga ini?")) return;

    try {
        const res = await fetch(`${API_BASE}/service-prices/${id}`, {
            method: "DELETE",
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const result = await res.json();

        if (!res.ok) {
            alert(result.message || "Gagal menghapus data");
            return;
        }

        alert("Harga berhasil dihapus!");
        loadPrices(); // refresh table

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
}
