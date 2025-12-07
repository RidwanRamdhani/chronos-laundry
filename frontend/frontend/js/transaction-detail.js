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
const updateStatusBtn = document.getElementById("updateStatusBtn");
const newStatusSelect = document.getElementById("newStatusSelect");
const statusReason = document.getElementById("statusReason");

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

        // Set status with appropriate class
        statusText.textContent = data.status;
        statusText.className = 'status-badge ' + getStatusClass(data.status);
        
        totalPrice.textContent = data.total_price.toLocaleString();
        
        // Set payment status with appropriate class
        paymentStatus.textContent = data.is_paid ? "Sudah Dibayar" : "Belum Dibayar";
        paymentStatus.className = 'payment-badge ' + (data.is_paid ? 'payment-paid' : 'payment-unpaid');
        
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

// Helper function to get status class
function getStatusClass(status) {
    const statusMap = {
        'Queued': 'status-queued',
        'Washing': 'status-washing',
        'Ironing': 'status-ironing',
        'Ready to pick up': 'status-ready',
        'Completed': 'status-completed'
    };
    return statusMap[status] || 'status-queued';
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
updateStatusBtn.addEventListener("click", async () => {
    const newStatus = newStatusSelect.value;
    const reason = statusReason.value.trim();

    if (!newStatus) {
        alert("Pilih status baru terlebih dahulu");
        return;
    }

    if (!reason) {
        alert("Masukkan alasan perubahan status");
        return;
    }

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
        
        // Reset form
        newStatusSelect.value = "";
        statusReason.value = "";
        
        // Reload data
        loadTransactionDetail();

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});


// Start
loadTransactionDetail();
