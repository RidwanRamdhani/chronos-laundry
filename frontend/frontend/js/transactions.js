import { API_BASE, requireAuth, getToken, logout } from "./src/main.js";

requireAuth();

// UI elements
const adminName = document.getElementById("adminName");
adminName.textContent = localStorage.getItem("admin_fullname") || "";

document.getElementById("logoutBtn").addEventListener("click", () => logout());

const tableBody = document.getElementById("transactionTable");
const statusFilter = document.getElementById("statusFilter");
const applyFilterBtn = document.getElementById("applyFilterBtn");

const prevBtn = document.getElementById("prevBtn");
const nextBtn = document.getElementById("nextBtn");
const pageInfo = document.getElementById("pageInfo");

let currentPage = 1;
let currentStatus = "";
let totalPages = 1;

// =====================
// Load Transactions
// =====================
async function loadTransactions() {
    tableBody.innerHTML = `
        <tr><td colspan="6" class="text-center">Loading...</td></tr>
    `;

    try {
        const url = `${API_BASE}/transactions?page=${currentPage}&limit=10&status=${currentStatus}`;

        const response = await fetch(url, {
            headers: {
                "Authorization": "Bearer " + getToken()
            }
        });

        const result = await response.json();

        if (!response.ok) {
            tableBody.innerHTML = `
                <tr><td colspan="6" class="text-center text-danger">Gagal mengambil data.</td></tr>
            `;
            return;
        }

        const data = result.data;
        totalPages = data.total_pages;

        renderTable(data.data);
        renderPagination();

    } catch (err) {
        console.error(err);
        tableBody.innerHTML = `
            <tr><td colspan="6" class="text-center text-danger">Server error.</td></tr>
        `;
    }
}

// =====================
// Render Table
// =====================
function renderTable(rows) {
    if (rows.length === 0) {
        tableBody.innerHTML = `
            <tr><td colspan="6" class="text-center">Tidak ada data.</td></tr>
        `;
        return;
    }

    tableBody.innerHTML = rows
        .map(
            (trx) => `
        <tr>
            <td>${trx.transaction_code}</td>
            <td>${trx.customer_name}</td>
            <td class="text-capitalize">${trx.status}</td>
            <td>Rp ${trx.total_price.toLocaleString()}</td>
            <td>${new Date(trx.created_at).toLocaleString()}</td>
            <td>
                <a href="./transaction-detail.html?id=${trx.id}" class="btn btn-sm btn-primary">Detail</a>
                <a href="./transaction-update.html?id=${trx.id}" class="btn btn-sm btn-warning">Edit</a>
                <button class="btn btn-sm btn-danger" onclick="deleteTransaction(${trx.id})">Hapus</button>
            </td>
        </tr>
        `
        )
        .join("");
}

// =====================
// Pagination
// =====================
function renderPagination() {
    pageInfo.textContent = `Halaman ${currentPage} dari ${totalPages}`;

    prevBtn.disabled = currentPage <= 1;
    nextBtn.disabled = currentPage >= totalPages;
}

prevBtn.addEventListener("click", () => {
    if (currentPage > 1) {
        currentPage--;
        loadTransactions();
    }
});

nextBtn.addEventListener("click", () => {
    if (currentPage < totalPages) {
        currentPage++;
        loadTransactions();
    }
});

// =======================
// Filter
// =======================
applyFilterBtn.addEventListener("click", () => {
    currentStatus = statusFilter.value;
    currentPage = 1;
    loadTransactions();
});

// =======================
// Delete Transaction
// =======================
window.deleteTransaction = async function (id) {
    if (!confirm("Yakin ingin menghapus transaksi?")) return;

    try {
        const response = await fetch(`${API_BASE}/transactions/${id}`, {
            method: "DELETE",
            headers: { "Authorization": "Bearer " + getToken() }
        });

        if (response.ok) {
            alert("Transaksi berhasil dihapus");
            loadTransactions();
        } else {
            alert("Gagal menghapus transaksi");
        }

    } catch (err) {
        alert("Server error");
    }
};

// Start
loadTransactions();
