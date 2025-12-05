import { API_BASE, requireAuth, getToken, logout } from "./src/main.js";

requireAuth();

// Get transaction ID
const params = new URLSearchParams(window.location.search);
const trxId = params.get("id");
document.getElementById("backBtn").href = `./transaction-detail.html?id=${trxId}`;

const itemTable = document.getElementById("itemTable");
const totalPriceEl = document.getElementById("totalPrice");
const addItemBtn = document.getElementById("addItemBtn");

let serviceTypes = [];
let servicePrices = {};
let itemIndex = 0;

// ==============================
// Load service types & prices
// ==============================
async function loadServices() {
    const typesRes = await fetch(`${API_BASE}/service-types`);
    const typesData = await typesRes.json();
    serviceTypes = typesData.data;

    const priceRes = await fetch(`${API_BASE}/service-prices`);
    const priceData = await priceRes.json();

    priceData.data.forEach((p) => {
        if (!servicePrices[p.service_type]) {
            servicePrices[p.service_type] = [];
        }
        servicePrices[p.service_type].push(p);
    });
}

await loadServices();

// ==============================
// Load transaction detail
// ==============================
async function loadTransaction() {
    try {
        const response = await fetch(`${API_BASE}/transactions/${trxId}`, {
            headers: { "Authorization": "Bearer " + getToken() }
        });

        const result = await response.json();

        if (!response.ok) {
            alert("Gagal mengambil data transaksi");
            return;
        }

        const trx = result.data;

        document.getElementById("customerName").value = trx.customer_name;
        document.getElementById("customerPhone").value = trx.customer_phone;
        document.getElementById("customerAddress").value = trx.customer_address;
        document.getElementById("notes").value = trx.notes;
        document.getElementById("pickupDate").value = trx.pickup_date;
        document.getElementById("paymentStatus").value = trx.is_paid ? "true" : "false";

        itemTable.innerHTML = "";
        trx.items.forEach(addItemRowFromData);

        refreshTotal();

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
}

await loadTransaction();

// ==============================
// Add row with existing data
// ==============================
function addItemRowFromData(item) {
    addItemRow(item.service_type, item.item_name, item.quantity, item.unit_price);
}

// ==============================
// Add new blank row
// ==============================
addItemBtn.addEventListener("click", () => addItemRow());

// ==============================
// Add item row (reusable)
// ==============================
function addItemRow(type = "", itemName = "", qty = 1, price = 0) {
    const rowId = "row-" + itemIndex++;

    const row = document.createElement("tr");
    row.setAttribute("data-row", rowId);

    row.innerHTML = `
        <td>
            <select class="form-select service-type" required>
                <option value="">Pilih</option>
                ${serviceTypes
                    .map((t) => `<option value="${t}" ${t === type ? "selected" : ""}>${t}</option>`)
                    .join("")}
            </select>
        </td>

        <td>
            <select class="form-select item-name" required>
                <option value="">Pilih item</option>
            </select>
        </td>

        <td>
            <input type="number" class="form-control quantity" min="1" value="${qty}" />
        </td>

        <td>
            <input type="text" class="form-control price" value="${price}" readonly />
        </td>

        <td>
            <input type="text" class="form-control subtotal" value="0" readonly />
        </td>

        <td>
            <button class="btn btn-danger btn-sm remove-btn">X</button>
        </td>
    `;

    itemTable.appendChild(row);

    const typeSelect = row.querySelector(".service-type");
    const itemSelect = row.querySelector(".item-name");
    const qtyInput = row.querySelector(".quantity");
    const priceInput = row.querySelector(".price");
    const subtotalInput = row.querySelector(".subtotal");
    const removeBtn = row.querySelector(".remove-btn");

    // load items based on type
    function populateItems() {
        const t = typeSelect.value;

        itemSelect.innerHTML = `<option value="">Pilih item</option>`;

        if (servicePrices[t]) {
            servicePrices[t].forEach((p) => {
                itemSelect.innerHTML += `
                    <option value="${p.item_name}" data-price="${p.price}">
                        ${p.item_name}
                    </option>
                `;
            });
        }

        // if editing existing data
        if (itemName) {
            itemSelect.value = itemName;
            const selectedOpt = itemSelect.selectedOptions[0];
            if (selectedOpt) priceInput.value = selectedOpt.dataset.price;
            subtotalInput.value = qty * priceInput.value;
        }
    }

    typeSelect.addEventListener("change", populateItems);
    populateItems();

    itemSelect.addEventListener("change", () => {
        const selected = itemSelect.selectedOptions[0];
        priceInput.value = selected.dataset.price;
        subtotalInput.value = priceInput.value * qtyInput.value;
        refreshTotal();
    });

    qtyInput.addEventListener("input", () => {
        subtotalInput.value = priceInput.value * qtyInput.value;
        refreshTotal();
    });

    removeBtn.addEventListener("click", () => {
        row.remove();
        refreshTotal();
    });
}

// ==============================
// Recalculate total
// ==============================
function refreshTotal() {
    let total = 0;

    document.querySelectorAll(".subtotal").forEach((el) => {
        total += parseInt(el.value || 0);
    });

    totalPriceEl.textContent = total.toLocaleString();
}

// ==============================
// Submit update
// ==============================
document.getElementById("updateForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const items = [];

    document.querySelectorAll("tr[data-row]").forEach((row) => {
        const type = row.querySelector(".service-type").value;
        const name = row.querySelector(".item-name").value;
        const qty = parseInt(row.querySelector(".quantity").value);
        const price = parseInt(row.querySelector(".price").value);

        items.push({
            service_type: type,
            item_name: name,
            quantity: qty,
            unit_price: price,
        });
    });

    const payload = {
        customer_name: document.getElementById("customerName").value.trim(),
        customer_phone: document.getElementById("customerPhone").value.trim(),
        customer_address: document.getElementById("customerAddress").value.trim(),
        notes: document.getElementById("notes").value.trim(),
        is_paid: document.getElementById("paymentStatus").value === "true",
        pickup_date: document.getElementById("pickupDate").value,
        items: items,
    };

    try {
        const response = await fetch(`${API_BASE}/transactions/${trxId}`, {
            method: "PUT",
            headers: {
                "Authorization": "Bearer " + getToken(),
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });

        const result = await response.json();

        if (!response.ok) {
            alert(result.message || "Gagal update transaksi");
            return;
        }

        alert("Transaksi berhasil diperbarui!");
        window.location.href = `./transaction-detail.html?id=${trxId}`;

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});
