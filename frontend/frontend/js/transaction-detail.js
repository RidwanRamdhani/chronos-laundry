import { API_BASE, requireAuth, getToken, logout } from "/src/main.js";

requireAuth();

const urlParams = new URLSearchParams(window.location.search);
const trxId = urlParams.get("id");

if (!trxId) {
    alert("ID transaksi tidak ditemukan");
    window.location.href = "./transactions.html";
}

// Elements - Updated to match HTML IDs
const statusText = document.getElementById("statusText");
const totalPrice = document.getElementById("totalPrice");
const paymentStatus = document.getElementById("paymentStatus");
const pickupDate = document.getElementById("pickupDate");

const customerName = document.getElementById("customerName");
const customerPhone = document.getElementById("customerPhone");
const customerAddress = document.getElementById("customerAddress");
const notes = document.getElementById("notes");

const itemTable = document.getElementById("itemTable");

const editBtn = document.getElementById("editBtn");
const deleteBtn = document.getElementById("deleteBtn");
const advanceStatusBtn = document.getElementById("advanceStatusBtn");

// ===============================
// Load Transaction Detail
// ===============================
async function loadTransactionDetail() {
    try {
        const response = await fetch(`${API_BASE}/transactions/${trxId}`, {
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const result = await response.json();

        if (!response.ok) {
            alert("Gagal mengambil detail transaksi");
            return;
        }

        const data = result.data;

        statusText.textContent = data.status;
        totalPrice.textContent = data.total_price.toLocaleString();
        paymentStatus.textContent = data.is_paid ? "Sudah Dibayar" : "Belum Dibayar";
        pickupDate.textContent = data.pickup_date;

        customerName.textContent = data.customer_name;
        customerPhone.textContent = data.customer_phone;
        customerAddress.textContent = data.customer_address;
        notes.textContent = data.notes || "-";

        renderItems(data.items);

        editBtn.href = `./transaction-update.html?id=${data.id}`;

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
}

function renderItems(items) {
    itemTable.innerHTML = items.map((item) => {
        return `
            <tr>
                <td>${item.service_type}</td>
                <td>${item.item_name}</td>
                <td>${item.quantity}</td>
                <td>Rp ${item.unit_price.toLocaleString()}</td>
                <td>Rp ${item.subtotal.toLocaleString()}</td>
            </tr>
        `;
    }).join("");
}


// ===============================
// Delete Transaction
// ===============================
deleteBtn.addEventListener("click", async () => {
    if (!confirm("Yakin ingin menghapus transaksi ini?")) return;

    try {
        const response = await fetch(`${API_BASE}/transactions/${trxId}`, {
            method: "DELETE",
            headers: { "Authorization": "Bearer " + getToken() }
        });

        if (!response.ok) {
            alert("Gagal menghapus transaksi");
            return;
        }

        alert("Transaksi berhasil dihapus");
        window.location.href = "./transactions.html";

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});

// ===============================
// Update Status
// ===============================
advanceStatusBtn.addEventListener("click", async () => {
    const statuses = ["Queued", "Washing", "Ironing", "Ready to pick up", "Completed"];

    const newStatus = prompt("Masukkan status baru:\n" + statuses.join(", "));

    if (!newStatus) return;
    if (!statuses.includes(newStatus)) {
        alert("Status tidak valid");
        return;
    }

    const reason = prompt("Alasan perubahan status:");
    if (!reason) return;

    try {
        const response = await fetch(`${API_BASE}/transactions/${trxId}/status`, {
            method: "PUT",
            headers: {
                "Authorization": "Bearer " + getToken(),
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ new_status: newStatus, reason })
        });

        const result = await response.json();

        if (!response.ok) {
            alert(result.message || "Gagal update status");
            return;
        }

        alert("Status berhasil diperbarui");
        loadTransactionDetail();

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});


// Start
loadTransactionDetail();
