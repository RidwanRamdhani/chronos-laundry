import { API_BASE, requireAuth, getToken, logout } from "/src/main.js";

requireAuth();

const itemTable = document.getElementById("itemTable");
const addItemBtn = document.getElementById("addItemBtn");
const totalPriceEl = document.getElementById("totalPrice");

let serviceTypes = [];
let servicePrices = {}; // { reguler: [...], express: [...] }
let itemIndex = 0;

// ==============================
// Load service types + prices
// ==============================
async function loadServices() {
    const typesRes = await fetch(`${API_BASE}/service-types`);
    const typesData = await typesRes.json();
    serviceTypes = typesData.data;

    const priceRes = await fetch(`${API_BASE}/service-prices`);
    const priceData = await priceRes.json();

    servicePrices = {};

    priceData.data.forEach((p) => {
        if (!servicePrices[p.service_type]) servicePrices[p.service_type] = [];
        servicePrices[p.service_type].push(p);
    });
}

await loadServices();

// ==============================
// Add Item Row
// ==============================
function addItemRow() {
    const rowId = "item-" + itemIndex++;

    const row = document.createElement("tr");
    row.setAttribute("data-row", rowId);

    row.innerHTML = `
        <td>
            <select class="form-select service-type" required>
                <option value="">Pilih</option>
                ${serviceTypes
                    .map((t) => `<option value="${t}">${t}</option>`)
                    .join("")}
            </select>
        </td>

        <td>
            <select class="form-select item-name" required disabled>
                <option value="">Pilih layanan dulu</option>
            </select>
        </td>

        <td>
            <input type="number" class="form-control quantity" value="1" min="1" required />
        </td>

        <td>
            <input type="text" class="form-control price" value="0" readonly />
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

    // Load item list based on service type
    typeSelect.addEventListener("change", () => {
        const type = typeSelect.value;
        itemSelect.innerHTML = `<option value="">Pilih item</option>`;

        if (!servicePrices[type]) {
            itemSelect.disabled = true;
            return;
        }

        itemSelect.disabled = false;

        servicePrices[type].forEach((p) => {
            itemSelect.innerHTML += `
                <option value="${p.item_name}" data-price="${p.price}">
                    ${p.item_name}
                </option>
            `;
        });
    });

    // Update price & subtotal when item changes
    itemSelect.addEventListener("change", () => {
        const price = parseInt(itemSelect.selectedOptions[0].dataset.price || 0);
        priceInput.value = price;
        subtotalInput.value = price * parseInt(qtyInput.value);
        refreshTotal();
    });

    // Recalculate subtotal when quantity changes
    qtyInput.addEventListener("input", () => {
        const price = parseInt(priceInput.value);
        subtotalInput.value = price * parseInt(qtyInput.value);
        refreshTotal();
    });

    // Remove row
    removeBtn.addEventListener("click", () => {
        row.remove();
        refreshTotal();
    });
}

addItemBtn.addEventListener("click", addItemRow);

// Add one default row
addItemRow();

// ==============================
// Calculate Total
// ==============================
function refreshTotal() {
    let total = 0;

    document.querySelectorAll(".subtotal").forEach((el) => {
        total += parseInt(el.value || 0);
    });

    totalPriceEl.textContent = total.toLocaleString();
}

// ==============================
// Submit
// ==============================
document.getElementById("transactionForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const items = [];

    document.querySelectorAll("tr[data-row]").forEach((row) => {
        const serviceType = row.querySelector(".service-type").value;
        const itemName = row.querySelector(".item-name").value;
        const qty = parseInt(row.querySelector(".quantity").value);
        const price = parseInt(row.querySelector(".price").value);

        items.push({
            service_type: serviceType,
            item_name: itemName,
            quantity: qty,
            unit_price: price,
        });
    });

    const payload = {
        customer_name: document.getElementById("customerName").value.trim(),
        customer_phone: document.getElementById("customerPhone").value.trim(),
        customer_address: document.getElementById("customerAddress").value.trim(),
        notes: document.getElementById("notes").value.trim(),
        pickup_date: document.getElementById("pickupDate").value,
        items: items,
    };

    try {
        const response = await fetch(`${API_BASE}/transactions`, {
            method: "POST",
            headers: {
                "Authorization": "Bearer " + getToken(),
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });

        const result = await response.json();

        if (!response.ok) {
            alert(result.message || "Gagal membuat transaksi");
            return;
        }

        alert("Transaksi berhasil dibuat!");
        window.location.href = "./transactions.html";

    } catch (err) {
        console.error(err);
        alert("Server error");
    }
});
