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
                <option value="">Pilih Layanan</option>
                ${serviceTypes
                    .map((t) => `<option value="${t}">${t.charAt(0).toUpperCase() + t.slice(1)}</option>`)
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
            <input type="text" class="form-control price" value="Rp 0" readonly />
        </td>

        <td>
            <input type="text" class="form-control subtotal" value="Rp 0" readonly />
        </td>

        <td class="text-center">
            <button type="button" class="btn-remove-item remove-btn">
                <i class="fas fa-times"></i>
            </button>
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
        priceInput.value = `Rp ${price.toLocaleString('id-ID')}`;
        const qty = parseInt(qtyInput.value);
        subtotalInput.value = `Rp ${(price * qty).toLocaleString('id-ID')}`;
        refreshTotal();
    });

    // Recalculate subtotal when quantity changes
    qtyInput.addEventListener("input", () => {
        const priceText = priceInput.value.replace(/[^0-9]/g, '');
        const price = parseInt(priceText || 0);
        const qty = parseInt(qtyInput.value);
        subtotalInput.value = `Rp ${(price * qty).toLocaleString('id-ID')}`;
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
        const value = el.value.replace(/[^0-9]/g, '');
        total += parseInt(value || 0);
    });

    totalPriceEl.textContent = total.toLocaleString('id-ID');
}

// ==============================
// Submit
// ==============================
document.getElementById("createForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const items = [];

    document.querySelectorAll("tr[data-row]").forEach((row) => {
        const serviceType = row.querySelector(".service-type").value;
        const itemName = row.querySelector(".item-name").value;
        const qty = parseInt(row.querySelector(".quantity").value);
        const priceText = row.querySelector(".price").value.replace(/[^0-9]/g, '');
        const price = parseInt(priceText);

        if (serviceType && itemName && qty > 0 && price > 0) {
            items.push({
                service_type: serviceType,
                item_name: itemName,
                quantity: qty,
                unit_price: price,
            });
        }
    });

    if (items.length === 0) {
        alert("Harap tambahkan minimal 1 item cucian!");
        return;
    }

    const paymentStatus = document.querySelector('input[name="paymentStatus"]:checked').value === 'true';

    const payload = {
        customer_name: document.getElementById("customerName").value.trim(),
        customer_phone: document.getElementById("customerPhone").value.trim(),
        customer_address: document.getElementById("customerAddress").value.trim(),
        notes: document.getElementById("notes").value.trim(),
        pickup_date: document.getElementById("pickupDate").value,
        payment_status: paymentStatus,
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

        alert("✅ Transaksi berhasil dibuat!");
        window.location.href = "./transactions.html";

    } catch (err) {
        console.error(err);
        alert("❌ Server error: " + err.message);
    }
});
