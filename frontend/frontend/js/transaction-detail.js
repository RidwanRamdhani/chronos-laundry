import { API_BASE, requireAuth, getToken, logout } from "/src/main.js";

requireAuth();

const urlParams = new URLSearchParams(window.location.search);
const trxId = urlParams.get("id");

if (!trxId) {
    alert("ID transaksi tidak ditemukan");
    window.location.href = "./transactions.html";
}

// Elements
const trxCode = document.getElementById("trxCode");
const trxStatus = document.getElementById("trxStatus");
const trxTotal = document.getElementById("trxTotal");
const trxPaid = document.getElementById("trxPaid");
const trxPickup = document.getElementById("trxPickup");

const custName = document.getElementById("custName");
const custPhone = document.getElementById("custPhone");
const custAddress = document.getElementById("custAddress");
const custNotes = document.getElementById("custNotes");

const itemTable = document.getElementById("itemTable");
const statusHistory = document.getElementById("statusHistory");

const editBtn = document.getElementById("editBtn");
const deleteBtn = document.getElementById("deleteBtn");
const updateStatusBtn = document.getElementById("updateStatusBtn");

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

        trxCode.textContent = data.transaction_code;
        trxStatus.textContent = data.status;
        trxTotal.textContent = data.total_price.toLocaleString();
        trxPaid.textContent = data.is_paid ? "Sudah Dibayar" : "Belum Dibayar";
        trxPickup.textContent = data.pickup_date;

        custName.textContent = data.customer_name;
        custPhone.textContent = data.customer_phone;
        custAddress.textContent = data.customer_address;
        custNotes.textContent = data.notes || "-";

        renderItems(data.items);
        renderStatusHistory(data.status_history);

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

function renderStatusHistory(history) {
    if (history.length === 0) {
        statusHistory.innerHTML = "<li class='list-group-item'>Belum ada riwayat</li>";
        return;
    }

    statusHistory.innerHTML = history.map((row) => {
        return `
            <li class="list-group-item">
                <strong>${row.previous_status}</strong> → 
                <strong>${row.new_status}</strong> <br />
                <small>Alasan: ${row.reason}</small><br />
                <small>Oleh: ${row.changed_by}</small><br />
                <small>${new Date(row.created_at).toLocaleString()}</small>
            </li>
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
updateStatusBtn.addEventListener("click", async () => {
    const statuses = ["antrian", "mencuci", "menyetrika", "siap_diambil", "selesai"];

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
